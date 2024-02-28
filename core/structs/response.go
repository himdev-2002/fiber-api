package structs

type DataResponse struct {
	Schema map[string]int `json:"s"`
	Data   []map[int]any  `json:"d"`
	Total  int            `json:"t"`
}

type Paginate struct {
	Page      int `json:"p"`
	PerPage   int `json:"pp"`
	Total     int `json:"t"`
	TotalPage int `json:"tp"`
}

type ResponseWithPagingData struct {
	Code     int       `json:"c"`
	Status   bool      `json:"s"`
	Message  string    `json:"m"`
	Paginate *Paginate `json:"p,omitempty"`
	Data     any       `json:"d"`
}

type ResponseWithoutPagingData struct {
	Code    int    `json:"c"`
	Status  bool   `json:"s"`
	Message string `json:"m"`
	Data    any    `json:"d"`
}

type ResponseWithoutData struct {
	Code    int    `json:"c"`
	Status  bool   `json:"s"`
	Message string `json:"m"`
}
