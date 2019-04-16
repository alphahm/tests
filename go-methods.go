// Methods on structs
package main

import "fmt"

func main() {

	developer := Developer{}
	developer.Name = "Ahmed"
	developer.Language = "Golang"

	developer.SetHobby("Golf", "Beat Tiger Woods")
	developer.SetHobby("Football", "The next Zidane")

	fmt.Println(developer.GetName())
	fmt.Println(developer.GetHobbies())
}

type Developer struct {
	Name, Language string
	Hobbies        map[string]string
}

func (a *Developer) GetName() string {
	return a.Name
}

func (a *Developer) GetHobbies() map[string]string {
	return a.Hobbies
}

func (a *Developer) SetHobby(hobby, description string) {
	if len(a.Hobbies) < 1 {
		a.Hobbies = make(map[string]string)
	}

	a.Hobbies[hobby] = description
}
