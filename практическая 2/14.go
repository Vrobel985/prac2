package main
import "fmt"
type I struct{Name string;Weight float64;IsQuest bool}
func total(i []I)float64{
	t:=0.0;for _,x:=range i{t+=x.Weight};return t
}
func main() {
	inv:=[]I{{"Меч",2.5,false},{"Щит",4.0,false},{"Зелье",0.2,true},{"Лук",1.8,false},{"Стрелы",0.5,false}}
	fmt.Println(total(inv))
}