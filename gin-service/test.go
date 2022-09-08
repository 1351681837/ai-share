package main

import (
	"fmt"
	"github.com/putyy/ai-share/app/models"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	paramsNum := len(os.Args)
	if paramsNum < 2 {
		fmt.Println("请输入参数")
		return
	}
	table := os.Args[1]

	if "all" == table {
		var allTables []string
		err := models.Db().Raw("show tables;").Scan(&allTables).Error
		checkErr(err)
		for _, table1 := range allTables {
			createModel(table1)
		}
	} else {
		createModel(table)
	}

}

type tableInfoStruct struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func createModel(table string) {
	var tableInfo []tableInfoStruct
	err := models.Db().Raw("show columns from " + table).Scan(&tableInfo).Error
	checkErr(err)
	tableName := strings.Replace(table, "a_", "", -1)
	//驼峰命名
	ttableName := camelString(tableName)
	//创建文件内容
	str := "package models\n\t" +
		"type " + ttableName + " struct {\n\t"
	for _, item := range tableInfo {
		isInt := strings.Index(item.Type, "int")
		if isInt == -1 {
			isTimestamp := strings.Index(item.Type, "timestamp")
			if isTimestamp == -1 {
				str += "  " + camelString(item.Field) + "   string "
			} else {
				str += "  " + camelString(item.Field) + "   JsonTime "
			}

		} else {
			str += "  " + camelString(item.Field) + "   int"
		}
		if item.Key == "PRI" {
			str += " `gorm:\"primaryKey\" json:\"" + item.Field + ",omitempty\"` "
		} else {
			str += "  `json:\"" + item.Field + "\"` "
		}
		str += "\n\t"
	}
	str += " }\n\t"
	str += "func (" + ttableName + ") TableName() string {\n\treturn \"" + table + "\"\n}"
	//创建文件
	dir, _ := filepath.Abs(`.`)
	fileName := strings.Replace(tableName, "_", "-", -1)
	filepaths := dir + "/app/models/" + fileName + ".go"
	outputFile, outputError := os.OpenFile(filepaths, os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(outputError)
	outputFile.WriteString(str)
	//格式化代码
	cmd := exec.Command("gofmt", "-w", filepaths)
	_, err3 := cmd.CombinedOutput()
	checkErr(err3)
}

func camelString(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
