# Slice
**1. 数据结构**
```
type slice struct{
    array unsafe.Pointer  // 指向底层数据
    len int               // 切片长度
    cap int               // 切片容量
}
```

**2. slice的创建**  
内置函数make()  
数据创建slice（注意）  
nil切片和空切片  
**3. 扩容(append)**   
3.1 append:  
  向slice追加元素，_如果slice空间（cap）不足，则会触发slice扩容_，会重新分配一块更大的内存，将源slice的数据拷贝过  
3.2 基本规则：  
如果原slice的容量小于1024，则新的slice的容量将扩大为原来的2倍  
如果原slice的容量大于1024，则新的slice的容量将扩大为原来的1.25倍  
**4. 切片拷贝**  
将源切片拷贝到目的切片，拷贝数量取两个切片长度的最小值
```
func copySlice(){
    s1 := []int{1,2,3,4,5,6,7}
	s2 := []int{2,3,4}  
  
    //  s2 source s1 dest  
	copy(s1,s2)
	fmt.Println(s1,s2) // s1 [2,3,4,4,5,6,7]   s2 [2,3,4]
	
	//  s1 source s2 dest  
	copy(s2,s1)
	fmt.Println(s1,s2)  // s1 [1,2,3,4,5,6,7] s2 [1,2,3]
}

```
**5. 切片表达式**   
- 简单表达式：  
      底层数组共享  
      存在问题：覆盖问题（如下）
```
    a := [5]int{1,2,3,4,5}
    b := a[1:4]
    b = append(b,0) // 此时元素a[4]将由5变为0
```  
- 扩展表达式  
a[low:high:max] :限制容量，避免出现覆盖问题

作用于 string，数组，切片


