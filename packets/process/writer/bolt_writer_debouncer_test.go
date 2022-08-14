package writer

import (
	"testing"

	"github.com/jakoblorz/autofone/packets/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/jakoblorz/autofone/packets"
)

func applyMotion21(m packets.PacketMotionData21, fn func(m *packets.PacketMotionData21) *packets.PacketMotionData21) *packets.PacketMotionData21 {
	return fn(&m)
}

func applyMotion22(m packets.PacketMotionData22, fn func(m *packets.PacketMotionData22) *packets.PacketMotionData22) *packets.PacketMotionData22 {
	return fn(&m)
}

func Test_averageAndLastPlayerCarMotion21_AverageAndLastPlayerCarMotion(t *testing.T) {
	tests := []struct {
		name   string
		a      averageAndLastPlayerCarMotion21
		wantFn func(t *testing.T, p *packets.PacketMotionData21) bool
	}{
		{
			name: "should correctly average 3 packet's uint16 values",
			a: averageAndLastPlayerCarMotion21{
				applyMotion21(mocks.PacketMotionData21, func(m *packets.PacketMotionData21) *packets.PacketMotionData21 {
					m.CarMotionData[0].WorldForwardDirX = 2
					m.CarMotionData[1].WorldForwardDirX = 10
					return m
				}),
				applyMotion21(mocks.PacketMotionData21, func(m *packets.PacketMotionData21) *packets.PacketMotionData21 {
					m.CarMotionData[0].WorldForwardDirX = 4
					m.CarMotionData[1].WorldForwardDirX = 20
					return m
				}),
				applyMotion21(mocks.PacketMotionData21, func(m *packets.PacketMotionData21) *packets.PacketMotionData21 {
					m.CarMotionData[0].WorldForwardDirX = 6
					m.CarMotionData[1].WorldForwardDirX = 40
					return m
				}),
			},
			wantFn: func(t *testing.T, p *packets.PacketMotionData21) bool {
				return assert.Equal(t, uint16(4), p.CarMotionData[0].WorldForwardDirX) &&
					assert.Equal(t, uint16(23), p.CarMotionData[1].WorldForwardDirX)
			},
		},
		{
			name: "should correctly average 3 packet's float32 values",
			a: averageAndLastPlayerCarMotion21{
				applyMotion21(mocks.PacketMotionData21, func(m *packets.PacketMotionData21) *packets.PacketMotionData21 {
					m.CarMotionData[0].WorldPositionX = 2
					m.CarMotionData[1].WorldPositionX = 10
					return m
				}),
				applyMotion21(mocks.PacketMotionData21, func(m *packets.PacketMotionData21) *packets.PacketMotionData21 {
					m.CarMotionData[0].WorldPositionX = 4
					m.CarMotionData[1].WorldPositionX = 20
					return m
				}),
				applyMotion21(mocks.PacketMotionData21, func(m *packets.PacketMotionData21) *packets.PacketMotionData21 {
					m.CarMotionData[0].WorldPositionX = 6
					m.CarMotionData[1].WorldPositionX = 40
					return m
				}),
			},
			wantFn: func(t *testing.T, p *packets.PacketMotionData21) bool {
				return assert.Equal(t, float32(4), p.CarMotionData[0].WorldPositionX) &&
					assert.Equal(t, float32((40+20+10)/3.0), p.CarMotionData[1].WorldPositionX)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AverageAndLastPlayerCarMotion(); !tt.wantFn(t, got) {
				t.Errorf("averageAndLastPlayerCarMotion21.AverageAndLastPlayerCarMotion() = %v", got)
			}
		})
	}
}

func Test_averageAndLastPlayerCarMotion22_AverageAndLastPlayerCarMotion(t *testing.T) {
	tests := []struct {
		name   string
		a      averageAndLastPlayerCarMotion22
		wantFn func(t *testing.T, p *packets.PacketMotionData22) bool
	}{
		{
			name: "should correctly average 3 packet's uint16 values",
			a: averageAndLastPlayerCarMotion22{
				applyMotion22(mocks.PacketMotionData22, func(m *packets.PacketMotionData22) *packets.PacketMotionData22 {
					m.CarMotionData[0].WorldForwardDirX = 2
					m.CarMotionData[1].WorldForwardDirX = 10
					return m
				}),
				applyMotion22(mocks.PacketMotionData22, func(m *packets.PacketMotionData22) *packets.PacketMotionData22 {
					m.CarMotionData[0].WorldForwardDirX = 4
					m.CarMotionData[1].WorldForwardDirX = 20
					return m
				}),
				applyMotion22(mocks.PacketMotionData22, func(m *packets.PacketMotionData22) *packets.PacketMotionData22 {
					m.CarMotionData[0].WorldForwardDirX = 6
					m.CarMotionData[1].WorldForwardDirX = 40
					return m
				}),
			},
			wantFn: func(t *testing.T, p *packets.PacketMotionData22) bool {
				return assert.Equal(t, int16(4), p.CarMotionData[0].WorldForwardDirX) &&
					assert.Equal(t, int16(23), p.CarMotionData[1].WorldForwardDirX)
			},
		},
		{
			name: "should correctly average 3 packet's float32 values",
			a: averageAndLastPlayerCarMotion22{
				applyMotion22(mocks.PacketMotionData22, func(m *packets.PacketMotionData22) *packets.PacketMotionData22 {
					m.CarMotionData[0].WorldPositionX = 2
					m.CarMotionData[1].WorldPositionX = 10
					return m
				}),
				applyMotion22(mocks.PacketMotionData22, func(m *packets.PacketMotionData22) *packets.PacketMotionData22 {
					m.CarMotionData[0].WorldPositionX = 4
					m.CarMotionData[1].WorldPositionX = 20
					return m
				}),
				applyMotion22(mocks.PacketMotionData22, func(m *packets.PacketMotionData22) *packets.PacketMotionData22 {
					m.CarMotionData[0].WorldPositionX = 6
					m.CarMotionData[1].WorldPositionX = 40
					return m
				}),
			},
			wantFn: func(t *testing.T, p *packets.PacketMotionData22) bool {
				return assert.Equal(t, float32(4), p.CarMotionData[0].WorldPositionX) &&
					assert.Equal(t, float32((40+20+10)/3.0), p.CarMotionData[1].WorldPositionX)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AverageAndLastPlayerCarMotion(); !tt.wantFn(t, got) {
				t.Errorf("averageAndLastPlayerCarMotion22.AverageAndLastPlayerCarMotion() = %v", got)
			}
		})
	}
}
