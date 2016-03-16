package lesson

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"
)

const PI = 3.14

var name = "gobar"

type newType int

type gopher struct{}

type golang interface{}

var a, b, c, d = 1, 2, 3, 4

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func DefaultTest() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func forTest() {
	a = 3
	for i := 0; a < 3; i++ {
		a++
		fmt.Println(a)
	}
	fmt.Println("over")
}

func switchTest() {
	switch a {
	case 0:
		fmt.Println("a=0")
	case 3:
		fmt.Println("a=3")
	default:
		fmt.Println("default")
	}
}

func BitTest() {
	x := 6
	y := 10
	/*
		6:  0110
		10: 1010
		--------
		&:  0010=2   //逻辑与，两者都为1
		|:  1110=14  //逻辑或，只要有一个1
		^:  1100=12  //逻辑非，两者不同为1则为1，相同则为0
		&^: 0100=4   //逻辑异或，后面有1前面就是0
	*/
	fmt.Println(x, y)
	fmt.Println(x & y)
	fmt.Println(x | y)
	fmt.Println(x ^ y)
	fmt.Println(x &^ y)
}

func labelTest() {
LABEL1:
	for {
		for k := 0; k < 10; k++ {
			if k > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("ok")
}

func arrTest() {
	//x, y := 1, 2
	//a := [...]*int{&x, &y}
	//var p *[100]int = &a

	//a := [2]int{1, 2}
	//b := [2]int{1, 2}

	//p := new([10]int)

	//a := [2][3]int{{1, 1, 1}, {2, 2, 2}}
	arrSort()
}

func arrSort() {
	a := [...]int{5, 2, 6, 3, 9}

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}

	fmt.Println(a)
}

func sliceTest() {
	//a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(a)

	//var s1 = a[:5]
	//fmt.Println(s1)

	//s1 := make([]int, 3, 10)
	//fmt.Println(len(s1), cap(s1))
	//fmt.Println(s1)

	/*
		a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
		sa := a[2:5]
		sb := sa[1:3]
		fmt.Println(string(sa))
		fmt.Println(string(sb))
	*/

	/*
		a := []int{1, 2, 3, 4, 5}
		s1 := make([]int, 3, 6)
		fmt.Printf("%p", s1)
		s1 = append(s1, 1, 2, 3)
		fmt.Printf("%v %p \n", s1, s1)
		s1 = append(s1, 1, 2, 3)
		fmt.Printf("%v %p \n", s1, s1)

		s2 := a[2:5]
		s3 := a[1:3]
		fmt.Println(s2, s3)
		s2[0] = 9
		fmt.Println(s2, s3)
	*/

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9, 10, 11, 12, 13}
	s3 := s1[:]
	copy(s2[2:4], s1[1:3])

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func mapTest() {
	/*
			m := make(map[int]string)
			m[1] = "1"
			delete(m, 1)
			fmt.Println(m)


		var m map[int]map[int]string
		m = make(map[int]map[int]string)
		var a, ok = m[2][1]
		if !ok {
			m[2] = make(map[int]string)
		}
		m[2][1] = "good"
		a, ok = m[2][1]
		fmt.Println(a, ok)
	*/

	/*
		sm := make([]map[int]string, 5)
		for i := range sm {
			sm[i] = make(map[int]string, 1)
			sm[i][1] = "ok"
			fmt.Println(sm[i])
		}
		fmt.Println(sm)
	*/

	//根据键值排序
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	m2 := map[string]int{}
	s := make([]int, len(m))
	i := 0
	for k, v := range m {
		s[i] = k
		i++
		m2[v] = k
	}
	sort.Ints(s)
	fmt.Println(s)

	fmt.Println(m)
	fmt.Println(m2)
}

func funcTest() {
	//slice 是地址参数
	//int string 是值传递
	fmt.Println("Fuc")

	b := func() {
		fmt.Println("func b")
	}
	b()
}

//闭包 地址拷贝
func closure(x int) func(int) int {
	fmt.Println("%p\n", &x)
	return func(y int) int {
		fmt.Println("%p\n", &x)
		return x + y
	}
}

func DeferTest() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("func a")
}

func DeferB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func DeferC() {
	fmt.Println("Func C")
}

