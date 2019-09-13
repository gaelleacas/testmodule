package main
import "github.com/gaelleacas/gostrings"
func main() {
    text:= gostrings.Squish(" foo   bar    \n   \t   boo")
    println(text)   
}
