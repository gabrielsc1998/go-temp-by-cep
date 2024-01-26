package temp

type Temperature struct {
	TempC float64
	TempF float64
	TempK float64
}

func New(tempC float64, tempF float64) *Temperature {
	return &Temperature{
		TempC: tempC,
		TempF: tempF,
		TempK: tempC + 273.15,
	}
}
