package main

import (
    "fmt"
    "sync"
    //"time"
)

type People interface {
	SayHello()
	ToString()
}
type Person struct {
	name  string
	age   int
	phone string
}

type Developer struct {
	Person
	company string
	lang    string
}
type Student struct {
	Person
	university string
}

//A person method
func (p *Person) SayHello() {
	fmt.Printf("Hi, I am %s, %d years old\n", p.name, p.age)
}

//A person method
func (p *Person) ToString() {
	fmt.Printf("[Name: %s, Age: %d, Phone: %s]\n", p.name, p.age, p.phone)
}

//A Developer method override
func (p *Developer) SayHello() {
	p.Person.SayHello()
	fmt.Printf("... I come from %s and I like to program in %s\n", p.company, p.lang)
}

//A Student method override
func (p *Student) SayHello() {
	p.Person.SayHello()
	fmt.Printf("... I study at %s\n", p.university)
}

func introducePeople(ppl []People) {
	for n, _ := range ppl {
		ppl[n].SayHello()
		ppl[n].ToString()
	}
}
func main() {
    var wg sync.WaitGroup
	john := Developer{Person{"John", 35, "111-222-XXX"}, "Accel North America", "Golang"}
	alice := Student{Person{"Alice", 42, "111-222-XXX"}, "UM"}
	peopleArr := [...]People{&john, &alice}
    go func() {
        defer wg.Done()
        introducePeople(peopleArr[0:len(peopleArr)])
    }()
    wg.Add(1)
    wg.Wait()
}
