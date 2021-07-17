package csv2json

import (
	"fmt"
)

func ExampleToJson() {
	query_result_0_path := "assests/query_result_0.csv"
	res := ToJson(query_result_0_path)
	fmt.Print(string(res))
}
