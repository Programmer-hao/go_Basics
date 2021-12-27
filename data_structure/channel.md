Channel

**1. 创建 内置函数make()**  
    - 无缓冲管道 :make(chan string)  
    - 带缓冲管道 : make(chan string,5)

**2.函数传递时 chan int 、 <- chan int、 chan <- int区别**  
 (1) 可读可写管道  chan int  
 单向管道：  
 (2) 可读管道 <- chan int ，不能close  
 (3) 可写管道 chan <- int  

**3. 数据读写的四种情况**
- 无缓冲channel  
- 有缓冲channel  
- nil管道  
- close 管道 

**4.数据结构**
```
type hchan struct {
   qcount 		uint 			//当前队列中剩余的元素个数
   dataqsize 	uint			//环形队列长度，即可以存放的元素个数 cap（）
   buf 		    unsafe.Pointer 	//环形队列指针
   elemsize 	uint16 			//每个元素的大小
   closed 		uint32 			//标识关闭状态
   elemtype 	*_type 			//元素类型
   sendx 		uint 			//队列下标，指示元素写入时存放到队列中的位置
   recvx 		uint 			//队列下标，指示下一个被读取的元素在队列中的位置
   sendq 		waitq 			//等待写消息的协程队列
   recvq 		waitq 			//等待读消息的协程队列
   lock 		mutex  			//互斥所，chan不允许并发读写
}
```
1）环形队列  
2）等待队列（重要）  
sendq:写协程队列  
recvq:读协程队列  
处于等待队列中的协程会在其他协程操作管道时被唤醒
  
3）类型信息  
4）互斥锁  

**5. 管道操作**   
关闭管道  
- 关闭管道时会把所有recvq中的协程全部唤醒，赋予类型对应零值
- 同时sendq队列中的协程也会唤醒，但协程会触发panic

**6. select**  
goroutine-select使用  
管道读取数据的顺序是随机的，case的执行也是随机的  
select的case语句读取管道时不会阻塞  
**7. for-range**  
可持续的想管道中读取数据，管道没有数据时会阻塞当前协程。即使管道被关闭，for-range也会结束。
```
func chanRange(ch chan int)  {
	for i := range ch{
		fmt.Printf("get element for ch : %d",i)
	}
}
```
**8 .阻塞条件以及造成panic的操作**
阻塞条件:  
- 读取时： 
    - 管道无缓冲区  
    - 管道缓冲区中没数据  
    - nil管道  
- 写入时： 
    - 管道无缓冲区  
    - 管道缓冲区已满  
    - nil管道

panic :  
- 关闭管道时，sendq等待写协程队列里的协程会被唤醒，触发panic
- 关闭值为nil的管道  
- 关闭已经被关闭的管道  
- 向已经关闭的管道写入数据  
