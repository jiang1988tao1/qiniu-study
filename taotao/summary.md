2016-01-04

1、搭建学习环境，并学习开发环境的基本使用 

2、学习golang编程入门，实现hello world

3、学习编写基本的单元测试

4、学习使用git的基本操作，完成一个demo并使用git提交代码

------------------------------------------------------------

2016-01-05



1、学习常量定义  const

   变量定义 的不同定义方式,及定义时可能遇到的问题,如 x:= 2   此时,x被自动定义为整形,不能x+0.1

2、学习golang类型系统,认识类型系统中字符串、数字类型、数组、切片、map、指针、结构体

1⃣️、字符串 var s string=“hdfs"  s[0] = ‘h’  //不可修改

2⃣️、数字类型:无符号，有符号，浮点数，复数等其他：byte（uint8），rune（int32）

3⃣️  数组 var a [5]int   len(a) //求长度

4⃣️  切片 底层为数组，但可变长 s := make([]string,3);  s[0] = “a”;  s[1]=“b”;    s[2]=“c"

5⃣️  map m := make(map[stirng]int);  m[“k1”]=7

6⃣️  指针  go有指针，但没有指针运算，不能用指针变量遍历字符串的各个字节  go指针指示一种引用    取址操作符 &

7⃣️  创建自己定义的类型</br>
     type Person struc {</br>
     	Name string;</br>
     	age int</br>
     }

3、 了解方法定义 func (p *Person) Study(a int){}

4、 匿名字段  type B struct{ </br>
        	Person  //匿名字段，只有类型Person，字段名为person </br>
           	age int</br>
           	int  //无意义</br>
           }

5、方法继承</br>
  type S struct{</br>
  	a,b string</br>
  }</br>

  type B sturct{</br>
   S</br>
  }

  func(s *S) String(){</br>
  	fmt.println(“qqqq")</br>
  }

  var b B</br>
  b.String()  等价于 b.S.String()</br>

6、字节对齐</br>
  type A struct{</br>
  	a string</br>
  	b byte</br>
  	c byte</br>
  }

  type B struct{</br>
  	a byte</br>
  	b string</br>
  	c byte</br>
  }

  A.size = 24</br>
  B.size =32

----------------------------------------------

2016-01-06

1、控制语句 了解 goto,if,else,for,switch

2、了解错误处理 对编程中错误的日志打印,及了解使用defer控制return 后的 操作

3、了解文件的读写及复制修改

作业:一.求小于 N 的最大素数，以及程序中所有if语句（含if、switch、for循环条件）的总执行次数。不使用第三方库及math等标准库

    二.实现一个程序，限速拷贝文件

    三.实现一个程序，验证 defer 语句和 return 语句的执行顺序

-------------------------------------------------------

2016-01-07

1、了解go面向对象编程,没有类的概念，只有结构体的概念,结构体＝成员变量＋成员方法,结构体之间没有继承／析构函数

2、go 面向对象的语法详解,包括结构体,接口,any类型,接口查询

作业:实现一个程序,更真实的模拟软件公司的运营

