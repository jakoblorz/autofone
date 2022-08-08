package writer

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/pkg/privateapi"
	"github.com/jakoblorz/autofone/pkg/streamdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bolt struct {
	*process.P
	privateapi.Client

	DB streamdb.I
}

func (ch *Bolt) Write(m *process.M) {
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
