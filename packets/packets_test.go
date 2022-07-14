package packets_test

import (
	"testing"

	"github.com/jakoblorz/autofone/packets"
	"github.com/jakoblorz/autofone/packets/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCarSetup(t *testing.T) {
	cs := packets.PacketCarSetupData21{}
	err := packets.Read_LE(mocks.PacketCarSetupData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack car status data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketCarSetupData21, cs, "should correctly read car status data") {
		t.Fail()
	}
}

func TestCarStatus(t *testing.T) {
	cs := packets.PacketCarStatusData21{}
	err := packets.Read_LE(mocks.PacketCarStatusData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack car status data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketCarStatusData21, cs, "should correctly read car status data") {
		t.Fail()
	}
}

func TestCarTelemetry(t *testing.T) {
	cs := packets.PacketCarTelemetryData21{}
	err := packets.Read_LE(mocks.PacketCarTelemetryData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack car telemetry data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketCarTelemetryData21, cs, "should correctly read car telemetry data") {
		t.Fail()
	}
}

func TestEventButtons(t *testing.T) {
	cs := packets.PacketEventButtons21{}
	err := packets.Read_LE(mocks.PacketEventButtons21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack event data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketEventButtons21, cs, "should correctly read event data") {
		t.Fail()
	}
}

func TestLap(t *testing.T) {
	cs := packets.PacketLapData21{}
	err := packets.Read_LE(mocks.PacketLapData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack lap data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketLapData21, cs, "should correctly read lap data") {
		t.Fail()
	}
}

func TestMotion(t *testing.T) {
	cs := packets.PacketMotionData21{}
	err := packets.Read_LE(mocks.PacketMotionData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack motion data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketMotionData21, cs, "should correctly read motion data") {
		t.Fail()
	}
}

func TestParticipants(t *testing.T) {
	cs := packets.PacketParticipantsData21{}
	err := packets.Read_LE(mocks.PacketParticipantsData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack participants data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketParticipantsData21, cs, "should correctly read participants data") {
		t.Fail()
	}
}

func TestSession(t *testing.T) {
	cs := packets.PacketSessionData21{}
	err := packets.Read_LE(mocks.PacketSessionData21Bytes, &cs)
	if !assert.NoError(t, err, "should correctly unpack session data") {
		t.Fail()
	}
	if !assert.Equal(t, mocks.PacketSessionData21, cs, "should correctly read session data") {
		t.Fail()
	}
}
