package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func expand_text_lua(filed, text string) string {
	return text
}

func expand_text_json(filed, text string) string {
	return text
}

func main() {
	args := len(os.Args)
	if args < 3 {
		fmt.Println("缺少解析文件名!")
		fmt.Println("usage: " + os.Args[0] + " jl filename")
		return
	}

	output := os.Args[1]
	excel_name := os.Args[2]
	excel, err := xlsx.OpenFile(excel_name)
	if err != nil {
		fmt.Printf("无法解析文件[%s]，是xlsx格式吗？", excel_name)
		fmt.Println(err)
		return
	}

	sheet := excel.Sheets[0]
	row := sheet.Rows[0]

	cols := len(row.Cells)
	rows := len(sheet.Rows)

	name := sheet.Name

	sides := make([]string, cols)
	field := make([]string, cols)
	types := make([]string, cols)

	// sides
	row = sheet.Rows[1]
	for i := 0; i < cols; i++ {
		sides[i] = row.Cells[i].Value
	}

	// field
	row = sheet.Rows[2]
	for i := 0; i < cols; i++ {
		field[i] = row.Cells[i].Value
	}

	// types
	row = sheet.Rows[3]
	for i := 0; i < cols; i++ {
		types[i] = row.Cells[i].Value
	}

	w := &sync.WaitGroup{}
	w.Add(2)

	// export json
	go func() {
		defer func() {
			w.Done()
		}()

		if !strings.Contains(output, "j") {
			return
		}

		file_name := "json/" + name + ".json"
		f, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Println("write file failed:", err)
			return
		}
		defer f.Close()

		f.WriteString("[\n")
		for i := 4; i < rows; i++ {
			row := sheet.Rows[i]
			f.WriteString("\t{\n")
			l := len(row.Cells) - 1
			for k, cell := range row.Cells {
				f.WriteString("\t\t\"" + field[k] + "\" : ")
				text := expand_text_json(field[k], cell.String())
				if types[k] == "string" {
					f.WriteString("\"" + text + "\"")
				} else {
					f.WriteString(text)
				}
				if k == l {
					f.WriteString("\n")
				} else {
					f.WriteString(",\n")
				}
			}

			if i != rows-1 {
				f.WriteString("\t},\n\n")
			} else {
				f.WriteString("\t}\n\n")
			}
		}
		f.WriteString("]\n")

		// valid checking
		data, err := ioutil.ReadFile(file_name)
		if err != nil {
			fmt.Println("ioutil.ReadFile:", err)
		}

		// v := map[string]interface{}{}
		v := []interface{}{}
		err = json.Unmarshal(data, &v)
		if err == nil {
			fmt.Println(file_name + " Checking: OK")
		} else {
			fmt.Println(file_name+" Checking: Invalid & ", err)
		}
	}()

	// export lua
	go func() {
		defer func() {
			w.Done()
		}()

		if !strings.Contains(output, "l") {
			return
		}

		file_name := "lua/" + name + ".lua"
		f, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Println("write file failed:", err)
			return
		}
		defer f.Close()

		f.WriteString("local data = \n{\n")

		idx := 1
		for i := 4; i < rows; i++ {
			row := sheet.Rows[i]
			key := fmt.Sprintf("\t[%d] = {\n", idx)
			idx++
			f.WriteString(key)

			for k, cell := range row.Cells {
				f.WriteString("\t\t" + field[k] + " = ")
				text := expand_text_json(field[k], cell.String())
				if types[k] == "string" {
					f.WriteString("\"" + text + "\"")
				} else {
					f.WriteString(text)
				}
				f.WriteString(",\n")
			}

			f.WriteString("\t},\n\n")
		}
		f.WriteString("}\n\n")
		f.WriteString("return data\n")
	}()

	w.Wait()
}
