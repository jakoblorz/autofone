package packets

type PacketEvent interface {
	EventCodeString() string
}

// This packet gives details of events that happen during the course of a session.

// Frequency: When the event occurs
// Size: 35 bytes (F1 2023: 45 bytes)
// Version: 1

type FastestLap21 struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float32 // Lap time is in seconds
}

type PacketEventFastestLap21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8     // Event string code, see below
	EventDetails    FastestLap21 // Event details - should be interpreted differently
}

type FastestLap22 struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float32 // Lap time is in seconds
}

type PacketEventFastestLap22 struct {
	Header          PacketHeader22
	EventStringCode [4]uint8     // Event string code, see below
	EventDetails    FastestLap22 // Event details - should be interpreted differently
}

type FastestLap23 struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float32 // Lap time is in seconds
}

type PacketEventFastestLap23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8     // Event string code, see below
	EventDetails    FastestLap23 // Event details - should be interpreted differently
}

type Penalty21 struct {
	PenaltyType      uint8 // Penalty type – see docs/TYPES.md#penalty-types
	InfringementType uint8 // Infringement type – see docs/TYPES.md#infringement-types
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

type PacketEventPenalty21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Penalty21 // Event details - should be interpreted differently
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
	Header          PacketHeader22
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Penalty22 // Event details - should be interpreted differently
}

type Penalty23 struct {
	PenaltyType      uint8 // Penalty type – see docs/TYPES.md#penalty-types
	InfringementType uint8 // Infringement type – see docs/TYPES.md#infringement-types
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

type PacketEventPenalty23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Penalty23 // Event details - should be interpreted differently
}

type SpeedTrap21 struct {
	VehicleIdx                uint8   // Vehicle index of the vehicle triggering speed trap
	Speed                     float32 // Top speed achieved in kilometres per hour
	IsOverallFastestInSession uint8   // Overall fastest speed in session = 1, otherwise 0
	IsDriverFastestInSession  uint8   // Fastest speed for driver in session = 1, otherwise 0
}

type PacketEventSpeedTrap21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    SpeedTrap21 // Event details - should be interpreted differently
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
	Header          PacketHeader22
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    SpeedTrap22 // Event details - should be interpreted differently
}

type SpeedTrap23 struct {
	VehicleIdx                 uint8   // Vehicle index of the vehicle triggering speed trap
	Speed                      float32 // Top speed achieved in kilometres per hour
	IsOverallFastestInSession  uint8   // Overall fastest speed in session = 1, otherwise 0
	IsDriverFastestInSession   uint8   // Fastest speed for driver in session = 1, otherwise 0
	FastestVehicleIdxInSession uint8   // Vehicle index of the vehicle that is the fastest in this session
	FastestSpeedInSession      float32 // Speed of the vehicle that is the fastest in this session
}

type PacketEventSpeedTrap23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    SpeedTrap23 // Event details - should be interpreted differently
}

type StartLights21 struct {
	NumLights uint8 // Number of lights showing
}

type PacketEventStartLights21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8      // Event string code, see below
	EventDetails    StartLights21 // Event details - should be interpreted differently
}

type StartLights22 struct {
	NumLights uint8 // Number of lights showing
}

type PacketEventStartLights22 struct {
	Header          PacketHeader22
	EventStringCode [4]uint8      // Event string code, see below
	EventDetails    StartLights22 // Event details - should be interpreted differently
}

type StartLights23 struct {
	NumLights uint8 // Number of lights showing
}

type PacketEventStartLights23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8      // Event string code, see below
	EventDetails    StartLights23 // Event details - should be interpreted differently
}

type Flashback21 struct {
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float32 // Session time flashed back to
}

type PacketEventFlashback21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    Flashback21 // Event details - should be interpreted differently
}

type Flashback22 struct {
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float32 // Session time flashed back to
}

type PacketEventFlashback22 struct {
	Header          PacketHeader22
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    Flashback22 // Event details - should be interpreted differently
}

type Flashback23 struct {
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float32 // Session time flashed back to
}

type PacketEventFlashback23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8    // Event string code, see below
	EventDetails    Flashback23 // Event details - should be interpreted differently
}

type Buttons21 struct {
	ButtonStatus uint32 // Bit flags specifying which buttons are being pressed currently - see appendices
}

type PacketEventButtons21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Buttons21 // Event details - should be interpreted differently
}

type Buttons22 struct {
	ButtonStatus uint32 // Bit flags specifying which buttons are being pressed currently - see appendices
}

type PacketEventButtons22 struct {
	Header          PacketHeader22
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Buttons22 // Event details - should be interpreted differently
}

type Buttons23 struct {
	ButtonStatus uint32 // Bit flags specifying which buttons are being pressed currently - see appendices
}

type PacketEventButtons23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8  // Event string code, see below
	EventDetails    Buttons23 // Event details - should be interpreted differently
}

type Overtake23 struct {
	OvertakingVehicleIdx     uint8 // Vehicle index of the vehicle overtaking
	BeingOvertakenVehicleIdx uint8 // Vehicle index of the vehicle being overtaken
}

type PacketEventOvertake23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8   // Event string code, see below
	EventDetails    Overtake23 // Event details - should be interpreted differently
}

type GenericVehicleEvent21 struct {
	VehicleIdx uint8 // Vehicle index
}

type PacketEventGenericVehicleEvent21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8              // Event string code, see below
	EventDetails    GenericVehicleEvent21 // Event details - should be interpreted differently
}

type GenericVehicleEvent22 struct {
	VehicleIdx uint8 // Vehicle index
}

type PacketEventGenericVehicleEvent22 struct {
	Header          PacketHeader22
	EventStringCode [4]uint8              // Event string code, see below
	EventDetails    GenericVehicleEvent22 // Event details - should be interpreted differently
}

type GenericVehicleEvent23 struct {
	VehicleIdx uint8 // Vehicle index
}

type PacketEventGenericVehicleEvent23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8              // Event string code, see below
	EventDetails    GenericVehicleEvent23 // Event details - should be interpreted differently
}

type PacketEventGenericSessionEvent21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8 // Event string code, see below
}

type PacketEventGenericSessionEvent22 struct {
	Header          PacketHeader22
	EventStringCode [4]uint8 // Event string code, see below
}

type PacketEventGenericSessionEvent23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8 // Event string code, see below
}

type PacketEventHeader21 struct {
	Header          PacketHeader21
	EventStringCode [4]uint8
}

func (p *PacketEventHeader21) EventCodeString() string {
	return string(p.EventStringCode[:])
}

type PacketEventHeader22 struct {
	Header          PacketHeader22
	EventStringCode [4]uint8
}

func (p *PacketEventHeader22) EventCodeString() string {
	return string(p.EventStringCode[:])
}

type PacketEventHeader23 struct {
	Header          PacketHeader23
	EventStringCode [4]uint8
}

func (p *PacketEventHeader23) EventCodeString() string {
	return string(p.EventStringCode[:])
}
