package tool

import (
	"strconv"

	"github.com/tealeg/xlsx"
)

// 单元格是否为空
func Iscellempty(sheet *xlsx.Sheet, row int, col int) bool {
	cell := sheet.Cell(row-1, col-1)
	return cell.Value == ""
}

// 获取指定单元格的文本内容，如果是数字格式，则仍然转为文本
func Getcellstr(sheet *xlsx.Sheet, row int, col int) string {
	cell := sheet.Cell(row-1, col-1)
	str, _ := cell.String()
	return str
}

// 获取表格的数字
func Getcellint(sheet *xlsx.Sheet, row int, col int) int {
	cell := sheet.Cell(row-1, col-1)
	cell.String()
	str, _ := cell.String()
	i, err := strconv.Atoi(str)
	if err != nil {
		// fmt.Println("字符串转换成整数失败 = ", cell.String(), err)
		i = 0
	}

	return i
}

// 获取工作表的行数
func Getrowcnt(sheet *xlsx.Sheet) int {
	return sheet.MaxRow
}

// 获取工作表的列数
func Getcolcnt(sheet *xlsx.Sheet) int {
	return sheet.MaxCol
}
