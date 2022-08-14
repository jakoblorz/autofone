package writer

import (
	"sync"
	"time"

	"github.com/jakoblorz/autofone/pkg/log"

	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets"
	"github.com/jakoblorz/autofone/packets/process"
)

const (
	motionDebouncerInterval = 250 * time.Millisecond
	packetDebouncerInterval = 1 * time.Second
)

type writer interface {
	write(m *process.M)
}

type packetDebouncer struct {
	mx       sync.Locker
	timer    *time.Timer
	interval time.Duration

	currentPacket *process.M
}

func NewPacketDebouncer(ch writer, interval time.Duration) *packetDebouncer {
	if interval == 0 {
		interval = packetDebouncerInterval
	}
	pdc := &packetDebouncer{
		mx:       &sync.Mutex{},
		interval: packetDebouncerInterval,
	}
	pdc.timer = time.AfterFunc(interval, func() { pdc.WriteTo(ch) })
	return pdc
}

func NewMotionDebouncer(ch writer, interval time.Duration) *motionDebouncer {
	if interval == 0 {
		interval = motionDebouncerInterval
	}
	mdc := &motionDebouncer{
		mx: &sync.Mutex{},

		packages_21: make([]*packets.PacketMotionData21, 0),
		packages_22: make([]*packets.PacketMotionData22, 0),
	}
	mdc.timer = time.AfterFunc(motionDebouncerInterval, func() { mdc.WriteTo(ch) })
	return mdc
}

func (dbc *packetDebouncer) Write(m *process.M) {
	dbc.mx.Lock()
	defer dbc.mx.Unlock()

	dbc.currentPacket = m
}

func (dbc *packetDebouncer) WriteTo(ch writer) {
	dbc.mx.Lock()
	defer dbc.mx.Unlock()
	defer func() {
		dbc.timer = time.AfterFunc(dbc.interval, func() {
			dbc.WriteTo(ch)
		})
	}()
	if dbc.currentPacket == nil {
		return
	}

	ch.write(dbc.currentPacket)
	dbc.currentPacket = nil
}

type motionDebouncer struct {
	mx    sync.Locker
	timer *time.Timer

	h *packets.PacketHeader

	packages_21 []*packets.PacketMotionData21
	packages_22 []*packets.PacketMotionData22
}

func (dbc *motionDebouncer) Write(m *process.M) {
	dbc.mx.Lock()
	defer dbc.mx.Unlock()

	dbc.h = &m.Header

	if m.Header.PacketFormat == constants.PacketFormat_2021 {
		dbc.packages_21 = append(dbc.packages_21, m.Pack.(*packets.PacketMotionData21))
	}
	if m.Header.PacketFormat == constants.PacketFormat_2022 {
		dbc.packages_22 = append(dbc.packages_22, m.Pack.(*packets.PacketMotionData22))
	}
}

func (dbc *motionDebouncer) WriteTo(ch writer) {
	dbc.mx.Lock()
	defer dbc.mx.Unlock()
	defer func() {
		dbc.timer = time.AfterFunc(motionDebouncerInterval, func() {
			dbc.WriteTo(ch)
		})
	}()
	if dbc.h == nil || (len(dbc.packages_21) == 0 && len(dbc.packages_22) == 0) {
		return
	}

	var pack interface{}
	if dbc.h.PacketFormat == constants.PacketFormat_2021 {
		if len(dbc.packages_21) == 0 {
			return
		}
		if len(dbc.packages_21) == 1 {
			pack = dbc.packages_21[0]
		} else {
			pack = averageAndLastPlayerCarMotion21(dbc.packages_21).AverageAndLastPlayerCarMotion()
		}
	}
	if dbc.h.PacketFormat == constants.PacketFormat_2022 {
		if len(dbc.packages_22) == 0 {
			return
		}
		if len(dbc.packages_22) == 1 {
			pack = dbc.packages_22[0]
		} else {
			pack = averageAndLastPlayerCarMotion22(dbc.packages_22).AverageAndLastPlayerCarMotion()
		}
	}

	m := &process.M{
		Header: *dbc.h,
		Pack:   pack,
	}

	var err error
	m.Buffer, err = packets.Write_LE(pack)
	if err != nil {
		log.Printf("error encoding averaged packet: %v", err)
		return
	}

	ch.write(m)
	dbc.h = nil
	dbc.packages_21 = make([]*packets.PacketMotionData21, 0)
	dbc.packages_22 = make([]*packets.PacketMotionData22, 0)
}

