package requests

type CreateItem struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type CreateOrder struct {
	CustomerName string       `json:"customerName"`
	Items        []CreateItem `json:"items"`
}

type UpdateItem struct {
	ItemId      int    `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type UpdateOrder struct {
	CustomerName string       `json:"customerName"`
	Items        []UpdateItem `json:"items"`
}
