# map 
**1. 数据结构**
```
type hmap struct {
    count     int // # 元素个数
    B         uint8  // 说明包含2^B个bucket
    buckets    unsafe.Pointer // buckets的数组指针
    oldbuckets unsafe.Pointer // 结构扩容的时候用于复制的buckets数组
    }
```

```
type bmap struct{
    tophash     [8]uint8    // 存储Hash值的高8位
    data        []byte      // key value 数据
    overflow    *bmap       // 溢出bucket的地址
}

每个bucket可以存储8个键值对  


- tophash是一个长度为8的整型数据，Hash值相同的键（Hash值低位相同的键）存入当前bucket时会将Hash值的高位存储在该数组中，以便后续匹配  
- data区存放的是key-value数据，存放顺序是key/key/key/.../value/value/value
- overflow指针指向下一个bucket,将所有冲突的键连接起来

```
**2. hash冲突**  
两个或多个键值hash匹配到同一bucket上，称作hash冲突。go使用链地址法来解决冲突，当bucket超过8个键值对，再创建一个键值对时，overflow指针指向溢出的bucket. 

**3. 负载因子**  
负载因子 = 键数量/bucket数量   

负载因子过大或过小都不理想：  
- 负载因子过小，说明空间利用率低  
- 负载因子过大，说明冲突严重，存取效率低  

go 语言中的map则在负载因子达到6.5时才会触发rehash  

**4. 扩容**
- 扩容条件
  - 负载因子大于6.5时，即平时每个bucket存储的键值对达到6.5以上
  - overflow的数量大于2^15,即overflow数量超过32768
- 增量扩容（当负载因子过大时）
  - hamp数据结构中oldbuckets成员指向原buckets数据
  - 申请新的buckets数据（长度是原来的2倍），并将数组指针保存到hmap数据结构的buckets成员中
  - 由于一次性搬迁会造成比较大的延时，Go采取逐步搬迁策略，即每次访问map时会触发一次搬迁，每次搬迁2个键值对
- 等量扩容
  - 问题：经过大量的元素增删之后，键值对刚好集中在一小部分的bucket中，会造成溢出的bucket数量增多
  - 解决：重新做一次类似增量扩容的搬迁动作，把松散的键值对重新排列一次，经过重新组织后的overflow的bucket数量会减少，即节省空间又提高范文效率