type averageAndLastPlayerCarMotion21 []*packets.PacketMotionData21

func (a averageAndLastPlayerCarMotion21) AverageAndLastPlayerCarMotion() *packets.PacketMotionData21 {
	pmd := *a[len(a)-1]
	pmd.CarMotionData = [22]packets.CarMotionData21{}
	for i := 0; i < len(pmd.CarMotionData); i++ {
		for _, p := range a {
			pmd.CarMotionData[i].WorldPositionX += p.CarMotionData[i].WorldPositionX
			pmd.CarMotionData[i].WorldPositionY += p.CarMotionData[i].WorldPositionY
			pmd.CarMotionData[i].WorldPositionZ += p.CarMotionData[i].WorldPositionZ
			pmd.CarMotionData[i].WorldVelocityX += p.CarMotionData[i].WorldVelocityX
			pmd.CarMotionData[i].WorldVelocityY += p.CarMotionData[i].WorldVelocityY
			pmd.CarMotionData[i].WorldVelocityZ += p.CarMotionData[i].WorldVelocityZ

			pmd.CarMotionData[i].WorldForwardDirX += uint16(p.CarMotionData[i].WorldForwardDirX)
			pmd.CarMotionData[i].WorldForwardDirY += uint16(p.CarMotionData[i].WorldForwardDirY)
			pmd.CarMotionData[i].WorldForwardDirZ += uint16(p.CarMotionData[i].WorldForwardDirZ)
			pmd.CarMotionData[i].WorldRightDirX += uint16(p.CarMotionData[i].WorldRightDirX)
			pmd.CarMotionData[i].WorldRightDirY += uint16(p.CarMotionData[i].WorldRightDirY)
			pmd.CarMotionData[i].WorldRightDirZ += uint16(p.CarMotionData[i].WorldRightDirZ)

			pmd.CarMotionData[i].GForceLateral += p.CarMotionData[i].GForceLateral
			pmd.CarMotionData[i].GForceLongitudinal += p.CarMotionData[i].GForceLongitudinal
			pmd.CarMotionData[i].GForceVertical += p.CarMotionData[i].GForceVertical
			pmd.CarMotionData[i].Yaw += p.CarMotionData[i].Yaw
			pmd.CarMotionData[i].Pitch += p.CarMotionData[i].Pitch
			pmd.CarMotionData[i].Roll += p.CarMotionData[i].Roll
		}

		pmd.CarMotionData[i].WorldPositionX /= float32(len(a))
		pmd.CarMotionData[i].WorldPositionY /= float32(len(a))
		pmd.CarMotionData[i].WorldPositionZ /= float32(len(a))
		pmd.CarMotionData[i].WorldVelocityX /= float32(len(a))
		pmd.CarMotionData[i].WorldVelocityY /= float32(len(a))
		pmd.CarMotionData[i].WorldVelocityZ /= float32(len(a))

		pmd.CarMotionData[i].WorldForwardDirX /= uint16(len(a))
		pmd.CarMotionData[i].WorldForwardDirY /= uint16(len(a))
		pmd.CarMotionData[i].WorldForwardDirZ /= uint16(len(a))
		pmd.CarMotionData[i].WorldRightDirX /= uint16(len(a))
		pmd.CarMotionData[i].WorldRightDirY /= uint16(len(a))
		pmd.CarMotionData[i].WorldRightDirZ /= uint16(len(a))

		pmd.CarMotionData[i].GForceLateral /= float32(len(a))
		pmd.CarMotionData[i].GForceLongitudinal /= float32(len(a))
		pmd.CarMotionData[i].GForceVertical /= float32(len(a))
		pmd.CarMotionData[i].Yaw /= float32(len(a))
		pmd.CarMotionData[i].Pitch /= float32(len(a))
		pmd.CarMotionData[i].Roll /= float32(len(a))
	}

	return &pmd
}

