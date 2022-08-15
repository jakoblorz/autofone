package writer

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/pkg/privateapi"
	"github.com/jakoblorz/autofone/pkg/streamdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bolt struct {
	*process.P
	privateapi.Client

	Motion         *motionDebouncer
	Lap            *packetDebouncer
	CarTelemetry   *packetDebouncer
	CarStatus      *packetDebouncer
	SessionHistory *sessionHistoryDebouncer

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
	return nil
}

func (ch *Bolt) write(m *process.M) {
	h := ch.DB.Get()
	defer ch.DB.Put(h)

	id := fmt.Sprintf("%s-%011d", primitive.NewObjectID().Hex(), time.Now().Unix())

	h.Batch(func(tx *bolt.Tx) (err error) {
		bkt, err := tx.CreateBucketIfNotExists([]byte{m.Header.PacketID})
		if err != nil {
			return
		}

		err = bkt.Put([]byte(id), m.Buffer)
		return
	})
}

func (ch *Bolt) Write(m *process.M) {
	if ch.Motion != nil && m.Header.PacketID == constants.PacketMotion {
		ch.Motion.Write(m)
		return
	}
	if ch.Lap != nil && m.Header.PacketID == constants.PacketLap {
		ch.Lap.Write(m)
		return
	}
	if ch.CarTelemetry != nil && m.Header.PacketID == constants.PacketCarTelemetry {
		ch.CarTelemetry.Write(m)
		return
	}
	if ch.CarStatus != nil && m.Header.PacketID == constants.PacketCarStatus {
		ch.CarStatus.Write(m)
		return
	}
	if ch.SessionHistory != nil && m.Header.PacketID == constants.PacketSessionHistory {
		ch.SessionHistory.Write(m)
		return
	}
	ch.write(m)
}
