## os 包相关的内容

注意: os 包其实底层封装了很多和系统相关的底层操作，比如抽象了unix,windows的系统调用

### os package


### os subPackage 

**os/exec**

该子包主要实现了os内部的命令执行。

**os/signal**

该子包实现了对进来的信号的访问控制(`注意:主要是类Unix 系统，Windows 之类的系统好像后期的版本会支持`)

`注意事项:` 
- SIGKILL(`9`) 和 SIGSTOP(`19`) 不能被程序捕获，因此该包不会对该信号有啥影响
- 同步信号是由程序执行中的错误触发的: SIGBUS SIGFPE SIGSEGV 
- 其他异步信号都不是由于程序错误触发，而是由于内核或者其他程序发送的
- - SIGHUP(`1`) 其实是在当程序的控制终端丢失时发送的信号 
- - SIGINT(`2`) 是当控制终端的用户结束终端程序时收到的信号，通常在执行`Control-C` 的时候就是该信号  
- - SIGQUIT(`3`) 是当控制终端的用户按下退出字符时发送的字段，通常在执行`Control-Backslash` 时的发送信号

Go 程序中，信号的默认行为是: 
- 同步信号被转换为运行时恐慌(panic), 像 SIGHUP, SIGINT, or SIGTERM(`15`) 信号会直接导致程序退出
- 异步信号SIGQUIT, SIGILL, SIGTRAP, SIGABRT, SIGSTKFLT, SIGEMT or SIGSYS 会导致程序退出并将堆栈转储
- SIGTSTP、SIGTTIN或SIGTTOU信号 会执行系统的默认行为
- SIGPROF  信号由go 底层的runtime 处理去实现`runtime.CPUProfile` 
- 其他的信号将会被程序捕获到，但是不会做任何措施

在Go 程序中改变信号的默认行为: 
- notify 禁用了一个给定同步信号的默认行为,然后通过一个或多个已注册的channels来传递信号，此时系统的默认行为将不会发生
- 适用于这些信号: SIGHUP, SIGINT, SIGQUIT, SIGABRT, and SIGTERM 以及工作控制信号 SIGTSTP, SIGTTIN, and SIGTTOU 
- 同样适合但是不会触发动作的信号: SIGUSR1, SIGUSR2, SIGPIPE, SIGALRM, SIGCHLD, SIGCONT, SIGURG, SIGXCPU, SIGXFSZ, SIGVTALRM, SIGWINCH, SIGIO, SIGPWR, SIGSYS, SIGINFO, SIGTHR, SIGWAITING, SIGLWP, SIGFREEZE, SIGTHAW, SIGLOST, SIGXRES, SIGJVM1, SIGJVM2
- 如果为该信号调用Reset或Ignore，或者在为该信号传递给Notify的所有通道上调用Stop，该信号将再次被忽略。Reset将恢复系统对该信号的默认行为，而Ignore将导致系统完全忽略该信号

核心的方法或函数: 

- signal.Ignore(sig ...os.Signal): os.Signal是一个需要实现String()和Signal()的接口
- signal.Notify(c chan<- os.Signal,sig ...os.Signal): 实现信号通知 而实现信号删除的唯一方法是调用Stop()方法
- signal.Reset(sig ...os.Signal): 充值信号
- signal.Stop(c chan<- os.Signal): 删除通知的信号

允许使用不同的通道和相同的信号多次调用Notify:每个通道独立接收传入信号的副本




**os/user**

该子包允许根据名称或者id来查找账户，其实就是基本的用户相关的查询(id)

该包中会使用两种方式来调用底层的用户信息: 纯go写的程序来解析`/etc/passwd`和`/etc/group`; 一个基于 cgo 的C library 编写的函数。

默认会使用cgo的程序，当然用户可以编译时选择使用pure go 的实现。








