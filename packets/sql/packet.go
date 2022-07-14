package sql

import (
	"context"

	"github.com/jakoblorz/autofone/packets"
	"github.com/jmoiron/sqlx"
)

type Packet struct {
	Hostname                string  `db:"hostname"`
	PacketFormat            uint16  `db:"packet_format"`              // 2021 or 2022
	GameMajorVersion        uint8   `db:"game_major_version"`         // Game major version - "X.00"
	GameMinorVersion        uint8   `db:"game_minor_version"`         // Game minor version - "1.XX"
	PacketVersion           uint8   `db:"packet_version"`             // Version of this packet type, all start from 1
	PacketID                uint8   `db:"packet_id"`                  // Identifier for the packet type, see below
	SessionUID              uint64  `db:"session_uid"`                // Unique identifier for the session
	SessionTime             float32 `db:"session_time"`               // Session timestamp
	FrameIdentifier         uint32  `db:"frame_identifier"`           // Identifier for the frame the data was retrieved on
	PlayerCarIndex          uint8   `db:"player_car_index"`           // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   `db:"secondary_player_car_index"` // Index of secondary player's car in the array (split screen)

	Data []byte `db:"data"`
}

func (p Packet) WithPacketHeader(ph *packets.PacketHeader) *Packet {
	p.PacketFormat = ph.PacketFormat
	p.GameMajorVersion = ph.GameMajorVersion
	p.GameMinorVersion = ph.GameMinorVersion
	p.PacketVersion = ph.PacketVersion
	p.PacketID = ph.PacketID
	p.SessionUID = ph.SessionUID
	p.SessionTime = ph.SessionTime
	p.FrameIdentifier = ph.FrameIdentifier
	p.PlayerCarIndex = ph.PlayerCarIndex
	p.SecondaryPlayerCarIndex = ph.SecondaryPlayerCarIndex
	return &p
}

func Init(db *sqlx.DB) (err error) {
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS packets (
		hostname TEXT,
		packet_format TEXT,
		game_major_version INTEGER,
		game_minor_version INTEGER,
		packet_version INTEGER,
		packet_id INTEGER,
		session_uid INTEGER,
		session_time REAL,
		frame_identifier INTEGER,
		player_car_index INTEGER,
		secondary_player_car_index INTEGER,
		data BLOB
	)`)
	return
}

func (p *Packet) Write(ctx context.Context, tx *sqlx.Tx) (err error) {
	_, err = tx.ExecContext(ctx, `INSERT INTO packets (
		hostname,
		packet_format,
		game_major_version,
		game_minor_version,
		packet_version,
		packet_id,
		session_uid,
		session_time,
		frame_identifier,
		player_car_index,
		secondary_player_car_index,
		data
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
		p.PacketFormat,
		p.GameMajorVersion,
		p.GameMinorVersion,
		p.PacketVersion,
		p.PacketID,
		p.SessionUID,
		p.SessionTime,
		p.FrameIdentifier,
		p.PlayerCarIndex,
		p.SecondaryPlayerCarIndex,
		p.Data,
	)
	return
}
