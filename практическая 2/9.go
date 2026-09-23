package main
import "fmt"
const (D="double";S="suite";F="free";B="booked")
type R struct{Type,Status string;Cost float64}
var rooms=map[string]R{}
func book(n string)error{
	r,ok:=rooms[n];if!ok||r.Status==B{return fmt.Errorf("no book")}
	r.Status=B;rooms[n]=r;return nil
}
func main() {
	rooms["101"]=R{D,F,3500}
	book("101")
	fmt.Println(rooms["101"].Status)
}