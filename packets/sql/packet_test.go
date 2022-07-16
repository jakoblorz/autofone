package sql_test

import (
	"context"
	"testing"

	"github.com/jakoblorz/autofone/packets/mocks"
	"github.com/jakoblorz/autofone/packets/sql"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func TestPacket_Write(t *testing.T) {
	ctx := context.Background()
	db, err := sqlx.ConnectContext(ctx, "sqlite3", ":memory:")
	if !assert.NoError(t, err) {
		t.Fail()
		return
	}

	sql.Init(db)

	tests := []struct {
		name string
		p    *sql.Packet
	}{
		{
			name: "should write to db and be able to retrieve it",
			p: sql.Packet{
				Hostname: "test-hostname",
				Data:     mocks.PacketCarSetupData21Bytes,
			}.WithPacketHeader(
				&mocks.PacketCarSetupData21.Header,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := db.Beginx()
			if !assert.NoError(t, err) {
				t.Fail()
				return
			}
			err = tt.p.Write(ctx, tx)
			if !assert.NoError(t, err) {
				t.Fail()
				return
			}
			err = tx.Commit()
			if !assert.NoError(t, err) {
				t.Fail()
				return
			}

			c := sql.Packet{}
			err = db.Get(&c, "SELECT * FROM packets WHERE hostname=$1", tt.p.Hostname)
			if !assert.NoError(t, err) {
				t.Fail()
				return
			}

			if !assert.Equal(t, tt.p, &c) {
				t.Fail()
				return
			}
		})
	}
}
