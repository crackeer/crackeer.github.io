## 调度器
- [GO SCHEDULER](https://www.dazhuanlan.com/2019/12/12/5df121e8a157b/)
- [Golang调度器源码分析](http://ga0.github.io/golang/2015/09/20/golang-runtime-scheduler.html)

## 知识网站图谱站点

- [https://draveness.me/golang/](https://draveness.me/golang/)
- [https://github.com/qcrao/Go-Questions](https://github.com/qcrao/Go-Questions)
- [https://github.com/cch123/golang-notes](https://github.com/cch123/golang-notes)
- [https://www.cnblogs.com/qcrao-2018](https://www.cnblogs.com/qcrao-2018)
- [https://xargin.com](https://xargin.com)
- [https://www.linkinstar.wiki](https://www.linkinstar.wiki)

---

## map打印为什么是无序的？

- 根据随机数，选择一个桶位置作为起始点进行遍历迭代，因此每次重新 for range map，你见到的结果都是不一样的。那是因为它的起始位置根本就不固定


## map的底层实现

- 参考[Go语言map底层实现](https://i6448038.github.io/2018/08/26/map-secret/)
- [大话图解golang map源码详解](https://www.linkinstar.wiki/2019/06/03/golang/source-code/graphic-golang-map/)
- [https://github.com/qcrao/Go-Questions/tree/master/map](https://github.com/qcrao/Go-Questions/tree/master/map)
- [https://github.com/cch123/golang-notes/blob/master/map.md](https://github.com/cch123/golang-notes/blob/master/map.md)
- [https://draveness.me/golang-hashmap](https://draveness.me/golang-hashmap)
- [https://lukechampine.com/hackmap.html](https://lukechampine.com/hackmap.html)


## golang的垃圾回收机制

### 逃逸分析

> 逃逸分析是编译器用来确定由程序创建的值所处位置的过程。具体来说，编译器执行静态代码分析，以确定是否可以将值放在构造函数的栈(帧)上，或者该值是否必须“逃逸”到堆上。在Go中，没有关键字或函数可以用于在此决策中指导编译器。只有通过你写的代码来分析这一点

- 逃逸分析的好处是为了减少gc的压力，不逃逸的对象分配在栈上，当函数返回时就回收了资源，不需要gc标记清除。
- 逃逸分析完后可以确定哪些变量可以分配在栈上，栈的分配比堆快，性能好
- 同步消除，如果你定义的对象的方法上有同步锁，但在运行时，却只有一个线程在访问，此时逃逸分析后的机器码，会去掉同步锁运行。

## defer panic recover

> 在golang当中，defer代码块会在函数调用链表中增加一个函数调用。这个函数调用不是普通的函数调用，而是会在函数正常返回，也就是return之后添加一个函数调用。因此，defer通常用来释放函数内部变量

- 当defer被声明时，其参数就会被实时解析
> 函数返回的过程是这样子的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
- defer执行顺序为先进后出

## golang中的值类型和引用类型

- 值类型 ： 基本数据类型int、float、bool、string以及数组和struct
- 值类型：变量直接存储，内存通常在栈中分配。
- 引用类型：指针、slice、map、chan等都是引用类型
- 引用类型：变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配。通过GC回收

---

### rune和byte的区别

- [https://www.jianshu.com/p/4fbf529926ca](https://www.jianshu.com/p/4fbf529926ca)
