package views

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type Movie struct {
	Title string `json:"title"`
	Year string `json:"year"`
}
