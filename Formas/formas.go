package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}
func main() {

	generica("Johnny Malta Kalfeltz")
	generica(41)
	generica("Gemeos")
	generica(true)

}
