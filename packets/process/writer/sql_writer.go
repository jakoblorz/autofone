package writer

import (
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/packets/sql"
	"github.com/jakoblorz/autofone/pkg/log"
	"github.com/jakoblorz/autofone/pkg/streamdb"
)

type SQL struct {
	*process.P

	DB *streamdb.I
}

func (ch *SQL) Write(m *process.M) {
	tx, err := ch.DB.Beginx()
	if err != nil {
		log.Printf("tx begin() error: %+v", err)
		return
	}
	defer tx.Rollback()

	err = sql.Packet{
		Hostname: ch.Hostname,
		Data:     m.Buffer,
	}.WithPacketHeader(&m.Header).Write(ch.Context, tx)
	if err != nil {
		log.Printf("tx write() error: %+v", err)
		return
	}

	err = ch.DB.SoftSync(ch.Context)
	if err != nil {
		log.Printf("tx sync(1) error: %+v", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("tx commit() error: %+v", err)
		return
	}

	err = ch.DB.SoftSync(ch.Context)
	if err != nil {
		log.Printf("tx sync(2) error: %+v", err)
		return
	}
}
