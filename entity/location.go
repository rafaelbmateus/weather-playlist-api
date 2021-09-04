package entity

type Location struct {
	City string
	Lat  string
	Lon  string
}

func NewLocation(city, lat, lon string) (*Location, error) {
	c := &Location{
		City: city,
		Lat:  lat,
		Lon:  lon,
	}

	err := c.Validade()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (l *Location) Validade() error {
	if l.City == "" {
		if l.Lat == "" || l.Lon == "" {
			return ErrInvalidEntity
		}
	}

	return nil
}
