package main  // GO PROGRAM
import "fmt"
import "strconv"

func main() {
	for j:=0; j<5; j++ {
		even_counter := 0
		odd_counter := 0

		for i:= 0; i<50000001; i++ {
			if i % 2 == 0 {
				even_counter += 1
			} else {
				odd_counter += 1
			}
		}

		fmt.Println("even_counter "+strconv.Itoa(even_counter))
		fmt.Println("odd_counter "+strconv.Itoa(odd_counter))
	}
}
