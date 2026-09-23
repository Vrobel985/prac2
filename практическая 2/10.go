package main
import ("fmt";"strings")
type S struct{C,W,Sen int}
func stats(t string)S{
	w:=len(strings.Fields(t));s:=0
	for _,c:=range t{if c=='.'||c=='!'||c=='?'{s++}}
	return S{len(t),w,s}
}
func main() {
	x:=stats("Привет! Как дела?")
	fmt.Println(x)
}