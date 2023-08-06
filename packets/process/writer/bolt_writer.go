package writer

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/pkg/streamdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bolt struct {
	*process.P

	Motion         *motionDebouncer
	Lap            *packetDebouncer
	CarTelemetry   *packetDebouncer
	CarStatus      *packetDebouncer
	SessionHistory *sessionHistoryDebouncer
	TyreSets       *tyreSetsDebouncer
	MotionEx       *motionExDebouncer

	DB streamdb.I
}

func (ch *Bolt) Close() error {
	if ch.Motion != nil {
		ch.Motion.timer.Stop()
	}
	if ch.Lap != nil {
		ch.Lap.timer.Stop()
	}
	if ch.CarTelemetry != nil {
		ch.CarTelemetry.timer.Stop()
	}
	if ch.CarStatus != nil {
		ch.CarStatus.timer.Stop()
	}
	if ch.SessionHistory != nil {
		ch.SessionHistory.Stop()
	}
	if ch.TyreSets != nil {
		ch.TyreSets.Stop()
	}
	if ch.MotionEx != nil {
		ch.MotionEx.timer.Stop()
	}
	return nil
}

func (ch *Bolt) write(m *process.M) {
	h := ch.DB.Get()
	defer ch.DB.Put(h)

	id := fmt.Sprintf("%s-%011d", primitive.NewObjectID().Hex(), time.Now().Unix())

	h.Batch(func(tx *bolt.Tx) (err error) {
		bkt, err := tx.CreateBucketIfNotExists([]byte{m.Header.GetPacketID()})
		if err != nil {
			return
		}

		err = bkt.Put([]byte(id), m.Buffer)
		return
	})
}

func (ch *Bolt) Write(m *process.M) {
	if ch.Motion != nil && m.Header.GetPacketID() == constants.PacketMotion {
		ch.Motion.Write(m)
		return
	}
	if ch.Lap != nil && m.Header.GetPacketID() == constants.PacketLap {
		ch.Lap.Write(m)
		return
	}
	if ch.CarTelemetry != nil && m.Header.GetPacketID() == constants.PacketCarTelemetry {
		ch.CarTelemetry.Write(m)
		return
	}
	if ch.CarStatus != nil && m.Header.GetPacketID() == constants.PacketCarStatus {
		ch.CarStatus.Write(m)
		return
	}
	if ch.SessionHistory != nil && m.Header.GetPacketID() == constants.PacketSessionHistory {
		ch.SessionHistory.Write(m)
		return
	}
	if ch.TyreSets != nil && m.Header.GetPacketID() == constants.PacketTyreSets {
		ch.TyreSets.Write(m)
		return
	}
	if ch.MotionEx != nil && m.Header.GetPacketID() == constants.PacketMotionEx {
		ch.MotionEx.Write(m)
		return
	}
	ch.write(m)
}
