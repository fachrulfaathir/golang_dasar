package main

import (
	"fmt"
)

func sumAll(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total;
}

func main() {
    slice := []int{10,2,2,22};

    // sumAll := sumAll(10,5,6,7,4)

    sumAllSlice := sumAll(slice...)

    fmt.Println(sumAllSlice)
    // fmt.Println("sum nya : ", sumAll)

    // s := []int{}

    // person := map[string]string{
    //     "name" : "fachrul",
    //     "address" : "jl bba cijeruk",
    // }

    // person["phone"] = "085171187250"

    // fmt.Println(person["name"]);


    // for k,v := range person{
    //     fmt.Println(k," : ",v)
    // }


    // s= append(s, 1,2,3,4)
    // s= slices.Delete(s,0,2)
    
    // fmt.Println(s)
}