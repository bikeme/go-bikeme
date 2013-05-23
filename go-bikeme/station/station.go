package station

type Position struct {
	Longitude string
	Latitude  string
}

type Address struct {
	DistrictCode string
	ZipCode      string
	Street       string
	Number       string
}

type Status struct {
	AvailableBikes int64
	AvailableDocs  int64
}

type Station struct {
	StationId   string
	StationName string
	Position    Position
	Status      Status
	Address     Address
}
