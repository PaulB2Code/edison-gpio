package edisongpio

import (
	"log"
	"testing"
	"time"
)

const (
	PINOUT1 = 12
	PINOUT2 = 165
	PINOUT3 = 45
	PININ1  = 48
	PININ2  = 49
)

func TestSimpleOutput(t *testing.T) {
	//Enable direction
	ExportPin(PINOUT1)
	ModePin("0", PINOUT1)
	DirectionPin("out", PINOUT1)
	ExportPin(PINOUT2)
	DirectionPin("out", PINOUT2)
	time.Sleep(1 * time.Second)

	ValuePin("0", PINOUT1)
	ValuePin("0", PINOUT2)
	ValuePin("1", PINOUT1)
	time.Sleep(2 * time.Second)
	ValuePin("0", PINOUT1)
	time.Sleep(2 * time.Second)
	ValuePin("1", PINOUT2)

}
func TestReading(t *testing.T) {

	//Enable pin to One
	ExportPin(PINOUT3)
	ModePin("0", PINOUT3)
	DirectionPin("out", PINOUT3)
	ValuePin("1", PINOUT3)

	ExportPin(PININ1)
	DirectionPin("in", PININ1)
	ExportPin(PININ2)
	DirectionPin("in", PININ2)

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		val, _ := ReadPinState(PININ1)
		log.Println("[INFO], Value ", PININ1, " : ", val)
		val, _ = ReadPinState(PININ2)
		log.Println("[INFO], Value ", PININ2, " :", val)

	}
}
