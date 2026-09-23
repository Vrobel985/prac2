package main
import ("fmt";"time")
type D struct{SensorID string;Temp,Hum float64;Ts time.Time}
func avg(d []D)float64{
	t:=0.0;for _,x:=range d{t+=x.Temp};return t/float64(len(d))
}
func main() {
	data:=[]D{{"S1",22.5,40,time.Now()},{"S2",23.0,42,time.Now()}}
	fmt.Println(avg(data))
}