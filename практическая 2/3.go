package main
type Order struct{ID,Total int;Items[]int;Address string;IsCompleted bool}
var orders=map[int]Order{}
func addOrder(id int,items[]int,total int,addr string,done bool){orders[id]=Order{ID:id,Items:items,Total:total,Address:addr,IsCompleted:done}}
func main(){addOrder(1,[]int{101},5999,"ул. Ленина",false)}