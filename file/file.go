package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
)

// SaveToJson 将内容存入json文件
func SaveToJson(filePath string, content interface{}, flag int) error {
	data, err := json.MarshalIndent(content, "", "	")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, flag, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//_, err := io.Write
	if err = ioutil.WriteFile(filePath, data, 0644); err != nil {
		return err
	}

	return nil
}

// ExistsFile 判断文件是否存在
func ExistsFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return true
	}

	return true
}

// GetValidFilename 若所给字符中含有Winodws文件名非法字符
// 返回删除非法字符后的文件名和是否含有非法字符
// Windows现在已知的文件名非法字符有 \ / : * ? " < > |
func GetValidFilename(str string) (string, bool) {
	re, _ := regexp.Compile("[\\\\/:*?\"<>|]")
	res := re.ReplaceAllString(str, "")
	if res == str {
		return str, true
	}
	return res, false
}
