package webhooks

type List struct {
	Meta struct {
		Count      int         `json:"count"`
		PageCount  int         `json:"pageCount"`
		TotalCount int         `json:"totalCount"`
		Next       interface{} `json:"next"`
		Previous   interface{} `json:"previous"`
		Self       string      `json:"self"`
		First      string      `json:"first"`
		Last       string      `json:"last"`
	} `json:"meta"`
	Data []struct {
		ID         int    `json:"id"`
		Event      string `json:"event"`
		Target     string `json:"target"`
		Enabled    bool   `json:"enabled"`
		Conditions []struct {
			Key       string `json:"key"`
			Condition string `json:"condition"`
			Value     string `json:"value"`
		} `json:"conditions"`
	} `json:"data"`
}

type Retrieve struct {
	Data struct {
		ID         int    `json:"id"`
		Event      string `json:"event"`
		Target     string `json:"target"`
		Enabled    bool   `json:"enabled"`
		Conditions []struct {
			Key       string `json:"key"`
			Condition string `json:"condition"`
			Value     string `json:"value"`
		} `json:"conditions"`
	} `json:"data"`
}

type CreateBody struct {
	Event      string      `json:"event"`
	Target     string      `json:"target"`
	Enabled    string      `json:"enabled"`
	Conditions []Condition `json:"Conditions"`
}

type Create struct {
	Data struct {
		ID         int    `json:"id"`
		Event      string `json:"event"`
		Target     string `json:"target"`
		Token      string `json:"token"`
		Enabled    bool   `json:"enabled"`
		Conditions []struct {
			Key       string `json:"key"`
			Condition string `json:"condition"`
			Value     string `json:"value"`
		} `json:"conditions"`
	} `json:"data"`
}
