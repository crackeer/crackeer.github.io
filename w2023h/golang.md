#### 面试八股文

- https://blog.csdn.net/qq_43716830/article/details/124405506
- https://blog.csdn.net/swg0110/article/details/114242776

#### new() 与 make() 的区别

> new 的作用是初始化一个指向类型的指针(*T)，make 的作用是为 slice，map 或 chan 初始化并返回引用(T)

#### gc垃圾回收，三色标记

> https://juejin.im/post/5d56b47a5188250541792ede

#### 逃逸分析

> https://blog.csdn.net/weixin_44065217/article/details/122482368

#### 调度模型、调度器，GMP，netpoller

1. https://wudaijun.com/2018/01/go-scheduler/
2. https://zhuanlan.zhihu.com/p/108603451
3. http://hushi55.github.io/2015/12/08/golang-G-P-M
4. https://pandaychen.github.io/2020/02/28/GOMAXPROCS-POT/

#### gin源码，前缀树路由

#### chan原理

- https://blog.csdn.net/weixin_42309691/article/details/125694412
- https://www.cnblogs.com/haiyux/p/15161495.html

```go
type hchan struct {
    qcount uint // 队列中的总元素个数
    dataqsiz uint // 环形队列大小，即可存放元素的个数
    buf unsafe.Pointer // 环形队列指针
    elemsize uint16 //每个元素的大小
    closed uint32 //标识关闭状态
    elemtype *_type // 元素类型
    sendx uint // 发送索引，元素写入时存放到队列中的位置

    recvx uint // 接收索引，元素从队列的该位置读出
    recvq waitq // 等待读消息的goroutine队列
    sendq waitq // 等待写消息的goroutine队列
    lock mutex //互斥锁，chan不允许并发读写
}
```

#### init函数要点知识
> https://blog.csdn.net/liuyuede123/article/details/127394496
- init函数先于main函数自动执行，不能被其他函数调用；
- init函数没有输入参数、返回值；
- 每个包可以有多个init函数；
- 包的每个源文件也可以有多个init函数，这点比较特殊；
- 同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个执行顺序。
- 不同包的init函数按照包导入的依赖关系决定执行顺序。
