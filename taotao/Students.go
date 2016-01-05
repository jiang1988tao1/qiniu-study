package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name string
	age      int
	stuID    int
}

type StudentWrapper struct {
	students []Student
	by func(p, q * Student) bool
}

func (sw StudentWrapper) Len() int {
	return len(sw.students)
}

func (sw StudentWrapper) Less(i, j int) bool {
	return sw.by(&sw.students[i], &sw.students[j])
}

func (sw StudentWrapper) Swap(i, j int){
	sw.students[i], sw.students[j] = sw.students[j], sw.students[i]
}



func main() {

	 studentSlice := make([]Student,0)

	for {


		var sname string
		fmt.Printf("please input name: ")
		fmt.Scanf("%s", &sname)
		var stage int
		fmt.Printf("please input age: ")
		fmt.Scanf("%d", &stage)
		var sstuID int
		fmt.Printf("please input id: ")
		fmt.Scanf("%d", &sstuID)
		var student Student
		student.name = sname
		student.age = stage
		student.stuID = sstuID
		studentSlice=append(studentSlice,student)

		var  flag int
		fmt.Printf("will you input next? 1 or 0: ")
		fmt.Scanf("%d",&flag)
		if flag==0{
			break
		}

	}



	fmt.Println(studentSlice)

	sort.Sort(StudentWrapper{studentSlice, func (p, q *Student) bool {
		return q.stuID > p.stuID
	}})

	fmt.Println(studentSlice)


}
