package main

import "fmt"

func main(){
	var arr1 [2] int = [2]int{10,20}
	arr2 := [3]int{ 30, 40, 50}
	fmt.Println(arr1,arr2)

	arr3 := [...]int{60,70,80,90,100}
	fmt.Println(arr3)
	fmt.Println("Total no. of elemnts in arr3= ", len(arr3))
	fmt.Println("4th position element= ", arr3[4])

	Two_Dim_Arr := [2][3]int{
    	{10, 20, 30},
    	{40, 50, 60},
	}
	fmt.Println(Two_Dim_Arr)
	fmt.Println("Req element= ",Two_Dim_Arr[1][1])
	Two_Dim_Arr[1][1] = 0
	fmt.Println("New value of Req element= ",Two_Dim_Arr[1][1])
}