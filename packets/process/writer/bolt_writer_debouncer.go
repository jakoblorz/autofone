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
	motionDebouncerInterval         = 250 * time.Millisecond
	motionExDebouncerInterval       = 100 * time.Millisecond
	packetDebouncerInterval         = 1 * time.Second
	sessionHistoryDebouncerInterval = 10 * time.Second
	tyreSetsDebouncerInterval       = 10 * time.Second
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
		packages_23: make([]*packets.PacketMotionData23, 0),
	}
	mdc.timer = time.AfterFunc(motionDebouncerInterval, func() { mdc.WriteTo(ch) })
	return mdc
}

func NewMotionExDebouncer(ch writer, interval time.Duration) *motionExDebouncer {
	if interval == 0 {
		interval = motionExDebouncerInterval
	}
	mdc := &motionExDebouncer{
		mx: &sync.Mutex{},

		packages_23: make([]*packets.PacketMotionExData23, 0),
	}
	mdc.timer = time.AfterFunc(motionExDebouncerInterval, func() { mdc.WriteTo(ch) })
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

type motionExDebouncer struct {
	mx    sync.Locker
	timer *time.Timer

	h packets.PacketHeader

	packages_23 []*packets.PacketMotionExData23
}

func (dbc *motionExDebouncer) Write(m *process.M) {
	dbc.mx.Lock()
	defer dbc.mx.Unlock()

	dbc.h = m.Header

	if m.Header.GetPacketFormat() == constants.PacketFormat_2023 {
		dbc.packages_23 = append(dbc.packages_23, m.Pack.(*packets.PacketMotionExData23))
	}
}

func (dbc *motionExDebouncer) WriteTo(ch writer) {
	dbc.mx.Lock()
	defer dbc.mx.Unlock()
	defer func() {
		dbc.timer = time.AfterFunc(motionDebouncerInterval, func() {
			dbc.WriteTo(ch)
		})
	}()
	if dbc.h == nil || len(dbc.packages_23) == 0 {
		return
	}

	var pack interface{}
	if dbc.h.GetPacketFormat() == constants.PacketFormat_2023 {
		if len(dbc.packages_23) == 0 {
			return
		}
		if len(dbc.packages_23) == 1 {
			pack = dbc.packages_23[0]
		} else {
			pack = averageAndLastPlayerCarMotionEx23(dbc.packages_23).AverageAndLastPlayerCarMotionEx()
		}
	}

	m := &process.M{
		Header: dbc.h,
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
	dbc.packages_23 = make([]*packets.PacketMotionExData23, 0)
}

type motionDebouncer struct {
	mx    sync.Locker
	timer *time.Timer

	h packets.PacketHeader

	packages_21 []*packets.PacketMotionData21
	packages_22 []*packets.PacketMotionData22
	packages_23 []*packets.PacketMotionData23
}

func (dbc *motionDebouncer) Write(m *process.M) {
	dbc.mx.Lock()
	defer dbc.mx.Unlock()

	dbc.h = m.Header

	if m.Header.GetPacketFormat() == constants.PacketFormat_2021 {
		dbc.packages_21 = append(dbc.packages_21, m.Pack.(*packets.PacketMotionData21))
	}
	if m.Header.GetPacketFormat() == constants.PacketFormat_2022 {
		dbc.packages_22 = append(dbc.packages_22, m.Pack.(*packets.PacketMotionData22))
	}
	if m.Header.GetPacketFormat() == constants.PacketFormat_2023 {
		dbc.packages_23 = append(dbc.packages_23, m.Pack.(*packets.PacketMotionData23))
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
	if dbc.h == nil || (len(dbc.packages_21) == 0 && len(dbc.packages_22) == 0 && len(dbc.packages_23) == 0) {
		return
	}

	var pack interface{}
	if dbc.h.GetPacketFormat() == constants.PacketFormat_2021 {
		if len(dbc.packages_21) == 0 {
			return
		}
		if len(dbc.packages_21) == 1 {
			pack = dbc.packages_21[0]
		} else {
			pack = averageAndLastPlayerCarMotion21(dbc.packages_21).AverageAndLastPlayerCarMotion()
		}
	}
	if dbc.h.GetPacketFormat() == constants.PacketFormat_2022 {
		if len(dbc.packages_22) == 0 {
			return
		}
		if len(dbc.packages_22) == 1 {
			pack = dbc.packages_22[0]
		} else {
			pack = averageAndLastPlayerCarMotion22(dbc.packages_22).AverageAndLastPlayerCarMotion()
		}
	}
	if dbc.h.GetPacketFormat() == constants.PacketFormat_2023 {
		if len(dbc.packages_23) == 0 {
			return
		}
		if len(dbc.packages_23) == 1 {
			pack = dbc.packages_23[0]
		} else {
			pack = averageAndLastPlayerCarMotion23(dbc.packages_23).AverageAndLastPlayerCarMotion()
		}
	}

	m := &process.M{
		Header: dbc.h,
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
	dbc.packages_23 = make([]*packets.PacketMotionData23, 0)
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

type averageAndLastPlayerCarMotion23 []*packets.PacketMotionData23

func (a averageAndLastPlayerCarMotion23) AverageAndLastPlayerCarMotion() *packets.PacketMotionData23 {
	pmd := *a[len(a)-1]
	pmd.CarMotionData = [22]packets.CarMotionData23{}
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

type averageAndLastPlayerCarMotionEx23 []*packets.PacketMotionExData23

func (a averageAndLastPlayerCarMotionEx23) AverageAndLastPlayerCarMotionEx() *packets.PacketMotionExData23 {
	pmd := *a[len(a)-1]
	for _, p := range a {
		pmd.SuspensionPosition[0] += p.SuspensionPosition[0]
		pmd.SuspensionPosition[1] += p.SuspensionPosition[1]
		pmd.SuspensionPosition[2] += p.SuspensionPosition[2]
		pmd.SuspensionPosition[3] += p.SuspensionPosition[3]

		pmd.SuspensionVelocity[0] += p.SuspensionVelocity[0]
		pmd.SuspensionVelocity[1] += p.SuspensionVelocity[1]
		pmd.SuspensionVelocity[2] += p.SuspensionVelocity[2]
		pmd.SuspensionVelocity[3] += p.SuspensionVelocity[3]

		pmd.SuspensionAcceleration[0] += p.SuspensionAcceleration[0]
		pmd.SuspensionAcceleration[1] += p.SuspensionAcceleration[1]
		pmd.SuspensionAcceleration[2] += p.SuspensionAcceleration[2]
		pmd.SuspensionAcceleration[3] += p.SuspensionAcceleration[3]

		pmd.WheelSpeed[0] += p.WheelSpeed[0]
		pmd.WheelSpeed[1] += p.WheelSpeed[1]
		pmd.WheelSpeed[2] += p.WheelSpeed[2]
		pmd.WheelSpeed[3] += p.WheelSpeed[3]

		pmd.WheelSlipRatio[0] += p.WheelSlipRatio[0]
		pmd.WheelSlipRatio[1] += p.WheelSlipRatio[1]
		pmd.WheelSlipRatio[2] += p.WheelSlipRatio[2]
		pmd.WheelSlipRatio[3] += p.WheelSlipRatio[3]

		pmd.WheelLatForce[0] += p.WheelLatForce[0]
		pmd.WheelLatForce[1] += p.WheelLatForce[1]
		pmd.WheelLatForce[2] += p.WheelLatForce[2]
		pmd.WheelLatForce[3] += p.WheelLatForce[3]

		pmd.WheelLongForce[0] += p.WheelLongForce[0]
		pmd.WheelLongForce[1] += p.WheelLongForce[1]
		pmd.WheelLongForce[2] += p.WheelLongForce[2]
		pmd.WheelLongForce[3] += p.WheelLongForce[3]

		pmd.HeightOfCOGAboveGround += p.HeightOfCOGAboveGround
		pmd.LocalVelocityX += p.LocalVelocityX
		pmd.LocalVelocityY += p.LocalVelocityY
		pmd.LocalVelocityZ += p.LocalVelocityZ
		pmd.AngularVelocityX += p.AngularVelocityX
		pmd.AngularVelocityY += p.AngularVelocityY
		pmd.AngularVelocityZ += p.AngularVelocityZ
		pmd.AngularAccelerationX += p.AngularAccelerationX
		pmd.AngularAccelerationY += p.AngularAccelerationY
		pmd.AngularAccelerationZ += p.AngularAccelerationZ
		pmd.FrontWheelAngle += p.FrontWheelAngle

		pmd.WheelVertForce[0] += p.WheelVertForce[0]
		pmd.WheelVertForce[1] += p.WheelVertForce[1]
		pmd.WheelVertForce[2] += p.WheelVertForce[2]
		pmd.WheelVertForce[3] += p.WheelVertForce[3]
	}

	pmd.SuspensionPosition[0] /= float32(len(a))
	pmd.SuspensionPosition[1] /= float32(len(a))
	pmd.SuspensionPosition[2] /= float32(len(a))
	pmd.SuspensionPosition[3] /= float32(len(a))

	pmd.SuspensionVelocity[0] /= float32(len(a))
	pmd.SuspensionVelocity[1] /= float32(len(a))
	pmd.SuspensionVelocity[2] /= float32(len(a))
	pmd.SuspensionVelocity[3] /= float32(len(a))

	pmd.SuspensionAcceleration[0] /= float32(len(a))
	pmd.SuspensionAcceleration[1] /= float32(len(a))
	pmd.SuspensionAcceleration[2] /= float32(len(a))
	pmd.SuspensionAcceleration[3] /= float32(len(a))

	pmd.WheelSpeed[0] /= float32(len(a))
	pmd.WheelSpeed[1] /= float32(len(a))
	pmd.WheelSpeed[2] /= float32(len(a))
	pmd.WheelSpeed[3] /= float32(len(a))

	pmd.WheelSlipRatio[0] /= float32(len(a))
	pmd.WheelSlipRatio[1] /= float32(len(a))
	pmd.WheelSlipRatio[2] /= float32(len(a))
	pmd.WheelSlipRatio[3] /= float32(len(a))

	pmd.WheelLatForce[0] /= float32(len(a))
	pmd.WheelLatForce[1] /= float32(len(a))
	pmd.WheelLatForce[2] /= float32(len(a))
	pmd.WheelLatForce[3] /= float32(len(a))

	pmd.WheelLongForce[0] /= float32(len(a))
	pmd.WheelLongForce[1] /= float32(len(a))
	pmd.WheelLongForce[2] /= float32(len(a))
	pmd.WheelLongForce[3] /= float32(len(a))

	pmd.HeightOfCOGAboveGround /= float32(len(a))
	pmd.LocalVelocityX /= float32(len(a))
	pmd.LocalVelocityY /= float32(len(a))
	pmd.LocalVelocityZ /= float32(len(a))
	pmd.AngularVelocityX /= float32(len(a))
	pmd.AngularVelocityY /= float32(len(a))
	pmd.AngularVelocityZ /= float32(len(a))
	pmd.AngularAccelerationX /= float32(len(a))
	pmd.AngularAccelerationY /= float32(len(a))
	pmd.AngularAccelerationZ /= float32(len(a))
	pmd.FrontWheelAngle /= float32(len(a))

	pmd.WheelVertForce[0] /= float32(len(a))
	pmd.WheelVertForce[1] /= float32(len(a))
	pmd.WheelVertForce[2] /= float32(len(a))
	pmd.WheelVertForce[3] /= float32(len(a))

	return &pmd
}

func NewSessionHistoryDebouncer(ch writer, interval time.Duration) *sessionHistoryDebouncer {
	if interval == 0 {
		interval = sessionHistoryDebouncerInterval
	}
	pdc := &sessionHistoryDebouncer{
		ch:         ch,
		interval:   packetDebouncerInterval,
		debouncers: make(map[uint8]*packetDebouncer),
	}
	return pdc
}

type sessionHistoryDebouncer struct {
	ch         writer
	interval   time.Duration
	debouncers map[uint8]*packetDebouncer
}

func (dbc *sessionHistoryDebouncer) Stop() {
	for _, d := range dbc.debouncers {
		if d.timer != nil {
			d.timer.Stop()
		}
	}
}

func (dbc *sessionHistoryDebouncer) Write(m *process.M) {
	if sh21, ok := m.Pack.(*packets.PacketSessionHistoryData21); ok {
		if dbc.debouncers[sh21.CarIdx] == nil {
			dbc.debouncers[sh21.CarIdx] = NewPacketDebouncer(dbc.ch, dbc.interval)
		}
		dbc.debouncers[sh21.CarIdx].Write(m)
		return
	}
	if sh22, ok := m.Pack.(*packets.PacketSessionHistoryData22); ok {
		if dbc.debouncers[sh22.CarIdx] == nil {
			dbc.debouncers[sh22.CarIdx] = NewPacketDebouncer(dbc.ch, dbc.interval)
		}
		dbc.debouncers[sh22.CarIdx].Write(m)
		return
	}
	if sh23, ok := m.Pack.(*packets.PacketSessionHistoryData23); ok {
		if dbc.debouncers[sh23.CarIdx] == nil {
			dbc.debouncers[sh23.CarIdx] = NewPacketDebouncer(dbc.ch, dbc.interval)
		}
		dbc.debouncers[sh23.CarIdx].Write(m)
		return
	}
}

func NewTyreSetsDebouncer(ch writer, interval time.Duration) *tyreSetsDebouncer {
	if interval == 0 {
		interval = tyreSetsDebouncerInterval
	}
	pdc := &tyreSetsDebouncer{
		ch:         ch,
		interval:   packetDebouncerInterval,
		debouncers: make(map[uint8]*packetDebouncer),
	}
	return pdc
}

type tyreSetsDebouncer struct {
	ch         writer
	interval   time.Duration
	debouncers map[uint8]*packetDebouncer
}

func (dbc *tyreSetsDebouncer) Stop() {
	for _, d := range dbc.debouncers {
		if d.timer != nil {
			d.timer.Stop()
		}
	}
}

func (dbc *tyreSetsDebouncer) Write(m *process.M) {
	if sh23, ok := m.Pack.(*packets.PacketTyreSetsData23); ok {
		if dbc.debouncers[sh23.CarIdx] == nil {
			dbc.debouncers[sh23.CarIdx] = NewPacketDebouncer(dbc.ch, dbc.interval)
		}
		dbc.debouncers[sh23.CarIdx].Write(m)
		return
	}
}
