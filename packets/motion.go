package packets

// The motion packet gives physics data for all the cars being driven.
// There is additional data for the car being driven with the goal of being able to drive a motion platform setup.

// Frequency: Rate as specified in menus
// Size: 1464 bytes
// Version: 1

type CarMotionData21 struct {
	WorldPositionX     float32 // World space X position
	WorldPositionY     float32 // World space Y position
	WorldPositionZ     float32 // World space Z position
	WorldVelocityX     float32 // Velocity in world space X
	WorldVelocityY     float32 // Velocity in world space Y
	WorldVelocityZ     float32 // Velocity in world space Z
	WorldForwardDirX   uint16  // World space forward X direction (normalised)
	WorldForwardDirY   uint16  // World space forward Y direction (normalised)
	WorldForwardDirZ   uint16  // World space forward Z direction (normalised)
	WorldRightDirX     uint16  // World space right X direction (normalised)
	WorldRightDirY     uint16  // World space right Y direction (normalised)
	WorldRightDirZ     uint16  // World space right Z direction (normalised)
	GForceLateral      float32 // Lateral G-Force component
	GForceLongitudinal float32 // Longitudinal G-Force component
	GForceVertical     float32 // Vertical G-Force component
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

type PacketMotionData21 struct {
	Header        PacketHeader21      // Header
	CarMotionData [22]CarMotionData21 // Data for all cars on track

	// Extra player car ONLY data
	SuspensionPosition     [4]float32 // Note: All wheel arrays have the following order:
	SuspensionVelocity     [4]float32 // RL, RR, FL, FR
	SuspensionAcceleration [4]float32 // RL, RR, FL, FR
	WheelSpeed             [4]float32 // Speed of each wheel
	WheelSlip              [4]float32 // Slip ratio for each wheel
	LocalVelocityX         float32    // Velocity in local space
	LocalVelocityY         float32    // Velocity in local space
	LocalVelocityZ         float32    // Velocity in local space
	AngularVelocityX       float32    // Angular velocity x-component
	AngularVelocityY       float32    // Angular velocity y-component
	AngularVelocityZ       float32    // Angular velocity z-component
	AngularAccelerationX   float32    // Angular velocity x-component
	AngularAccelerationY   float32    // Angular velocity y-component
	AngularAccelerationZ   float32    // Angular velocity z-component
	FrontWheelsAngle       float32    // Current front wheels angle in radians
}

// The motion packet gives physics data for all the cars being driven.
// There is additional data for the car being driven with the goal of being able to drive a motion platform setup.

// Frequency: Rate as specified in menus
// Size: 1464 bytes
// Version: 1

type CarMotionData22 struct {
	WorldPositionX     float32 // World space X position
	WorldPositionY     float32 // World space Y position
	WorldPositionZ     float32 // World space Z position
	WorldVelocityX     float32 // Velocity in world space X
	WorldVelocityY     float32 // Velocity in world space Y
	WorldVelocityZ     float32 // Velocity in world space Z
	WorldForwardDirX   int16   // World space forward X direction (normalised)
	WorldForwardDirY   int16   // World space forward Y direction (normalised)
	WorldForwardDirZ   int16   // World space forward Z direction (normalised)
	WorldRightDirX     int16   // World space right X direction (normalised)
	WorldRightDirY     int16   // World space right Y direction (normalised)
	WorldRightDirZ     int16   // World space right Z direction (normalised)
	GForceLateral      float32 // Lateral G-Force component
	GForceLongitudinal float32 // Longitudinal G-Force component
	GForceVertical     float32 // Vertical G-Force component
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

type PacketMotionData22 struct {
	Header        PacketHeader22      // Header
	CarMotionData [22]CarMotionData22 // Data for all cars on track

	// Extra player car ONLY data
	SuspensionPosition     [4]float32 // Note: All wheel arrays have the following order:
	SuspensionVelocity     [4]float32 // RL, RR, FL, FR
	SuspensionAcceleration [4]float32 // RL, RR, FL, FR
	WheelSpeed             [4]float32 // Speed of each wheel
	WheelSlip              [4]float32 // Slip ratio for each wheel
	LocalVelocityX         float32    // Velocity in local space
	LocalVelocityY         float32    // Velocity in local space
	LocalVelocityZ         float32    // Velocity in local space
	AngularVelocityX       float32    // Angular velocity x-component
	AngularVelocityY       float32    // Angular velocity y-component
	AngularVelocityZ       float32    // Angular velocity z-component
	AngularAccelerationX   float32    // Angular velocity x-component
	AngularAccelerationY   float32    // Angular velocity y-component
	AngularAccelerationZ   float32    // Angular velocity z-component
	FrontWheelsAngle       float32    // Current front wheels angle in radians
}

// The motion packet gives physics data for all the cars being driven.
// N.B. For the normalised vectors below, to convert to float values divide by 32767.0f – 16-bit signed values are used to pack the data and on the assumption that direction values are always between -1.0f and 1.0f.

// Frequency: Rate as specified in menus
// Size: 1349 bytes
// Version: 1

type CarMotionData23 struct {
	WorldPositionX     float32 // World space X position - metres
	WorldPositionY     float32 // World space Y position
	WorldPositionZ     float32 // World space Z position
	WorldVelocityX     float32 // Velocity in world space X - metres per second
	WorldVelocityY     float32 // Velocity in world space Y
	WorldVelocityZ     float32 // Velocity in world space Z
	WorldForwardDirX   int16   // World space forward X direction (normalised)
	WorldForwardDirY   int16   // World space forward Y direction (normalised)
	WorldForwardDirZ   int16   // World space forward Z direction (normalised)
	WorldRightDirX     int16   // World space right X direction (normalised)
	WorldRightDirY     int16   // World space right Y direction (normalised)
	WorldRightDirZ     int16   // World space right Z direction (normalised)
	GForceLateral      float32 // Lateral G-Force component
	GForceLongitudinal float32 // Longitudinal G-Force component
	GForceVertical     float32 // Vertical G-Force component
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

type PacketMotionData23 struct {
	Header        PacketHeader23      // Header
	CarMotionData [22]CarMotionData23 // Data for all cars on track
}
