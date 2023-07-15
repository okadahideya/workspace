package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"protobuf-lesson/pd"
)

func main() {
	employee := &pd.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pd.Occipation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pd.Company_Project{"ProjectX": &pd.Company_Project{}},
		Profile: &pd.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pd.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	//binData, err := proto.Marshal(employee)
	//if err != nil {
	//	log.Fatalln("Can't serialize", err)
	//}
	//
	//if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
	//	log.Fatalln("Can't write", err)
	//}
	//
	//in, err := ioutil.ReadFile("test.bin")
	//if err != nil {
	//	log.Fatalln("Can't read file", err)
	//}
	//
	//readEmployee := &pd.Employee{}
	//
	//err = proto.Unmarshal(in, readEmployee)
	//if err != nil {
	//	log.Fatalln("Can't deserialize", err)
	//}
	//
	//fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Can't marshal to json", err)
	}

	//fmt.Println(out)

	readEmployee := &pd.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Can't unmarshal from json", err)
	}

	fmt.Println(readEmployee)
}
