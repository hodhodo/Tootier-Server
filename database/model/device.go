package model

type Device struct {
	DeviceID           uint64 `json:"device_id" gorm:"primaryKey"`
	Hardware           `gorm:"embedded"`
	Token              string     `json:"token" faker:"jwt"`
	LastTimeVisit      string  `json:"last_time_visit" faker:"timestamp"`
	RegisterTimeDevice string  `json:"register_time_device" faker:"timestamp"`
	UserID             uint64     `json:"user_id" faker:"oneof: 1, 6"`
}

type Hardware struct {
	DeviceName string `json:"device_name"  faker:"oneof:Sony , Sumsung , LG"`
	MacAddress string `json:"mac_address" gorm:"unique"  faker:"mac_address"`
	OS         string `json:"os" faker:"oneof:Android,IOS"`
	IP         string `json:"ip" faker:"ipv4"`
	DeviceUUID string `json:"device_uuid" gorm:"unique" faker:"uuid_digit"` // FireBase
}
