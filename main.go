package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"math"
	"math/cmplx"
	"os"
	"runtime"
	"time"
	"sort"
	"strings"
	"sync"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"golang.org/x/tour/reader"
	"golang.org/x/tour/tree"
)

import (
	"io"
)

func add(a int, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool = true, false, true
// d := 10

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5+12i)
)

func basic1to12() {
	fmt.Println("Hello, 世界")
	fmt.Println("Hello, Pieter")
	rand.Seed(33)
	fmt.Println("Random nr:", rand.Intn(10))
	fmt.Printf("Square root of 7: %g\n", math.Sqrt(7))
	fmt.Println("Square root of 7:", math.Sqrt(7))
	// fmt.Printfln("Pi: %g", math.Pi)
	fmt.Printf("Square of Pi: %g\n", math.Pow(math.Pi, 2))
	fmt.Println(add(4,8))
	fmt.Println(swap("Hello", "Earth"))
	fmt.Println(split(10))

	var i int = 1
	j := 3
	fmt.Println(i, j, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var k int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", k, f, b, s)


	x := 42
	f = float64(x)
	u := uint(f)
	fmt.Printf("%v %v %v", x, f, u)
}

func basic13() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println("")
	fmt.Println("basic13:")
	fmt.Println(x, y, z)
}

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt (x int) int { return x*10 + 1 }
func needFloat (x float64) float64 { return x*0.1 }

func basic16() {
	// fmt.Printf("Small: %g", Small)
	// // fmt.Printf("Big: %x", Big)
	// fmt.Printf("needInt(Small): %g", needInt(Small))
	// fmt.Printf("needFloat(Small): %g", needFloat(Small))
	// fmt.Printf("needFloat(Big): %g", needFloat(Big))

	fmt.Println("")
	fmt.Println("basic16:")
	fmt.Println(Small)
	// fmt.Println(Big)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func flowcontrol1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("")
	fmt.Println("flowcontrol1")
	fmt.Println("Sum:", sum)
}

func flowcontrol2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	} 
	fmt.Println("")
	fmt.Println("flowcontrol2")
	fmt.Println("Sum:", sum)
}

func flowcontrol4() {
	for {}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func flowcontrol5() {
	fmt.Println("")
	fmt.Println("flowcontrol5:")
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x,n); v < lim {
		return v
	}
	return lim
} 

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x,n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func flowcontrol6() {
	fmt.Println("")
	fmt.Println("flowcontrol6:")
	fmt.Println(pow(3,2,10), pow(3,3,20))
}

func flowcontrol7() {
	fmt.Println("")
	fmt.Println("flowcontrol7:")
	fmt.Println(pow2(3,2,10), pow2(3,3,20))
}

func flowcontrol9() {
	fmt.Println()
	fmt.Println("flowcontrol9:")
	// fmt.Println("Runtime:", runtime.GOOS)
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func flowcontrol10() {
	fmt.Println()
	fmt.Println("flowcontrol10:")

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("today: ", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func flowcontrol12() {
	fmt.Println()
	fmt.Println("flowcontrol12:")

	// defer uses LIFO!
	defer fmt.Println("!")
	defer fmt.Println("world")
	fmt.Println("hello")
}

func moretypes1() {
	fmt.Println()
	fmt.Println("moretypes1:")

	i, j := 42, 2701
	fmt.Printf("i: %v, j: %v\n", i, j)

	p := &i
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)
	*p = 21
	fmt.Println("i:", i)

	p = &j
	fmt.Println("p:",p)
	*p = *p / 37
	fmt.Println("j:", j)
}


type Vertex struct {
	X int
	Y int
}

func moretypes2() {
	fmt.Println()
	fmt.Println("moretypes2:")
	fmt.Println(Vertex{1,2})

	v := Vertex{1,2}
	fmt.Println("v.X:",v.X)

	p := &v
	(*p).X = 1e8
	fmt.Println("v:", v)
	p.X = 1e6
	fmt.Println("v:", v)
}

func moretypes5() {
	fmt.Println()
	fmt.Println("moretypes5:")

	v1 := Vertex{1,2}
	v2 := Vertex{X:1}
	v3 := Vertex{}
	p := &Vertex{1,2}
	fmt.Println(v1, v2, v3, p)
}

func moretypes6() {
	fmt.Println()
	fmt.Println("moretypes6:")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}

func moretypes7() {
	fmt.Println()
	fmt.Println("moretypes7:")

	primes := [6]int{2,3,5,7,11,13}
	s := primes[1:4]
	fmt.Println(s)
}

func moretypes9() {
	fmt.Println()
	fmt.Println("moretypes9:")

	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	} {
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func moretypes11() {
	fmt.Println()
	fmt.Println("moretypes11:")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = s[:4]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func moretypes12() {
	fmt.Println()
	fmt.Println("moretypes12:")

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func moretypes13() {
	fmt.Printf("\nmoretypes13:\n")

	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func moretypes14() {
	fmt.Println("moretypes14:")
	
	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	board[0][0] = "X"
	board[2][2] = "X"
	board[0][1] = "X"
	board[2][1] = "X"
	fmt.Println(board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i]," "))
	}
}

func moretypes15() {
	fmt.Println("moretypes15:")

	var s []int
	printSlice(s)

	s = append(s,0)
	printSlice(s)

	s = append(s,1)
	printSlice(s)

	s = append(s,2,3,4)
	printSlice(s)

	s = append(s, s...)
	printSlice(s)
}

func moretypes16() {
	var pow = []int{1,2,4,8,16,32,64}
	for i,v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for _,v := range pow {
		fmt.Printf("v=%d\n", v)
	}
	for i := range pow {
		fmt.Printf("i=%d\n", i)
	}
}

func moretypes17() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("value=%d\n", value)
	}
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8((i+j)/2)
		}
	}
	return pic
}

