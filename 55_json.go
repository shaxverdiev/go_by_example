package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"pppage"`
	Fruits []string `json:"fffruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	// превращаем структуру в JSON
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))  // {"pppage":1,"fffruits":["apple","peach","pear"]}


	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}
	// из JSON в структуру из byt в &dat
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// fmt.Printf("%T\n", dat["num"])
	num := dat["num"].(float64) // Так как dat имеет тип map[string]interface{}, значение по ключу "num" имеет тип interface{}. Мы знаем, что это число, поэтому приводим его к типу float64 с помощью конструкции приведения типа (.(float64))
	fmt.Println(num)


	strs := dat["strs"].([]interface{})  // Мы приводим его к типу []interface{}, так как JSON-массивы в Go при демаршалинге автоматически преобразуются в срез типа []interface{} (срез значений неизвестного типа).
	str1 := strs[0].(string)
	fmt.Println(str1)


    str := `{"pppage": 1, "fffruits": ["apple", "peach"]}` // ПОЛЯ ИЗ JSON ДОЛЖНЫ ТОЧНО СООТВЕТСТВОВАТЬ СТРУТУРЕ response2
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

	// Кодирование карты в JSON и вывод её в консоль:
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
