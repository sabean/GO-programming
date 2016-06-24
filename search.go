package main

import "fmt"

func binsearch(arr[]int, searchnum int, first int, last int)int{
	if last<first{
		fmt.Println("Invalid")
		return -1
	}
	mid := (first + last)/2
	if arr[mid] > searchnum {
		return binsearch(arr, searchnum, first, mid-1)
	}else if arr[mid] < searchnum{
		return binsearch(arr, searchnum, mid+1, last)
	} else {
		return mid
	}
}

func insort(arr[]int){
        for j :=1; j<len(arr); j++ {
                key := arr[j]
                i := j -1
                for i>=0 && arr[i]>key{
                        arr[i+1] = arr[i]
                        i--
                }
                arr[i+1] = key
        }
}

func main(){
        numList :=[]int{5, 12, 9, 11, 8, 2, 7}
        fmt.Println("The given list: ", numList)
        insort(numList)
        fmt.Println("The sorted list: ", numList)
	fmt.Println("The number 9 is at position",binsearch(numList, 9, 0, len(numList)-1)+1)
}
