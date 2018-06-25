package icobench

type SearchResponse struct {
	ICOs        int   `json:"icos"`
	Pages       int   `json:"pages"`
	CurrentPage int   `json:"currentPage"`
	Results     []ICO `json:"results"`
}

type SearchRequest struct {
	OrderDesc       string `json:"orderDesc"`
	OrderAsc        string `json:"orderAsc"`
	Page            int    `json:"page"`
	Category        int    `json:"category"`
	Platform        string `json:"platform"`
	Accepting       string `json:"accepting"`
	Country         string `json:"country"`
	Status          string `json:"status"`
	Search          string `json:"search"`
	Bonus           string `json:"bonus"`
	Bounty          string `json:"bounty"`
	Team            string `json:"team"`
	Expert          string `json:"expert"`
	Rating          string `json:"rating"`
	StartAfter      string `json:"startAfter"`
	EndBefore       string `json:"endBefore"`
	Registration    int    `json:"registration"`
	ExclRestCountry string `json:"excludeRestrictedCountry"`
}

type TrendingResponse struct {
	Results []ICO `json:"results"`
}

type ICO struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	URL     string  `json:"url"`
	Logo    string  `json:"logo"`
	Desc    string  `json:"desc"`
	Rating  float64 `json:"rating"`
	Premium int     `json:"premium"`
	Dates   struct {
		PreICOStart string `json:"preIcoStart"`
		PreICOEnd   string `json:"preIcoEnd"`
		ICOStart    string `json:"icoStart"`
		ICOEnd      string `json:"icoEnd"`
	} `json:"dates"`
}

type ProfileResponse struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Rating        float64 `json:"rating"`
	RatingTeam    float64 `json:"ratingTeam"`
	RatingVision  float64 `json:"ratingVision"`
	RatingProduct float64 `json:"ratingProduct"`
	RatingProfile float64 `json:"ratingProfile"`
	URL           string  `json:"url"`
	Tagline       string  `json:"tagline"`
	Intro         string  `json:"intro"`
	About         string  `json:"about"`
	Logo          string  `json:"logo"`
	Country       string  `json:"country"`
	Milestones    []struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"milestones"`
	TeamIntro    string `json:"teamIntro"`
	Notification string `json:"notification"`
	Registration string `json:"registration"`
	Restrictions []struct {
		Country string `json:"country"`
	} `json:"restrictions"`
	Links struct {
		Twitter     string `json:"twitter"`
		Slack       string `json:"slack"`
		Telegram    string `json:"telegram"`
		Facebook    string `json:"facebook"`
		Medium      string `json:"medium"`
		Bitcointalk string `json:"bitcointalk"`
		Github      string `json:"github"`
		Reddit      string `json:"reddit"`
		Discord     string `json:"discord"`
		Youtube     string `json:"youtube"`
		Www         string `json:"www"`
		Bounty      string `json:"bounty"`
		Whitepaper  string `json:"whitepaper"`
	} `json:"links"`
	Finance struct {
		Token       string  `json:"token"`
		Price       string  `json:"price"`
		Bonus       bool    `json:"bonus"`
		Tokens      int     `json:"tokens"`
		Tokentype   string  `json:"tokentype"`
		Hardcap     string  `json:"hardcap"`
		Softcap     string  `json:"softcap"`
		Raised      float64 `json:"raised"`
		Platform    string  `json:"platform"`
		Distributed string  `json:"distributed"`
		Minimum     string  `json:"minimum"`
		Accepting   string  `json:"accepting"`
	} `json:"finance"`
	Dates struct {
		PreIcoStart string `json:"preIcoStart"`
		PreIcoEnd   string `json:"preIcoEnd"`
		IcoStart    string `json:"icoStart"`
		IcoEnd      string `json:"icoEnd"`
	} `json:"dates"`
	Team []struct {
		Name    string `json:"name"`
		Title   string `json:"title"`
		Links   string `json:"links"`
		URL     string `json:"url"`
		Socials []struct {
			Site string `json:"site"`
			URL  string `json:"url"`
		} `json:"socials"`
		Group string `json:"group"`
		Photo string `json:"photo"`
		Iss   string `json:"iss"`
	} `json:"team"`
	Ratings []struct {
		Date    string  `json:"date"`
		Name    string  `json:"name"`
		URL     string  `json:"url"`
		Title   string  `json:"title"`
		Photo   string  `json:"photo"`
		Team    int     `json:"team"`
		Vision  int     `json:"vision"`
		Product int     `json:"product"`
		Profile float64 `json:"profile"`
		Review  string  `json:"review"`
		Weight  string  `json:"weight"`
		Agree   int     `json:"agree"`
	} `json:"ratings"`
	Categories []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
	Exchanges []struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Logo     string  `json:"logo"`
		Price    float64 `json:"price"`
		Currency string  `json:"currency"`
		Roi      string  `json:"roi"`
	} `json:"exchanges"`
	Kyc struct {
		Invited []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"invited"`
		Succeed int `json:"succeed"`
		Failed  int `json:"failed"`
	} `json:"kyc"`
}

type FiltersResponse struct {
	Categories []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
	Platforms []struct {
		Name string `json:"name"`
	} `json:"platforms"`
	Accepting []struct {
		Name string `json:"name"`
	} `json:"accepting"`
	Countries []struct {
		Name string `json:"name"`
	} `json:"countries"`
}

type RatingsResponse struct {
	ICOs    int `json:"icos"`
	Results []struct {
		ID     int     `json:"id"`
		Name   string  `json:"name"`
		URL    string  `json:"url"`
		Logo   string  `json:"logo"`
		Rating float64 `json:"float64"`
	} `json:"results"`
}
