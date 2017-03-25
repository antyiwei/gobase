package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//加密手机和身份证号为星号
func EncryptData(data string) string {
	var str string
	length := len(data)
	if 0 == length {
		return ""
	}
	if 11 == length {
		str = sToS(data, '*', 4, 4)
	} else if 15 == length {
		str = sToS(data, '*', 5, 6)
	} else if 18 == length {
		str = sToS(data, '*', 6, 8)
	}
	return str
}

//加密电子邮箱为星号
func EncryptEmail(email string) string {
	length := len(email)
	if 0 == length {
		return ""
	}
	strs := strings.Split(email, "@")
	return sToS(strs[0], '*', 1, len(strs[0])-2) + "@" + strs[1]
}

//加密姓名
func EncryptName(name string) string {
	r := []rune(name)
	r[0] = '*'
	return string(r)
}

func sToS(data string, b byte, position, length int) string {
	t := make([]byte, len(data))
	copy(t[:position], data[:position])
	for i := 0; i < length; i++ {
		t[position+i] = b
	}
	copy(t[position+length:], data[position+length:])
	return string(t[0:len(data)])
}

func RoundMoney(money int) int {
	//TODO 8400 ~9000元?   11400 ~ 1.2万?
	return 0
}

func ToString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

//获取银行卡后四位
func GetBankCardLast4(cardNumber string) string {
	return string(cardNumber[len(cardNumber)-4 : len(cardNumber)])
}

func KSort(params map[string]interface{}) {
	//keys 排序
	// 呵呵 ，现在没写
}

//对字符串进行截取
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

//格式化float类型，截断，不四舍五入
func FormatFloat64(f float64, l int) float64 {
	str1 := fmt.Sprintf("%f", f)
	strl := strings.Split(str1, ".")
	strl[1] = Substr(strl[1], 0, l)
	strre := strl[0] + "." + strl[1]
	fre, _ := strconv.ParseFloat(strre, 64)
	return fre
}

// 返回版本信息 如V1.2.3返回123
func VersionToInt(mobileversion string) int {
	if mobileversion == "" {
		return 0
	}
	versionint := 0
	if strings.Contains(mobileversion, "PC") {
		versionint, _ = strconv.Atoi(strings.Replace(strings.Replace(mobileversion, "PC", "", -1), ".", "", -1))
	} else if strings.Contains(mobileversion, "WAP") {
		versionint, _ = strconv.Atoi(strings.Replace(strings.Replace(mobileversion, "WAP", "", -1), ".", "", -1))
	} else if strings.Contains(mobileversion, "zcm") {
		versionint, _ = strconv.Atoi(strings.Replace(strings.Replace(mobileversion, "zcm", "", -1), ".", "", -1))
	} else if strings.Contains(mobileversion, "wap") {
		versionint, _ = strconv.Atoi(strings.Replace(strings.Replace(mobileversion, "wap", "", -1), ".", "", -1))
	} else {
		versionint, _ = strconv.Atoi(strings.Replace(strings.Replace(mobileversion, "V", "", -1), ".", "", -1))
	}
	return versionint
}
func VersionToString(mobileversion string) string {

	if strings.Contains(mobileversion, "PC") {
		mobileversion = strings.Replace(mobileversion, "PC", "", -1)
	} else if strings.Contains(mobileversion, "WAP") {
		mobileversion = strings.Replace(mobileversion, "WAP", "", -1)
	} else if strings.Contains(mobileversion, "zcm") {
		mobileversion = strings.Replace(mobileversion, "zcm", "", -1)
	} else if strings.Contains(mobileversion, "wap") {
		mobileversion = strings.Replace(mobileversion, "wap", "", -1)
	} else if strings.Contains(mobileversion, "V") {
		mobileversion = strings.Replace(mobileversion, "V", "", -1)
	} else if strings.Contains(mobileversion, "v") {
		mobileversion = strings.Replace(mobileversion, "v", "", -1)
	}
	return mobileversion
}

//============================================字符串数字保留小数=========================================
func ChangeNumber1(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

// 如 “1333.34434343”
func ChangeNumber(n string, m int) string {
	if n == "" {
		return ""
	}

	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

//截取小数点后几位
func SubFloatToString(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

//截取小数点后几位
func SubFloatToFloat(f float64, m int) float64 {
	newn := SubFloatToString(f, m)
	newf, _ := strconv.ParseFloat(newn, 64)
	return newf
}

// 计算每日收益
func GetDayProfit(capital int, annual_rate float32) float64 {
	return SubFloatToFloat(float64(capital)*float64(annual_rate)/100/365, 3)
}

func test(){
fmt.Prnt("OOOdfsdddddOiOOOOOOOOOOOOOOOOOOO")


}
