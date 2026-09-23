package main
import "fmt"
func main() {
	days := []string{"ПН","ВТ","СР","ЧТ","ПТ","СБ","ВС"}
	s := 0
	for _, d := range days {
		if d=="ПТ"||d=="СБ"||d=="ВС" { s += 2850 } else { s += 2100 }
	}
	fmt.Println(s)
}