package main;

import "fmt";

func main() {
    // one
    // var initial [4]int;
    // fmt.Println(initial);   //[0 0 0 0]
    // var numarray = [3]int{1,2,3}
    // fmt.Println(numarray);  //[1 2 3]
    // var stringarray = [3]string{"卡卡西","我爱罗","薛之谦"};
    // fmt.Println(stringarray);   // [卡卡西 我爱罗 薛之谦]

    //two
    // 按照上面的方法每次都要确保提供的初始值和数组长度一致
    // 一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：
    var testArray [4]int
    fmt.Println(testArray);
    fmt.Printf("type of testArray:%T",testArray);
    var numArray = [...]int{996,717}
    fmt.Printf("type of numArray:%T",numArray)
    var StringArray = [...]string{"卡卡西","我爱罗","薛之谦"}
    fmt.Printf("type of StringArray:%T",StringArray)

    //three
    // 我们还可以使用指定索引值的方式来初始化数组，例如:
    arr := [...]int{1:1,3:3}
    fmt.Println(arr);
    fmt.Printf("type of arr:%T",arr);
}
