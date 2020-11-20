package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"go-pythagoras/pythagoras"
)

// PythagorasRequest is the input into the lambda function.
type PythagorasRequest struct {
	Side1 float64 `json:"side1"`
	Side2 float64 `json:"side2"`
}

// PythagorasResponse is the output from the lambda function.
type PythagorasResponse struct {
	Hypotenuse float64 `json:"hypotenuse"`
	Area       float64 `json:"area"`
	Perimeter  float64 `json:"perimeter"`
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var pythagorasRequest PythagorasRequest
	var err1, err2 error

	switch request.HTTPMethod {
	case "GET":
		pythagorasRequest.Side1, err1 = strconv.ParseFloat(request.QueryStringParameters["side1"], 64)
		pythagorasRequest.Side2, err2 = strconv.ParseFloat(request.QueryStringParameters["side2"], 64)

		if err1 != nil || err2 != nil {
			return apiResponse(http.StatusBadRequest, nil)
		}
		break

	case "POST":
		err1 = json.Unmarshal([]byte(request.Body), &pythagorasRequest)

		if err1 != nil {
			return apiResponse(http.StatusBadRequest, nil)
		}
		break
	}

	hypotenuse := pythagoras.GetHypotenuse(pythagorasRequest.Side1, pythagorasRequest.Side2)
	area := pythagoras.GetArea(pythagorasRequest.Side1, pythagorasRequest.Side2)
	perimeter := pythagoras.GetPerimeter(pythagorasRequest.Side1, pythagorasRequest.Side2)

	return apiResponse(http.StatusOK, PythagorasResponse{
		Hypotenuse: hypotenuse,
		Area:       area,
		Perimeter:  perimeter})
}

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}
