package main

import (
	"fmt"
	"reflect"
)

/*
map[_id:map[year:2019] counter:1]
map[_id:map[year:2018] counter:13]
map[_id:map[year:2017] counter:5]
map[_id:map[year:2015] counter:1]
map[_id:map[year:2005] counter:1]
map[_id:map[year:2004] counter:1]
map[_id:map[year:1980] counter:1]
*/

func main() {
	//n := map[string]string{"year": "2019"}
	m := map[string]interface{}{
		"_id":     map[string]string{"year": "2019"},
		"counter": 13,
	}

	for _, v := range m {
		reflect.ValueOf(v)
		fmt.Println(v)

		/*
			switch val := e.(type) {
			case string:
				fmt.Println(k, "is string", val)
			case int:
				fmt.Println(k, "is int", val)
			case []interface{}:
				fmt.Println(k, "is an array")
				for i, v := range val {
					fmt.Println(i, v)
				}
			default:
				fmt.Println(k, "is unknown type")
			}
		*/
	}
}

func test(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		// how do I iterate here?
		for _, value := range t {
			fmt.Println(value)
		}
	}
}
