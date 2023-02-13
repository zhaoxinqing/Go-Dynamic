#

```go
package script

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// JsonInfo ...
type JsonInfo struct {
    Name           string      `json:"name"`
    Image          string      `json:"image"`
    Attributes     []Attribute `json:"attributes"`
    CollectionName string      `json:"collectionName"`
    Description    string      `json:"description"`
    Compiler       string      `json:"compiler"`
}

// Attribute ...
type Attribute struct {
    Trait_type string `json:"trait_type"`
    Value      string `json:"value"`
}

type TargetJsonInfo struct {
    Name                 string      `json:"name"`
    Description          string      `json:"description"`
    ExternalURL          string      `json:"external_url"`
    Image                string      `json:"image"`
    Attributes           []Attribute `json:"attributes"`
    Properties           Properties  `json:"properties"`
    AnimationURL         string      `json:"animation_url"`
    SellerFeeBasisPoints int         `json:"seller_fee_basis_points"`
    Collection           Collection  `json:"collection"`
    Symbol               string      `json:"symbol"`
}

type Properties struct {
    Files    []File   `json:"files"`
    Creators []string `json:"creators"`
}

type File struct {
    URI  string `json:"uri"`
    Type string `json:"type"`
}

type Collection struct {
    Name   string `json:"name"`
    Family string `json:"family"`
}

// 主程序 ...
func JsonConv() {
    var (
        dirList   = []string{"1-C罗", "2-富安健洋", "3-凯恩", "4-莱万", "5-梅西", "6-门迪", "7-莫德里奇", "8-姆巴佩", "9-内马尔", "10-苏亚雷斯", "11-孙兴愍"}
        originDir = "./docs/CryptoSportMeta-Json/"
        newDir    = "./docs/CryptoSportMeta-Json(11.14)/"
    )

    for k, v := range dirList {
    //
    var (
        theOriginDir = originDir + v + "/"
        theNewDir    = newDir + v + "/"
        fromIndex    = k*700 + 1
        endIndex     = k*700 + 700
    )
    imagePrefix := getImagePrefix(v)
    // 新文件存储目录
    os.MkdirAll(theNewDir, 0777)

    // 循环读取文件、转换、保存新文件
    for i := fromIndex; i <= endIndex; i++ {

        // 1、读取原文件
        info := ReadOriginFile(fmt.Sprintf("%s%d.json", theOriginDir, i))

        // 2、执行转换逻辑
        newInfo := transform(info, i, imagePrefix)

        // 3、保存新数据
        SaveNewFile(fmt.Sprintf("%s%d.json", theNewDir, i), newInfo)

        }
    }

    fmt.Println("MISSION SUCCESS !!!")
}

// getImagePrefix
func getImagePrefix(dir string) (imagePrefix string) {
    switch dir {
        
        case "1-C罗":
            imagePrefix = "ipfs://bafybeidv4retbtt2m6o7fnshskh2azb72mi3jb7y2oqjfbs5oy2hkrcz7q"

        
        case "2-富安健洋":
            imagePrefix = "ipfs://bafybeibfvzgwc73chfby3hrzwpu3wpt4fzctyhoi3iegxbrf4766p2mari"
            
        
        case "3-凯恩":
            imagePrefix = "ipfs://bafybeigeq5e54p5ubiypqiyqlnunfgpxzbdadcwyn3p2mkmxaevotxo2oq"

        
        case "4-莱万":
            imagePrefix = "ipfs://bafybeihi72grk4lzjwis62kjinso555673btcdnsgsnd7ud2mzes6tlo2e"

        
        case "5-梅西":
            imagePrefix = "ipfs://bafybeifvxxd7x7774wl6mtlxiqltizpsm6q2tm3tfgin7fk7wa4ql67nki"

        
        case "6-门迪":
            imagePrefix = "ipfs://bafybeihvnslx4xwicr2lkbpiyyis2hp5ur3m2oerqdelfw52jrt3gtzzmu"

        
        case "7-莫德里奇":
            imagePrefix = "ipfs://bafybeif5cek6kl5ekw7ck7r6ahotnmvtpnorkucwlwu7u776cntr3abahy"

        
        case "8-姆巴佩":
            imagePrefix = "ipfs://bafybeib2ans6yefurabteysiwjnshaj6zhy2esmqhvbx5fsg77hfvxy2ma"

        
        case "9-内马尔":
            imagePrefix = "ipfs://bafybeiguyabbrt5g5wrs7ldmcmiwbr6l7hpxquakbrvoovohkeva33en2e"

        
        case "10-苏亚雷斯":
            imagePrefix = "ipfs://bafybeih4wqab5wa26heilys7gufyfgmwoxor5b7ptkddgrtpbagjv45oii"

        
        case "11-孙兴愍":
            imagePrefix = "ipfs://bafybeid5arfkkoexjwdgatpy6pvq7kg5ho2oljiwnbnc27vnwzm6swt6gm"
    }
    return
}

// ReadOriginFile ...
func ReadOriginFile(filePath string) (info JsonInfo) {

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
    fd, err := ioutil.ReadAll(fi)
    return string(fd)
}

// SaveNewFile ...
func SaveNewFile(newFilePath string, info TargetJsonInfo) {
    // marshal
    res, _ := json.Marshal(info)

    // new file
    file, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("文件打开失败", err)
    }
    defer file.Close()

    // json format
    var out bytes.Buffer
    _ = json.Indent(&out, res, "", "\t")

    // write
    out.WriteTo(file)
}

// // collectionName
// switch oldInfo.CollectionName {
// case "C罗":
//
// newInfo.CollectionName = "Portugal-7"
//
// case "富安健洋":
//
// newInfo.CollectionName = "Japan-16"
//
// case "凯恩":
//
// newInfo.CollectionName = "England-9"
//
// case "莱万":
//
// newInfo.CollectionName = "Poland-9"
//
// case "梅西":
//
// newInfo.CollectionName = "Argentina-10"
//
// case "门迪":
//
// newInfo.CollectionName = "Senegal-16"
//
// case "莫德里奇":
//
// newInfo.CollectionName = "Croatia-10"
//
// case "姆巴佩":
//
// newInfo.CollectionName = "France-7"
//
// case "内马尔":
//
// newInfo.CollectionName = "Brazil-10"
//
// case "苏亚雷斯":
//
// newInfo.CollectionName = "Uruguay-9"
//
// case "孙兴愍":
//
//  newInfo.CollectionName = "Korea-7"
// }
//

// transform ...
func transform(oldInfo JsonInfo, i int, imagePrefix string) (newInfo TargetJsonInfo) {
    // 不变之赋值
    newInfo.Name = fmt.Sprintf("Ordinary#%d", i)
    newInfo.Description = "Ordinary"
    newInfo.ExternalURL = ""
    newInfo.Image = fmt.Sprintf("%s/%d.png", imagePrefix, i)
    var properties Properties

    var files []File
    var file File
    file.URI = fmt.Sprintf("%d.png", i)
    file.Type = "image/png"
    files = append(files, file)

    properties.Files = files
    properties.Creators = []string{}
    newInfo.Properties = properties

    newInfo.AnimationURL = ""
    newInfo.SellerFeeBasisPoints = 0

    var collection Collection
    collection.Name = "Ordinary NFT"
    collection.Family = "Ordinary NFT"
    newInfo.Collection = collection

    newInfo.Symbol = ""

    // Attributes
    var newAttributes []Attribute
    for _, v := range oldInfo.Attributes {
    var newAttribute Attribute
    switch v.Trait_type {
        
        case "背景":
            newAttribute.Trait_type = "Backgroud"
            newAttribute.Value = transformBackgroud(v.Value) // 背景
            newAttributes = append(newAttributes, newAttribute)
        
        case "耳饰":
            newAttribute.Trait_type = "Earring"
            newAttribute.Value = transformEarring(v.Value) // 耳饰
            newAttributes = append(newAttributes, newAttribute)
        
        case "头饰":
            newAttribute.Trait_type = "Hair Accessories"
            newAttribute.Value = transformHairAccessories(v.Value) // 头饰
            newAttributes = append(newAttributes, newAttribute)
        
        case "眼睛":
            newAttribute.Trait_type = "Eye"
            newAttribute.Value = transformEye(v.Value) // 眼睛
            newAttributes = append(newAttributes, newAttribute)
        
        case "衣服":
            newAttribute.Trait_type = "Costume"
            newAttribute.Value = transformCostumeh(v.Value) // 衣服
            newAttributes = append(newAttributes, newAttribute)
        
        case "嘴巴":
            newAttribute.Trait_type = "mouth"
            newAttribute.Value = transformMouth(v.Value) // 嘴巴
            newAttributes = append(newAttributes, newAttribute)
        default: // Nude
            newAttribute.Trait_type = "Nude"
            newAttribute.Value = transformNude(v.Trait_type) // 人体
            newAttributes = append(newAttributes, newAttribute)
        }
    }
    newInfo.Attributes = newAttributes
    return
}

// transformNude ...
func transformNude(in string) (out string) {
    switch in {
        
        case "C落":
            out = "Portugal-7"
        
        case "富安-裸 3":
            out = "Japan-16"
        
        case "凯恩-裸 5":
            out = "England-9"
        
        case "莱万":
            out = "Poland-9"
        
        case "梅西":
            out = "Argentina-10"
        
        case "门迪-裸 4":
            out = "Senegal-16"
        
        case "莫德里奇-裸 4":
            out = "Croatia-10"
        
        case "姆巴佩-裸 5":
            out = "France-7"
        
        case "内马尔":
            out = "Portugal-7"
        
        case "苏亚雷斯":
            out = "Uruguay-9"
        
        case "孙兴慜-裸 4":
            out = "Korea-7"
    }
    return
}

// transformBackgroud ...
func transformBackgroud(in string) (out string) {
    switch in {
        
        case "背景9":
            out = "Dark Black"
        
        case "背景1":
            out = "Olive Green"
        
        case "背景6":
            out = "Dingy Yellow"
        
        case "背景7":
            out = "Teal"
        
        case "背景2":
            out = "Bright Orange"
        
        case "背景4":
            out = "Light Grey"
        
        case "背景3":
            out = "Dark Blue"
        
        case "背景8":
            out = "Grey Green"
        
        case "背景5":
            out = "Sky Blue"
    }
    return
}

// transformEarring ...
func transformEarring(in string) (out string) {
    switch in {
        
        case "白耳钉":
            out = "White Stud"
        
        case "红耳钉":
            out = "Red Stud"
        
        case "黄耳钉":
            out = "Yellow Stud"
        
        case "耳环":
            out = "Gold Earrings"
        
        case "蓝耳钉":
            out = "Blue Stud"
        
        case "空耳饰":
            out = "Blank"
        
        case "绿耳钉":
            out = "Green Stud"
    }
    return
}

// transformHairAccessories ...
func transformHairAccessories(in string) (out string) {
    switch in {
        
        case "分耳机":
            out = "Pink Headset"
        
        case "光环":
            out = "Halo"
        
        case "黑耳机":
            out = "Black Headset"
        
        case "红耳机":
            out = "Red Headset"
        
        case "红头巾":
            out = "Red Kerchief"
        
        case "皇冠":
            out = "Crown"
        
        case "黄头巾":
            out = "Yellow Kerchief"
        
        case "蓝耳机":
            out = "Blue Headset"
        
        case "蓝头巾":
            out = "Blue Kerchief"
        
        case "绿耳机":
            out = "Green Headset"
        
        case "绿头巾":
            out = "Green Kerchief"
        
        case "紫头巾":
            out = "Purple Kerchief"
        
        case "空头饰":
            out = "Blank"
    }
    return
}

// transformEye ...
func transformEye(in string) (out string) {
    switch in {
        
        case "墨镜":
            out = "Sunglasses"
        
        case "眼罩":
            out = "Eye Patch"
        
        case "空眼镜":
            out = "Blank"
    }
    return
}

// transformMouth ...
func transformMouth(in string) (out string) {
    switch in {
        
        case "白牙":
            out = "White Teeth"
        
        case "粉泡":
            out = "Pink Bubble"
        
        case "黄泡":
            out = "Yellow Bubble"
        
        case "黄牙":
            out = "Yellow Teeth"
        
        case "口罩":
            out = "Mask"
        
        case "蓝牙":
            out = "Teal Teeth"
        
        case "蓝泡":
            out = "Blue Bubble"
        
        case "绿泡":
            out = "Green Bubble"
        
        case "绿牙":
            out = "Green Teeth"
        
        case "烟":
            out = "Cigarette"
        
        case "粉牙":
            out = "Purple Teeth"
        
        case "烟斗":
            out = "Pipe"
        
        case "空嘴巴":
            out = "Blank"
    }
    return
}

// transformCostumeh ...
func transformCostumeh(in string) (out string) {
    switch in {
        
        case "礼服1":
            out = "White Tuxedo"
        
        case "上衣2":
            out = "Pink Top"
        
        case "礼服6":
            out = "Black Tuxedo"
        
        case "上衣6":
            out = "Black Top"
        
        case "背心1":
            out = "Red Vest"
        
        case "礼服3":
            out = "Red Tuxedo"
        
        case "连帽衫3":
            out = "Red Hoodie"
        
        case "上衣1":
            out = "Red Top"
        
        case "背心3":
            out = "Yellow Vest"
        
        case "礼服5":
            out = "Yellow Tuxedo"
        
        case "连帽衫4":
            out = "Yellow Hoodie"
        
        case "上衣5":
            out = "Yellow Top"
        
        case "背心2":
            out = "Blue Vest"
        
        case "礼服4":
            out = "Blue Tuxedo"
        
        case "连帽衫5":
            out = "Blue Hoodie"
        
        case "上衣4":
            out = "Blue Top"
        
        case "背心4":
            out = "Green Vest"
        
        case "连帽衫2":
            out = "Green Hoodie"
        
        case "上衣3":
            out = "Green Top"
        
        case "背心5":
            out = "Purple Vest"
        
        case "礼服2":
            out = "Purple Tuxedo"
        
        case "连帽衫1":
            out = "Purple Hoodie"
        
        case "空衣服":
            out = "Blank"
    }
    return
}

// {
//     "name": "4",
//     "image": "4.png",
//     "attributes": [
//         {
//             "trait_type": "背景",
//             "value": "背景7"
//         },
//         {
//             "trait_type": "孙兴慜-裸 4",
//             "value": "孙兴慜-裸 3"
//         },
//         {
//             "trait_type": "嘴巴",
//             "value": "空嘴巴"
//         },
//         {
//             "trait_type": "眼睛",
//             "value": "眼罩"
//         },
//         {
//             "trait_type": "耳饰",
//             "value": "白耳钉"
//         },
//         {
//             "trait_type": "头饰",
//             "value": "红头巾"
//         },
//         {
//             "trait_type": "衣服",
//             "value": "背心2"
//         }
//     ],
//     "collectionName": "孙兴愍",
//     "description": "",
//     "compiler": "NFT Generator for Figma, https://www.figma.com/community/plugin/1052280914006102970"
// }

```
