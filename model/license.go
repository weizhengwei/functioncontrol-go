package model

import (
	"time"
)

type TbLicenseModule struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	LicenseNumber  string    `xorm:"VARCHAR(32)"`
	Functions      string    `xorm:"VARCHAR(512)"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LicenseState   int       `xorm:"default 0 INT(11)"`
	Sn             string    `xorm:"VARCHAR(128)"`
	VerifyTime     time.Time `xorm:"DATETIME"`
	ExpirationTime time.Time `xorm:"DATETIME"`
	Verified       int       `xorm:"default 0 TINYINT(1)"`
}
