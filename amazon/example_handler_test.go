package amazon_test

import (
	"github.com/Golamu/core/amazon"
)

type RequestBody struct {
	Message string `json:"message"`
}

type Result struct {
	ItWorked string `json:"it_worked"`
}

func ExampleStartHTTP() {

	// endpoint that just gives back a JSON body and a 201 code
	endpoint := func(ctx *amazon.Context) error {
		res := ctx.GetResponse()

		result := Result{ItWorked: "Yes, it sure did"}
		res.SetBody(result)
		res.SetCode(201) // whatever HTTP code you can think of

		return nil
	}

	// this will boot with the lambda.Start that the docs tell you
	// to use and automatically give an APIGatewayProxyResponse for
	// you this goes in the `func main` of your serverless function
	amazon.StartHTTP(endpoint)
}
