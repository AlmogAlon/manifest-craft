package providers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type CountriesProvider struct{}

type APIResponse struct {
	Data map[string]map[string]string
}

func (o *CountriesProvider) fetch() ([]string, error) {
	res, err := http.Get("https://api.first.org/data/v1/countries")
	if err != nil {
		log.Error("could not get countries")
		return []string{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var response APIResponse
	err = decoder.Decode(&response)

	if err != nil {
		log.Error("could not get countries")
		return []string{}, err
	}

	v := make([]string, 0, len(response.Data))

	for _, value := range response.Data {
		v = append(v, value["country"])
	}

	return v, nil
}

func (o *CountriesProvider) GetComponentOptions() (*ComponentTypes, error) {
	res, err := o.fetch()

	return &ComponentTypes{
		"MultiComboBox": res,
	}, err
}
