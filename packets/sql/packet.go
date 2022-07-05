package sql

import (
	"context"

	"github.com/jakoblorz/autofone/packets"
	"github.com/jmoiron/sqlx"
)

type Packet struct {
	Hostname string
	packets.PacketHeader
	Data []byte
}

func Init(db *sqlx.DB) (err error) {
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS packets (
		Hostname TEXT,
		PacketFormat TEXT,
		GameMajorVersion INTEGER,
		GameMinorVersion INTEGER,
		PacketVersion INTEGER,
		PacketID INTEGER,
		SessionUID INTEGER,
		SessionTime REAL,
		FrameIdentifier INTEGER,
		PlayerCarIndex INTEGER,
		SecondaryPlayerCarIndex INTEGER,
		Data BLOB
	)`)
	return
}

func (p *Packet) Write(ctx context.Context, tx *sqlx.Tx) (err error) {
	_, err = tx.ExecContext(ctx, `INSERT INTO packets (
		Hostname,
		PacketFormat,
		GameMajorVersion,
		GameMinorVersion,
		PacketVersion,
		PacketID,
		SessionUID,
		SessionTime,
		FrameIndentifier,
		PlayerCarIndex,
		SecondaryPlayerCarIndex,
		Data
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12
	)`,
		p.Hostname,
		p.PacketHeader.PacketFormat,
		p.PacketHeader.GameMajorVersion,
		p.PacketHeader.GameMinorVersion,
		p.PacketHeader.PacketVersion,
		p.PacketHeader.PacketID,
		p.PacketHeader.SessionUID,
		p.PacketHeader.SessionTime,
		p.PacketHeader.FrameIdentifier,
		p.PacketHeader.PlayerCarIndex,
		p.PacketHeader.SecondaryPlayerCarIndex,
		p.Data,
	)
	return
}
