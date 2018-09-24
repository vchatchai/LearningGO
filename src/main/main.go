package main

import (
	"fmt"
	"network"
)

type Name struct {
	FirstName string
	LastName  string
}

func (name *Name) String() string {
	fmt.Println(name)
	return fmt.Sprintln("Mr.", name.FirstName, name.LastName)
}
func main() {

	// name := Name{
	// 	FirstName: "Chatchai",
	// 	LastName: "Vichai",
	// }

	// fmt.Println(name)

	// fmt.Println("Start Testing")
	// testFor()
	// testForNoInitial()

	// s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)

	// s = s[1:]
	// fmt.Println(s)

	// s = s[2:1]
	// fmt.Println(s)

	// pic.Show(Pic)
	// testMap()
	// echoServer()
	// deferFunction()
	// runGoRoutine()
	// channel()

	// var c Circle
	// c = new(Circle)

	// c := Circle{x: 0, y: 0, r: 5}

	// fmt.Println(c)
	// fmt.Println(c.r)
	// fmt.Println(c.area())
	// MaxParallelism()

	// c2 := circle{radius: 10}
	// r2 := rect{height: 10, width: 10}
	// measure(c2)
	// measure(r2)

	// word := "Chatchai Vichai"

	// fmt.Println(strings.Contains("test", "es"))
	// fmt.Println(strings.Count(strings.ToLower(word), "c"))
	// fmt.Println(strings.Join(strings.Split(word, ""), "-"))

	// arr := []byte("atest")
	// str := string([]byte{98, 't', 'e', 's', 't'})
	// fmt.Println(arr)
	// fmt.Println(str)

	// file, err := os.Open("test.txt")
	// if err != nil {

	// 	file, err := os.Create("test.txt")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	file.WriteString("test for golang")

	// }
	// defer file.Close()

	// //get the file size
	// stat, err := file.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// //read the file
	// bs := make([]byte, stat.Size())
	// _, err = file.Read(bs)
	// if err != nil {
	// 	return
	// }

	// str = string(bs)
	// fmt.Println(str)

	// bs, err = ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// str = string(bs)
	// fmt.Println(str)

	// dir, err := os.Open(".")
	// if err != nil {
	// 	return
	// }

	// defer dir.Close()

	// fileInfos, err := dir.Readdir(-1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, fi := range fileInfos {
	// 	fmt.Println(fi.Name(), fi.Mode(), fi.IsDir())
	// }

	// filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
	// 	fmt.Println(path)
	// 	return nil
	// })

	// h := crc32.NewIEEE()
	// h.Write([]byte("test"))
	// v := h.Sum32()
	// fmt.Println(v)

	// runcrc32()
	// runHash()
	// runSHA1()

	// // serverRun()
	// // clientRun()
	// // httpRun();
	// go rpcserver()
	// go rpcclient()

	// var input string
	// fmt.Scanln(&input)

	// go f(0)

	// for i := 0; i < 10; i++ {
	// 	routine.F2(i)
	// }

	// routine.Channel()

	// var input string
	// fmt.Scanln(&input)

	// http.HandleFunc("/test", TestHandler)
	// http.HandleFunc("/static", serveStatic)
	// http.HandleFunc("/", serveDynamic)
	// http.FileServer(http.Dir("D:\\"))
	// http.ListenAndServe(PORT, nil)
	// fmt.Println("Everything is set up")

	// web.GenerateSSL()
	// web.ListenAndServeTLS()
	// web.Httprouter()
	// web.Uploading()
	// web.ResponseWriter()

	// Compute()
	// Assertions()
	// Empty()
	// Nil()
	// InterfaceValues()
	// TypeSwitch()
	// Stringers();
	// IPAddrStringer()
	// Errors2()
	// Readers()
	// Rot13Reader()
	// Images()
	// routine.Saying()
	// routine.Channels()
	// routine.BinaryTree()
	// routine.Mutex()
	// web.Template()
	// storage.Gob()
	// storage.CRUD()
	// storage.CRUD2()
	// storage.SQLX()
	// storage.GORM()
	// webservice.Unmarshal()
	// webservice.Marshal()
	// webservice.XMLDecode()
	// webservice.XMLEncode()
	// webservice.JsonUnmarshal()
	// webservice.JsonDecode()

	// webservice.JsonMarshal()
	// webservice.JsonEcoder();
	// webservice.WebService()
	// web.Cookies()
	// status, err := os.Stat("key.pem")
	// fmt.Println("isExist", status.IsDir(), "err", os.IsExist(err))

	// inputoutput.Stdout()
	// inputoutput.Stdin()
	// osfunc.Cla()
	// inputoutput.Stderr()
	// loging.Syslog()
	// loging.FatalLog()
	// error.NewError()
	// inputoutput.SUM()
	// inputoutput.Average()
	// inputoutput.StopBreak()

	// garbage.GColl()
	// cgo.CGO()
	// osfunc.GetENV()
	// osfunc.SetENV("FOO", "FUCKYOU")
	// cgo.CallC()
	// deferer.Deferer()
	// panicRecover.PanicAndRecover()
	// osfunc.RuntimeEnv()
	// osfunc.RequiredVersion()
	// fmt.Println("NodeTree")
	// slice.Copy()

	// compile.A()
	// compile.B()
	// signals.HandleTwo()
	// signals.HandleAll()
	// files.Cat()
	// files.GoFind()

	// filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
	// 	fmt.Println(path)
	// 	return nil
	// })

	// err := syscaller.Ptrace("echo", "Mastering GO!")
	// fmt.Println(err)

	// syscaller.Trace("ls", "/tmp/")
	// routine.RacingCondition(10)
	// contexts.SlowWWWW()
	// profiling.ProfileMe()
	// profiling.BetterProfile()
	// profiling.GOGC()
	// compilation.XCompile()(
	// web.WWW()
	// profiling.WWW()
	// web.KVWeb()
	// network.ServerTimeOut()
	// network.TCPserver(":8001")
	// network.TCPclient("localhost:8001", "test")
	// network.UDPClient("127.0.0.1:8001")
	// network.FiboTCP(":8001")
	// network.KVTCP(":8001")
	// network.RPCServer(":8001")
	network.LowLevel("127.0.0.1")
}
