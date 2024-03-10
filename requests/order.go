package requests

type Item struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type Order struct {
	OrderedAt    string `json:"orderedAt"`
	CustomerName string `json:"customerName"`
	Items        []Item `json:"items"`
}
