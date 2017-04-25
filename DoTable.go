package main

import (
	"fmt"
    "github.com/jbussdieker/golibxml"
)



func DoTable(fname string,outpath *string) {
	excelFileName := fname

    xsdSchema, err := xsd.ParseSchema([]byte(XSD))
	if err != nil {
		fmt.Println(err)
		return
	}

	doc := golibxml.ParseDoc(XML)
	if doc == nil {
		// TODO capture and display error - help please
		fmt.Println("Error parsing document")
		return
	}
	defer doc.Free()

	// golibxml._Ctype_xmlDocPtr can't be cast to xsd.DocPtr, even though they are both
	// essentially _Ctype_xmlDocPtr.  Using unsafe gets around this.
	if err := xsdSchema.Validate(xsd.DocPtr(unsafe.Pointer(doc.Ptr))); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("XML Valid as per XSD")

	// xlFile, error := xlsx.OpenFile(excelFileName)
	// if error != nil {
	// 	fmt.Println(excelFileName+":", error)
	// }
	// fmt.Printf("开始处理excel: %s\n", excelFileName)
	// for _, sheet := range xlFile.Sheets {
	// 	fmt.Printf("文件: %s\n", sheet.Name)
	// 	Label1, _ := sheet.Cell(0, 0).String()
	// 	Label2, _ := sheet.Cell(1, 0).String()
	// 	Label3, _ := sheet.Cell(2, 0).String()
	// 	//复合标准头才处理
	// 	if Label1 == "文件名：" && Label2 == "类  名：" && Label3 == "导出类型:" {
	// 		fmt.Printf("开始处理: %s\n", sheet.Name)
	// 		outname, _ := sheet.Cell(0, 1).String()
	// 		//outclass, _ := sheet.Cell(1, 1).String()
	// 		outtype, _ := sheet.Cell(2, 1).String()
	// 		switch outtype {
	// 		case "1":
	// 			//data := []interface{}
	// 			data := Data2Array(sheet)
	// 			bytes, err := json.Marshal(data)
	// 			if err == nil {
	// 				os.MkdirAll(*outpath+"/", 0666)
	// 				ioutil.WriteFile(*outpath+"/"+outname+".json", bytes, 0666)
	// 			}
	// 		case "2":
	// 			//data := map[string]interface{}
	// 			data := Data2Map(sheet)
	// 			bytes, err := json.Marshal(data)
	// 			if err == nil {
	// 				ioutil.WriteFile(*outpath+"/"+outname+".json", bytes, 0666)
	// 			}
	// 		default:

	// 		}
	// 	}
	// }
}