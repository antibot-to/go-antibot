package kasada

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGenCT(t *testing.T) {

	kasada := NewKasadaAntibot("YOUR KEY")

	CD, err := kasada.GenerateCD()
	if err != nil {

	}

	fmt.Println(CD)

	//Returns x-kpsdk-cd header
	//
	// {"workTime":1653514742302,"id":"e6519b0e64e3a0ef28455606af0d971c","answers":[3,2]}
	//

}

func TestGenCustomCT(t *testing.T) {

	kasada := NewKasadaAntibot("YOUR KEY")

	CD, err := kasada.GenerateCustomCD(150)
	if err != nil {

	}

	fmt.Println(CD)

	//Returns full x-kpsdk-cd header with custom "d" value
	//
	// {"workTime":1653515068107,"id":"c574fca7990e4302eb29bf4e0df22ef0","answers":[1,8],"duration":"1.5","d":150,"st":1653515068257,"rst":1653515068407}
	//

}

func TestGenerateCT(t *testing.T) {

	kasada := NewKasadaAntibot("YOUR KEY")

	CD, err := kasada.GenerateFullCD()
	if err != nil {

	}

	fmt.Println(CD)

	//Returns full x-kpsdk-cd header with random "d" value
	//
	// {"workTime":1653515921266,"id":"cdb996af3700f6ccf48b57c3c934ecad","answers":[3,2],"duration":"1.4","d":221,"st":1653515921488,"rst":1653515921709}
	//

}

func TestSolveTl(t *testing.T) {

	kasada := NewKasadaAntibot("YOUR KEY")

	//&http.Response{} = Response of => https://api.nike.com/149e9513-01fa-4fb0-aad4-566afd725d1b/2d206a39-8ed7-437e-a3be-862e0f06eea3/ips.js?ak_bmsc_nke-2.3=02oSiT0I6j1FAbzgHKc83ufq8W45CZTVfDngqLGn9OVkHmav3ITCBpQMilJNeTiGHx5romNAOhprOrW7E838VUJ1woNuJFV0bA3B9CheSU3yqPSZS8RcESJaNb9Fl6pDPNEwuJnD3B5KZIMvncI18W4cHAf

	Tl, err := kasada.SolveTl(&http.Response{}, "https://api.nike.com/149e9513-01fa-4fb0-aad4-566afd725d1b/2d206a39-8ed7-437e-a3be-862e0f06eea3/fp", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36", "Windows")
	if err != nil {

	}

	fmt.Println(Tl)

	//Returns Solved kasada body
	//
	//  kasada body..
	//

}
