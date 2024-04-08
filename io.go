package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/liserjrqlxue/goUtil/osUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
	"github.com/liserjrqlxue/goUtil/textUtil"
)

func loadData() (data []map[string]string, title []string) {
	for _, f := range snvs {
		if isGz.MatchString(f) {
			d, t := textUtil.Gz2MapArray(f, "\t", isComment)
			title = t
			data = append(data, d...)
		} else {
			d, t := textUtil.File2MapArray(f, "\t", isComment)
			title = t
			fmt.Printf("输入列 loadData 顺序: %v\n", title)

			data = append(data, d...)
		}
	}
	logTime("load anno file")
	return
}

func check_transverTitle_relation_map(coltitles []string) {
	allTitleCheck := false
	var miss_title []string
	index := 1
	for acmg_title, input_title := range transverTitle_relation_map {
		index++
		check_tag := true
		for _, coltitle := range coltitles {
			// if acmg_title == coltitle {
			// 	fmt.Printf("Pre input Title check %d : %s is found .\n", index, coltitle)
			// 	delete(transverTitle_relation_map, acmg_title)
			// 	check_tag = false
			// 	break
			// }
			if input_title == coltitle {
				log.Printf("Pre input Title check %d :  pass : %s is found .\n", index, coltitle)
				fmt.Printf("Pre input Title check %d : \033[32m %s is found .\033[0m\n", index, coltitle)
				check_tag = false
			}
		}
		if check_tag {
			allTitleCheck = true
			delete(transverTitle_relation_map, acmg_title)
			miss_title = append(miss_title, acmg_title)
			log.Printf("Pre input Title check %d :  unpass %s is not found .\n", index, acmg_title)
			fmt.Printf("Pre input Title check %d :\033[31m %s is not found .\033[0m\n", index, acmg_title)

		}
	}
	if allTitleCheck {
		fmt.Printf("\033[31mTitle check : Fail .\033[0m\n")
		log.Printf("Title check : Fail .\n")
		fmt.Printf("\033[31mWarning!!!!! Run acmg annotation without : \"%s\"\033[0m\n", strings.Join(miss_title, "\",\""))
		log.Printf("Warning!!!!! Run acmg annotation without : \"%s\"\n", strings.Join(miss_title, "\",\""))
	} else {
		fmt.Printf("Title check : Pass .\n\n")
		log.Printf("Title check : Pass .\n\n")
	}

}
func transverTitle(raw_item map[string]string) map[string]string {
	item := make(map[string]string)
	for key, value := range transverTitle_relation_map {
		item[key] = raw_item[value]
	}
	return item

}

// mapArray2tsv 将 []map[string]string 转换为 TSV 格式并写入到文件或标准输出流(未提供output文件名时)
func mapArray2tsv(data []map[string]string, output_columns []string, output string) error {
	// 如果output是文件名，则打开文件用于写入
	var writer *os.File
	if output != "" {
		file, err := os.Create(output)
		if err != nil {
			return fmt.Errorf("无法创建输出文件: %w", err)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	// 定义列顺序，假设我们想要按照某个固定的顺序输出键对应的值
	var columns []string
	if len(output_columns) > 0 {
		columns = output_columns
	} else {
		if len(data) > 0 {
			for col := range data[0] {
				columns = append(columns, col)
			}
		}
	}
	// 首行输出列标题
	for i, col := range columns {
		if i > 0 {
			_, _ = writer.Write([]byte("\t"))
		}
		_, _ = writer.WriteString(col)
	}
	_, _ = writer.Write([]byte("\n"))

	// 遍历数据，并按列顺序输出
	for _, item := range data {
		for i, col := range columns {
			value, exists := item[col]
			if exists {
				if i > 0 {
					_, _ = writer.Write([]byte("\t"))
				}
				fit_value := strings.Replace(value, "\n", "^", -1)
				_, _ = writer.WriteString(fit_value)
			} else {
				// 若列不存在于item中，可以填充默认值或者错误信息
				if i > 0 {
					_, _ = writer.Write([]byte("\t"))
				}
				_, _ = writer.WriteString("(Column not found)")
			}
		}
		_, _ = writer.Write([]byte("\n"))
	}

	return nil
}

func mapArray2jsonList(data []map[string]string, output string) {
	var f = osUtil.Create(output)
	defer simpleUtil.DeferClose(f)

	for _, datum := range data {
		simpleUtil.HandleError(f.Write(jsonMarshal(datum)))
	}
}

func jsonMarshal(t interface{}) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	simpleUtil.CheckErr(encoder.Encode(t))
	return buffer.Bytes()
}
