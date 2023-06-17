package main

import (
	"fmt"
	"log"
	"protobuf-lesson/pb"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My Name is Suzuki",
		},
		BirthDay: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	m := protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		AllowPartial:    false,
	}

	out, err := m.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't marshal", err)
	}

	readEmployee := &pb.Employee{}

	unmarshal := protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}

	if err := unmarshal.Unmarshal(out, readEmployee); err != nil {
		log.Fatalln("Can't unmarshal", err)
	}

	fmt.Println(readEmployee)
}
