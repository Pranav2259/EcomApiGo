package orders

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParam struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}

// func jsonify(data any) ([]byte, error) {

// 	jsonData, err := json.MarshalIndent(data, " ", "  ")
// 	if err != nil {
// 		return byte(data), err
// 	}
// 	return jsonData, nil
// }
