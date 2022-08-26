package pokemon

type Response struct {
	Data Card `json:"data"`
}

type Card struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Supertype string   `json:"supertype"`
	Subtypes  []string `json:"subtypes"`
	Hp        string   `json:"hp"`
	Types     []string `json:"types"`
	EvolvesTo []string `json:"evolvesTo"`
	Rules     []string `json:"rules"`
	Attacks   []struct {
		Name                string   `json:"name"`
		Cost                []string `json:"cost"`
		ConvertedEnergyCost int      `json:"convertedEnergyCost"`
		Damage              string   `json:"damage"`
		Text                string   `json:"text"`
	} `json:"attacks"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`
	RetreatCost          []string `json:"retreatCost"`
	ConvertedRetreatCost int      `json:"convertedRetreatCost"`
	Set                  struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Series       string `json:"series"`
		PrintedTotal int    `json:"printedTotal"`
		Total        int    `json:"total"`
		Legalities   struct {
			Unlimited string `json:"unlimited"`
			Expanded  string `json:"expanded"`
		} `json:"legalities"`
		PtcgoCode   string `json:"ptcgoCode"`
		ReleaseDate string `json:"releaseDate"`
		UpdatedAt   string `json:"updatedAt"`
		Images      struct {
			Symbol string `json:"symbol"`
			Logo   string `json:"logo"`
		} `json:"images"`
	} `json:"set"`
	Number                 string `json:"number"`
	Artist                 string `json:"artist"`
	Rarity                 string `json:"rarity"`
	NationalPokedexNumbers []int  `json:"nationalPokedexNumbers"`
	Legalities             struct {
		Unlimited string `json:"unlimited"`
		Expanded  string `json:"expanded"`
	} `json:"legalities"`
	Images struct {
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"images"`
	Tcgplayer struct {
		Url       string `json:"url"`
		UpdatedAt string `json:"updatedAt"`
		Prices    struct {
			Holofoil struct {
				Low    float64 `json:"low"`
				Mid    float64 `json:"mid"`
				High   float64 `json:"high"`
				Market float64 `json:"market"`
			} `json:"holofoil"`
		} `json:"prices"`
	} `json:"tcgplayer"`
	Cardmarket struct {
		Url       string `json:"url"`
		UpdatedAt string `json:"updatedAt"`
		Prices    struct {
			AverageSellPrice float64 `json:"averageSellPrice"`
			LowPrice         float64 `json:"lowPrice"`
			TrendPrice       float64 `json:"trendPrice"`
			ReverseHoloTrend float64 `json:"reverseHoloTrend"`
			LowPriceExPlus   float64 `json:"lowPriceExPlus"`
			Avg1             float64 `json:"avg1"`
			Avg7             float64 `json:"avg7"`
			Avg30            float64 `json:"avg30"`
			ReverseHoloAvg1  float64 `json:"reverseHoloAvg1"`
			ReverseHoloAvg7  float64 `json:"reverseHoloAvg7"`
			ReverseHoloAvg30 float64 `json:"reverseHoloAvg30"`
		} `json:"prices"`
	} `json:"cardmarket"`
}

// TODO:
// 	Place this functio in a more appropiate place
// 	Make method

// Every card (of any game) will have this method, so that they can all be treated the same way
// func (c Card) ToMessage() {}
