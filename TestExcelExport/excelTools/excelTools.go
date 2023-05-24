package excelTools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

var (
	defaultSheetName   = "Sheet" //默认Sheet名称
	defaultHeight      = 25.0    //默认行高度
	sheetNameIncrement = make(map[string]int, 10)
)

type lkExcelExport struct {
	file       *excelize.File
	SheetNames []string //多工作表名
}

type sheetModel struct {
	sheetName string
	title     []map[string]string
	data      []map[string]interface{}
}

func NewMyExcel() *lkExcelExport {
	var excel lkExcelExport
	excel.file = excelize.NewFile()
	sheetName := defaultSheetName + strconv.Itoa(1)
	excel.SheetNames = append(excel.SheetNames, sheetName)
	index := excel.file.NewSheet(sheetName)
	// 设置工作簿的默认工作表
	excel.file.SetActiveSheet(index)
	return &excel
}

func (l *lkExcelExport) ExportToPath(sheets []sheetModel, path string, fileName string) (string, error) {
	l.exports(sheets)
	name := createFileName(fileName)
	filePath := path + "/" + name
	err := l.file.SaveAs(filePath)
	return filePath, err
}

func (l *lkExcelExport) ExportToWeb(sheets []sheetModel, ctx *gin.Context, fileName string) {
	l.exports(sheets)
	buffer, _ := l.file.WriteToBuffer()
	//设置文件类型
	ctx.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	//设置文件名称
	ctx.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(createFileName(fileName)))
	_, _ = ctx.Writer.Write(buffer.Bytes())
}

func (l *lkExcelExport) exports(sheets []sheetModel) {
	file := l.file
	SheetNames := l.SheetNames
	for i, sheet := range sheets {
		//1.创建数据表
		var sheetName string
		//1.1.如果未设置表名，则生成默认的表名：sheet$n
		var suffix int
		if sheet.sheetName != "" {
			sheetNameIncrement[sheet.sheetName]++
			suffix = sheetNameIncrement[sheet.sheetName]
			sheetName = sheet.sheetName + strconv.Itoa(suffix)
		} else {
			sheetNameIncrement[defaultSheetName]++
			suffix = sheetNameIncrement[defaultSheetName]
			sheetName = defaultSheetName + strconv.Itoa(suffix)
		}
		//1.2.如果是首张表，则复用默认表而不是创建
		if i == 0 {
			file.SetSheetName(l.SheetNames[0], sheetName)
			l.SheetNames[0] = sheetName
		} else {
			file.NewSheet(sheetName)
			SheetNames = append(SheetNames, sheetName)
		}
		//2.设置顶栏
		l.writeTops(sheet.title, sheetName)
		//3.写数据
		l.writeDatas(sheet.title, sheet.data, sheetName)
	}
}

func createFileName(fileName string) string {
	nowTime := time.Now().Format("2006-01-02-15-04-05")
	if fileName == "" {
		fileName = "excle"
	}
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf(fileName+"-%v-%v.xlsx", nowTime, rand.Int63n(time.Now().Unix()))
}

// 设置首行
func (l *lkExcelExport) writeTops(params []map[string]string, sheetName string) {
	topStyle, _ := l.file.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center","vertical":"center"}}`)
	var word = 'A'
	//首行写入
	for _, conf := range params {
		title := conf["title"]
		width, _ := strconv.ParseFloat(conf["width"], 64)
		line := fmt.Sprintf("%c1", word)
		//设置标题
		_ = l.file.SetCellValue(sheetName, line, title)
		//列宽
		_ = l.file.SetColWidth(sheetName, fmt.Sprintf("%c", word), fmt.Sprintf("%c", word), width)
		//设置样式
		_ = l.file.SetCellStyle(sheetName, line, line, topStyle)
		word++
	}
}

// 写入数据
func (l *lkExcelExport) writeDatas(params []map[string]string, data []map[string]interface{}, sheetName string) {
	lineStyle, _ := l.file.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	//lineStyle, _ := l.file.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"fill":{"type":"gradient","color":"FFFFFF","shading":"1"}`)
	//数据写入
	var j = 2 //数据开始行数
	for i, val := range data {
		//设置行高
		_ = l.file.SetRowHeight(sheetName, i+1, defaultHeight)
		//逐列写入
		var word = 'A'
		for _, conf := range params {
			valKey := conf["key"]
			line := fmt.Sprintf("%c%v", word, j)
			isNum := conf["is_num"]

			//设置值
			if isNum != "0" {
				valNum := fmt.Sprintf("'%v", val[valKey])
				_ = l.file.SetCellValue(sheetName, line, valNum)
			} else {
				_ = l.file.SetCellValue(sheetName, line, val[valKey])
			}

			//设置样式
			_ = l.file.SetCellStyle(sheetName, line, line, lineStyle)
			word++
		}
		j++
	}
	//设置行高 尾行
	_ = l.file.SetRowHeight(sheetName, len(data)+1, defaultHeight)
}
