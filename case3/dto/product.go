package dto

type (
	ProductListResponseElement struct {
		ID           int    `json:"id"`
		DeviceName   string `json:"name"`
		Manufacturer string `json:"manufacturer"`
		Price        int    `json:"price"`
		Image        string `json:"image"`
		Stock        int    `json:"stock"`
	}

	ProductListResponse []ProductListResponseElement
)
