package lib

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

// StrToInt64 ...
func StrToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

// StrToUint64 ...
func StrToUint64(str string) uint64 {
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}

// StrToFloat64 ...
func StrToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

// IsContain ...
func IsContain(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// IsContainStr ...
func IsContainStr(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// RemoveDuplicate ... int64 array deduplication
func RemoveDuplicate(arr []uint64) []uint64 {
	var (
		result  = make([]uint64, 0, len(arr))
		tempMap = map[uint64]byte{}
	)
	for _, e := range arr {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// InitDir 初始化创建必要文件夹
func InitDir(dirArr []string) {
	for _, dir := range dirArr {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// GetCurrDir 获取当前目录
func GetCurrDir() string {
	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)
	return dir
}

// HideEmail 匹配邮箱 进行脱敏处理
func HideEmail(str string) string {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			return resString + "@" + res[1]
		} else {
			res2 := substr2(str, 0, 3)
			resString := res2 + "***"
			return resString + "@" + res[1]
		}
	}
	return "***"
}

func substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}

// Generate6RandomNumbers  随机生成一个 6 位数字的验证码
func Generate6RandomNumbers() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}

// GenerateRandomNumbersByLength   按照 位数随机生成一个数字加字母的字符串
func GenerateRandomNumbersByLength(strLen int) string {
	// 使用纳秒作为随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成随机字符串
	var bytes = make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		// 随机生成一个字符
		bytes[i] = byte(97 + rand.Intn(26)) // 97 是 a 的 ASCII 码值
	}
	str := string(bytes)
	return str
}

func RemoveRepeatedElementByUint64(arr []uint64) (newArr []uint64) {
	newArr = make([]uint64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func BuildSignStr(p map[string]string, bizKey string) string {
	keys := make([]string, 0, len(p))
	for k := range p {
		if k == "sign" || k == "sigh" {
			continue
		}
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if p[k] == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')

		buf.WriteString(fmt.Sprint(p[k]))
	}
	buf.WriteString("&key=" + bizKey)
	return buf.String()
}

// SignWithMd5 生成 MD5 签名
func SignWithMd5(p map[string]string, bizKey string) string {
	signStr := BuildSignStr(p, bizKey)
	d := []byte(signStr)
	md5str := fmt.Sprintf("%x", md5.Sum(d))

	return strings.ToUpper(md5str)
}

func GetHeaderLanguage(c *gin.Context) string {
	return c.Request.Header.Get("Accept-Language")
}

func CheckKeyExist(needle string) bool {
	allowedCryptoSlice := []string{"TRON_USDT", "ETH_USDT", "BSC_USDT"}
	set := make(map[string]struct{})
	for _, e := range allowedCryptoSlice {
		set[e] = struct{}{}
	}
	_, ok := set[needle]
	return ok
}

func SnowflakeRandom() string {
	// Create a new Node with a Node number of 1
	node, _ := snowflake.NewNode(1)

	// Generate a snowflake CollectibleItemID.
	id := node.Generate()

	return time.Now().Format("20060102") + id.String()
	//return id.String()
}

// Float2Percentage .... 小数转百分数
func Float2Percentage(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64) + "%"
}

// GetAppPath ...
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func RemoveDuplicateElementByUint64(languages []uint64) []uint64 {
	result := make([]uint64, 0, len(languages))
	temp := map[uint64]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
