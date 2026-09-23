package main
import "fmt"
func main() {
	m:=map[string]float64{"Еда":15000,"Транспорт":5000,"Развлечения":3000}
	m["Еда"]+=2000
	t:=0.0
	for _,v:=range m{t+=v}
	fmt.Println(t)
}