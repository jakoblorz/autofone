package cmd

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets"
	"github.com/jakoblorz/autofone/pkg/log"
	"github.com/spf13/cobra"
)

var (
	to      string
	sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send a test package",
		Run: func(cmd *cobra.Command, args []string) {
			addr, err := net.ResolveUDPAddr("udp", to)
			if err != nil {
				log.Printf("%+v", err)
				return
			}

			log.Verbosef("dialing %s", addr.String())
			conn, err := net.DialUDP("udp", nil, addr)
			if err != nil {
				log.Printf("%+v", err)
				return
			}
			defer conn.Close()

			buf := new(bytes.Buffer)
			err = binary.Write(buf, binary.LittleEndian, &motionData)
			if err != nil {
				err = nil
			} else {
				log.Verbosef("sending packet with id %d: %s", constants.PacketMotion, buf.String())
			}

			err = binary.Write(conn, binary.LittleEndian, &motionData)
			if err != nil {
				log.Printf("error occured during send: %+v", err)
				return
			}

			log.Printf("done")
		},
	}

	motionData = packets.PacketMotionData21{
		Header: packets.PacketHeader{
			PacketFormat:            2021,
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
		CarMotionData:          [22]packets.CarMotionData21{}, // empty
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

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVar(&to, "to", "localhost:20777", "Address to send the package to")
}
