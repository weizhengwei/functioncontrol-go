package model

import (
	"time"
	"github.com/go-xorm/xorm"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	 "github.com/elgs/gostrgen"
)

type TbLicenseModule struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	LicenseNumber  string    `xorm:"VARCHAR(32)"`
	Functions      string    `xorm:"VARCHAR(512)"`
	CreateTime     time.Time `xorm:"created"`
	LicenseState   int       `xorm:"default 0 INT(11)"`
	Sn             string    `xorm:"VARCHAR(128)"`
	VerifyTime     time.Time `xorm:"DATETIME"`
	ExpirationTime time.Time `xorm:"DATETIME"`
	Verified       int       `xorm:"default 0 TINYINT(1)"`
}

type FunctionItem struct {
	FunctionNumber string `json:"functionnumber"`
	FunctionName string `json:"functionname"`
}

type Functions struct {
	Functions []FunctionItem `json:"functions"`
	Amount int `json:"amount"`
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
	ss, err := json.Marshal(funcs.Functions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ss))

	var licensemodule TbLicenseModule
	licensemodule.LicenseNumber, _ = GenerateRandomString()
	licensemodule.Functions = string(ss)


	affected, err := engine.Insert(&licensemodule)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = affected
	res.Write([]byte("add license ok"))
}


func VerifyLicense(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	
}


func GenerateRandomString() (string,error) {
	charsToGenerate := 16
	charSet := gostrgen.Upper | gostrgen.Digit
	includes := "" //"[]{}<>" // optionally include some additional letters
	excludes := "Ol"     //exclude big 'O' and small 'l' to avoid confusion with zero and one.

	str, err := gostrgen.RandGen(charsToGenerate, charSet, includes, excludes)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var ret string
	var k int
	for j := 0; j < 4; j++ {
		k = j*4
		if j < 3 {
			ret += str[k:k+4]+"-"
		}else{
			ret += str[k:k+4]
		}
	}
	return ret, nil
}