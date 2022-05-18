package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"
)

func main() {
	path := `./mijie3`
	GetMaxInfo(path)
}

func GetMaxInfo(path string) {
	//path := `.\info\1239875_low.txt`
	//fmt.Println(os.Getwd())
	//fmt.Println(filepath.Abs(path))

	//获取贴图
	//读取文件，返回整个内容
	//正则获取取想要的内容
	contentsByte, _ := ioutil.ReadFile(path)
	contentString := string(contentsByte)

	regMap := regexp.MustCompile(`ped_object_ids: [[0-9]+`)
	resultMap := regMap.FindAllStringSubmatch(contentString, -1)
	//maps := H_GetListFromString(resultMap)
	//fmt.Println(resultMap)

	sliceRes := make([]string, 0)
	for _, v := range resultMap {
		regMap1 := regexp.MustCompile(`[0-9]+`)
		result := regMap1.FindAllStringSubmatch(v[0], -1)
		fmt.Println(result[0][0])
		if result[0][0] != "" {
			sliceRes = append(sliceRes, result[0][0])
		}
	}

	yiyuanb, _ := json.Marshal(sliceRes)
	fmt.Println(string(yiyuanb))

	for _, v := range sliceRes {
		convert(v)
	}

	//regMat := regexp.MustCompile("(?s:Materials(.*))")
	//resultMat := regMat.FindAllStringSubmatch(contentString,-1)[0][1]
	//mats := H_GetListFromString(resultMat)
	//fmt.Println(mats)

	//打开文件
	//声明bufio.Reader
	//读取文件

	//rw, err := os.Open(path)
	//if err != nil {
	//	panic(err)
	//}
	//defer rw.Close()

	//rb := bufio.NewReader(rw)
	//getMaxInfoFromContents(rb)
}

func convert(idd string) {

	id, _ := strconv.Atoi(idd)
	//id:=942868822014558  //objectid
	//626208950289641
	//626208924486353

	Sequence := int32(id & 0x3ff)

	CameraIdx := int32((id >> 42) & 0x7f)
	RegionId := int32((id >> 49) & 0x7fffffff)

	timestamps := int64((id >> 10) & 0xffffffff)
	tm := time.Unix(timestamps, 0)

	fmt.Println("CameraIdx:", CameraIdx)

	fmt.Println("RegionId:", RegionId)

	fmt.Println("Sequence:", Sequence)

	fmt.Println("time:", tm)
	fmt.Println("timestamp:", timestamps)

	fmt.Println("====================:")

}
