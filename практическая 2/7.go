package main
import "fmt"
type E struct{ID int;Name string;Pos string;Sal float64}
func payroll(e []E)(t,a float64){for x:=range e{t+=e[x].Sal};if len(e)>0{a=t/float64(len(e))};return}
func main() {
	e := []E{{1,"Анна","Dev",120000},{2,"Борис","QA",90000}}
	t,a := payroll(e)
	fmt.Printf("Total: %.2f, Avg: %.2f\n",t,a)
}