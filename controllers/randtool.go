package controllers

import (
	"github.com/astaxie/beego"
	"math/rand"
	"time"
)

var rows, lengths int

//48-57 数字
//65-90 大写
//97-122 小写
//33-47 58-64 91-96 123-126 特殊字符
func Ascii2String(start, end int, flag bool) (res string) {
	asciiSlice := []byte{}
	if flag {
		for i := start; i <= end; i++ {
			asciiSlice = append(asciiSlice, byte(i))
		}
	}
	res = string(asciiSlice)
	return res
}

func ChoiceFlag(FlagDigits, FlagUppercase, FlagLowercase, FlagPunctuation bool) (str string) {
	strDigits := Ascii2String(48, 57, FlagDigits)
	strUppercase := Ascii2String(65, 90, FlagUppercase)
	strLowercase := Ascii2String(97, 122, FlagLowercase)
	strPunctuation := Ascii2String(33, 47, FlagPunctuation) +
		Ascii2String(58, 64, FlagPunctuation) +
		Ascii2String(91, 96, FlagPunctuation) +
		Ascii2String(123, 126, FlagPunctuation)
	//利用二进制方法，跟linux用户权限一样 FlagDigits(0001) FlagUppercase(0010) FlagLowercase(0100) FlagPunctuation(1000)
	var bitDigits, bitUppercase, bitLowercase, bitPunctuation, sumBitFlag int
	if FlagDigits {
		bitDigits = 1
	} else {
		bitDigits = 0
	}

	if FlagUppercase {
		bitUppercase = 2
	} else {
		bitUppercase = 0
	}

	if FlagLowercase {
		bitLowercase = 4
	} else {
		bitLowercase = 0
	}

	if FlagPunctuation {
		bitPunctuation = 8
	} else {
		bitPunctuation = 0
	}
	sumBitFlag = bitDigits + bitUppercase + bitLowercase + bitPunctuation
	//          0001      0010      0100         1000
	switch sumBitFlag {
	case 1:
		str = strDigits
	case 2:
		str = strUppercase
	case 3:
		str = strDigits + strUppercase
	case 4:
		str = strLowercase
	case 5:
		str = strDigits + strLowercase
	case 6:
		str = strUppercase + strLowercase
	case 7:
		str = strDigits + strUppercase + strLowercase
	case 8:
		str = strPunctuation
	case 9:
		str = strDigits + strPunctuation
	case 10:
		str = strUppercase + strPunctuation
	case 11:
		str = strDigits + strUppercase + strPunctuation
	case 12:
		str = strLowercase + strPunctuation
	case 13:
		str = strDigits + strLowercase + strPunctuation
	case 14:
		str = strUppercase + strLowercase + strPunctuation
	case 15:
		str = strDigits + strUppercase + strLowercase + strPunctuation
	}
	return
}

func RandomString(rows, lengths int, strIn string) (sliceOut []string) {
	bytes := []byte(strIn) //字节切片
	res := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < rows; i++ {
		for j := 0; j < lengths; j++ {
			res = append(res, bytes[r.Intn(len(bytes))])
		}
		sliceOut = append(sliceOut, string(res))
		res = res[0:0]
	}
	return
}

type RandToolController struct {
	beego.Controller
}

func (c *RandToolController) Get() {
	c.Layout = "layout.html"
	c.TplName = "randtool.html"
}

//根据官方的增加了on的判断
func MyParseBool(str string) bool {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "on", "ON":
		return true
	case "0", "f", "F", "false", "FALSE", "False", "off", "OFF":
		return false
	}
	return false
}

func (c *RandToolController) Post() {
	rows, _ = c.GetInt("rows")
	lengths, _ = c.GetInt("lengths")
	if rows == 0 {
		rows = 1
	}
	if lengths == 0 {
		lengths = 16
	}
	flagDigits := MyParseBool(c.GetString("chk_digtis"))
	flagUppercase := MyParseBool(c.GetString("chk_upper"))
	flagLowercase := MyParseBool(c.GetString("chk_lower"))
	flagPunctuation := MyParseBool(c.GetString("chk_punc"))

	strChoice := ChoiceFlag(flagDigits, flagUppercase, flagLowercase, flagPunctuation)
	sliceOut := RandomString(rows, lengths, strChoice)
	c.Data["outputpwd"] = sliceOut

	c.Layout = "layout.html"
	c.TplName = "randtool.html"
}
