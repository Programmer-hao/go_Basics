struct结构体：
1.方法受体
```
type Student struct{
    Name string
}

// 作用于Student的拷贝对像，修改不会反映到原对象
type (s Student) SetName(name string){
    s.Name = name
}

// 作用于Student的拷贝对像，修改会反映到原对象
func (s *Student) UpdateName(){
    s.Name = name
}

```
2. 字段标签  
   Tag约定:
   `key:"value"`  
   Tag 的意义:  
   Go语言的反射特性可以动态地给结构体成员复制，可以用Tag来决定赋值前的动作

3. 嵌套和字段提升