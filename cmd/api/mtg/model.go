package mtg

type CardsResponse struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	Name          string   `json:"name"`
	ManaCost      string   `json:"manaCost"`
	Cmc           float64  `json:"cmc"`
	Colors        []string `json:"colors"`
	ColorIdentity []string `json:"colorIdentity"`
	Type          string   `json:"type"`
	Types         []string `json:"types"`
	Subtypes      []string `json:"subtypes,omitempty"`
	Rarity        string   `json:"rarity"`
	Set           string   `json:"set"`
	SetName       string   `json:"setName"`
	Text          string   `json:"text"`
	Artist        string   `json:"artist"`
	Number        string   `json:"number"`
	Power         string   `json:"power,omitempty"`
	Toughness     string   `json:"toughness,omitempty"`
	Layout        string   `json:"layout"`
	Multiverseid  string   `json:"multiverseid,omitempty"`
	ImageUrl      string   `json:"imageUrl,omitempty"`
	Variations    []string `json:"variations,omitempty"`
	ForeignNames  []struct {
		Name         string      `json:"name"`
		Text         string      `json:"text"`
		Type         string      `json:"type"`
		Flavor       interface{} `json:"flavor"`
		ImageUrl     string      `json:"imageUrl"`
		Language     string      `json:"language"`
		Multiverseid interface{} `json:"multiverseid"`
	} `json:"foreignNames,omitempty"`
	Printings    []string `json:"printings"`
	OriginalText string   `json:"originalText,omitempty"`
	OriginalType string   `json:"originalType,omitempty"`
	Legalities   []struct {
		Format   string `json:"format"`
		Legality string `json:"legality"`
	} `json:"legalities"`
	Id      string `json:"id"`
	Flavor  string `json:"flavor,omitempty"`
	Rulings []struct {
		Date string `json:"date"`
		Text string `json:"text"`
	} `json:"rulings,omitempty"`
	Supertypes []string `json:"supertypes,omitempty"`
}
