package slice

func AveSlice(slice []float64)(average float64){
	var sum float64
	for _,v := range slice{
		sum += v
	}
	return sum/float64(len(slice))
}
