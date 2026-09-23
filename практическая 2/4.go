package main
import "fmt"
func main() {
	v := []string{"Анна","Борис","Анна","Виктор","Анна","Борис"}
	c := map[string]int{"Анна":0,"Борис":0,"Виктор":0}
	for _,n := range v {c[n]++}
	t := 0
	for _,x := range c {t+=x}
	for n,x := range c {fmt.Printf("%s: %d (%.2f%%)\n",n,x,float64(x)/float64(t)*100)}
}