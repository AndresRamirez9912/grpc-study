package models

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Exam struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