func DeferD() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i=", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}
	for _, f := range fs {
		f()
	}
}

//捕获因未知输入导致的程序异常
func CatchError(nums ...int) int {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	return nums[1] * nums[2] * nums[3] //index out of range
}

//主动抛出 panic，不推荐使用，可能会导致性能问题
func ToFloat64(num string) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[W]", r)
		}
	}()
	if num == "" {
		panic("param is null") //主动抛出 panic
	}
	return strconv.ParseFloat(num, 10)
}

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

type uperson struct {
	string
	int
}

func strTest(per *person) {
	per.Age = 13
	fmt.Println("a", per)
}

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func structTest() {
	/*
		a := &person{
			Name: "joe",
			Age:  19,
		}
		a.Name = "ok"
		fmt.Println(a)
		strTest(a)
		fmt.Println(a)
	*/

	/*
		a := struct {
			Name string
			Age  int
		}{
			Name: "joe",
			Age:  15,
		}
	*/

	/*
		a := person{Name: "joe", Age: 19}
		a.Contact.Phone = "158"
		a.Contact.City = "gz"
		fmt.Println(a)
	*/

	/*
		a := uperson{"joe", 19}
		var b uperson
		b = a
		fmt.Println(a)
		fmt.Print(b)
		fmt.Println(b == a)
	*/
	a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
	b := student{Name: "joe", Age: 19, human: human{Sex: 1}}
	a.Name = "joe2"
	a.Age = 13
	a.Sex = 100
	fmt.Println(a)
	fmt.Println(b)
}

type ATest struct {
	Name string
}

type BTest struct {
	Name string
}

func (a *ATest) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

func (b BTest) Print() {
	b.Name = "BB"
	fmt.Println("B")
}

func methodTest() {
	a := ATest{}
	a.Print()
	fmt.Println(a.Name)

	b := BTest{}
	b.Print()
	fmt.Println(b.Name)
}

type TZ int

func (a *TZ) Print() {
	fmt.Println("TZ")
}

func (t *TZ) Increase(num int) {
	*t += TZ(num)
}

// interface 接口
type empty interface {
}

type Connecter interface {
	Connect()
}

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func Disconnect(usb interface{}) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnenct.", pc.name)
		return
	}
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnect:", v.name)
	default:
		fmt.Println("Unkown device.")
	}
}

// reflect 反射
type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) typeTest() {
	fmt.Println("hello world.")
}

func (u User) Test2(name string) {
	fmt.Println("Hello", name, ",my name is", u.Name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("xx")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func reflectOut() {
	m := Manager{User: User{1, "ok", 12}, title: "123"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}

func reflectTest2() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)

	fmt.Println(x)
}

func SetUser(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Bad")
	}

	if f.Kind() == reflect.String {
		f.SetString("ByeBye")
	}
}

func reflectOut1() {
	u := User{1, "OK", 12}
	SetUser(&u)
	fmt.Println(u)
}

func reflectOut2() {
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Test2")
	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}

// conncurrency 并发 goroutine
func Go() {
	fmt.Println("Go Go Go!!!")
}

func chanTest() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func chanTest1() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go !!!")
		<-c
	}()
	c <- true
}

func chanTest3(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += 1
	}
	fmt.Println(index, a)

	c <- true
}

func chanTest3Out() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10) // 设置缓存保证所有的chan都能执行
	for i := 0; i < 10; i++ {
		go chanTest3(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

// 通过同步包实现10个goroutine打印
func chanTest4(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += 1
	}
	fmt.Println(index, a)

	wg.Done()
}

func chanTest4Out() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go chanTest4(&wg, i)
	}
	wg.Wait()
}

func chanTest5() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}
}

// 会无限执行下去
func chanTest6() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}

func chanTest7() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}

// 消息发送
var ch chan string

func chanSever() {
	i := 0
	for {
		fmt.Println(<-ch)
		ch <- fmt.Sprintf("From chanSever: Hi, #%d", i)
		i++
	}
}
func chanClient() {
	ch = make(chan string)
	go chanSever()
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("From chanClient: Hello, #%d", i)
		fmt.Println(<-ch)
	}
}

func httpSingeCPU() {
	// 限制为CPU的数量减一
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum - 1)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world.")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	fmt.Println("package lesson.init().")
}
