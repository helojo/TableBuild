package main

import(
    "flag"
    "fmt"
    "path/filepath"
    "os"
	"strings"
	"helojo.com/libs"
)

// var (
// 	outpath = flag.String("outpath", "./out", "outpath json path default = ./out")
// )

func walkFunc(path string, info os.FileInfo, err error) error {
	//fmt.Printf("%s\n", path)
	fext := filepath.Ext(path)
	//排除打开文件
    // fmt.Print("\n"+fext)
	if strings.Index(path, "~$") != -1 {
		return nil
	}
	if fext == ".table" {
		libs.DoTable(path,outpath)
	}
	return nil
}

func main()  {
	outpath := flag.String("out","./out","outpath json path default = ./out")
	inpath := flag.String("in","./in","inpath xls path default = ./in")
    flag.Parse()

	fmt.Println(*inpath)
	fmt.Println(*outpath)

	exist,_ := libs.PathExists(*inpath)

	if exist == false{
		fmt.Println(libs.Colorize("打表目录不存在,请检查是否存在打表目录："+*inpath,libs.Error))
		return
	}

    // fmt.Println(*outpath)
    filepath.Walk(*inpath,walkFunc)
}