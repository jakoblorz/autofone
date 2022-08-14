package privateapi

type DevicesClient struct {
	*client
}

type DeviceRefreshTokenResponse struct {
	Token string `json:"token"`
}

func (d *DevicesClient) RefreshToken() (r *DeviceRefreshTokenResponse, err error) {
	r = new(DeviceRefreshTokenResponse)
	err = d.JSON("POST", "/api/v1/devices/refresh", nil, &r)
	return
}