type averageAndLastPlayerCarMotion22 []*packets.PacketMotionData22

func (a averageAndLastPlayerCarMotion22) AverageAndLastPlayerCarMotion() *packets.PacketMotionData22 {
	pmd := *a[len(a)-1]
	pmd.CarMotionData = [22]packets.CarMotionData22{}
	for i := 0; i < len(pmd.CarMotionData); i++ {
		for _, p := range a {
			pmd.CarMotionData[i].WorldPositionX += p.CarMotionData[i].WorldPositionX
			pmd.CarMotionData[i].WorldPositionY += p.CarMotionData[i].WorldPositionY
			pmd.CarMotionData[i].WorldPositionZ += p.CarMotionData[i].WorldPositionZ
			pmd.CarMotionData[i].WorldVelocityX += p.CarMotionData[i].WorldVelocityX
			pmd.CarMotionData[i].WorldVelocityY += p.CarMotionData[i].WorldVelocityY
			pmd.CarMotionData[i].WorldVelocityZ += p.CarMotionData[i].WorldVelocityZ

			pmd.CarMotionData[i].WorldForwardDirX += int16(p.CarMotionData[i].WorldForwardDirX)
			pmd.CarMotionData[i].WorldForwardDirY += int16(p.CarMotionData[i].WorldForwardDirY)
			pmd.CarMotionData[i].WorldForwardDirZ += int16(p.CarMotionData[i].WorldForwardDirZ)
			pmd.CarMotionData[i].WorldRightDirX += int16(p.CarMotionData[i].WorldRightDirX)
			pmd.CarMotionData[i].WorldRightDirY += int16(p.CarMotionData[i].WorldRightDirY)
			pmd.CarMotionData[i].WorldRightDirZ += int16(p.CarMotionData[i].WorldRightDirZ)

			pmd.CarMotionData[i].GForceLateral += p.CarMotionData[i].GForceLateral
			pmd.CarMotionData[i].GForceLongitudinal += p.CarMotionData[i].GForceLongitudinal
			pmd.CarMotionData[i].GForceVertical += p.CarMotionData[i].GForceVertical
			pmd.CarMotionData[i].Yaw += p.CarMotionData[i].Yaw
			pmd.CarMotionData[i].Pitch += p.CarMotionData[i].Pitch
			pmd.CarMotionData[i].Roll += p.CarMotionData[i].Roll
		}

		pmd.CarMotionData[i].WorldPositionX /= float32(len(a))
		pmd.CarMotionData[i].WorldPositionY /= float32(len(a))
		pmd.CarMotionData[i].WorldPositionZ /= float32(len(a))
		pmd.CarMotionData[i].WorldVelocityX /= float32(len(a))
		pmd.CarMotionData[i].WorldVelocityY /= float32(len(a))
		pmd.CarMotionData[i].WorldVelocityZ /= float32(len(a))

		pmd.CarMotionData[i].WorldForwardDirX /= int16(len(a))
		pmd.CarMotionData[i].WorldForwardDirY /= int16(len(a))
		pmd.CarMotionData[i].WorldForwardDirZ /= int16(len(a))
		pmd.CarMotionData[i].WorldRightDirX /= int16(len(a))
		pmd.CarMotionData[i].WorldRightDirY /= int16(len(a))
		pmd.CarMotionData[i].WorldRightDirZ /= int16(len(a))

		pmd.CarMotionData[i].GForceLateral /= float32(len(a))
		pmd.CarMotionData[i].GForceLongitudinal /= float32(len(a))
		pmd.CarMotionData[i].GForceVertical /= float32(len(a))
		pmd.CarMotionData[i].Yaw /= float32(len(a))
		pmd.CarMotionData[i].Pitch /= float32(len(a))
		pmd.CarMotionData[i].Roll /= float32(len(a))
	}

	return &pmd
}
