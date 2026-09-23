package main
import "fmt"
func main() {
	posts := [][]string{{"go","backend"},{"git","go","tools"}}
	tags := map[string]bool{}
	for _,p := range posts {for _,t := range p {tags[t]=true}}
	for t := range tags {fmt.Print(t," ")}
}