package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	//"reflect"
	//"syscall"
	//"reflect"
	"io/ioutil"
	"encoding/json"
)
var testPath string

type ImageInfo struct {
	OriginalName string `json:fromName`
	CovertName string `json:toName`

}

type Config struct {
	Name string `json:packageName`
	Path string`json:packagePath`
	Images []ImageInfo`json:images`

}

//func getConfigData() (data string)  {
//	var config Config
//	config.Name="configName"
//	config.Path="\\home\\hu\\desktop\\config.json"
//	config.Images=append(config.Images,ImageInfo{
//		OriginalName:"1.png",
//		CovertName:"convertedName1.png"})
//	config.Images=append(config.Images,ImageInfo{
//		OriginalName:"1.png",
//		CovertName:"convertedName1.png"})
//	byteStr,_:=json.Marshal(config);
//	fmt.Printf("json.m,%s\n",byteStr)
//
//	var cg Config
//	err:=json.Unmarshal(byteStr,&cg);
//	if err!=nil {
//		fmt.Println(err)
//	}
//	fmt.Println(cg)
//	return ""
//
//
//}
func main() {
	//getConfigData()
	fileExePath := readPath()
	//partenPaths := getPatrenPath(fileExePath)
	fmt.Println("exepath=====", fileExePath)
	testPath ="/home/hu/work/gocode/bin"

	config:=parseFile(testPath)
	fmt.Println(config)




}

func parseFile(partenPaths string) (Config){
	var config Config
	f, err := os.Open(partenPaths) //打开一个目录
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("parseFile path>>>>>>",partenPaths)
	defer f.Close()
	ff, _ := f.Readdir(0) //设置读取的数量 <=0是读取所有的文件 返回的[]fileinfo
	for i, fi := range ff {
		if fi.Name() == "myImageRemaneConfig.json" {
			fmt.Printf("filename %d: %+v\n", i, fi.Name())

			configPaht :=partenPaths+"/"+fi.Name()
			fmt.Println("file name>>>>>>>>>>",configPaht)
			//readConfig(configPaht)
			config,er:=readConfig(configPaht)
			if er!=nil {
				fmt.Println("readConfig err",er.Error())
			}
			return config

			break
		}
	}
	return  config
}
func readConfig(path string) (config Config,err error){

	//读取配置文件
	fmt.Println("读取配置文件")

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile error: ", err.Error())
		return config,err
	}
	if err := json.Unmarshal(bytes, &config); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return config,err
	}

	fmt.Println(config)
	return config,nil
}
func getPatrenPath(dir string) string {
	if path.IsAbs(dir) {

		return path.Dir(dir)
	}
	return ""
}
func readPath() string {
	file, _ := exec.LookPath(os.Args[0])
	dir, _ := filepath.Abs(file)
	println(dir)
	return dir
}