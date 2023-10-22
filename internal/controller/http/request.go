package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"users-services/internal/types"
)

func RequestAge(name string, ch chan types.UserAgeChan) {
	req, err := http.Get(types.AgeAPI + name)
	if err != nil {
		log.Println("Error get age API: ", err)
		ch <- types.UserAgeChan{Err: err}
		return
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error read body age API: ", err)
		ch <- types.UserAgeChan{Err: err}
		return
	}

	userAgeApi := &types.UserAgeApi{}
	err = json.Unmarshal(body, &userAgeApi)
	if err != nil {
		log.Println("Error unmarshal body age API: ", err)
		ch <- types.UserAgeChan{Err: err}
		return
	}

	ch <- types.UserAgeChan{Age: userAgeApi, Err: nil}
}

func RequestGender(name string, ch chan types.UserGenChan) {
	req, err := http.Get(types.GenAPI + name)
	if err != nil {
		log.Println("Error get gender API: ", err)
		ch <- types.UserGenChan{Err: err}
		return
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error read body gender API: ", err)
		ch <- types.UserGenChan{Err: err}
		return
	}

	userGenderApi := &types.UserGenderApi{}
	err = json.Unmarshal(body, &userGenderApi)
	if err != nil {
		log.Println("Error unmarshal body gender API: ", err)
		ch <- types.UserGenChan{Err: err}
		return
	}

	ch <- types.UserGenChan{Gen: userGenderApi, Err: nil}
}

func RequestCountry(name string, ch chan types.UserCountryChan) {
	req, err := http.Get(types.CountryAPI + name)
	if err != nil {
		log.Println("Error get country API: ", err)
		ch <- types.UserCountryChan{Err: err}
		return
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error read body country API: ", err)
		ch <- types.UserCountryChan{Err: err}
		return
	}

	userCountryApi := &types.UserCountryApi{}
	err = json.Unmarshal(body, &userCountryApi)
	if err != nil {
		log.Println("Error unmarshal body country API: ", err)
		ch <- types.UserCountryChan{Err: err}
		return
	}

	ch <- types.UserCountryChan{Cnt: userCountryApi, Err: nil}
}
