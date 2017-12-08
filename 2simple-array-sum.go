package main
import "fmt"

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    s := sum(read())
    fmt.Printf("%d", s)
}

func read() []int {
    var n int
    fmt.Scan(&n)
    
    a := make([]int, n)
    for i := range a {
        fmt.Scanf("%d", &a[i])   
    }
    return a
}

func sum(a []int) int {
    var sum int
    for i := range a {
        sum += a[i]
        //fmt.Printf("i:%d, a%d:%d, sum:%d\n", i, i, a[i], sum)
    }  
    return sum
}