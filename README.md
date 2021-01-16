 # Golang学习

### 1.goModulePackageSetting

> golang的模块和包的使用设置
>
> >GoModule的简单介绍以及作用【个人理解】
> >
> >>1. **Go Moudule 即 Go 1.11** 之后官方支持的版本管理工具 mod
> >>
> >>2. 通过Go Mod 减轻了对GOPATH的依赖可以方便的安装从其他途径获取的包（一般是github的一些项目）
> >>
> >>3. 例如：
> >>
> >>   --bin
> >>
> >>   ​			main.exe
> >>
> >>   --pkg
> >>
> >>   --src
> >>
> >>   ​		--main
> >>
> >>   ​						go.mod
> >>
> >>   ​						go.sum
> >>
> >>   ​						main.go
> >>
> >>4. 通过在main中使用go mod init main 会在当前文件夹下生成go.mod文件，里面记录了所引用的package
> >>
> >>5. 当通过go build -o 【当前目录】 【输出目录 】 windows下可以生成main.exe文件
> >>
> >>  linux,macos等同样会生成。
> >
> >>1. Go Package 使用 
> >>
> >>2. --mypackage
> >>
> >>   ​				go.mod
> >>
> >>   ​				mypackage.go
> >>
> >>   ​				mypackagetwo.go
> >>
> >>3. 以上图为例：golang中的package 一般对应用 每一个目录名，同一个目录中的package名一样都是对于mypackage来说mypackage.go和mypackagetwo.go都是其中的*[portion](javascript:;)*都可以直接使用。

### 2.Go通道的使用demo

>1.Channel 是 Go 语言中被用来实现并行计算方程之间通信的类型。其功能是允许线程间通过发送和接收来传输指定类型的数据。其初始值是 nil。
>
>2.创建channel的方法如下：
>
>```
>var c1 chan [value type]
>c1 = make([channel type] [value type], [capacity])
>[value type] 定义的是 Channel 中所传输数据的类型。
>[channel type] 定义的是 Channel 的类型，其类型有以下三种：
>“chan” 可读可写——“chan int” 则表示可读写 int 数据的 channel
>"chan<-" 仅可写——“chan<- float64” 则表示仅可写64位 float 数据的 channel
>“<-chan” 仅可读——“<-chan int” 则表示仅可读 int 数据的 channel
>[capacity] 是一个可选参数，其定义的是 channel 中的缓存区 (buffer)。如果不填则默认该 channel 没有缓冲区 (unbuffered)。对于没有缓冲区的 channel，消息的发送和收取必须能同时完成，否则会造成阻塞并提示死锁错误。对于 channel 的阻塞和非阻塞将在后面详细介绍。
>```
>
>```go
>func chanTest() {
>    //声明一个通道
>	var mychan chan int
>    //准备接收数据
>	var myrecive int
>    //构造通道容量
>	mychan = make(chan int, 100)
>    //将数据写入通道
>	mychan <- 100
>    //将数据从通道中取出
>	myrecive = <-mychan
>	fmt.Println(myrecive)
>}
>```
>
>>使用channel发生死锁
>>
>>```go
>>func chanLockTest() {
>>	var mychan chan string
>>	var myrecive string
>>	mychan = make(chan string)
>>	func() {
>>		time.Sleep(2)
>>		mychan <- "hello channel"
>>	}()
>>	myrecive = <-mychan
>>	fmt.Println(myrecive)
>>}
>>```
>>
>>以上代码会报错：fatal error: all goroutines are asleep - deadlock! 
>>
>>😀可以自己试试
>>
>>使用go语句进行并行计算
>>
>>```go
>>func chanLockTest() {
>>	var mychan chan string
>>	var myrecive string
>>	mychan = make(chan string)
>>	go func() {
>>		time.Sleep(2)
>>		mychan <- "hello channel"
>>	}()
>>	myrecive = <-mychan
>>	fmt.Println(myrecive)
>>}
>>```
>>
>>😁可以解决死锁了吧
>>
>>同时还可以使用buffer做缓冲
>>
>>```go
>>func chanLockTest() {
>>	var mychan chan string
>>	var myrecive string
>>	mychan = make(chan string, 1)
>>	func() {
>>		time.Sleep(2)
>>		mychan <- "hello channel"
>>	}()
>>	myrecive = <-mychan
>>	fmt.Println(myrecive)
>>}
>>```
>>
>>当然Buff不是万能的
>>
>>```go
>>func chanLockTest() {
>>	var mychan chan string
>>	var myrecive string
>>	mychan = make(chan string, 1)
>>	func() {
>>		time.Sleep(2)
>>	}()
>>	myrecive = <-mychan
>>	fmt.Println(myrecive)
>>}
>>```
>>
>>当没有数据拉取的时候直接推依然会发生死锁

### 3.Go中的面向对象

> 1. golang中并没有自带传统中高级语言的继承、多态而是提供了一种更有意思的方式，一般称为组合

### 4.Go中的New和Make的区别

> 1. golang中New分配的空间被清零,返回的是指针类型*Type
> 2. Make会分配空间并且初始化，并且返回的是引用Type

###  Go中的接口

>1. golang中只要接口定义的方法被实现则实现
>
>2. ```go
>   type Handler interface {
>       ServeHTTP(ResponseWriter, *Request)   //路由具体实现
>   }
>   
>   func (engin *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
>   	key := req.Method + "-" + req.URL.Path
>   	if handleFunc, ok := engin.router[key]; ok {
>   		handleFunc(w, req)
>   	} else {
>   		fmt.Fprintf(w, "404 not found\n %s", req.URL.Path)
>   	}
>   }
>   
>   ```
>
>3. 上面代码则实现了对Handler接口中定义的方法实现