package main
import "fmt"
type L struct{IP string;Code int;Ts string}
func errs(l []L)[]L{
	r:=[]L{}
	for _,x:=range l{if (x.Code>=400&&x.Code<500)||(x.Code>=500&&x.Code<600){r=append(r,x)}}
	return r
}
func main() {
	l := []L{{"1.1.1.1",200,"t1"},{"1.1.1.2",404,"t2"},{"1.1.1.3",500,"t3"}}
	for _,x:=range errs(l){fmt.Println(x)}
}