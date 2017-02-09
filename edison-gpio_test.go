package edisongpio

import (
	"log"
	"testing"
	"time"
)

func TestSimpleOutput(t *testing.T) {
	//Enable direction
	ExportPin(44)
	DirectionPin("out", 44)
	ExportPin(165)
	DirectionPin("out", 165)
	time.Sleep(1 * time.Second)

	ValuePin("0", 44)
	ValuePin("0", 165)
	ValuePin("1", 44)
	time.Sleep(2 * time.Second)
	ValuePin("0", 44)
	time.Sleep(2 * time.Second)
	ValuePin("1", 165)

}
func TestReading(t *testing.T) {
	ExportPin(12)
	DirectionPin("in", 12)
	ExportPin(13)
	DirectionPin("in", 13)

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		val, _ := ReadPinState(12)
		log.Println("[INFO], Value 12", val)
		val, _ = ReadPinState(13)
		log.Println("[INFO], Value 13", val)

	}
}
