package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

func expand_json_map(field, text string) (string, string) {
	idx := strings.Index(field, "#")
	if idx == -1 {
		return field, text
	}

	field_split := strings.Split(field, "#")
	key := field_split[0]
	sub_field := field_split[1]

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
	str := "[ "
	vals := strings.Split(text, "|")
	for k, val := range vals {
		str += val
		if k != len(vals)-1 {
			str += ", "
		}
	}
	str += " ]"
	return field, str
}

func expand_json_array_string(field, text string) (string, string) {
	text = strings.Trim(text, " ")
	vals := strings.Split(text, "|")

	if len(vals) == 0 || (len(vals) == 1 && vals[0] == "") {
		return field, "[]"
	}

	str := "["
	for k, val := range vals {
		str += "\"" + val + "\""
		if k != len(vals)-1 {
			str += ", "
		}
	}
	str += "]"

	return field, str
}

func expand_lua_map(field, text string) (string, string) {
	idx := strings.Index(field, "#")
	if idx == -1 {
		return field, text
	}

	field_split := strings.Split(field, "#")
	key := field_split[0]
	sub_field := field_split[1]

	str := "\n\t{\n"

	fields := strings.Split(sub_field, "_")

	values := strings.Split(text, "|")
	for _, value := range values {
		sub_values := strings.Split(value, "_")
		if len(sub_values) != len(fields) {
			panic("expand_json_map error:" + field + " & " + text)
		}
		s := "\t\t{\n"
		for k, sub_value := range sub_values {
			s = s + "\t\t\t" + fields[k] + " = " + sub_value
			s = s + ",\n"
		}
		s += "\t\t},\n"
		str += s
	}
	str += "\t}"
	return key, str
}

func expand_lua_array(field, text string) (string, string) {
	text = strings.Trim(text, " ")
	vals := strings.Split(text, "|")

	if len(text) == 0 || len(vals) == 0 {
		return field, "{}"
	}

	str := "{ "
	for _, val := range vals {
		if val != "" {
			str += val + ", "
		}
	}
	str += "}"

	return field, str
}

func expand_lua_array_string(field, text string) (string, string) {
	text = strings.Trim(text, " ")
	vals := strings.Split(text, "|")

	if len(text) == 0 || len(vals) == 0 {
		return field, "{}"
	}

	str := "{ "
	for _, val := range vals {
		if val != "" {
			str += "\"" + val + "\", "
		}
	}
	str += "}"

	return field, str
}

func main() {
	args := len(os.Args)
	if args < 3 {
		fmt.Println("缺少解析文件名!")
		fmt.Println("usage: " + os.Args[0] + " jl filename [sheet]")
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

	var sheet *xlsx.Sheet
	if args >= 4 {
		for _, s := range excel.Sheets {
			if s.Name == os.Args[3] {
				sheet = s
				break
			}
		}
	} else {
		sheet = excel.Sheets[0]
	}
	if sheet == nil {
		fmt.Println("未找到sheet", os.Args[3])
		return
	}

	name := sheet.Name

	rows := len(sheet.Rows)

	field := make([]string, rows)
	types := make([]string, rows)
	value := make([]string, rows)
	sides := make([]string, rows)

	for i := 1; i < rows; i++ {
		row := sheet.Rows[i]
		field[i] = strings.Trim(row.Cells[0].Value, "")
		types[i] = strings.Trim(row.Cells[1].Value, "")
		value[i] = strings.Trim(row.Cells[2].Value, "")
		sides[i] = strings.Trim(row.Cells[3].Value, "")
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
		f.WriteString("\t{\n")

		for i := 1; i < rows; i++ {
			if !strings.Contains(sides[i], "j") {
				continue
			}

			key := field[i]
			val := value[i]

			if len(val) == 0 {
				continue
			}

			if types[i] == "string" {
				f.WriteString("\t\t\"" + key + "\" : ")
				f.WriteString(fmt.Sprintf("\"%s\"", val))
			} else if types[i] == "array" {
				key, val = expand_json_array(key, val)
				f.WriteString("\t\t\"" + key + "\" : ")
				f.WriteString(val)
			} else if types[i] == "array_string" {
				key, val = expand_json_array_string(key, val)
				f.WriteString("\t\t\"" + key + "\" : ")
				f.WriteString(val)
			} else if types[i] == "map" {
				key, val = expand_json_map(key, val)
				f.WriteString("\t\t\"" + key + "\" : ")
				f.WriteString(val)
			} else if types[i] == "float" {
				f.WriteString("\t\t\"" + key + "\" : ")
				v, _ := strconv.ParseFloat(val, 32)
				f.WriteString(fmt.Sprintf("%.2f", v))
			} else {
				f.WriteString("\t\t\"" + key + "\" : ")
				f.WriteString(val)
			}

			if i != rows-1 {
				f.WriteString(",\n")
			} else {
				f.WriteString("\n")
			}
		}

		f.WriteString("\t}\n")
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
			switch err.(type) {
			case *json.SyntaxError:
				e := err.(*json.SyntaxError)
				s := fmt.Sprintf("[%s] SyntaxError: %s. /Offset: %d(0x%X)", file_name, e.Error(), e.Offset, e.Offset)
				fmt.Println(s)
			case *json.UnmarshalTypeError:
				fmt.Println("Unsurpported format", err)
			default:
				fmt.Println("Untreated err:", err)
			}
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

		f.WriteString("local data =\n")
		f.WriteString("{\n")

		for i := 1; i < rows; i++ {
			if !strings.Contains(sides[i], "j") {
				continue
			}

			key := field[i]
			val := value[i]

			if len(val) == 0 {
				continue
			}

			if types[i] == "string" {
				f.WriteString("\t" + key + " = ")
				f.WriteString(fmt.Sprintf("\"%s\"", val))
			} else if types[i] == "array" {
				key, val = expand_lua_array(key, val)
				f.WriteString("\t" + key + " = ")
				f.WriteString(val)
			} else if types[i] == "array_string" {
				key, val = expand_lua_array_string(key, val)
				f.WriteString("\t" + key + " = ")
				f.WriteString(val)
			} else if types[i] == "map" {
				key, val = expand_lua_map(key, val)
				f.WriteString("\t" + key + " = ")
				f.WriteString(val)
			} else if types[i] == "float" {
				f.WriteString("\t" + key + " = ")
				v, _ := strconv.ParseFloat(val, 32)
				f.WriteString(fmt.Sprintf("%.2f", v))
			} else {
				f.WriteString("\t" + key + " = ")
				f.WriteString(val)
			}

			f.WriteString(",\n")
		}

		f.WriteString("}\n")
		f.WriteString("return data\n")
	}()

	w.Wait()
}
