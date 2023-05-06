package test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var ageMap = make(map[string]int)
	ageMap["张三"] = 18
	ageMap["李四"] = 19

	ageMap = map[string]int{"张三": 18, "李四": 19}
	fmt.Println(ageMap)

	age := ageMap["张三"]
	age1, exist := ageMap["李四"]

	fmt.Println("age := ageMap[\"张三\"]", age)
	fmt.Println("age1,exist := ageMap[\"李四\"]", "age1:", age1, "exist:", exist)
	age2, exist2 := ageMap["王五"]
	fmt.Println("age2,exist1 := ageMap[\"王五\"]", "age2:", age2, "exist2:", exist2)

	fmt.Println()

	// map的长度
	fmt.Println("len(ageMap)", len(ageMap))

	fmt.Println()

	// map的遍历
	for key := range ageMap {
		fmt.Println("key:", key, "ageMap[key]:", ageMap[key])
	}

	for key, value := range ageMap {
		fmt.Println("key:", key, "value:", value, "ageMap[key]:", ageMap[key])
	}

	fmt.Println()

	ageMap["王五"] = 20
	fmt.Println("ageMap[\"王五\"]", ageMap)
	delete(ageMap, "王五")
	fmt.Println("ageMap[\"王五\"]", ageMap)
	delete(ageMap, "赵六")
	fmt.Println("ageMap[\"赵六\"]", ageMap)
}
