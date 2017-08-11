package model

import (
	"time"
)

type TbFunctionModule struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	FunctionNumber string    `xorm:"VARCHAR(32)"`
	FunctionType   int       `xorm:"INT(11)"`
	FunctionName   string    `xorm:"VARCHAR(255)"`
	State          int       `xorm:"default 0 INT(11)"`
	CreateTime     time.Time `xorm:"DATETIME"`
}
