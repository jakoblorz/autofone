package packets

type HeaderProvider interface {
	GetPacketID() uint8
	GetPacketFormat() uint16
}

// 24 byte

type PacketHeader21 struct {
	PacketFormat            uint16  // 2021
	GameMajorVersion        uint8   // Game major version - "X.00"
	GameMinorVersion        uint8   // Game minor version - "1.XX"
	PacketVersion           uint8   // Version of this packet type, all start from 1
	PacketID                uint8   // Identifier for the packet type, see below
	SessionUID              uint64  // Unique identifier for the session
	SessionTime             float32 // Session timestamp
	FrameIdentifier         uint32  // Identifier for the frame the data was retrieved on
	PlayerCarIndex          uint8   // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   // Index of secondary player's car in the array (split screen)
}

func (p *PacketHeader21) GetPacketFormat() uint16 {
	return p.PacketFormat
}

func (p *PacketHeader21) GetPacketID() uint8 {
	return p.PacketID
}

// 24 byte

type PacketHeader22 struct {
	PacketFormat            uint16  // 2022
	GameMajorVersion        uint8   // Game major version - "X.00"
	GameMinorVersion        uint8   // Game minor version - "1.XX"
	PacketVersion           uint8   // Version of this packet type, all start from 1
	PacketID                uint8   // Identifier for the packet type, see below
	SessionUID              uint64  // Unique identifier for the session
	SessionTime             float32 // Session timestamp
	FrameIdentifier         uint32  // Identifier for the frame the data was retrieved on
	PlayerCarIndex          uint8   // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   // Index of secondary player's car in the array (split screen)
}

func (p *PacketHeader22) GetPacketFormat() uint16 {
	return p.PacketFormat
}

func (p *PacketHeader22) GetPacketID() uint8 {
	return p.PacketID
}

type PacketHeader23 struct {
	PacketFormat            uint16  // 2023
	GameYear                uint8   // Game year - last two digits e.g. 23
	GameMajorVersion        uint8   // Game major version - "X.00"
	GameMinorVersion        uint8   // Game minor version - "1.XX"
	PacketVersion           uint8   // Version of this packet type, all start from 1
	PacketID                uint8   // Identifier for the packet type, see below
	SessionUID              uint64  // Unique identifier for the session
	SessionTime             float32 // Session timestamp
	FrameIdentifier         uint32  // Identifier for the frame the data was retrieved on
	OverallFrameIdentifier  uint32  // Overall identifier for the frame the data was retrieved
	PlayerCarIndex          uint8   // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   // Index of secondary player's car in the array (split screen)
}

func (p *PacketHeader23) GetPacketFormat() uint16 {
	return p.PacketFormat
}

func (p *PacketHeader23) GetPacketID() uint8 {
	return p.PacketID
}
