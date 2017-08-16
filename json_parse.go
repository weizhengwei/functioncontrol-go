package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
)

/*
sample json
{
"Humans":{
    "People":[
     {
        "Name":"John",
         "City":"Rome",
        "Pin":675675
     }
   ],
   "Animal":[
   {
        "Name":"Julie",
        "Type":"Dog",
        "Age":5,
        "Owner":"John"
    }
   ]
  }
}
*/


type jsonData struct {
    Humans `json:"Humans"`
}
type Humans struct {
   PeopleData []People `json:"People"`
   AnimalData []Animal `json:"Animal"`
}

type People struct {
  Name string `json:"Name"`
  City string `json:"City"`
  Pin  int    `json:"Pin"`
}

type Animal struct {
   Name  string `json:"Name"`
   Type  string `json:"Type"`
   Age   int    `json:"Age"`
   Owner string `json:"Owner"`
}

func main() {

    file, err := ioutil.ReadFile("./test.json")
    if err != nil {
        fmt.Println("Error while opening file")
        return
    }

    people := make([]People, 0)
    animal := make([]Animal, 0)
    data := jsonData{Humans{people,animal}}

    err = json.Unmarshal(file, &data)
    if err != nil {
        fmt.Println("Error while parsing file")
        return
    }
    
    fmt.Println(data)

}