package akamai

import (
	"fmt"
	"testing"
)

func TestGenerateSensorData(t *testing.T) {

	akamai := NewAkamaiAntibot("YOUR APIKEY")

	sensor, err := akamai.GenerateSensor("02A7BCB3460DCE17F0B30CCF9C86EFB9~-1~YAAQIQUXAsxdiuiAAQAAh/xQ/QeI3VQbEVyGxy2DE2LCs5SsxrV6XOxuSrrsdNVs6VtCO96eX2Ip3xMppv/nu6jrzaJeco1Hh21S9ZN1AsO3QIO9WROgZtp9esGzxQ/lg/A6H9VXOlIUCHqyep1NGEYVMQCIRAb/nItSFj3OBS+N4U6Y6TODIQuJ4jb4D4xlJiuDMRJmUKHadPx8z3kymfrDDXttsQpg6j13z+02Mw9dUXTba4ApNkc8EkaU+QouE/dPs2ysopn96Zxqgd05mSpnoru7RyzZtkbrN/9QzGAodNfCWtSbH7PXC6JypQe1AEsWost4ci+lnf/5ZFiDBEpXw/hLmaE0hQzFX3mvdQ12k+qYF9WeAw1E/jm54Jx7wW3j1Vp5QuRY+rYCrm/+YnFRAFKrGacjBnlj4j358T67Bzp/gVYhVj5MuhwpzLfpVkWI9Eo+uV/71YfjfsXvE2ZrnFGD3s7NvNAOsoyUsJ6DRMIfBjW9OO+BagNQl+SpjB/PzVGcwZxrYw8mKeZqBlxUrWQryQU=~-1~-1~-1", "https://www.nike.com/", "9536C6634F69DFE6C54A8499431F3288~YAAQIQUXAl9ciuiAAQAAUfBQ/Q8C5s7xXT7eh0y4ISyGi9yOy5t6P614sktkiSlD1KjviIlyWhscf7rps2l0Frhvx2Qchl4pPcimSMibH4n54mHe4CDds5QCIaxGWQZieikMA0R98vvTXi/WQdWz+XxZUmSXdLrP9NTfYsBeuSfDNbjFJTUbSUWhC0SrRy+D4zLelH9WE2qd46KAxZH8PBcKpQeWfEaSo84HnKerWLpns9rGpiEOoGhtIzUeppLbzR8R3pnaTn8nvqFUpl2z1qmZVsmpJAwQPe+4n2Cwk65yHc64eweORRbJPIVL4fFGWiULw19ZSH60Rp1wuesBcxhYfEB/K5i6OkhdvEfnj8yTsIRhJO9kVCBVwcWAkrhNLYEleGMIX/EHiNHHY1T94I7UCEvjvE8tBy1pMfhobRogYzJAAXOzxFlhi9c=~3748145~3162421", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36")
	if err != nil {
	}

	fmt.Println(sensor)

	//Returns Sensor Data
	//
	//2;3748145;3162421;4,16,0,0,1,16;A}jP`AdG...
	//
}

func TestGenerateV1SensorData(t *testing.T) {

	akamai := NewAkamaiAntibot("YOUR APIKEY")

	sensor, err := akamai.GenerateV1Sensor("02A7BCB3460DCE17F0B30CCF9C86EFB9~-1~YAAQIQUXAsxdiuiAAQAAh/xQ/QeI3VQbEVyGxy2DE2LCs5SsxrV6XOxuSrrsdNVs6VtCO96eX2Ip3xMppv/nu6jrzaJeco1Hh21S9ZN1AsO3QIO9WROgZtp9esGzxQ/lg/A6H9VXOlIUCHqyep1NGEYVMQCIRAb/nItSFj3OBS+N4U6Y6TODIQuJ4jb4D4xlJiuDMRJmUKHadPx8z3kymfrDDXttsQpg6j13z+02Mw9dUXTba4ApNkc8EkaU+QouE/dPs2ysopn96Zxqgd05mSpnoru7RyzZtkbrN/9QzGAodNfCWtSbH7PXC6JypQe1AEsWost4ci+lnf/5ZFiDBEpXw/hLmaE0hQzFX3mvdQ12k+qYF9WeAw1E/jm54Jx7wW3j1Vp5QuRY+rYCrm/+YnFRAFKrGacjBnlj4j358T67Bzp/gVYhVj5MuhwpzLfpVkWI9Eo+uV/71YfjfsXvE2ZrnFGD3s7NvNAOsoyUsJ6DRMIfBjW9OO+BagNQl+SpjB/PzVGcwZxrYw8mKeZqBlxUrWQryQU=~-1~-1~-1", "https://www.nike.com/", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36", 1.75)
	if err != nil {
	}

	fmt.Println(sensor)

	//Returns Sensor Data
	//
	// 7a74G7m23Vrp0o5c9245251.7-
	//

}
