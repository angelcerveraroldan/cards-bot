package pokemon

// Response This will be used for generics, can make functions to handle both responses that have various cards, or just one card
type Response interface {
	CardsResponse | CardResponse
}

// CardsResponse
// The response given when there are multiple cards
// E.g. Get card by name
type CardsResponse struct {
	Cards      []Card `json:"data"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	Count      int    `json:"count"`
	TotalCount int    `json:"totalCount"`
}

// CardResponse
// The response given when there is onlt one card in the response
// E.g. Get card by id
type CardResponse struct {
	Card Card `json:"data"`
}

// Card
// Pokemon card response
type Card struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Supertype string   `json:"supertype"`
	Subtypes  []string `json:"subtypes"`
	Level     string   `json:"level,omitempty"`
	Hp        string   `json:"hp"`
	Types     []string `json:"types"`
	Attacks   []struct {
		Name                string   `json:"name"`
		Cost                []string `json:"cost"`
		ConvertedEnergyCost int      `json:"convertedEnergyCost"`
		Damage              string   `json:"damage"`
		Text                string   `json:"text"`
	} `json:"attacks,omitempty"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses,omitempty"`
	Resistances []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"resistances,omitempty"`
	RetreatCost          []string `json:"retreatCost,omitempty"`
	ConvertedRetreatCost int      `json:"convertedRetreatCost,omitempty"`
	Set                  struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Series       string `json:"series"`
		PrintedTotal int    `json:"printedTotal"`
		Total        int    `json:"total"`
		Legalities   struct {
			Unlimited string `json:"unlimited"`
			Expanded  string `json:"expanded,omitempty"`
			Standard  string `json:"standard,omitempty"`
		} `json:"legalities"`
		PtcgoCode   string `json:"ptcgoCode,omitempty"`
		ReleaseDate string `json:"releaseDate"`
		UpdatedAt   string `json:"updatedAt"`
		Images      struct {
			Symbol string `json:"symbol"`
			Logo   string `json:"logo"`
		} `json:"images"`
	} `json:"set"`
	Number                 string `json:"number"`
	Artist                 string `json:"artist"`
	Rarity                 string `json:"rarity,omitempty"`
	NationalPokedexNumbers []int  `json:"nationalPokedexNumbers"`
	Legalities             struct {
		Unlimited string `json:"unlimited"`
		Expanded  string `json:"expanded,omitempty"`
		Standard  string `json:"standard,omitempty"`
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
				Low       float64 `json:"low"`
				Mid       float64 `json:"mid"`
				High      float64 `json:"high"`
				Market    float64 `json:"market"`
				DirectLow float64 `json:"directLow,omitempty"`
			} `json:"holofoil,omitempty"`
			ReverseHolofoil struct {
				Low       float64 `json:"low"`
				Mid       float64 `json:"mid"`
				High      float64 `json:"high"`
				Market    float64 `json:"market"`
				DirectLow float64 `json:"directLow,omitempty"`
			} `json:"reverseHolofoil,omitempty"`
			Normal struct {
				Low       float64 `json:"low"`
				Mid       float64 `json:"mid"`
				High      float64 `json:"high"`
				Market    float64 `json:"market"`
				DirectLow float64 `json:"directLow,omitempty"`
			} `json:"normal,omitempty"`
			StEditionHolofoil struct {
				Low       float64 `json:"low"`
				Mid       float64 `json:"mid"`
				High      float64 `json:"high"`
				Market    float64 `json:"market"`
				DirectLow float64 `json:"directLow,omitempty"`
			} `json:"1stEditionHolofoil,omitempty"`
			UnlimitedHolofoil struct {
				Low       float64 `json:"low"`
				Mid       float64 `json:"mid"`
				High      float64 `json:"high"`
				Market    float64 `json:"market"`
				DirectLow float64 `json:"directLow,omitempty"`
			} `json:"unlimitedHolofoil,omitempty"`
		} `json:"prices"`
	} `json:"tcgplayer"`
	Cardmarket struct {
		Url       string `json:"url"`
		UpdatedAt string `json:"updatedAt"`
		Prices    struct {
			AverageSellPrice float64 `json:"averageSellPrice,omitempty"`
			LowPrice         float64 `json:"lowPrice"`
			TrendPrice       float64 `json:"trendPrice"`
			ReverseHoloSell  float64 `json:"reverseHoloSell,omitempty"`
			ReverseHoloLow   float64 `json:"reverseHoloLow,omitempty"`
			ReverseHoloTrend float64 `json:"reverseHoloTrend,omitempty"`
			LowPriceExPlus   float64 `json:"lowPriceExPlus,omitempty"`
			Avg1             float64 `json:"avg1"`
			Avg7             float64 `json:"avg7"`
			Avg30            float64 `json:"avg30"`
			ReverseHoloAvg1  float64 `json:"reverseHoloAvg1,omitempty"`
			ReverseHoloAvg7  float64 `json:"reverseHoloAvg7,omitempty"`
			ReverseHoloAvg30 float64 `json:"reverseHoloAvg30,omitempty"`
		} `json:"prices"`
	} `json:"cardmarket"`
	Abilities []struct {
		Name string `json:"name"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"abilities,omitempty"`
	EvolvesFrom    string   `json:"evolvesFrom,omitempty"`
	EvolvesTo      []string `json:"evolvesTo,omitempty"`
	FlavorText     string   `json:"flavorText,omitempty"`
	Rules          []string `json:"rules,omitempty"`
	RegulationMark string   `json:"regulationMark,omitempty"`
}
