package packets

// The motion packet gives extended data for the car being driven with the goal of being able to drive a motion platform setup.

// Frequency: Rate as specified in menus
// Size: 217 bytes
// Version: 1

type PacketMotionExData23 struct {
	Header PacketHeader23

	// Note: All wheel arrays have the following
	// RL, RR, FL, FR
	SuspensionPosition     [4]float32 // RL, RR, FL, FR
	SuspensionVelocity     [4]float32 // RL, RR, FL, FR
	SuspensionAcceleration [4]float32 // RL, RR, FL, FR
	WheelSpeed             [4]float32 // Speed of each wheel
	WheelSlipRatio         [4]float32 // Slip angles for each wheel
	WheelLatForce          [4]float32 // Lateral force on each wheel
	WheelLongForce         [4]float32 // Longitudinal force on each wheel
	HeightOfCOGAboveGround float32    // Height of centre of gravity above ground
	LocalVelocityX         float32    // Velocity in local space – metres/s
	LocalVelocityY         float32    // Velocity in local space – metres/s
	LocalVelocityZ         float32    // Velocity in local space – metres/s
	AngularVelocityX       float32    // Angular velocity x-component – radians/s
	AngularVelocityY       float32    // Angular velocity y-component – radians/s
	AngularVelocityZ       float32    // Angular velocity z-component – radians/s
	AngularAccelerationX   float32    // Angular acceleration x-component – radians/s/s
	AngularAccelerationY   float32    // Angular acceleration y-component – radians/s/s
	AngularAccelerationZ   float32    // Angular acceleration z-component – radians/s/s
	FrontWheelAngle        float32    // Current front wheels angle in radians
	WheelVertForce         [4]float32 // Vertical force on each wheel
}
