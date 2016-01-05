package main

import "fmt"

func stringindex(str1 string,str2 string){
	strrune1 := []rune(str1)
	strrune2 := []rune(str2)
	position := -1
	flag := false

	for i:=0; i < len(strrune2); i++{
			if strrune2[i]==strrune1[0]{
				j:=0
				for j=0; j < len(strrune1); j++{
					if strrune2[i+j] == strrune1[j]{
						if j>=len(strrune1)-1{
							position = i;
							fmt.Printf("from the position the str is same: %d\n",position)
							flag = true
						}
					}else{
						break
					}
				}

			}
	}
	if !flag {
		     fmt.Printf("from the position the str is same: %d",position)
	}
}


func main()  {
	 stringindex("xyx","asdxyxyxfsdfxyxsdf")

}


