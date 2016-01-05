package stringutils

func StringReverse(str string)(result string){
	r := []rune(str)
	for i,j := 0,len(str)-1;i<j;i,j = i+1,j-1{
		r[i],r[j] = r[j],r[i]
	}
	return string(r)
}
