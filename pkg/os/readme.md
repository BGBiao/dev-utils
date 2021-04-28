## os 包相关的内容

注意: os 包其实底层封装了很多和系统相关的底层操作，比如抽象了unix,windows的系统调用

### os package

os 包提供倆独立于平台的操作系统功能接口，其实底层会封装了`sys/unix`和`sys/windows`相关的操作，从而对于上层的系统调用而言实现了通用接口。

常用方法或者函数: 
- os.Chdir(dir string): 切换工作目录
- os.Chmod(name string,mode FileMode): 修改权限; FileMode 一般就是Unix 系统中的权限，比如0644
- os.Chown(name string,uid,gid int) : 修改文件的属主和属组
- os.Chtimes(name string,atime,mtime time.Time): 修改文件的访问时间和修改时间
- os.Clearenv(): 清理全部的环境变量
- os.DirFS(dir string) fs.FS: 返回以dir 为目录的文件树的文件系统
- os.Environ(): 以k=v形式返回系统的全部环境变量
- os.Executable(): 启动进程的可执行文件 
- os.Exit(code int): 制定状态码推出
- os.Expand(s string,map func(string) string) string: 使用给定的map函数中的变量来替换s 字符串中的变量内容 `os.ExpandEnv(s) 等同于 os.Expand(s,os.Getenv)`
- os.ExpandEnv(s string) string: 直接使用系统环境变量中的 key 来渲染整个s
- os.Getpagesize(): 返回底层系统的内存页大小
- os.Link(old,new string): 创建一个链接文件 
- os.Symlink(old,new string): 创建软链接
- os.Mkdir(name string,perm FileMode): 和os.MkdirAll 的区别是后者直接可以创建父级别目录
- os.WriteFile(name string,data []byte,perm FileMode): 直接写入内容到文件，会自动创建文件
- os.ReadFile(name string): 读取文件内容到字节切片
- os.Create(name string): 创建指定名称的文件描述符
- os.Open(name string): 直接打开一个文件，返回fd
- os.OpenFile(name string,flag int,perm FileMode): 指定全县直接打开文件，并返回fd
- os.SameFile(f1,f2 FileInfo): 判断两个文件是否相同，在Unix中会使用device或者inode的值进行比较
- os.ReadDir(name string): 读取目录名称，返回排序后的全部子目录; 返回的`io/fs.DirEntry` 是一个需要实现 Name(),IsDir(),Type(),Info() 的接口


- 

- os.Stdout.Write(d []byte):

### os subPackage 

**os/exec**

该子包主要实现了os内部的命令执，主要包装了`os.StartProcess` 使之可以轻松的重新映射到stdin和stdout ，使用pipes 链接到 I/O 设备，以及其他的一些调整。

需要注意的是，该包不像其他系统的C调用，它不会注入shell 来处理扩展，比如 pipelines ,重定向等等。该包的行为更像是 C 语言的 `exec` 函数系列。

如果想要扩展 glob模式，官方建议的是直接使用 shell 进行调用，但是需要注意一些危险的输入操作，比如万恶的`/bin/rm -rf `，或者可以使用`path/filepath` 包的 glob 函数。

如果想要扩展环境变量，可以直接使用 os 包的 ExpandEnv。

- exec.Cmd: Unix 执行 command 的结构体类型{Path,Args,Env,Dir,Stdin,Stdout,Stderr,ExtraFiles,SysProcAttr,Process,ProcessState}
- exec.Command(cmd string,arg ...string): 初始化一个可执行的 Cmd 的指针
- exec.CommandContext(ctx context.Context,name string,arg ...string): 使用带Context的 Cmd 初始化
- - cmd.CombinedOutput(): 将Command中的指令执行并返回其执行结果，包含标准输出和标准错误输出
- - cmd.Output(): 同上，但是仅返回标准输出
- - cmd.Run(): 运行一个指定的命令并等待它完成, 相关的输出并不会有反馈
- - cmd.Start(): 启动一个命令执行，但是不会等待明令执行完成，如果成功执行后，将设置c.Process 字段; 而cmd.Wait() 防范将在命令推出后返回状态码并释放相关资源
- - cmd.String(): 用于调试的返回人类可读描述 



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








