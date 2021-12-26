 
1. 数据结构
```
type hmap struct {
    count     int // # 元素个数
    B         uint8  // 说明包含2^B个bucket
    buckets    unsafe.Pointer // buckets的数组指针
    oldbuckets unsafe.Pointer // 结构扩容的时候用于复制的buckets数组
    }
```
