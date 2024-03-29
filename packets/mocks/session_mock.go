package mocks

import (
	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets"
)

var (
	PacketSessionData21Bytes = []byte{229, 7, 1, 5, 1, 1, 87, 47, 83, 174, 11, 32, 201, 212, 27, 79, 153, 68, 104, 98, 0, 0, 19, 255, 1, 30, 22, 200, 185, 22, 1, 13, 0, 69, 9, 16, 14, 80, 0, 0, 255, 0, 19, 54, 173, 235, 61, 0, 65, 49, 27, 62, 0, 40, 80, 63, 62, 0, 24, 235, 87, 62, 0, 131, 165, 115, 62, 0, 82, 33, 154, 62, 0, 77, 130, 173, 62, 0, 27, 19, 188, 62, 0, 5, 109, 217, 62, 0, 33, 12, 244, 62, 0, 34, 175, 11, 63, 0, 95, 75, 27, 63, 0, 136, 84, 40, 63, 0, 58, 176, 48, 63, 0, 116, 10, 58, 63, 0, 245, 169, 68, 63, 0, 154, 25, 85, 63, 0, 152, 237, 99, 63, 0, 28, 74, 118, 63, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 48, 1, 0, 1, 31, 2, 22, 2, 14, 1, 5, 1, 31, 2, 22, 2, 15, 1, 10, 1, 31, 2, 22, 2, 15, 1, 15, 2, 31, 2, 22, 2, 15, 1, 30, 2, 30, 1, 22, 2, 15, 1, 45, 2, 30, 2, 22, 2, 23, 1, 60, 2, 29, 1, 22, 2, 43, 2, 0, 2, 27, 1, 19, 1, 17, 2, 5, 2, 27, 2, 19, 2, 17, 2, 10, 2, 27, 2, 19, 2, 17, 2, 15, 2, 27, 2, 19, 2, 17, 2, 30, 2, 26, 1, 19, 2, 17, 2, 45, 2, 26, 2, 18, 1, 17, 2, 60, 2, 26, 2, 18, 2, 17, 3, 0, 1, 29, 1, 21, 1, 8, 3, 5, 1, 29, 2, 21, 2, 8, 3, 10, 1, 29, 2, 21, 2, 8, 3, 15, 1, 29, 2, 21, 2, 9, 3, 30, 1, 29, 2, 21, 2, 10, 3, 45, 1, 29, 2, 21, 2, 11, 3, 60, 1, 29, 2, 21, 2, 11, 5, 0, 1, 29, 1, 20, 1, 8, 5, 5, 1, 29, 2, 20, 2, 8, 5, 10, 1, 28, 1, 20, 2, 8, 5, 15, 1, 28, 2, 20, 2, 8, 5, 30, 1, 28, 2, 20, 2, 9, 5, 60, 1, 27, 1, 19, 1, 9, 6, 0, 1, 29, 1, 20, 1, 8, 6, 5, 1, 29, 2, 20, 2, 8, 6, 10, 1, 28, 1, 20, 2, 8, 6, 15, 1, 28, 2, 20, 2, 8, 6, 30, 1, 28, 2, 20, 2, 9, 6, 60, 1, 27, 1, 19, 1, 9, 7, 0, 1, 29, 1, 20, 1, 8, 7, 5, 1, 29, 2, 20, 2, 8, 7, 10, 1, 28, 1, 20, 2, 8, 7, 15, 1, 28, 2, 20, 2, 8, 7, 30, 1, 28, 2, 20, 2, 9, 7, 60, 1, 27, 1, 19, 1, 9, 10, 0, 0, 31, 2, 22, 2, 1, 10, 5, 0, 31, 2, 22, 2, 1, 10, 10, 0, 31, 2, 22, 2, 1, 10, 15, 0, 31, 2, 22, 2, 1, 10, 30, 0, 31, 2, 22, 2, 1, 10, 45, 0, 31, 2, 22, 2, 2, 10, 60, 0, 30, 1, 22, 2, 2, 10, 90, 0, 29, 1, 21, 1, 2, 10, 120, 0, 29, 2, 20, 1, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 31, 219, 229, 119, 168, 219, 229, 119, 168, 219, 229, 119, 168, 0, 0, 0, 0, 2, 3, 1, 1, 1, 1, 2, 1}
	PacketSessionData21      = packets.PacketSessionData21{
		Header: packets.PacketHeader21{
			PacketFormat:            constants.PacketFormat_2021,
			GameMajorVersion:        1,
			GameMinorVersion:        5,
			PacketVersion:           1,
			PacketID:                1,
			SessionTime:             1226.4720458984375,
			SessionUID:              15332821640900980567,
			FrameIdentifier:         25192,
			PlayerCarIndex:          19,
			SecondaryPlayerCarIndex: 255,
		},
		Weather:             1,
		TrackTemperature:    30,
		AirTemperature:      22,
		TotalLaps:           200,
		TrackLength:         5817,
		SessionType:         1,
		TrackID:             13,
		Formula:             0,
		SessionTimeLeft:     2373,
		SessionDuration:     3600,
		PitSpeedLimit:       80,
		GamePaused:          0,
		IsSpectating:        0,
		SpectatorCarIndex:   255,
		SliProNativeSupport: 0,
		NumMarshalZones:     19,
		MarshalZones: [21]packets.MarshalZone21{
			{ZoneStart: 0.1150764673948288, ZoneFlag: 0},
			{ZoneStart: 0.15155507624149323, ZoneFlag: 0},
			{ZoneStart: 0.18682920932769775, ZoneFlag: 0},
			{ZoneStart: 0.21085774898529053, ZoneFlag: 0},
			{ZoneStart: 0.23793606460094452, ZoneFlag: 0},
			{ZoneStart: 0.30103546380996704, ZoneFlag: 0},
			{ZoneStart: 0.33888474106788635, ZoneFlag: 0},
			{ZoneStart: 0.3673332631587982, ZoneFlag: 0},
			{ZoneStart: 0.4246598780155182, ZoneFlag: 0},
			{ZoneStart: 0.4766550362110138, ZoneFlag: 0},
			{ZoneStart: 0.5456410646438599, ZoneFlag: 0},
			{ZoneStart: 0.6066188216209412, ZoneFlag: 0},
			{ZoneStart: 0.6575398445129395, ZoneFlag: 0},
			{ZoneStart: 0.690189003944397, ZoneFlag: 0},
			{ZoneStart: 0.726722002029419, ZoneFlag: 0},
			{ZoneStart: 0.7682183384895325, ZoneFlag: 0},
			{ZoneStart: 0.8324218988418579, ZoneFlag: 0},
			{ZoneStart: 0.8903441429138184, ZoneFlag: 0},
			{ZoneStart: 0.9620683193206787, ZoneFlag: 0},
			{ZoneStart: 0, ZoneFlag: 0},
			{ZoneStart: 0, ZoneFlag: 0},
		},
		SafetyCarStatus:           0,
		NetworkGame:               0,
		NumWeatherForecastSamples: 48,
		WeatherForecastSamples: [56]packets.WeatherForecastSample21{
			{
				SessionType:            1,
				TimeOffset:             0,
				Weather:                1,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         14,
			},
			{
				SessionType:            1,
				TimeOffset:             5,
				Weather:                1,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         15,
			},
			{
				SessionType:            1,
				TimeOffset:             10,
				Weather:                1,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         15,
			},
			{
				SessionType:            1,
				TimeOffset:             15,
				Weather:                2,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         15,
			},
			{
				SessionType:            1,
				TimeOffset:             30,
				Weather:                2,
				TrackTemperature:       30,
				TrackTemperatureChange: 1,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         15,
			},
			{
				SessionType:            1,
				TimeOffset:             45,
				Weather:                2,
				TrackTemperature:       30,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         23,
			},
			{
				SessionType:            1,
				TimeOffset:             60,
				Weather:                2,
				TrackTemperature:       29,
				TrackTemperatureChange: 1,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         43,
			},
			{
				SessionType:            2,
				TimeOffset:             0,
				Weather:                2,
				TrackTemperature:       27,
				TrackTemperatureChange: 1,
				AirTemperature:         19,
				AirTemperatureChange:   1,
				RainPercentage:         17,
			},
			{
				SessionType:            2,
				TimeOffset:             5,
				Weather:                2,
				TrackTemperature:       27,
				TrackTemperatureChange: 2,
				AirTemperature:         19,
				AirTemperatureChange:   2,
				RainPercentage:         17,
			},
			{
				SessionType:            2,
				TimeOffset:             10,
				Weather:                2,
				TrackTemperature:       27,
				TrackTemperatureChange: 2,
				AirTemperature:         19,
				AirTemperatureChange:   2,
				RainPercentage:         17,
			},
			{
				SessionType:            2,
				TimeOffset:             15,
				Weather:                2,
				TrackTemperature:       27,
				TrackTemperatureChange: 2,
				AirTemperature:         19,
				AirTemperatureChange:   2,
				RainPercentage:         17,
			},
			{
				SessionType:            2,
				TimeOffset:             30,
				Weather:                2,
				TrackTemperature:       26,
				TrackTemperatureChange: 1,
				AirTemperature:         19,
				AirTemperatureChange:   2,
				RainPercentage:         17,
			},
			{
				SessionType:            2,
				TimeOffset:             45,
				Weather:                2,
				TrackTemperature:       26,
				TrackTemperatureChange: 2,
				AirTemperature:         18,
				AirTemperatureChange:   1,
				RainPercentage:         17,
			},
			{
				SessionType:            2,
				TimeOffset:             60,
				Weather:                2,
				TrackTemperature:       26,
				TrackTemperatureChange: 2,
				AirTemperature:         18,
				AirTemperatureChange:   2,
				RainPercentage:         17,
			},
			{
				SessionType:            3,
				TimeOffset:             0,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 1,
				AirTemperature:         21,
				AirTemperatureChange:   1,
				RainPercentage:         8,
			},
			{
				SessionType:            3,
				TimeOffset:             5,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         21,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            3,
				TimeOffset:             10,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         21,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            3,
				TimeOffset:             15,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         21,
				AirTemperatureChange:   2,
				RainPercentage:         9,
			},
			{
				SessionType:            3,
				TimeOffset:             30,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         21,
				AirTemperatureChange:   2,
				RainPercentage:         10,
			},
			{
				SessionType:            3,
				TimeOffset:             45,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         21,
				AirTemperatureChange:   2,
				RainPercentage:         11,
			},
			{
				SessionType:            3,
				TimeOffset:             60,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         21,
				AirTemperatureChange:   2,
				RainPercentage:         11,
			},
			{
				SessionType:            5,
				TimeOffset:             0,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 1,
				AirTemperature:         20,
				AirTemperatureChange:   1,
				RainPercentage:         8,
			},
			{
				SessionType:            5,
				TimeOffset:             5,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            5,
				TimeOffset:             10,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 1,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            5,
				TimeOffset:             15,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            5,
				TimeOffset:             30,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         9,
			},
			{
				SessionType:            5,
				TimeOffset:             60,
				Weather:                1,
				TrackTemperature:       27,
				TrackTemperatureChange: 1,
				AirTemperature:         19,
				AirTemperatureChange:   1,
				RainPercentage:         9,
			},
			{
				SessionType:            6,
				TimeOffset:             0,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 1,
				AirTemperature:         20,
				AirTemperatureChange:   1,
				RainPercentage:         8,
			},
			{
				SessionType:            6,
				TimeOffset:             5,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            6,
				TimeOffset:             10,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 1,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            6,
				TimeOffset:             15,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            6,
				TimeOffset:             30,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         9,
			},
			{
				SessionType:            6,
				TimeOffset:             60,
				Weather:                1,
				TrackTemperature:       27,
				TrackTemperatureChange: 1,
				AirTemperature:         19,
				AirTemperatureChange:   1,
				RainPercentage:         9,
			},
			{
				SessionType:            7,
				TimeOffset:             0,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 1,
				AirTemperature:         20,
				AirTemperatureChange:   1,
				RainPercentage:         8,
			},
			{
				SessionType:            7,
				TimeOffset:             5,
				Weather:                1,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            7,
				TimeOffset:             10,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 1,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            7,
				TimeOffset:             15,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         8,
			},
			{
				SessionType:            7,
				TimeOffset:             30,
				Weather:                1,
				TrackTemperature:       28,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   2,
				RainPercentage:         9,
			},
			{
				SessionType:            7,
				TimeOffset:             60,
				Weather:                1,
				TrackTemperature:       27,
				TrackTemperatureChange: 1,
				AirTemperature:         19,
				AirTemperatureChange:   1,
				RainPercentage:         9,
			},
			{
				SessionType:            10,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         1,
			},
			{
				SessionType:            10,
				TimeOffset:             5,
				Weather:                0,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         1,
			},
			{
				SessionType:            10,
				TimeOffset:             10,
				Weather:                0,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         1,
			},
			{
				SessionType:            10,
				TimeOffset:             15,
				Weather:                0,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         1,
			},
			{
				SessionType:            10,
				TimeOffset:             30,
				Weather:                0,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         1,
			},
			{
				SessionType:            10,
				TimeOffset:             45,
				Weather:                0,
				TrackTemperature:       31,
				TrackTemperatureChange: 2,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         2,
			},
			{
				SessionType:            10,
				TimeOffset:             60,
				Weather:                0,
				TrackTemperature:       30,
				TrackTemperatureChange: 1,
				AirTemperature:         22,
				AirTemperatureChange:   2,
				RainPercentage:         2,
			},
			{
				SessionType:            10,
				TimeOffset:             90,
				Weather:                0,
				TrackTemperature:       29,
				TrackTemperatureChange: 1,
				AirTemperature:         21,
				AirTemperatureChange:   1,
				RainPercentage:         2,
			},
			{
				SessionType:            10,
				TimeOffset:             120,
				Weather:                0,
				TrackTemperature:       29,
				TrackTemperatureChange: 2,
				AirTemperature:         20,
				AirTemperatureChange:   1,
				RainPercentage:         3,
			},
			{
				SessionType:            0,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       0,
				TrackTemperatureChange: 0,
				AirTemperature:         0,
				AirTemperatureChange:   0,
				RainPercentage:         0,
			},
			{
				SessionType:            0,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       0,
				TrackTemperatureChange: 0,
				AirTemperature:         0,
				AirTemperatureChange:   0,
				RainPercentage:         0,
			},
			{
				SessionType:            0,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       0,
				TrackTemperatureChange: 0,
				AirTemperature:         0,
				AirTemperatureChange:   0,
				RainPercentage:         0,
			},
			{
				SessionType:            0,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       0,
				TrackTemperatureChange: 0,
				AirTemperature:         0,
				AirTemperatureChange:   0,
				RainPercentage:         0,
			},
			{
				SessionType:            0,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       0,
				TrackTemperatureChange: 0,
				AirTemperature:         0,
				AirTemperatureChange:   0,
				RainPercentage:         0,
			},
			{
				SessionType:            0,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       0,
				TrackTemperatureChange: 0,
				AirTemperature:         0,
				AirTemperatureChange:   0,
				RainPercentage:         0,
			},
			{
				SessionType:            0,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       0,
				TrackTemperatureChange: 0,
				AirTemperature:         0,
				AirTemperatureChange:   0,
				RainPercentage:         0,
			},
			{
				SessionType:            0,
				TimeOffset:             0,
				Weather:                0,
				TrackTemperature:       0,
				TrackTemperatureChange: 0,
				AirTemperature:         0,
				AirTemperatureChange:   0,
				RainPercentage:         0,
			},
		},
		ForecastAccuracy:       1,
		AIDifficulty:           31,
		SeasonLinkIdentifier:   2826429915,
		WeekendLinkIdentifier:  2826429915,
		SessionLinkIdentifier:  2826429915,
		PitStopWindowIdealLap:  0,
		PitStopWindowLatestLap: 0,
		PitStopRejoinPosition:  0,
		SteeringAssist:         0,
		BreakingAssist:         2,
		GearboxAssist:          3,
		PitAssist:              1,
		PitReleaseAssist:       1,
		ERSAssist:              1,
		DRSAssist:              1,
		DynamicRacingLine:      2,
		DynamicRacingLineType:  1,
	}
)
