package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)
// 详细文档 https://xuri.me/excelize/zh-hans/
// 详细的请看以上地址。
// 注意添加图片需要添加image包
import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	createXlsxFile()
	//readXlsxFile()
	//addPictureInXlsx()
	//readAndWriteDemo()
}
// 创建一个简单的xlsx文件
func createXlsxFile () {
	// 创建一个文件
	xlsx := excelize.NewFile()
	// 创建sheet
	Sheet1 := xlsx.NewSheet("Sheet1")
	// 设置单元格的数据
	xlsx.SetCellValue("Sheet1", "A1", "title1")
	xlsx.SetCellValue("Sheet1", "B1", "title2")
	xlsx.SetCellValue("Sheet1", "C1", "title3")
	// 连续行赋值,列D1开始向后赋值
	columns := []interface{}{"title4","title5","title6","title7","title8"}
	xlsx.SetSheetRow("Sheet1", "D1", &columns)

	// 设置列宽 eg: B到C列宽度为20
	xlsx.SetColWidth("Sheet1","B","C",20)
	// 设置行高 eg: 第一行高度30
	xlsx.SetRowHeight("Sheet1", 1, 30)
	// 在某列前插入一列空白列
	//xlsx.InsertCol("Sheet1", "B")
	// 在某行前插入一列空白行
	//xlsx.InsertRow("Sheet1", 3)
	//合并单元格, 对角的两个单元格
	//xlsx.MergeCell("Sheet1", "B3","C4")
	// 设置工作簿的默认工作表
	xlsx.SetActiveSheet(Sheet1)
	// 生成文件并命名
	if err := xlsx.SaveAs("test.xlsx"); err != nil {
		fmt.Println("create xlsx file error is " + err.Error())
	}
}

// 读取excel 文件
func readXlsxFile () {
	// 打开文件
	xlsx, err := excelize.OpenFile("./test.xlsx")
	if err != nil {
		fmt.Println("read xlsx error is " + err.Error())
		return
	}
	// 获取工作表中指定单元格的值
	cell, err := xlsx.GetCellValue("Sheet1", "A1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func addPictureInXlsx () {
	//创建文件
	xlsx := excelize.NewFile()
	//创建sheet1
	xlsx.SetCellValue("Sheet1", "A1", "title1")
	xlsx.SetActiveSheet(xlsx.NewSheet("Sheet1"))
	//保存文件pic.xls，注意如果存在则会覆盖。
	if err := xlsx.SaveAs("pic.xls"); err != nil {
		fmt.Println("create pic xls error is " + err.Error())
	}
	//读取文件
	//插入图片
	readXlsx, err := excelize.OpenFile("./pic.xls")
	if err != nil {
		fmt.Println("read file error is " + err.Error())
	}
	fmt.Println(readXlsx.GetCellValue("Sheet1", "A1"))
	if err := readXlsx.AddPicture("Sheet1", "A2","./img.png",`{"x_scale": 1, "y_scale": 1}`); err != nil {
		fmt.Println("xlsx add picture error is " + err.Error())
	}
	if err := readXlsx.Save(); err != nil {
		fmt.Println("save pic file error is " + err.Error())
	}
}

func readAndWriteDemo () {
	//读取文件
	xlsx, err := excelize.OpenFile("./test.xlsx")
	if err != nil {
		fmt.Println("read File error is " + err.Error())
	}
	// 选择操作sheet
	xlsx.SetActiveSheet(xlsx.GetSheetIndex("Sheet1"))

	// 选择读取的列
	readColTitle := []string{"title1", "title3"}
	//遍历所有列
	rows, err :=xlsx.Cols("Sheet1")
	if err != nil {
		fmt.Println("read rows error" + err.Error())
	}
	cols, err :=xlsx.Rows("Sheet1")
	rowscopy := *cols
	for rows.Next() {
		// 获取列值
		row,_ := rows.Rows()
		// 遍历列的值
		for _,valid:=range row{
			// 遍历选择的列的值
			for _,title := range readColTitle{
				// 匹配就显示
				if valid == title {
					fmt.Println(row)
				}
			}
		}

	}
	// 使用获取到的值根据模板写入数据,模板使用testcopy.xlsx
	doWrite(rowscopy)
}

func doWrite (rows excelize.Rows) {
	// 读取文件
	writeXlsx,_ := excelize.OpenFile("./testTemplate.xlsx")
	// 设置默认操作sheet
	writeXlsx.SetActiveSheet(writeXlsx.GetSheetIndex("Sheet1"))
	// 获取第一行的title
	templateTitle,_ := writeXlsx.GetRows("Sheet1")
	fmt.Println(templateTitle[0])
	// 使用第一行作为map.key形成[]map[string]interface{}{} 用来对应template上的title
	newRow := []map[string]interface{}{}
	// 保存对应后的切片
	newSlice := []interface{}{}
	// 获取的第一行row作为title
	title := []string{}
	// index 目的是区分第一行和其他行
	i := 0
	// 遍历rows
	for rows.Next() {
		i++
		// 获取列
		row,_ := rows.Columns()
		// 忽略第一列
		if i > 1 {
			// index 目的是切割切片
			index := 0
			// 使用第一行作为map.key形成[]map[string]interface{}{} 用来对应template上的title
			for _,x:= range row {
				tt := title[index: index + 1]
				index++
				newRow = append(newRow, map[string]interface{}{tt[0]: x})
			}
			fmt.Println(newRow)
			// 对应title得到新的写入切片
			for _,vv := range templateTitle[0]{
				// 当 len(newSlice) 相等的时候说明这个title这一行的值是空的
				a := len(newSlice)
				for _,bb:=range newRow{
					if bb[vv] !=nil {
						newSlice = append(newSlice, bb[vv])
					}
				}
				b := len(newSlice)
				// 空的就写入空字符串“”
				if a == b  {
					newSlice = append(newSlice, "")
				}
			}
			fmt.Println(newSlice)
			// 按行写入
			writeXlsx.SetSheetRow("Sheet1", "A" + strconv.Itoa(i), &newSlice)
		} else {
			title = row
		}
	}
	writeXlsx.SaveAs("testCopy.xlsx")
}