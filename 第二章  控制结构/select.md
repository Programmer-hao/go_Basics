# Select

**1.select特性**
- 管道读写
- 返回值
  - 注意以关闭的管道可以读取到类型的零值
- default
  - select中的default语句不能处理管道读写操作，当select中的所有case语句都阻塞时，default语句将被执行
  
**2. 使用例子**
- select{} 用于永久阻塞
- 快速检错
- 限时等待
    ```
      func waitForStopOrTimeout(stopCh <- chan struct{},timeout time.Duration) <-chan struct{} {
        stopChWithTimeout := make(chan struct{})
	
        go func() {
            select {
            case <- stopCh:
            case <- time.After(timeout):
            close(stopChWithTimeout)
            }
		
        }()
        return stopChWithTimeout
}
    ```

**3. 实现原理**  
######3.1. 数据结构
```
// select中的case语句对于runtime包中的scase(select-case)数据结构：
type scase struct {
    c       *hchan          // 操作的管道
    kind    uint16          // case类型
    elem    unsafe.Pointer  // data element
    ...
}

```
- 管道
- 类型
  - scase中的成员kind表示case语句的类型
    ```
        const(
            caseNil  = iota // 管道的值为nil
            caseRecv        // 读管道的case
            caseSend        // 写管道的case
            caseDefault     // default
        )
    ```
    - caseNil  : 由于管道不可读也不可写，意味着case永远不会命中，运行时会忽略，这就是为什么在case语句中向值为nil的管道中写数据不会触发panic的原因
    - caseRecv : 表示其将从管道中读取数据
    - caseSend : 表示其将发送数据到管道
    - default  : 其不会操作管道
- 数据
  - caseRecv的case中，elem表示从管道读出的数据的存放地址
  - caseSend的case中，elem表示将写入管道的数据的存放地址

**4. 实现逻辑**
```
func selectgo(cas0 *scase,order0 *uint16, ncases int) (int, bool)
```
######4.1 参数
- cas0   : 存储select中case的数据的地址
- order0 : 一个整型数组地址，其长度为case个数的2倍
  - 数组前半部分存放case的随机顺序
  - 数组后半部分管道加锁的顺序
- ncases : case的个数（包括default）,即cas0数组的长度

######4.2 返回值
