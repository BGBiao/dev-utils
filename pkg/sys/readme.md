## 底层system交互的相关包

```
go get -v golang.org/x/sys 
```

subpackage:

- sys/cpu:
- sys/execabs:
- sys/unix: 需要注意的是，该包其实主要会用在其他上层的通用接口中，比如`os` 、`time` 、`net` 之类的包，一般而言使用上层封装的包即可，没必要使用底层的包 
- sys/windows:

`注意:` 在这些系统调用中，当`err != nil`时，将会返回 `syscall.Errno`

有几个核心的变量:

```
var (
	Stdin = 0 
	Stdout = 1
	Stderr = 2
 
)
```


