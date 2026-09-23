package main
import ("fmt";"strconv")
const (bin=2;dec=10;hex=16)
func conv(n int,from,to int)string{
	s:=strconv.FormatInt(int64(n),from)
	v,_:=strconv.ParseInt(s,from,64)
	return strconv.FormatInt(v,to)
}
func main() {
	n:=255
	fmt.Println(conv(n,dec,bin),conv(n,dec,hex))
}