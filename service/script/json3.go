package script

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Json3Info struct {
	Name                 string     `json:"name"`
	Description          string     `json:"description"`
	Image                string     `json:"image"`
	SellerFeeBasisPoints int        `json:"seller_fee_basis_points"`
	Collection           Collection `json:"collection"`
	Symbol               string     `json:"symbol"`
}

type Collection struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}

func JsonConv3() {
	// read file
	fi, err := os.Open("docs/11-24/1.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, _ := io.ReadAll(fi)

	os.MkdirAll("docs/11-24(new)", 0777)
	// json ——> struct
	var info Json3Info
	err = json.Unmarshal(fd, &info)
	if err != nil {
		fmt.Println("error:", err)
	}
	for i := 1; i <= 5000; i++ {
		newJson := Json3Info{
			Name:                 fmt.Sprintf("#%d", i),
			Description:          info.Description,
			Image:                fmt.Sprintf("ipfs://bafybeih4d54lyigpepl3rl22sqqr5237q4dxtrzs54mndczr5qpre566pq/%d.png", i),
			SellerFeeBasisPoints: 0,
			Collection: Collection{
				Name:   info.Collection.Name,
				Family: info.Collection.Family,
			},
			Symbol: info.Symbol,
		}
		newFilePath := fmt.Sprintf("docs/11-24(new)/%d.json", i)
		SaveNewFile3(newFilePath, newJson)
	}
	fmt.Println("MISSION SUCCESS !!!")
}

// SaveNewFile ...
func SaveNewFile3(newFilePath string, info Json3Info) {
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
