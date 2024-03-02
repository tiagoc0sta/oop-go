/*package main

import "fmt"

type Engine struct {
 Horsepower int
}

type Car struct {
 Make   string
 Model  string
 Engine Engine
}

func (e Engine) start() {
 fmt.Println("Engine started")
}

func (c Car) Drive() {
 fmt.Printf("Driving %s %s with %d horsepower engine \n", c.Make, c.Model, c.Engine.Horsepower)
}

func main() {

 car := Car{
  Make:  "Toyota",
  Model: "Camry",
  Engine: Engine{
   Horsepower: 200,
  },
 }

 car.Engine.start()
 car.Drive()

}*/

package main

import "fmt"


