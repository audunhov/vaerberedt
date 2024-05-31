package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func Load(request CompactRequest) (*CompactResponse, error) {

	query := url.Values{}

	query.Add("lat", strconv.FormatFloat(request.Lat, 'f', -1, 64))
	query.Add("lon", strconv.FormatFloat(request.Lon, 'f', -1, 64))

	url := fmt.Sprintf("https://api.met.no/weatherapi/locationforecast/2.0/compact?%s", query.Encode())

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, errors.New("Could not create request")
	}

	req.Header.Set("User-Agent", "VaerBeredt")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.New("Could not fetch forecast")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("Could not read data")
	}

	var response CompactResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, errors.New("Could not convert to Response")
	}

	return &response, nil
}
