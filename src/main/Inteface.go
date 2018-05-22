package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/tour/pic"
)

func Assertions() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)

}

func Empty() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func Nil() {
	var i I
	describe(i)
	i.M()
}

func InterfaceValues() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

}

func TypeSwitch() {
	t := T{"Chatchai"}
	do(21)
	do(t)
	do("hello")
	do(true)
}

func Stringers() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beebleebrox", 900}
	fmt.Println(a, z)
}

func IPAddrStringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string {
	//	fmt.Printf("%v.%v.%v.%v", i[0], i[1], i[2], i[3], )
	// result := strconv.Itoa(int(i[0]))
	// result = result + "."
	// result = result + strconv.Itoa(int(i[1]))
	// result = result + "."
	// result = result + strconv.Itoa(int(i[2]))
	// result = result + "."
	// result = result + strconv.Itoa(int(i[3]))
	result := fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
	return result
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type I interface {
	M()
}

type T struct {
	S string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func Errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number :%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return x, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func Errors2() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Readers() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

func (m MyReader) Read(bs []byte) (n int, err error) {
	for i := range bs {
		bs[i] = 'A'
	}
	return len(bs), nil
}

func ReadersExercise() {

}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(bs []byte) (n int, err error) {
	n, err = r.r.Read(bs)
	for i := 0; i < n; i++ {
		bs[i] = rot13(bs[i])
	}
	return
}

func Rot13Reader() {
	s := strings.NewReader("Xresdoni setierngpner z toliocetnnsla, wrhee svsp rtcz anicosnt ")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

type Image struct {
	rec   image.Rectangle
	model color.Model
	color color.Color
}

func (i Image) ColorModel() color.Model {
	return i.model
}

func (i Image) Bounds() image.Rectangle {
	return i.rec
}

func (i Image) At(x, y int) color.Color {
	return i.color
}

// func (i Image) Write(p []byte) (n int, err error) {
// 	file, err := os.Create("pic.png")
// 	if err != nil {
// 		log.Fatal("Cannot create file", err)
// 	}
// 	defer file.Close()
// 	file.Write
// 	// ioutil.WriteFile("pic.png")

// }

func Images() {
	m := Image{
		rec:   image.Rect(0, 0, 100, 200),
		model: color.RGBAModel,
		color: color.RGBA{255, 0, 0, 255},
	}
	fmt.Println("....")
	pic.ShowImage(m)

	file, err := os.Create("pic.png")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	png.Encode(file, m)

}
