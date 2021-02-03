[TOC]

更多参见：[从0开始学GO之目录](https://blog.csdn.net/leacock1991/article/details/112853343)



Go语言的接口并不是其他语言（ C++、 Java等）中所提供的接口概念 ，其他语言中实现类需要时要明确声明自己实现了某个接口。  

在Go语言中，接口(interface)是一个自定义类型，接口类型具体描述了一系列方法的集合。一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口  。Go通过接口实现了鸭子类型(duck-typing)，我们并不关心对象是什么类型，到底是不是鸭子，只关心行为。鸭子类型(duck-typing)  可参见 [C++泛型与多态（4）：Duck Typing](https://blog.csdn.net/leacock1991/article/details/111500530)







个人能力有限，如有错误或者其他建议，敬请告知欢迎探讨，谢谢!