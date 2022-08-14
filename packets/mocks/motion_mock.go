package mocks

import (
	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets"
)

var (
	PacketMotionData21Bytes = []byte{229, 7, 1, 5, 1, 0, 187, 134, 38, 178, 108, 178, 251, 17, 160, 1, 214, 67, 255, 32, 0, 0, 19, 255, 12, 20, 22, 68, 3, 11, 34, 192, 34, 35, 54, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 156, 96, 255, 254, 12, 172, 244, 83, 28, 0, 157, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 132, 81, 18, 64, 193, 7, 176, 59, 147, 199, 99, 186, 43, 53, 142, 67, 219, 250, 90, 65, 228, 237, 35, 195, 76, 207, 128, 194, 155, 208, 172, 63, 173, 152, 138, 192, 70, 128, 196, 2, 44, 248, 232, 7, 186, 3, 78, 128, 216, 223, 14, 64, 40, 228, 240, 62, 176, 189, 75, 190, 233, 249, 208, 191, 3, 6, 162, 60, 118, 199, 238, 188, 164, 125, 29, 68, 128, 36, 120, 192, 147, 26, 58, 193, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 78, 96, 0, 0, 177, 171, 79, 84, 0, 0, 78, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 68, 141, 18, 64, 198, 200, 157, 175, 1, 0, 128, 35, 82, 118, 13, 68, 68, 149, 162, 191, 213, 193, 169, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 95, 255, 254, 88, 171, 169, 84, 228, 255, 255, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 129, 201, 18, 64, 147, 75, 212, 59, 136, 199, 99, 58, 224, 161, 16, 68, 66, 128, 231, 191, 0, 162, 140, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 95, 255, 254, 88, 171, 169, 84, 228, 255, 255, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 129, 201, 18, 64, 147, 75, 212, 59, 136, 199, 99, 58, 90, 144, 36, 68, 0, 193, 41, 193, 55, 120, 146, 67, 193, 240, 14, 194, 159, 17, 2, 64, 78, 8, 180, 193, 44, 148, 42, 6, 81, 187, 198, 68, 64, 0, 14, 148, 30, 157, 255, 191, 144, 35, 25, 61, 144, 203, 188, 189, 60, 222, 8, 192, 127, 36, 34, 61, 14, 23, 1, 187, 175, 214, 17, 68, 170, 161, 231, 191, 25, 156, 129, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 140, 96, 225, 254, 250, 171, 6, 84, 0, 0, 141, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 148, 93, 18, 64, 99, 248, 216, 59, 254, 253, 87, 183, 58, 223, 20, 68, 3, 11, 34, 192, 104, 47, 76, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 95, 255, 254, 88, 171, 169, 84, 228, 255, 255, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 129, 201, 18, 64, 147, 75, 212, 59, 136, 199, 99, 58, 140, 66, 25, 68, 226, 129, 68, 192, 183, 187, 247, 193, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 156, 96, 255, 254, 12, 172, 244, 83, 28, 0, 157, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 132, 81, 18, 64, 193, 7, 176, 59, 147, 199, 99, 186, 204, 18, 6, 68, 157, 243, 154, 188, 172, 160, 237, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 95, 255, 254, 88, 171, 169, 84, 228, 255, 255, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 129, 201, 18, 64, 147, 75, 212, 59, 136, 199, 99, 58, 96, 237, 2, 68, 166, 200, 7, 63, 36, 249, 4, 195, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 95, 255, 254, 88, 171, 169, 84, 228, 255, 255, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 129, 201, 18, 64, 147, 75, 212, 59, 136, 199, 99, 58, 153, 116, 10, 68, 225, 215, 14, 191, 96, 119, 197, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 156, 96, 255, 254, 12, 172, 244, 83, 28, 0, 157, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 132, 81, 18, 64, 193, 7, 176, 59, 147, 199, 99, 186, 186, 13, 24, 68, 226, 129, 68, 192, 23, 234, 17, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 95, 255, 254, 88, 171, 169, 84, 228, 255, 255, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 129, 201, 18, 64, 147, 75, 212, 59, 136, 199, 99, 58, 34, 171, 14, 68, 68, 149, 162, 191, 211, 187, 158, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 156, 96, 255, 254, 12, 172, 244, 83, 28, 0, 157, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 132, 81, 18, 64, 193, 7, 176, 59, 147, 199, 99, 186, 146, 222, 36, 68, 154, 2, 162, 192, 70, 232, 177, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 156, 96, 255, 254, 12, 172, 244, 83, 28, 0, 157, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 132, 81, 18, 64, 193, 7, 176, 59, 147, 199, 99, 186, 144, 113, 31, 68, 20, 101, 138, 192, 235, 85, 41, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 95, 255, 254, 88, 171, 169, 84, 228, 255, 255, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 129, 201, 18, 64, 147, 75, 212, 59, 136, 199, 99, 58, 208, 72, 28, 68, 40, 122, 114, 192, 31, 34, 137, 193, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 35, 96, 33, 255, 129, 171, 127, 84, 160, 255, 36, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 102, 173, 18, 64, 58, 179, 231, 59, 94, 111, 64, 59, 199, 169, 35, 68, 20, 11, 162, 192, 252, 207, 133, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 96, 225, 254, 106, 171, 151, 84, 0, 0, 15, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 148, 189, 18, 64, 29, 106, 216, 59, 3, 1, 88, 55, 197, 2, 210, 67, 55, 231, 197, 63, 63, 111, 199, 66, 144, 116, 197, 193, 183, 77, 5, 191, 134, 38, 0, 194, 57, 178, 77, 254, 93, 154, 164, 101, 211, 252, 70, 178, 106, 215, 0, 192, 0, 246, 223, 60, 236, 170, 64, 188, 0, 66, 31, 192, 179, 193, 62, 60, 85, 121, 203, 60, 45, 34, 4, 68, 166, 200, 7, 63, 93, 236, 254, 194, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 156, 96, 255, 254, 12, 172, 244, 83, 28, 0, 157, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 132, 81, 18, 64, 193, 7, 176, 59, 147, 199, 99, 186, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 184, 219, 137, 193, 187, 219, 137, 193, 192, 64, 208, 192, 202, 64, 208, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128}
	PacketMotionData21      = packets.PacketMotionData21{
		Header: packets.PacketHeader{
			PacketFormat:            constants.PacketFormat_2021,
			GameMajorVersion:        1,
			GameMinorVersion:        5,
			PacketVersion:           1,
			PacketID:                0,
			SessionTime:             428.0126953125,
			SessionUID:              1295825497714230971,
			FrameIdentifier:         8447,
			PlayerCarIndex:          19,
			SecondaryPlayerCarIndex: 255,
		},
		CarMotionData: [22]packets.CarMotionData21{
			{
				WorldPositionX:     600.313232421875,
				WorldPositionY:     -2.5319221019744873,
				WorldPositionZ:     -45.53430938720703,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   44044,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     284.4153747558594,
				WorldPositionY:     13.686244010925293,
				WorldPositionZ:     -163.92926025390625,
				WorldVelocityX:     -64.40487670898438,
				WorldVelocityY:     1.3501161336898804,
				WorldVelocityZ:     -4.331137180328369,
				WorldForwardDirX:   32838,
				WorldForwardDirY:   708,
				WorldForwardDirZ:   63532,
				WorldRightDirX:     2024,
				WorldRightDirY:     954,
				WorldRightDirZ:     32846,
				GForceLateral:      2.232412338256836,
				GForceLongitudinal: 0.4704906940460205,
				GForceVertical:     -0.19896578788757324,
				Yaw:                -1.6326266527175903,
				Pitch:              0.019778257235884666,
				Roll:               -0.02914784476161003,
			},
			{
				WorldPositionX:     629.963134765625,
				WorldPositionY:     -3.877227783203125,
				WorldPositionZ:     -11.631487846374512,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24654,
				WorldForwardDirY:   0,
				WorldForwardDirZ:   43953,
				WorldRightDirX:     21583,
				WorldRightDirY:     0,
				WorldRightDirZ:     24654,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.289872169494629,
				Pitch:              -2.8700791743219156e-10,
				Roll:               1.3877789462175682e-17,
			},
			{
				WorldPositionX:     565.8487548828125,
				WorldPositionY:     -1.2701802253723145,
				WorldPositionZ:     -84.87857818603516,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   43864,
				WorldRightDirX:     21673,
				WorldRightDirY:     65508,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     578.529296875,
				WorldPositionY:     -1.8086016178131104,
				WorldPositionZ:     -70.31640625,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   43864,
				WorldRightDirX:     21673,
				WorldRightDirY:     65508,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     658.2554931640625,
				WorldPositionY:     -10.609619140625,
				WorldPositionZ:     292.9391784667969,
				WorldVelocityX:     -35.735111236572266,
				WorldVelocityY:     2.032325506210327,
				WorldVelocityZ:     -22.50405502319336,
				WorldForwardDirX:   37932,
				WorldForwardDirY:   1578,
				WorldForwardDirZ:   47953,
				WorldRightDirX:     17606,
				WorldRightDirY:     64,
				WorldRightDirZ:     37902,
				GForceLateral:      -1.9969823360443115,
				GForceLongitudinal: 0.0373874306678772,
				GForceVertical:     -0.0921851396560669,
				Yaw:                -2.138564109802246,
				Pitch:              0.03958558663725853,
				Roll:               -0.001969757955521345,
			},
			{
				WorldPositionX:     583.3544311523438,
				WorldPositionY:     -1.8096210956573486,
				WorldPositionZ:     -64.80487823486328,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24716,
				WorldForwardDirY:   65249,
				WorldForwardDirZ:   44026,
				WorldRightDirX:     21510,
				WorldRightDirY:     0,
				WorldRightDirZ:     24717,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.286961555480957,
				Pitch:              0.006621406879276037,
				Roll:               -0.000012874135791207664,
			},
			{
				WorldPositionX:     595.4879150390625,
				WorldPositionY:     -2.5319221019744873,
				WorldPositionZ:     -51.046295166015625,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   43864,
				WorldRightDirX:     21673,
				WorldRightDirY:     65508,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     613.039794921875,
				WorldPositionY:     -3.070427417755127,
				WorldPositionZ:     -30.966657638549805,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   44044,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     536.293701171875,
				WorldPositionY:     -0.01891499198973179,
				WorldPositionZ:     -118.81381225585938,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   43864,
				WorldRightDirX:     21673,
				WorldRightDirY:     65508,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     523.708984375,
				WorldPositionY:     0.5304054021835327,
				WorldPositionZ:     -132.97320556640625,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   43864,
				WorldRightDirX:     21673,
				WorldRightDirY:     65508,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     553.8218383789062,
				WorldPositionY:     -0.557981550693512,
				WorldPositionZ:     -98.733154296875,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   44044,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     608.2144775390625,
				WorldPositionY:     -3.070427417755127,
				WorldPositionZ:     -36.47860336303711,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   43864,
				WorldRightDirX:     21673,
				WorldRightDirY:     65508,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     570.6739501953125,
				WorldPositionY:     -1.2701802253723145,
				WorldPositionZ:     -79.3668441772461,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   44044,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     659.4776611328125,
				WorldPositionY:     -5.062817573547363,
				WorldPositionZ:     22.238414764404297,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   44044,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     637.7744140625,
				WorldPositionY:     -4.324838638305664,
				WorldPositionZ:     -2.645869016647339,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   43864,
				WorldRightDirX:     21673,
				WorldRightDirY:     65508,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     625.1376953125,
				WorldPositionY:     -3.788705825805664,
				WorldPositionZ:     -17.141660690307617,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24611,
				WorldForwardDirY:   65313,
				WorldForwardDirZ:   43905,
				WorldRightDirX:     21631,
				WorldRightDirY:     65440,
				WorldRightDirZ:     24612,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2918334007263184,
				Pitch:              0.007070926018059254,
				Roll:               0.002936325501650572,
			},
			{
				WorldPositionX:     654.6527709960938,
				WorldPositionY:     -5.063852310180664,
				WorldPositionZ:     16.72655487060547,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24590,
				WorldForwardDirY:   65249,
				WorldForwardDirZ:   43882,
				WorldRightDirX:     21655,
				WorldRightDirY:     0,
				WorldRightDirZ:     24591,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.292820930480957,
				Pitch:              0.006604446563869715,
				Roll:               0.000012874838830612134,
			},
			{
				WorldPositionX:     420.0216369628906,
				WorldPositionY:     1.5461186170578003,
				WorldPositionZ:     99.71727752685547,
				WorldVelocityX:     -24.681915283203125,
				WorldVelocityY:     -0.5207170844078064,
				WorldVelocityZ:     -32.037620544433594,
				WorldForwardDirX:   45625,
				WorldForwardDirY:   65101,
				WorldForwardDirZ:   39517,
				WorldRightDirX:     26020,
				WorldRightDirY:     64723,
				WorldRightDirZ:     45638,
				GForceLateral:      -2.0131478309631348,
				GForceLongitudinal: 0.02733898162841797,
				GForceVertical:     -0.011759500950574875,
				Yaw:                -2.4884033203125,
				Pitch:              0.011642861180007458,
				Roll:               0.024838129058480263,
			},
			{
				WorldPositionX:     528.5339965820312,
				WorldPositionY:     0.5304054021835327,
				WorldPositionZ:     -127.4616470336914,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   65279,
				WorldForwardDirZ:   44044,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     0,
				WorldPositionY:     0,
				WorldPositionZ:     0,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   0,
				WorldForwardDirY:   0,
				WorldForwardDirZ:   0,
				WorldRightDirX:     0,
				WorldRightDirY:     0,
				WorldRightDirZ:     0,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                0,
				Pitch:              0,
				Roll:               0,
			},
			{
				WorldPositionX:     0,
				WorldPositionY:     0,
				WorldPositionZ:     0,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   0,
				WorldForwardDirY:   0,
				WorldForwardDirZ:   0,
				WorldRightDirX:     0,
				WorldRightDirY:     0,
				WorldRightDirZ:     0,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                0,
				Pitch:              0,
				Roll:               0,
			},
		},
		SuspensionPosition: [4]float32{
			-17.232284545898438,
			-17.232290267944336,
			-6.507904052734375,
			-6.507908821105957,
		},
		SuspensionVelocity:     [4]float32{0, 0, 0, 0},
		SuspensionAcceleration: [4]float32{0, 0, 0, 0},
		WheelSpeed:             [4]float32{0, 0, 0, 0},
		WheelSlip:              [4]float32{0, 0, 0, 0},
		LocalVelocityX:         0,
		LocalVelocityY:         0,
		LocalVelocityZ:         0,
		AngularVelocityX:       0,
		AngularVelocityY:       0,
		AngularVelocityZ:       -0,
		AngularAccelerationX:   0,
		AngularAccelerationY:   0,
		AngularAccelerationZ:   0,
		FrontWheelsAngle:       -0,
	}

	PacketMotionData22 = packets.PacketMotionData22{
		Header: packets.PacketHeader{
			PacketFormat:            constants.PacketFormat_2021,
			GameMajorVersion:        1,
			GameMinorVersion:        5,
			PacketVersion:           1,
			PacketID:                0,
			SessionTime:             428.0126953125,
			SessionUID:              1295825497714230971,
			FrameIdentifier:         8447,
			PlayerCarIndex:          19,
			SecondaryPlayerCarIndex: 255,
		},
		CarMotionData: [22]packets.CarMotionData22{
			{
				WorldPositionX:     600.313232421875,
				WorldPositionY:     -2.5319221019744873,
				WorldPositionZ:     -45.53430938720703,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4404,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     284.4153747558594,
				WorldPositionY:     13.686244010925293,
				WorldPositionZ:     -163.92926025390625,
				WorldVelocityX:     -64.40487670898438,
				WorldVelocityY:     1.3501161336898804,
				WorldVelocityZ:     -4.331137180328369,
				WorldForwardDirX:   3283,
				WorldForwardDirY:   708,
				WorldForwardDirZ:   6353,
				WorldRightDirX:     2024,
				WorldRightDirY:     954,
				WorldRightDirZ:     3284,
				GForceLateral:      2.232412338256836,
				GForceLongitudinal: 0.4704906940460205,
				GForceVertical:     -0.19896578788757324,
				Yaw:                -1.6326266527175903,
				Pitch:              0.019778257235884666,
				Roll:               -0.02914784476161003,
			},
			{
				WorldPositionX:     629.963134765625,
				WorldPositionY:     -3.877227783203125,
				WorldPositionZ:     -11.631487846374512,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24654,
				WorldForwardDirY:   0,
				WorldForwardDirZ:   4395,
				WorldRightDirX:     21583,
				WorldRightDirY:     0,
				WorldRightDirZ:     24654,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.289872169494629,
				Pitch:              -2.8700791743219156e-10,
				Roll:               1.3877789462175682e-17,
			},
			{
				WorldPositionX:     565.8487548828125,
				WorldPositionY:     -1.2701802253723145,
				WorldPositionZ:     -84.87857818603516,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4386,
				WorldRightDirX:     21673,
				WorldRightDirY:     6550,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     578.529296875,
				WorldPositionY:     -1.8086016178131104,
				WorldPositionZ:     -70.31640625,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4386,
				WorldRightDirX:     21673,
				WorldRightDirY:     6550,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     658.2554931640625,
				WorldPositionY:     -10.609619140625,
				WorldPositionZ:     292.9391784667969,
				WorldVelocityX:     -35.735111236572266,
				WorldVelocityY:     2.032325506210327,
				WorldVelocityZ:     -22.50405502319336,
				WorldForwardDirX:   3793,
				WorldForwardDirY:   1578,
				WorldForwardDirZ:   4795,
				WorldRightDirX:     17606,
				WorldRightDirY:     64,
				WorldRightDirZ:     3792,
				GForceLateral:      -1.9969823360443115,
				GForceLongitudinal: 0.0373874306678772,
				GForceVertical:     -0.0921851396560669,
				Yaw:                -2.138564109802246,
				Pitch:              0.03958558663725853,
				Roll:               -0.001969757955521345,
			},
			{
				WorldPositionX:     583.3544311523438,
				WorldPositionY:     -1.8096210956573486,
				WorldPositionZ:     -64.80487823486328,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24716,
				WorldForwardDirY:   6524,
				WorldForwardDirZ:   4402,
				WorldRightDirX:     21510,
				WorldRightDirY:     0,
				WorldRightDirZ:     24717,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.286961555480957,
				Pitch:              0.006621406879276037,
				Roll:               -0.000012874135791207664,
			},
			{
				WorldPositionX:     595.4879150390625,
				WorldPositionY:     -2.5319221019744873,
				WorldPositionZ:     -51.046295166015625,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4386,
				WorldRightDirX:     21673,
				WorldRightDirY:     6550,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     613.039794921875,
				WorldPositionY:     -3.070427417755127,
				WorldPositionZ:     -30.966657638549805,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4404,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     536.293701171875,
				WorldPositionY:     -0.01891499198973179,
				WorldPositionZ:     -118.81381225585938,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4386,
				WorldRightDirX:     21673,
				WorldRightDirY:     6550,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     523.708984375,
				WorldPositionY:     0.5304054021835327,
				WorldPositionZ:     -132.97320556640625,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4386,
				WorldRightDirX:     21673,
				WorldRightDirY:     6550,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     553.8218383789062,
				WorldPositionY:     -0.557981550693512,
				WorldPositionZ:     -98.733154296875,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4404,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     608.2144775390625,
				WorldPositionY:     -3.070427417755127,
				WorldPositionZ:     -36.47860336303711,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4386,
				WorldRightDirX:     21673,
				WorldRightDirY:     6550,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     570.6739501953125,
				WorldPositionY:     -1.2701802253723145,
				WorldPositionZ:     -79.3668441772461,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4404,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     659.4776611328125,
				WorldPositionY:     -5.062817573547363,
				WorldPositionZ:     22.238414764404297,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4404,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     637.7744140625,
				WorldPositionY:     -4.324838638305664,
				WorldPositionZ:     -2.645869016647339,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24574,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4386,
				WorldRightDirX:     21673,
				WorldRightDirY:     6550,
				WorldRightDirZ:     24575,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.293548822402954,
				Pitch:              0.006478735711425543,
				Roll:               0.0008689095266163349,
			},
			{
				WorldPositionX:     625.1376953125,
				WorldPositionY:     -3.788705825805664,
				WorldPositionZ:     -17.141660690307617,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24611,
				WorldForwardDirY:   6531,
				WorldForwardDirZ:   4390,
				WorldRightDirX:     21631,
				WorldRightDirY:     6544,
				WorldRightDirZ:     24612,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2918334007263184,
				Pitch:              0.007070926018059254,
				Roll:               0.002936325501650572,
			},
			{
				WorldPositionX:     654.6527709960938,
				WorldPositionY:     -5.063852310180664,
				WorldPositionZ:     16.72655487060547,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24590,
				WorldForwardDirY:   6524,
				WorldForwardDirZ:   4388,
				WorldRightDirX:     21655,
				WorldRightDirY:     0,
				WorldRightDirZ:     24591,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.292820930480957,
				Pitch:              0.006604446563869715,
				Roll:               0.000012874838830612134,
			},
			{
				WorldPositionX:     420.0216369628906,
				WorldPositionY:     1.5461186170578003,
				WorldPositionZ:     99.71727752685547,
				WorldVelocityX:     -24.681915283203125,
				WorldVelocityY:     -0.5207170844078064,
				WorldVelocityZ:     -32.037620544433594,
				WorldForwardDirX:   4562,
				WorldForwardDirY:   6510,
				WorldForwardDirZ:   3951,
				WorldRightDirX:     26020,
				WorldRightDirY:     6472,
				WorldRightDirZ:     4563,
				GForceLateral:      -2.0131478309631348,
				GForceLongitudinal: 0.02733898162841797,
				GForceVertical:     -0.011759500950574875,
				Yaw:                -2.4884033203125,
				Pitch:              0.011642861180007458,
				Roll:               0.024838129058480263,
			},
			{
				WorldPositionX:     528.5339965820312,
				WorldPositionY:     0.5304054021835327,
				WorldPositionZ:     -127.4616470336914,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   24732,
				WorldForwardDirY:   6527,
				WorldForwardDirZ:   4404,
				WorldRightDirX:     21492,
				WorldRightDirY:     28,
				WorldRightDirZ:     24733,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                2.2862253189086914,
				Pitch:              0.005372018087655306,
				Roll:               -0.000868910166900605,
			},
			{
				WorldPositionX:     0,
				WorldPositionY:     0,
				WorldPositionZ:     0,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   0,
				WorldForwardDirY:   0,
				WorldForwardDirZ:   0,
				WorldRightDirX:     0,
				WorldRightDirY:     0,
				WorldRightDirZ:     0,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                0,
				Pitch:              0,
				Roll:               0,
			},
			{
				WorldPositionX:     0,
				WorldPositionY:     0,
				WorldPositionZ:     0,
				WorldVelocityX:     0,
				WorldVelocityY:     0,
				WorldVelocityZ:     0,
				WorldForwardDirX:   0,
				WorldForwardDirY:   0,
				WorldForwardDirZ:   0,
				WorldRightDirX:     0,
				WorldRightDirY:     0,
				WorldRightDirZ:     0,
				GForceLateral:      0,
				GForceLongitudinal: 0,
				GForceVertical:     0,
				Yaw:                0,
				Pitch:              0,
				Roll:               0,
			},
		},
		SuspensionPosition: [4]float32{
			-17.232284545898438,
			-17.232290267944336,
			-6.507904052734375,
			-6.507908821105957,
		},
		SuspensionVelocity:     [4]float32{0, 0, 0, 0},
		SuspensionAcceleration: [4]float32{0, 0, 0, 0},
		WheelSpeed:             [4]float32{0, 0, 0, 0},
		WheelSlip:              [4]float32{0, 0, 0, 0},
		LocalVelocityX:         0,
		LocalVelocityY:         0,
		LocalVelocityZ:         0,
		AngularVelocityX:       0,
		AngularVelocityY:       0,
		AngularVelocityZ:       -0,
		AngularAccelerationX:   0,
		AngularAccelerationY:   0,
		AngularAccelerationZ:   0,
		FrontWheelsAngle:       -0,
	}
)
