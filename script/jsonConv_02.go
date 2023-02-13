package script

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"os"
)

type OriginalJson struct {
	Name        string       `json:"name"`
	Image       string       `json:"image"`
	Description string       `json:"description"`
	ID          string       `json:"id"`
	Attributes  []Attributes `json:"attributes"`
}
type Attributes struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TargetJson struct {
	Name        string       `json:"name"`
	Image       string       `json:"image"`
	Description string       `json:"description"`
	Attributes  []Attributes `json:"attributes"`
}

// 主程序 ...
func JsonConv02() {
	var (
		originDir = "./docs/1000js文件/"
		newDir    = "./docs/1000js文件(new-11.19)/"
	)

	//
	// 新文件存储目录
	os.MkdirAll(newDir, 0777)

	// 循环读取文件、转换、保存新文件
	for i := 1; i <= 1000; i++ {
		// 1、读取原文件
		info := ReadOriginFile(fmt.Sprintf("%sLego%d.json", originDir, i))

		// 2、执行转换逻辑
		newInfo := transform(info, i)

		// 3、保存新数据
		SaveNewFile(fmt.Sprintf("%sLego%d.json", newDir, i), newInfo)

	}

	fmt.Println("MISSION SUCCESS !!!")
}

// ReadOriginFile ...
func ReadOriginFile(filePath string) (info OriginalJson) {

	// read file
	DataJsonStruct := ReadFile(filePath)

	// json ——> struct
	err := json.Unmarshal([]byte(DataJsonStruct), &info)
	if err != nil {
		fmt.Println("error:", err)
	}
	return
}

// ReadFile 读取测试数据的文件内容
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, _ := io.ReadAll(fi)
	return string(fd)
}

// SaveNewFile ...
func SaveNewFile(newFilePath string, info TargetJson) {
	// marshal

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(info)

	// res, _ := json.Marshal(info)

	// new file
	file, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()

	// json format
	var out bytes.Buffer
	_ = json.Indent(&out, bf.Bytes(), "", "\t")

	// write
	out.WriteTo(file)
}

// transform ...
func transform(oldInfo OriginalJson, i int) (newInfo TargetJson) {
	// 不变之赋值
	newInfo = TargetJson{
		Name:        fmt.Sprintf("Treasures FIFA Star Card #%d", i),
		Image:       fmt.Sprintf("https://ipfs.newland.club/ipfs/QmYStTjNrqU9TvixmNacf2nWXHVyMVX89DptEVEkvAWnRY/Lego%d.avif", i),
		Description: "Treasures FIFA Star Card",
		Attributes:  oldInfo.Attributes,
	}
	return
}
