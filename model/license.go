package model

import (
	"time"
	"github.com/go-xorm/xorm"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
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

type FunctionItem struct {
	FunctionNumber string
	FunctionName string
}

type Functions struct {
	functions []FunctionItem
	amount int
}

func GetAllLicense(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {

}

//curl -X POST -d "@addlicense.json" localhost:9090/api/license
func AddLicense(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	ret, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(ret))
	var funcs Functions
	if err := json.Unmarshal(ret, &funcs); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(funcs)
	fmt.Println(funcs.amount)

	// affected, err := engine.Insert(&licensemodule)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// _ = affected
	//res.Write([]byte("add license ok"))
}

func VerifyLicense(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	
}