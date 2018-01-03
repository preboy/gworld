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

// field := "Rewards#ItemID_ItemCt_Type"
// text := "2001_9_1|2007_3_2|2009_3_6"
func expand_json_map(field, text string) (string, string) {
	text = strings.Trim(text, " ")

	idx := strings.Index(field, "#")
	if idx == -1 {
		return field, text
	}

	field_split := strings.Split(field, "#")
	key := field_split[0]
	sub_field := field_split[1]

	if len(text) == 0 {
		return key, "[]"
	}

	str := "[\n"

	fields := strings.Split(sub_field, "_")

	values := strings.Split(text, "|")
	for i, value := range values {
		sub_values := strings.Split(value, "_")
		if len(sub_values) != len(fields) {
			panic("expand_json_map error:" + field + " & " + text)
		}
		s := "\t\t\t{\n"
		for k, sub_value := range sub_values {
			s = s + "\t\t\t\t\"" + fields[k] + "\" : " + sub_value
			if k != len(sub_values)-1 {
				s = s + ",\n"
			} else {
				s = s + "\n"
			}
		}
		if i != len(values)-1 {
			s += "\t\t\t},\n"
		} else {
			s += "\t\t\t}\n"
		}
		str += s
	}
	str += "\t\t]"
	return key, str
}

func expand_json_array(field, text string) (string, string) {
	text = strings.Trim(text, " ")
	vals := strings.Split(text, "|")

	if len(vals) == 0 {
		return field, "[]"
	}

	str := "[ "
	for k, val := range vals {
		str += val
		if k != len(vals)-1 {
			str += ", "
		}
	}
	str += "]"

	return field, str
}

func expand_json_normal(field, text string) (string, string) {
	return field, strings.Trim(text, " ")
}

func expand_lua_map(field, text string) (string, string) {
	text = strings.Trim(text, " ")

	idx := strings.Index(field, "#")
	if idx == -1 {
		return field, text
	}

	field_split := strings.Split(field, "#")
	key := field_split[0]
	sub_field := field_split[1]

	if len(text) == 0 {
		return key, "{}"
	}

	str := "{\n"

	fields := strings.Split(sub_field, "_")

	values := strings.Split(text, "|")
	for i, value := range values {
		sub_values := strings.Split(value, "_")
		if len(sub_values) != len(fields) {
			panic("expand_json_map error:" + field + " & " + text)
		}
		s := "\t\t\t{\n"
		for k, sub_value := range sub_values {
			s = s + "\t\t\t\t" + fields[k] + " = " + sub_value
			if k != len(sub_values)-1 {
				s = s + ",\n"
			} else {
				s = s + "\n"
			}
		}
		if i != len(values)-1 {
			s += "\t\t\t},\n"
		} else {
			s += "\t\t\t}\n"
		}
		str += s
	}
	str += "\t\t}"
	return key, str
}

func expand_lua_array(field, text string) (string, string) {
	text = strings.Trim(text, " ")
	vals := strings.Split(text, "|")

	if len(vals) == 0 {
		return field, "{}"
	}

	str := "{ "
	for k, val := range vals {
		str += val
		if k != len(vals)-1 {
			str += ", "
		}
	}
	str += "}"

	return field, str
}

func expand_lua_normal(field, text string) (string, string) {
	return field, strings.Trim(text, " ")
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
		sides[i] = strings.Trim(row.Cells[i].Value, "")
	}

	// field
	row = sheet.Rows[2]
	for i := 0; i < cols; i++ {
		field[i] = strings.Trim(row.Cells[i].Value, "")
	}

	// types
	row = sheet.Rows[3]
	for i := 0; i < cols; i++ {
		types[i] = strings.Trim(row.Cells[i].Value, "")
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

		jcnt := 0
		for _, v := range sides {
			if strings.Contains(v, "j") {
				jcnt++
			}
		}

		for i := 4; i < rows; i++ {
			row := sheet.Rows[i]
			f.WriteString("\t{\n")
			idx := 0
			for k, cell := range row.Cells {
				if !strings.Contains(sides[k], "j") {
					continue
				}

				idx++
				var key, text string
				if types[k] == "map" {
					key, text = expand_json_map(field[k], cell.String())
				} else if types[k] == "array" {
					key, text = expand_json_array(field[k], cell.String())
				} else {
					key, text = expand_json_normal(field[k], cell.String())
				}

				f.WriteString("\t\t\"" + key + "\" : ")
				if types[k] == "string" {
					f.WriteString("\"" + text + "\"")
				} else {
					f.WriteString(text)
				}

				if jcnt == idx {
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
			fmt.Println(file_name+" Checking Invalid: ", err)
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
				if !strings.Contains(sides[k], "l") {
					continue
				}

				var key, text string
				if types[k] == "map" {
					key, text = expand_lua_map(field[k], cell.String())
				} else if types[k] == "array" {
					key, text = expand_lua_array(field[k], cell.String())
				} else {
					key, text = expand_lua_normal(field[k], cell.String())
				}

				f.WriteString("\t\t" + key + " = ")
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