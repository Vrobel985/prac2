package main
import "fmt"
type M struct{Title string;Year int;Rating float64;Genres []string}
func best(m []M)M{
	b:=m[0]
	for _,x:=range m{if x.Rating>b.Rating{b=x}}
	return b
}
func byGenre(m []M,g string)[]M{
	r:=[]M{}
	for _,x:=range m{for _,y:=range x.Genres{if y==g{r=append(r,x);break}}}
	return r
}
func main() {
	movies:=[]M{{"Фильм1",2020,8.5,[]string{"драма","триллер"}},
		{"Фильм2",2019,7.8,[]string{"комедия"}},
		{"Фильм3",2021,9.0,[]string{"триллер"}}}
	fmt.Println(best(movies))
	fmt.Println(byGenre(movies,"триллер"))
}