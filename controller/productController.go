package controller

import (
	"ecommerce_go/products"
	"encoding/json"
	"net/http"
)

// import

// type allProductsHandler struct {
// 	allProductService products.Service
// }

// check params
// call the func
// return
func getProducts(response http.ResponseWriter, request *http.Request) {
	var httpError = products.ProductErrorResponse{
		Code: http.StatusInternalServerError, Message: "Http error.",
	}
	var productsResponse = []products.ProductResponse{}
	productsResponse = products.Service.getProducts()

	if productsResponse == nil {
		dataErrorResponse(response, request, httpError)
	} else {
		response.Header().Set("Content-Type", "application/json")
		// response.Write(productsResponse)
	}

}

func dataErrorResponse(response http.ResponseWriter, request *http.Request, errorMessage products.ProductErrorResponse) {
	httpResponse := &products.ProductErrorResponse{Code: errorMessage.Code, Message: errorMessage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMessage.Code)
	response.Write(jsonResponse)
}
