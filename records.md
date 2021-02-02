## GO踩坑记录  
  
1. slice  
  
> - slice底层是对于数组的饮用，长度和容量  
>    
> 
> - 向函数传递slice，实际是传slice值，也是传递数组的引用  
>   
>   
> - slice的声明方式  
> > var s []string         //s==nil  
> > var s := []string(nil) //s==nil  
> > var s := []string{}    //s!=nil  
> > var s := make([]string,len,cap)   //s!=nil 
>   
> - `s = append(s, news...)` 函数接受原值为nil的切片类型，news切片类型加...之后就类似python解包  
  
2. map  
  
> - map的本质是哈希表的引用，所以向函数传递map类型不用加指针就可以当传引用用  
  
3. error  
  
> - `fmt.Errorf()`可以构造error类型  
