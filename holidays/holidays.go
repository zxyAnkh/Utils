package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// used to delete holidays stored in map[]
func main() {
	// the day wants to delete
	var days_int []int = []int{1479992251, 1479667600, 1479647600}
	// store days
	var map_day map[string][]byte = make(map[string][]byte, 1)

	// init map
	map_day["holidays"] = append(map_day["holidays"], []byte(strconv.Itoa(1479992251))...)
	map_day["holidays"] = append(map_day["holidays"], []byte(",")...)
	map_day["holidays"] = append(map_day["holidays"], []byte(strconv.Itoa(1479667600))...)
	map_day["holidays"] = append(map_day["holidays"], []byte(",")...)
	map_day["holidays"] = append(map_day["holidays"], []byte(strconv.Itoa(1479571200))...)
	map_day["holidays"] = append(map_day["holidays"], []byte(",")...)
	map_day["holidays"] = append(map_day["holidays"], []byte(strconv.Itoa(1479484800))...)
	map_day["holidays"] = append(map_day["holidays"], []byte(",")...)
	map_day["holidays"] = append(map_day["holidays"], []byte(strconv.Itoa(1479647600))...)
	map_day["holidays"] = append(map_day["holidays"], []byte(",")...)
	map_day["holidays"] = append(map_day["holidays"], []byte(strconv.Itoa(1478793600))...)

	// tmp []string stroe days
	var day_str = strings.Split(string(map_day["holidays"]), ",")
	// store remained days
	var day_str_new []string

	for _, v1 := range day_str {
		// does or not find the same day in map
		var flag bool = false
		for _, v2 := range days_int {
			if isSameDay(v1, strconv.Itoa(v2)) {
				flag = true
				break
			}
		}
		if flag == false {
			day_str_new = append(day_str_new, v1)
		}

	}

	fmt.Println("before delete.")
	printDay(day_str)
	fmt.Println("after delete.")
	printDay(day_str_new)
}

// format print day
func printDay(days []string) {
	for _, v := range days {
		day, _ := strconv.Atoi(v)
		fmt.Println(time.Unix(int64(day), 0).Format(time.RFC3339)[:10])
	}
}

// judge day_1 and day_2 is or not same day
func isSameDay(day_1 string, day_2 string) bool {
	d1, _ := strconv.Atoi(day_1)
	d2, _ := strconv.Atoi(day_2)
	d_1 := time.Unix(int64(d1), 0).Format(time.RFC3339)[:10]
	d_2 := time.Unix(int64(d2), 0).Format(time.RFC3339)[:10]
	if d_1 == d_2 {
		fmt.Println(d_1, d_2)
	}
	return d_1 == d_2
}
