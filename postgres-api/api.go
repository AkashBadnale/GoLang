package main

type CreateRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ReadResponse struct {
	Success bool `json:"success"`
	User    struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	} `json:"user"`
}
