package songs

type Song struct {
	Id       float64
	Name     string
	Artist   string
	Duration float64
	Album    string
	Artwork  string
	Price    float64
	Origin   string
}

type Songs struct {
	Songs []Song
}
