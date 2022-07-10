package packets

// This packet gives details of events that happen during the course of a session.

// Frequency: When the event occurs
// Size: 35 bytes
// Version: 1

type FastestLap struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float32 // Lap time is in seconds
}

type PacketEventFastestLap struct {
	Header          PacketHeader
	EventStringCode [4]uint8   // Event string code, see below
	EventDetails    FastestLap // Event details - should be interpreted differently
}

type FastestLap22 struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float32 // Lap time is in seconds
}

type PacketEventFastestLap22 struct {
	Header          PacketHeader
	EventStringCode [4]uint8     // Event string code, see below
	EventDetails    FastestLap22 // Event details - should be interpreted differently
}

type Penalty struct {
	PenaltyType      uint8 // Penalty type – see docs/TYPES.md#penalty-types
	InfringementType uint8 // Infringement type – see docs/TYPES.md#infringement-types
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

type PacketEventPenalty struct {
	Header          PacketHeader
	EventStringCode [4]uint8 // Event string code, see below
	EventDetails    Penalty  // Event details - should be interpreted differently
}

type Penalty22 struct {
	PenaltyType      uint8 // Penalty type – see docs/TYPES.md#penalty-types
	InfringementType uint8 // Infringement type – see docs/TYPES.md#infringement-types
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

type PacketEventPenalty22 struct {
	Header          PacketHeader
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Penalty22 // Event details - should be interpreted differently
}

type SpeedTrap struct {
	VehicleIdx                uint8   // Vehicle index of the vehicle triggering speed trap
	Speed                     float32 // Top speed achieved in kilometres per hour
	IsOverallFastestInSession uint8   // Overall fastest speed in session = 1, otherwise 0
	IsDriverFastestInSession  uint8   // Fastest speed for driver in session = 1, otherwise 0
}

type PacketEventSpeedTrap struct {
	Header          PacketHeader
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    SpeedTrap // Event details - should be interpreted differently
}

type SpeedTrap22 struct {
	VehicleIdx                 uint8   // Vehicle index of the vehicle triggering speed trap
	Speed                      float32 // Top speed achieved in kilometres per hour
	IsOverallFastestInSession  uint8   // Overall fastest speed in session = 1, otherwise 0
	IsDriverFastestInSession   uint8   // Fastest speed for driver in session = 1, otherwise 0
	FastestVehicleIdxInSession uint8   // Vehicle index of the vehicle that is the fastest in this session
	FastestSpeedInSession      float32 // Speed of the vehicle that is the fastest in this session
}

type PacketEventSpeedTrap22 struct {
	Header          PacketHeader
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    SpeedTrap22 // Event details - should be interpreted differently
}

type StartLights struct {
	NumLights uint8 // Number of lights showing
}

type PacketEventStartLights struct {
	Header          PacketHeader
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    StartLights // Event details - should be interpreted differently
}

type StartLights22 struct {
	NumLights uint8 // Number of lights showing
}

type PacketEventStartLights22 struct {
	Header          PacketHeader
	EventStringCode [4]uint8      // Event string code, see below
	EventDetails    StartLights22 // Event details - should be interpreted differently
}

type Flashback struct {
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float32 // Session time flashed back to
}

type PacketEventFlashback struct {
	Header          PacketHeader
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Flashback // Event details - should be interpreted differently
}
type Flashback22 struct {
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float32 // Session time flashed back to
}

type PacketEventFlashback22 struct {
	Header          PacketHeader
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    Flashback22 // Event details - should be interpreted differently
}

type Buttons struct {
	ButtonStatus uint32 // Bit flags specifying which buttons are being pressed currently - see appendices
}

type PacketEventButtons struct {
	Header          PacketHeader
	EventStringCode [4]uint8 // Event string code, see below
	EventDetails    Buttons  // Event details - should be interpreted differently
}

type Buttons22 struct {
	ButtonStatus uint32 // Bit flags specifying which buttons are being pressed currently - see appendices
}

type PacketEventButtons22 struct {
	Header          PacketHeader
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Buttons22 // Event details - should be interpreted differently
}

type GenericVehicleEvent struct {
	VehicleIdx uint8 // Vehicle index
}

type PacketEventGenericVehicleEvent struct {
	Header          PacketHeader
	EventStringCode [4]uint8            // Event string code, see below
	EventDetails    GenericVehicleEvent // Event details - should be interpreted differently
}
type GenericVehicleEvent22 struct {
	VehicleIdx uint8 // Vehicle index
}

type PacketEventGenericVehicleEvent22 struct {
	Header          PacketHeader
	EventStringCode [4]uint8              // Event string code, see below
	EventDetails    GenericVehicleEvent22 // Event details - should be interpreted differently
}

type PacketEventGenericSessionEvent struct {
	Header          PacketHeader
	EventStringCode [4]uint8 // Event string code, see below
}

type PacketEventGenericSessionEvent22 struct {
	Header          PacketHeader
	EventStringCode [4]uint8 // Event string code, see below
}

type PacketEventHeader struct {
	Header          PacketHeader
	EventStringCode [4]uint8
}

func (p *PacketEventHeader) EventCodeString() string {
	return string(p.EventStringCode[:])
}
