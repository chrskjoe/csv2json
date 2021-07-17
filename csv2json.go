package csv2json

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
)

func ToJson(path string) (jsonBytes []byte) {
	// read the csv file
	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatal("The file is not found || wrong root")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	content, _ := reader.ReadAll()

	if len(content) < 1 {
		log.Fatal("Something wrong, the file maybe empty or length of the lines are not the same")
	}

	headersArr := make([]string, 0)
	headersArr = append(headersArr, content[0]...)

	//Remove the header row from csv content
	content = content[1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, row := range content {
		buffer.WriteString("{")
		for j, col := range row {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			_, fErr := strconv.ParseFloat(col, 32)
			_, bErr := strconv.ParseBool(col)
			if fErr == nil {
				buffer.WriteString(col)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(col))
			} else {
				buffer.WriteString((`"` + col + `"`))
			}
			//end of property
			if j < len(row)-1 {
				buffer.WriteString(",")
			}

		}
		//end of object of the array
		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	rawMessage := json.RawMessage(buffer.String())
	jsonBytes, _ = json.MarshalIndent(rawMessage, "", "  ")

	return jsonBytes
}
