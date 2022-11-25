package encode

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
encoding/json 包
核心函数：Marshal 、 Unmarshal
核心结构体: Decoder Encoder,用于结束流数据
*/

func TestJson() {
	// json1()
	// json2()
	// json3()
	json4()
}

type Person struct {
	Name string
	Age  int
}

func json1() {
	p := Person{
		Name: "Brightness",
		Age:  22,
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))

}

func json2() {
	s := []byte(`{"Name":"Brightness","Age":22}`)
	var p Person
	json.Unmarshal(s, &p)

	fmt.Printf("p: %v\n", p)
	fmt.Printf("T: %T\n", p)

	var v map[string]interface{}
	json.Unmarshal(s, &v)
	for k, v := range v {
		fmt.Printf("key:%v--value:%v\n", k, v)
	}

}

func json3() {
	f, _ := os.Open("a.json")
	defer f.Close()
	d := json.NewDecoder(f)

	var v map[string]interface{}
	err := d.Decode(&v)
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}

	fmt.Printf("v: %v\n", v)

}

func json4() {
	//写入json数据
	f, _ := os.OpenFile("a.json", os.O_WRONLY, os.ModePerm)
	defer f.Close()
	p := Person{
		Name: "Brightness",
		Age:  22,
	}

	e := json.NewEncoder(f)
	err := e.Encode(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
