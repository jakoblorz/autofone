package packets_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/jakoblorz/autofone/packets"
	"github.com/jakoblorz/autofone/packets/mocks"
	"github.com/stretchr/testify/assert"
)

func read(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}

func TestCarSetup(t *testing.T) {
	cs := packets.PacketCarSetupData21{}
	err := read(mocks.PacketCarSetupData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack car status data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketCarSetupData21, cs, "should correctly read car status data") {
		t.Fail()
	}
}

func TestCarStatus(t *testing.T) {
	cs := packets.PacketCarStatusData21{}
	err := read(mocks.PacketCarStatusData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack car status data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketCarStatusData21, cs, "should correctly read car status data") {
		t.Fail()
	}
}

func TestCarTelemetry(t *testing.T) {
	cs := packets.PacketCarTelemetryData21{}
	err := read(mocks.PacketCarTelemetry21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack car telemetry data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketCarTelemetry21, cs, "should correctly read car telemetry data") {
		t.Fail()
	}
}

func TestEventButtons(t *testing.T) {
	cs := packets.PacketEventButtons21{}
	err := read(mocks.PacketEventButtons21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack event data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketEventButtons21, cs, "should correctly read event data") {
		t.Fail()
	}
}

func TestLap(t *testing.T) {
	cs := packets.PacketLapData21{}
	err := read(mocks.PacketLapData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack lap data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketLapData21, cs, "should correctly read lap data") {
		t.Fail()
	}
}

func TestMotion(t *testing.T) {
	cs := packets.PacketMotionData21{}
	err := read(mocks.PacketMotionData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack motion data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketMotionData21, cs, "should correctly read motion data") {
		t.Fail()
	}
}

// func TestParticipants(t *testing.T) {
// 	cs := packets.PacketParticipantsData21{}
// 	err := read(mocks.PacketParticipantsData21Bytes, &cs)
// 	if !assert.NoError(t, err, "should correctly unpack motion data") {
// 		t.Fail()
// 	}
// 	if !assert.Equal(t, mocks.PacketParticipantsData21, cs, "should correctly read motion data") {
// 		t.Fail()
// 	}
// }
