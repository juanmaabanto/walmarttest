package response

type ProductResponse struct {
	Id           int64  `json:"id"`
	Brand        string `json:"brand"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	Price        int64  `json:"price"`
	ConDescuento bool   `json:"conDescuento"`
}
