package products

type Product struct {
	Id          int
	Name        string
	Image_url   string
	Description string
}

type ProductRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Image_url   string `json:"image_url"`
	Description string `json:"description"`
}

type ProductResponse struct {
	// Id          int    `json:"id"`
	Name string `json:"name"`
	// Image_url   string `json:"image_url"`
	// Description string `json:"description"`
}

type ProductErrorResponse struct {
	Code    int
	Message string
}
