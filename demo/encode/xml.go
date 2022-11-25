package encode

import (
	"encoding/xml"
	"fmt"
	"os"
)

func TestXml() {
	// xml1()
	// xml2()
	xml3()
	// xml4()
}

type Person2 struct {
	XMLName xml.Name `xml:"person2"`
	Name    string   `xml:"name`
	Age     int      `xml:"age"`
}

func xml1() {
	p := Person2{
		Name: "Brightness",
		Age:  22,
	}

	// b, _ := xml.Marshal(p)
	b, _ := xml.MarshalIndent(p, " ", " ") //带缩进
	fmt.Printf("b: %v\n", string(b))
}

func xml2() {
	s := []byte(`<person2><Name>Tom</Name><age>4</age></person2>`)
	var p Person2
	xml.Unmarshal(s, &p)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("p: %T\n", p)

}

func xml3() {
	b, _ := os.ReadFile("a.xml")
	var p Person2
	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("p: %T\n", p)

}

func xml4() {
	f, _ := os.OpenFile("a.xml", os.O_WRONLY, os.ModePerm)
	defer f.Close()
	p := Person2{
		Name: "Brightness",
		Age:  22,
	}

	e := xml.NewEncoder(f)
	e.Encode(p)

}
