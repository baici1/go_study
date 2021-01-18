# 坑点

## 操作数据库

### 错误点

![image-20210118232238407](C:\Users\yang\AppData\Roaming\Typora\typora-user-images\image-20210118232238407.png)

他总是报这样的错误.我一直没有懂这个错误.翻译是:==无效的内存地址或无指针取消引用== 然后百度发现

> 典型的Go 指针声明后没有对指针先初始化而直接赋值导致的错误。

我以为是我单行查询时候出现出现的错误,我看源码,以为是sql.Row导致,但是还是报错.

而且当你不运行你写的查询函数时候,他却不提示错误,导致我的目光都在查询函数上

### 解决点

[解决](https://studygolang.com/articles/19672?fr=sidebar)

其实他提示的错误的行数也是错的.

真正出错的是这一段代码(后面一个是正确的)

![image-20210118232839835](C:\Users\yang\AppData\Roaming\Typora\typora-user-images\image-20210118232839835.png)

正确的代码:

![image-20210118232820230](C:\Users\yang\AppData\Roaming\Typora\typora-user-images\image-20210118232820230.png)

发现仅仅是`=`和`:=`的一个小小区别

>原因就是
>当我们在使用 := 时 会创建一个新的db变量,新的db会把全局变量db覆盖掉

其实很好理解.后续的代码我们都将使用db变量里面的函数,但是当你运行上面一段代码(有`:=`),就把全局的db变量覆盖,导致后续用的是覆盖之后的db变量