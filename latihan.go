package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	hasil := FindNextSquare(n)
	fmt.Print(hasil)
}
func FindNextSquare(n int)int{
	var i int
	i = 1
	for i*i <n{
		i += 1
	}
	if i*i == n{
		i+= 1
		return i*i
	}else{
		return -1
	}
}

