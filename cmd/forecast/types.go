package forecast

import "time"

type CompactRequest struct {
	Lat float64
	Lon float64
}

type CompactRequestWithAltitude struct {
	CompactRequest
	Altitude int
}

type CompactResponse struct {
	Geometry struct {
		Coordinates []int  `json:"coordinates"`
		Type        string `json:"type"`
	} `json:"geometry"`
	Properties struct {
		Meta struct {
			Units struct {
				AirPressureAtSeaLevel       string `json:"air_pressure_at_sea_level"`
				AirTemperature              string `json:"air_temperature"`
				AirTemperatureMax           string `json:"air_temperature_max"`
				AirTemperatureMin           string `json:"air_temperature_min"`
				CloudAreaFraction           string `json:"cloud_area_fraction"`
				CloudAreaFractionHigh       string `json:"cloud_area_fraction_high"`
				CloudAreaFractionLow        string `json:"cloud_area_fraction_low"`
				CloudAreaFractionMedium     string `json:"cloud_area_fraction_medium"`
				DewPointTemperature         string `json:"dew_point_temperature"`
				FogAreaFraction             string `json:"fog_area_fraction"`
				PrecipitationAmount         string `json:"precipitation_amount"`
				PrecipitationAmountMax      string `json:"precipitation_amount_max"`
				PrecipitationAmountMin      string `json:"precipitation_amount_min"`
				ProbabilityOfPrecipitation  string `json:"probability_of_precipitation"`
				ProbabilityOfThunder        string `json:"probability_of_thunder"`
				RelativeHumidity            string `json:"relative_humidity"`
				UltravioletIndexClearSkyMax string `json:"ultraviolet_index_clear_sky_max"`
				WindFromDirection           string `json:"wind_from_direction"`
				WindSpeed                   string `json:"wind_speed"`
				WindSpeedOfGust             string `json:"wind_speed_of_gust"`
			} `json:"units"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"meta"`
		Timeseries []struct {
			Data struct {
				Instant struct {
					Details struct {
						AirPressureAtSeaLevel   float64 `json:"air_pressure_at_sea_level"`
						AirTemperature          float64 `json:"air_temperature"`
						CloudAreaFraction       float64 `json:"cloud_area_fraction"`
						CloudAreaFractionHigh   float64 `json:"cloud_area_fraction_high"`
						CloudAreaFractionLow    float64 `json:"cloud_area_fraction_low"`
						CloudAreaFractionMedium float64 `json:"cloud_area_fraction_medium"`
						DewPointTemperature     float64 `json:"dew_point_temperature"`
						FogAreaFraction         float64 `json:"fog_area_fraction"`
						RelativeHumidity        float64 `json:"relative_humidity"`
						WindFromDirection       float64 `json:"wind_from_direction"`
						WindSpeed               float64 `json:"wind_speed"`
						WindSpeedOfGust         float64 `json:"wind_speed_of_gust"`
					} `json:"details"`
				} `json:"instant"`
				Next12Hours struct {
					Details struct {
						AirTemperatureMax           float64 `json:"air_temperature_max"`
						AirTemperatureMin           float64 `json:"air_temperature_min"`
						PrecipitationAmount         float64 `json:"precipitation_amount"`
						PrecipitationAmountMax      float64 `json:"precipitation_amount_max"`
						PrecipitationAmountMin      float64 `json:"precipitation_amount_min"`
						ProbabilityOfPrecipitation  int     `json:"probability_of_precipitation"`
						ProbabilityOfThunder        float64 `json:"probability_of_thunder"`
						UltravioletIndexClearSkyMax int     `json:"ultraviolet_index_clear_sky_max"`
					} `json:"details"`
					Summary struct {
						SymbolCode string `json:"symbol_code"`
					} `json:"summary"`
				} `json:"next_12_hours"`
				Next1Hours struct {
					Details struct {
						AirTemperatureMax           float64 `json:"air_temperature_max"`
						AirTemperatureMin           float64 `json:"air_temperature_min"`
						PrecipitationAmount         float64 `json:"precipitation_amount"`
						PrecipitationAmountMax      float64 `json:"precipitation_amount_max"`
						PrecipitationAmountMin      float64 `json:"precipitation_amount_min"`
						ProbabilityOfPrecipitation  int     `json:"probability_of_precipitation"`
						ProbabilityOfThunder        float64 `json:"probability_of_thunder"`
						UltravioletIndexClearSkyMax int     `json:"ultraviolet_index_clear_sky_max"`
					} `json:"details"`
					Summary struct {
						SymbolCode string `json:"symbol_code"`
					} `json:"summary"`
				} `json:"next_1_hours"`
				Next6Hours struct {
					Details struct {
						AirTemperatureMax           float64 `json:"air_temperature_max"`
						AirTemperatureMin           float64 `json:"air_temperature_min"`
						PrecipitationAmount         float64 `json:"precipitation_amount"`
						PrecipitationAmountMax      float64 `json:"precipitation_amount_max"`
						PrecipitationAmountMin      float64 `json:"precipitation_amount_min"`
						ProbabilityOfPrecipitation  int     `json:"probability_of_precipitation"`
						ProbabilityOfThunder        float64 `json:"probability_of_thunder"`
						UltravioletIndexClearSkyMax int     `json:"ultraviolet_index_clear_sky_max"`
					} `json:"details"`
					Summary struct {
						SymbolCode string `json:"symbol_code"`
					} `json:"summary"`
				} `json:"next_6_hours"`
			} `json:"data"`
			Time time.Time `json:"time"`
		} `json:"timeseries"`
	} `json:"properties"`
	Type string `json:"type"`
}