// store output into image:
// 	go run main.go | sed -e s/^IMAGE:// | base64 -d > moretypes18.img
func moretypes18() {
	pic.Show(Pic)
}

type Vertex2 struct {
	Lat, Long float64
}

func moretypes19() {
	var m map[string]Vertex2

	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}

func moretypes20() {
	m := map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func moretypes22() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m)

	m["Answer"] = 48
	fmt.Println(m)

	delete(m, "Answer")
	fmt.Println(m)

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}


func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		wordCount[word]++
	}
	fmt.Println(wordCount)
	return wordCount
}

func moretypes23() {
	wc.Test(WordCount)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func moretypes24() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0 // function closure
	return func(x int) int {
		sum += x
		return sum
	}
}

func moretypes25() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	n0 := 0
	n1 := 1
	return func() int {
		oldN0 := n0
		oldN1 := n1

		n1 += n0
		n0 = oldN1

		return oldN0
	}
}

func moretypes26() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

type Vertex3 struct {
	X, Y float64
}

// extension method
func (v *Vertex3) Abs() float64 {
	// return math.Sqrt(v.X*v.X + v.Y*v.Y)
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func methods1() {
	v := Vertex3{3,4}
	fmt.Println(v.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// NOT ALLOWED!
// func (f float64) Abs() float64 {
// 	if f < 0 {
// 		return -f
// 	}
// 	return f
// }

func methods3() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methods4() {
	v := Vertex3{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

type Abser interface {
	Abs() float64
}

func methods9() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex3{3,4}

	a = f
	a = &v
	// a = v // ERROR

	fmt.Println(a)
}

type I interface {
	M()
}

type T struct {
	S string
}

// interfaces are implemented implicitly
func (t T) M() {
	fmt.Println(t.S)
}

func methods10() {
	var i I
	i = T{"hello"}
	i.M()
}

type I2 interface {
	M2()
}

func (t *T) M2() {
	fmt.Println(t.S)
}

type F float64

func (f F) M2() {
	fmt.Println(f)
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func methods11() {
	var i I2

	i = &T{"Hello"}
	describe(i)
	i.M2()

	i = F(math.Pi)
	describe(i)
	i.M2()
}

func describe14(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func methods14() {
	var i interface{}
	describe14(i)

	i = 42
	describe14(i)

	i = "hello"
	describe14(i)
}

func do (i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("v*2=%v\n",v*2)
	case string:
		fmt.Printf("len(%q)=%v\n", v, len(v))
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}

func methods16() {
	do(21)
	do("hello")
	do(uint(10))
	do(false)
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func methods17() {
	p1 := Person{"Pieter", 27}
	p2 := Person{"Juliëtte", 26}
	fmt.Println(p1, p2)
}


type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func methods18() {
	hosts := map[string]IPAddr {
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
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

func methods19() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt3(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return x, nil
} 

func methods20() {
	fmt.Println(Sqrt3(2))
	fmt.Println(Sqrt3(-2))
}

func methods21() {
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

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func methods22() {
	reader.Validate(MyReader{})
}

// Source: https://en.wikipedia.org/wiki/ROT13#Go
func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b = 'A' + (b - 'A' + 13) % 26
	} else if b >= 'a' && b <= 'z' {
		b = 'a' + (b - 'a' + 13) % 26
	}
	return b
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	r.r.Read(b)
	for i := 0; i < len(b); i++ {
		b[i] = rot13(b[i])
	}
	return len(b), io.EOF
}

func methods23() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func methods24() {
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
}

type Image struct {}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0,100,100)
}

func (i Image) At(x,y int) color.Color {
	return color.RGBA{255,50,150,255}
} 

func methods25() {
	m := Image{}
	pic.ShowImage(m)
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func generics1() {
	si := []int{10,12,13,14}
	fmt.Println(Index(si, 13))

	ss := []string{"a", "b", "d"}
	fmt.Println(Index(ss, "c"))
}

type List[T any] struct {
	next *List[T]
	val T
}

func (l *List[T]) Add(val T) {
	newList := List[T]{nil,val}
	iter := l
	for iter.next != nil {
		iter = iter.next
	}
	iter.next = &newList
}

func generics2() {
	l := List[int]{nil, 10}
	l.Add(11)
	l.Add(12)
	l.Add(14)

	fmt.Println(l)
	fmt.Println(*l.next)
	fmt.Println(*(*l.next).next)
	fmt.Println(*(*(*l.next).next).next)
	// fmt.Println(*(*(*(*l.next).next).next).next) // exception!
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func concurrency1() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func concurrency2() {
	s := []int{7,2,8,-9,4,0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func concurrency3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // Exception!

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // Exception!
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func concurrency4() {
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func concurrency5() {
	c, quit := make(chan int), make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci3(c, quit)
}

func concurrency6() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// type Tree struct {
// 	Left *Tree
// 	Value int
// 	Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)	
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// walk
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	// get values for ch1 and ch2
	var a1, a2 [10]int
	for i := 0; i < 10; i++ {
		a1[i] = <-ch1
		a2[i] = <-ch2
	}
	
	// sort values
	sort.Ints(a1[:])
	sort.Ints(a2[:])

	// compare values
	return a1 == a2
}

func concurrency8() {
	// Test Walk()
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// Test Same()
	fmt.Println(Same(tree.New(1), tree.New(1))) // should return true
	fmt.Println(Same(tree.New(1), tree.New(2))) // should return false
}

type SafeCounter struct {
	mu sync.Mutex
	v map[string]int	
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func concurrency9() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func contains (s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Crawl(url string, fetcher Fetcher, mu *sync.Mutex, cache *[]string) {
	mu.Lock()
	if (contains(*cache, url)) {
		mu.Unlock()
		return
	}
	*cache = append(*cache, url)
	body, urls, err := fetcher.Fetch(url)
	mu.Unlock()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, fetcher, mu, cache)
	}
	return
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher {
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func concurrency10() {
	mu, cache := sync.Mutex{}, []string{}
	go Crawl("https://golang.org/", fetcher, &mu, &cache)
	time.Sleep(5 * time.Millisecond)
}

func main() {
	// basic1to12()
	// basic13()
	// basic16()

	// flowcontrol1()
	// flowcontrol2()
	// // flowcontrol4() // infinite loop, hence disabled
	// flowcontrol5()
	// flowcontrol6()
	// flowcontrol7()
	// flowcontrol9()
	// flowcontrol10()
	// flowcontrol12()

	// moretypes1()
	// moretypes2()
	// moretypes5()
	// moretypes6()
	// moretypes7()
	// moretypes9()
	// moretypes11()
	// moretypes12()
	// moretypes13()
	// moretypes14()
	// moretypes15()
	// moretypes17()
	// moretypes18()
	// moretypes19()
	// moretypes20()
	// moretypes22()
	// moretypes23()
	// moretypes24()
	// moretypes25()
	// moretypes26()

	// methods1()
	// methods3()
	// methods4()
	// methods9()
	// methods10()
	// methods11()
	// methods14()
	// methods16()
	// methods17()
	// methods18()
	// methods19()
	// methods20()
	// methods21()
	// methods22()
	// methods23()
	// methods24()
	// methods25()

	// generics1()
	// generics2()

	// concurrency1()
	// concurrency2()
	// concurrency3()
	// concurrency4()
	// concurrency5()
	// // concurrency6() // infinite loop, hence disabled
	// concurrency8()
	// concurrency9()
	concurrency10()
}

