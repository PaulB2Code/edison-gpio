package edisongpio

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	SYSFS_CLASS_GPIO = "/sys/class/gpio"
)

type GpioEdi struct {
}

func NewGPIO() (*GpioEdi, error) {
	log.Println("[INFO], Start")

	return &GpioEdi{}, nil
}

//Export pin as gpio
func ExportPin(pinvalue int) error {
	exportPin := fmt.Sprintf("echo %v > %v/export", pinvalue, SYSFS_CLASS_GPIO)
	//log.Println("===> export command ", exportPin)
	_, err := exec.Command("sh", "-c", exportPin).Output()
	if err != nil {
		return (err)
	}
	return nil
}

//Give Pin Mode : 0,1,2
func ModePin(pinMode string, pinValue int) error {
	if pinMode != "1" && pinMode != "1" && pinMode != "0" {
		return errors.New("Mode is not correct, only 0 1 2 authorized")
	}
	cmd := fmt.Sprintf("echo mode%v > /sys/kernel/debug/gpio_debug/gpio%v/current_pinmux", pinMode, pinValue)
	//log.Println(cmd)
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return (err)
	}
	return nil
}

//Give Pin direction : in , out, low, high
func DirectionPin(pinDir string, pinValue int) error {
	if pinDir != "low" && pinDir != "high" && pinDir != "in" && pinDir != "out" {
		return errors.New("Direction is not correct, only low high in out authorized")
	}
	cmd := fmt.Sprintf("echo %v > %v/gpio%v/direction", pinDir, SYSFS_CLASS_GPIO, pinValue)
	//log.Println(cmd)
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return (err)
	}
	return nil
}

//Will send a value to a pin
func ValuePin(pinState string, pinValue int) error {
	cmd := fmt.Sprintf("echo %v > %v/gpio%v/value", pinState, SYSFS_CLASS_GPIO, pinValue)
	//log.Println(cmd)
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Println("[ERROR], ", err)
		return (err)
	}
	return nil
}

//Will send a value to a pin
func ReadPinState(pinValue int) (int, error) {
	cmd := fmt.Sprintf("cat %v/gpio%v/value", SYSFS_CLASS_GPIO, pinValue)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Println("[ERROR], ", err)
		return -1, err
	}
	i, err := strconv.Atoi(strings.Replace(string(out), "\n", "", -1))
	if err != nil {
		log.Println("[ERROR], ", err)
		return -1, err
	}
	return i, nil
}
