#声明 这些东西都我边看go 英文文档 边实验得出的结论，最近每有预算买书学习了，就这样自己自娱自乐来学习了，所以其中也许会有很多滑稽的错误，我会在我不断学习了解golang的过程中不断修正，也欢迎批评指正

package main //这是什么？

import "fmt"//这又是什么？

func main(){
	fmt.Println("Hello, world") //没有分号？
}

#首先还是看程序的结构，暂时不看上面的程序，先来了解go 文档中介绍的如何写go代码
##第一个讲的是工作空间 go 默认的工作空间 设计的可以直接来发布go 模块包，（也许是为了让模块包快速的涨起来），即使自己不想发布模块，但是了解一下总没有坏处
	src 目录下是Go的源代码
	pkg 目录下是模块
	bin 目录下是可执行命令
##下面是环境变量，肯定是要一个指向go的安装目录，另一个指向go的bin目录
##接下来是重点的模块包目录，像 fmt 就是个包，当然自己也可以写模块，放到自己的目录下 像$GOPATH/src/github.com/user这样的目录（当然，就现在来看 我还是有一个疑问，我能不能脱离go的环境变量，指定一个目录来当作模块包路径）
##哦，文档里下面说到了安装成可执行命令，这是在晒优势么？把上面的hello程序，放置到一个项目目录，然后 go install 这样就能在 $GOPATH/bin/看到hello 可以执行，也许能把hello重命名（在下面是。。git存回，算了，保存不保存的吧 木有意义）
##嗯哼，要写一个简单的库了，通过这一步就算是入门了，能把一个大程序拆分成小程序的集合，这就已经学会了这么编程，至于拆分的好不好，设计的巧不巧这是后话。$GOPATH/src/github.com/user/newmath 假设创建了这样一个模块包，里面写这样一个
	// Package newmath is a trivial example package.
package newmath

// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
## 将这个文件执行 go build 编译成库 然后使用它

package main//我现在怀疑这个东西就是限定了 库搜索目录的（当然这是猜测，我就是在瞎猜）

import (
	"fmt"

	"github.com/user/newmath"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))
}

##现在已经能用库和模块组合程序了， 真不赖啊
##好吧，是时候证实瞎猜这个问题了，文档里说，go 源码中第一部分必须先引入 package （我得做个使用，我就不在第一部分引入，看你能怎么样），文档还说包中所有的文件必须名字一样，爱咋咋地。重点：包引入会把最后一个元素名，当作模块名引用，这么说有点怪 就是"crypto/rot13" should be named rot13。文档说可执行命令必须总是使用 package main。既然是必须，那就以后看到不是必须的再说

#ok ,是时间总结结束并完成今天的学习了，首先，确认 $GOPATH 是一个很重要的环境变量，包文件搜索也会搜索$GOPATH/src/pkg 所以，可以扩展到自己的目录来书写模块包，其次对于 main 这个包，它默认引入了很多有用的东西，但是不是任何地方都需要引入，需要用到的时候再用也可以。好了

[golang 这样书写构建go 语言程序](http://www.siampple.com/241.html "golang 这样书写构建go 语言程序")
