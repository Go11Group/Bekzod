package main

import "fmt"

func main()  {
	k:=6

	func (k int)  {
		
		for i := 0; i < k; i++ {
			fmt.Println(i*k)
		}
	}(k)

}