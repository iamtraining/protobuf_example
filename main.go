package main

import (
	"fmt"
	"io/ioutil"
	"log"
	complexpb "protobuf/complex_example"
	enumpb "protobuf/enum_example"
	simple_proto "protobuf/simplepb"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWrite(sm)

	JSON(sm)

	doEnum()

	doComplex()
}

func doComplex() {
	complex := complexpb.ComplexMessage{
		Single: &complexpb.SimpleMessage{
			Id:   1,
			Name: "first message",
		},
		Complex: []*complexpb.SimpleMessage{
			&complexpb.SimpleMessage{
				Id:   2,
				Name: "second message",
			},
			&complexpb.SimpleMessage{
				Id:   3,
				Name: "third message",
			},
			&complexpb.SimpleMessage{
				Id:   4,
				Name: "fourth message",
			},
		},
	}

	fmt.Println(complex)
}

func doEnum() {
	enum := enumpb.EnumMessage{
		Id:           10,
		DayOfTheWeek: enumpb.DayOfTheWeek_sunday,
	}

	enum.DayOfTheWeek = enumpb.DayOfTheWeek_monday

	fmt.Println(enum.GetDayOfTheWeek())
}

func JSON(sm proto.Message) {
	smAsStr := toJSON(sm)

	fmt.Println("to json", smAsStr)

	sm2 := &simple_proto.SimpleMessage{}
	fromJSON(smAsStr, sm2)

	fmt.Println("from json", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("cant convert to json", err)
		return ""
	}

	return str
}

func fromJSON(s string, pb proto.Message) {
	if err := jsonpb.UnmarshalString(s, pb); err != nil {
		log.Fatalln("couldnt unmarshal JSON object", err)
	}
}

func readAndWrite(sm proto.Message) {
	writeToFile("simple.bin", sm)

	sm2 := &simple_proto.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	b, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("cant serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("cant write to file", err)
		return err
	}

	fmt.Println("data has been written")

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("cant read from file", err)
	}

	if err := proto.Unmarshal(b, pb); err != nil {
		log.Fatalln("couldnt pub the bytes into the protobuf struct")
		return err
	}

	fmt.Println("file has been read successfully")

	return nil
}

func doSimple() *simple_proto.SimpleMessage {
	sm := simple_proto.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SimpleList: []int32{1, 3, 5, 7, 9},
	}

	fmt.Println(sm)

	sm.Name = "i renamed you"

	fmt.Println(sm)

	fmt.Println("the id is", sm.GetId())

	return &sm
}
