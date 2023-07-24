package packets

type ParticipantData interface {
	NameToString() string
}

var (
	_ ParticipantData = (*ParticipantData21)(nil)
	_ ParticipantData = (*ParticipantData22)(nil)
)

// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name.
// If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.

// Frequency: Every 5 seconds
// Size: 1213 bytes
// Version: 1

type ParticipantData21 struct {
	AIControlled  uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID      uint8    // Driver id - see appendix
	NetworkID     uint8    // Network id – unique identifier for network players
	TeamID        uint8    // Team id - see appendix
	MyTeam        uint8    // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber    uint8    // Race number of the car
	Nationality   uint8    // Nationality of the driver
	Name          [48]byte // Name of participant in UTF-8 format – null terminated, Will be truncated with … (U+2026) if too long
	YourTelemetry uint8    // The player's UDP setting, 0 = restricted, 1 = public
}

type PacketParticipantsData21 struct {
	Header        PacketHeader21
	NumActiveCars uint8 // Number of active cars in the data – should match number of cars on HUD
	Participants  [22]ParticipantData21
}

func (p *ParticipantData21) NameToString() string {
	return string(p.Name[:])
}

// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name.
// If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.

// Frequency: Every 5 seconds
// Size: 1257 bytes
// Version: 1

type ParticipantData22 struct {
	AIControlled  uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID      uint8    // Driver id - see appendix
	NetworkID     uint8    // Network id – unique identifier for network players
	TeamID        uint8    // Team id - see appendix
	MyTeam        uint8    // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber    uint8    // Race number of the car
	Nationality   uint8    // Nationality of the driver
	Name          [48]byte // Name of participant in UTF-8 format – null terminated, Will be truncated with … (U+2026) if too long
	YourTelemetry uint8    // The player's UDP setting, 0 = restricted, 1 = public
}

type PacketParticipantsData22 struct {
	Header        PacketHeader22
	NumActiveCars uint8 // Number of active cars in the data – should match number of cars on HUD
	Participants  [22]ParticipantData22
}

func (p *ParticipantData22) NameToString() string {
	return string(p.Name[:])
}

// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name. If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.

// N.B. on Xbox One, the names will always be the driver name, on PS4 the name will be the LAN name if playing a LAN game, otherwise it will be the driver name.

// The array should be indexed by vehicle index.

// Frequency: Every 5 seconds
// Size: 1306 bytes
// Version: 1

type ParticipantData23 struct {
	AIControlled    uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID        uint8    // Driver id - see appendix
	NetworkID       uint8    // Network id – unique identifier for network players
	TeamID          uint8    // Team id - see appendix
	MyTeam          uint8    // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber      uint8    // Race number of the car
	Nationality     uint8    // Nationality of the driver
	Name            [48]byte // Name of participant in UTF-8 format – null terminated, Will be truncated with … (U+2026) if too long
	YourTelemetry   uint8    // The player's UDP setting, 0 = restricted, 1 = public
	ShowOnlineNames uint8    // The player's show online names setting, 0 = off, 1 = on
	Platform        uint8    // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
}

type PacketParticipantsData23 struct {
	Header        PacketHeader23
	NumActiveCars uint8 // Number of active cars in the data – should match number of cars on HUD
	Participants  [22]ParticipantData23
}

func (p *ParticipantData23) NameToString() string {
	return string(p.Name[:])
}
