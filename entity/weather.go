package entity

type Weather struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	}
}
