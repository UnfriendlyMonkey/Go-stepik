package main

import (
    "fmt"
)

type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
    Person
    Model string
}

func androids() {
    a := new(Android)
    a.Person.Talk()
    a.Talk()

    var b = Android{
        Model: "model",
        Person: Person{
            Name: "R2D2",
        },
    }
    b.Talk()
}

type terminator struct {
    On bool
    Ammo int
    Power int
}
func (ts *terminator) Shoot() bool {
    if ts.On == false || ts.Ammo <= 0 {
        return false
    }
    ts.Ammo--
    return true
}
func (ts *terminator) RideBike() bool {
    if ts.On == false || ts.Power <= 0{
        return false
    }
    ts.Power--
    return true
}

func main() {
    t := new(terminator)
    fmt.Println(t)
   // androids() 
}
