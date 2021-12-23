package pipe

import (
	"context"
	"encoding/binary"
	"net"
	"testing"
	"time"

	"github.com/jakoblorz/f1-metrics-transformer/constants"
	"github.com/jakoblorz/f1-metrics-transformer/packets"
	"github.com/jakoblorz/f1-metrics-transformer/pkg/udptest"
	"github.com/stretchr/testify/assert"
)

var (
	motionData = packets.PacketMotionData{
		Header: packets.PacketHeader{
			PacketFormat:            2020,
			GameMajorVersion:        1,
			GameMinorVersion:        18,
			PacketVersion:           1,
			PacketID:                0,
			SessionUID:              11855624319420004949,
			SessionTime:             1.5710948,
			FrameIdentifier:         43,
			PlayerCarIndex:          19,
			SecondaryPlayerCarIndex: 255,
		},
		CarMotionData:          [22]packets.CarMotionData{}, // empty
		SuspensionPosition:     [4]float32{-0.6432814, 0.049691506, -0.122375205, -0.11044062},
		SuspensionVelocity:     [4]float32{-5.7109776, -2.5368745, -0.33160102, 1.6033937},
		SuspensionAcceleration: [4]float32{-742.09515, -298.6635, -38.07375, 275.05835},
		WheelSpeed:             [4]float32{120.0, 120.0, 120.0, 120.0},
		WheelSlip:              [4]float32{10.5, 5.3, 9.5, 9.5},
		LocalVelocityX:         -0.00016372708,
		LocalVelocityY:         -0.0011803857,
		LocalVelocityZ:         0.0015708461,
		AngularVelocityX:       0.0113855535,
		AngularVelocityY:       -0.00053515914,
		AngularVelocityZ:       0.0073023424,
		AngularAccelerationX:   1.5726409,
		AngularAccelerationY:   0.047205403,
		AngularAccelerationZ:   1.0070984,
		FrontWheelsAngle:       0,
	}
)

func TestPacketReader_read(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientCh := make(chan net.Conn)
	go func() {
		defer close(clientCh)
		binary.Write(<-clientCh, binary.LittleEndian, &motionData)
	}()

	var received interface{}
	assert.NoError(t, udptest.NewConn(func(clientConn, serverConn net.Conn) error {
		clientCh <- clientConn
		reader := ReadUDPPackets(ctx, serverConn, []uint{uint(constants.PacketMotion)})
		go reader.Process()
		received = <-reader.Out()
		return nil
	}, 5*time.Second))
	assert.Equal(t, &motionData, received)
}
