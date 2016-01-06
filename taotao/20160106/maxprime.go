package main

import "fmt"

func main(){
	var input int
	fmt.Printf("input a num between 3 and 4294967296:")
	fmt.Scanf("%d",&input)

	i,j := maxprime(input)
	fmt.Printf("the max prime is :%d \nthe condition count is %d",i,j)

}



func maxprime(input int)(prime int,count int){
	count = 0
	prime =1
	for k:=input;k>1;k--{
		count++
		m := k/2+2
		fmt.Println(m)
		for i:=2;i<m;i++{
			count++
			x := k%i
			if x==0{
				count++
				break
			}
			if i==k/2+1{
				count++
				prime = k
				break
			}
		}
		if prime!=1{
			count++
			break
		}
	}
	return prime,count
}
