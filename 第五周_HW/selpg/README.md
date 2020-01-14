# Selpg

>Selpg is a utility that selects page range from text input. The input can come from the file specified as the last command line parameter, and can also be from standard input when no file name argument is given. selpg first handles all command line arguments. After scanning all the option parameters (that is, those with a hyphen), if selpg finds a parameter, it will accept the parameter as the name of the input file and try to open it for reading. If there are no other parameters, selpg assumes that the input comes from standard input.

>selpg 是从文本输入选择页范围的实用程序。该输入可以来自作为最后一个命令行参数指定的文件，在没有给出文件名参数时也可以来自标准输入。selpg首先处理所有的命令行参数。在扫描了所有的选项参数（也就是那些以连字符为前缀的参数）后，如果selpg发现还有一个参数，则它会接受该参数为输入文件的名称并尝试打开它以进行读取。如果没有其它参数，则 selpg 假定输入来自标准输入。
>selpg命令开发C语言版本：[selpg.c](https://www.ibm.com/developerworks/cn/linux/shell/clutil/selpg.c)

## Reference
1.[开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)
2.[Linux命令行程序设计](https://wenku.baidu.com/view/c7cf91ee5ef7ba0d4a733b58.html)
3.[Using Python to create UNIX command line tools](https://www.ibm.com/developerworks/aix/library/au-pythocli/index.html)


## design process

导入所需包
```go
packge main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)
```
定义 selpgargs 结构体，存储用户从命令行输入的各项参数等信息：
```go
type selpgargs struct {
	start_page int      //起始页面
	end_page   int      //终止页面
	input_file string   //读取的文件名
	page_len   int      //文件中每页行数，默认值72
	form_deli  bool    //
}
```
全局的 string 变量 progname ，在错误消息中显示，即使将 selpg 命令重命名为别的名称，新的名称也将在消息中显示。
```go
var progname string
```

函数接口
```go
func Usage() {}   // 向用户显示 selpg 指令的用法
func FlagInit(args *selpgargs) {} //初始定义一些标记变量
func ProcessArgs(args *selpgargs) {}
// 把用户输入的各个参数分割成 selpgArgs 结构体实例中的每个部分，并存储在结构体实例中
func ProcessInput(args *selpgargs) {}
// 根据用户输入的各个参数进行相应的操作
func printOrWrite(args *selpgargs, line string, stdin io.WriteCloser) {}
//写行
```


## Selpg usage

 >   selpg -s=Number -e=Number [options] [filename]
>
* `-s=Number` Specify the start number. Must be the first argument.
* `-e=Number` Specify the end number. Must be the second argument.
* `-l=number` [option] Set the number of line per page. Default is 72.
* `-f` [option] Pagers are separated by /f is true. Default is false.
* `filename` [option] Input from this file. Default is input from standard input.


## Test examples

#### 安装Go语言环境，然后将bin目录添加到路径下面，通过go get安装selpg。如果环境变量中已经设置过GOBIN，可以在系统其他位置执行selpg。


### 1.help
```
./selpg -h
```

```
Usage of ./selpg:

./selpg is a tool to select pages from what you want.

Usage:

        selpg -s=Number -e=Number [options] [filename]

The arguments are:

        -s=Number       Start from Page <number>.
        -e=Number       End to Page <number>.
        -l=Number       [options]Specify the number of line per page.Default is 72.
        -d=lp number    [options]Using cat to test.
        -f              [options]Specify that the pages are sperated by \f.
        [filename]      [options]Read input from the file.

If no file specified, ./selpg will read input from stdin. Control-D to end.

```

### 2.some tests

There is a text file named `test` and the content is 135 lines tests:
```
 test1
 test2
 test3
 test4
 ...
 test131
 test132
 test133
 test134
 test135
```
Then here are some commands to test the selpg.

1.以每页15行选择第一页输出

    $ selpg -s 1 -e 1  -l 15 test
    test16
    test17
    test18
    test19
    test20
    test21
    test22
    test23
    test24
    test25
    test26
    test27
    test28
    test29
    test30

2.以每页6行选择2-3页输出

    $ selpg -s=2 -e=3 -l 6 test
    test13
    test14
    test15
    test16
    test17
    test18
    test19
    test20
    test21
    test22
    test23
    test24
3.读它的标准输入，不过 shell／内核已将其重定向，所以标准输入来自test

    $ selpg -s=1 -e=1 -l 10 < test
    test11
    test12
    test13
    test14
    test15
    test16
    test17
    test18
    test19
    test20

4.cat test的标准输出被 shell／内核重定向至 selpg 的标准输入,将第2页到第3页写至 selpg 的标准输出（屏幕）

    $ cat test | selpg -s 2 -e 3 -l 10
    test21
    test22
    test23
    test24
    test25
    test26
    test27
    test28
    test29
    test30
    test31
    test32
    test33
    test34
    test35
    test36
    test37
    test38
    test39
    test40

5.仍然写至它的标准输出，不过 shell／内核将其重定向，所以输出写至output

    $ selpg -s 2 -e 3 -l 10 test >output
    test21
    test22
    test23
    test24
    test25
    test26
    test27
    test28
    test29
    test30
    test31
    test32
    test33
    test34
    test35
    test36
    test37
    test38
    test39
    test40

6.错误命令示例,正常输出在屏幕显示，错误消息被写至error

    $ selpg -s 2 -e 0 -l 5 test
    Invalid arguments
    Usage of selpg:
    
    selpg is a tool to select pages from what you want.

    Usage:

            selpg -s=Number -e=Number [options] [filename]

    The arguments are:

            -s=Number       Start from Page <number>.
            -e=Number       End to Page <number>.
            -l=Number       [options]Specify the number of line per page.Default is 72.
            -f              [options]Specify that the pages are sperated by \f.
            [filename]      [options]Read input from the file.

    If no file specified, selpg will read input from stdin. Control-D to end.
  

7.测试按默认每页行数输出

    $ selpg -s 0 -e 0 test 
    test1
    test2
    test3
    test4
    test5
    test6
    test7
    test8
    test9
    test10
    test11
    test12
    test13
    ...
    test68
    test69
    test70
    test71
    test72
8.第0页到第1页由管道输送至命令“lp -dlp1”，该命令将使输出在打印机 lp1 上打印

    selpg -s 0 -e 1 -l 5 -d=lp1 test
    1  test1
    2  test2
    3  test3
    4  test4
    5  test5
    6  test6
    7  test7
    8  test8
    9  test9 
    10 test10
    
---

---

---
