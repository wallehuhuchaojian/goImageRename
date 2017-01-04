package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func main() {
	fileExePath := readPath()
	partenPaths := getPatrenPath(fileExePath)
	fmt.Println("exepath=====", fileExePath)
	fmt.Println("partenPaths=====", partenPaths)

	parseFile(partenPaths)

}

func parseFile(partenPaths string) {

	f, err := os.Open(partenPaths) //打开一个目录
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	ff, _ := f.Readdir(0) //设置读取的数量 <=0是读取所有的文件 返回的[]fileinfo
	for i, fi := range ff {
		if fi.Name() == "myImageRemaneConfig.json" {
			fmt.Printf("filename %d: %+v\n", i, fi.Name())
			readConfig(fi.Name())
			fmt.Println(fi.IsDir())
			fmt.Println(fi.Sys())
			break
		}
		//我们输出文件的名称
	}
}

func readConfig(path string) {

	// lens, ok := ioutil.ReadFile(path)

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
