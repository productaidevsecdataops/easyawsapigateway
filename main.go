package easyawsapigateway

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

func AllowCors() map[string]string {
	return map[string]string{
		"X-Requested-With":             "*",
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,x-requested-with",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST,GET,DELETE,PUT,OPTIONS",
		"Content-Type":                 "application/json",
	}
}

// Returns a API gateway appropriate error to client
func ServerError(err error, hint string) events.APIGatewayProxyResponse {
	s := fmt.Sprintf(`{"error":"%s","hint":"%s"}`, err.Error(), hint)

	log.Println("ServerError reached with: ", s)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       s,
		Headers:    AllowCors(),
	}

}

// Returns a API unauthorized to client
func RequestUnauthorized(hint string) events.APIGatewayProxyResponse {
	//s := fmt.Sprintf(`{"hint":"%s"}`, hint)
	s := fmt.Sprintf(`{"error":"%s", "hint":"%s"}`, hint, hint)
	log.Println("ServerError reached with: ", s)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusUnauthorized,
		Body:       s,
		Headers:    AllowCors(),
	}

}

// Returns a API gateway appropriate error to client
func NotFound(m string) events.APIGatewayProxyResponse {

	s := fmt.Sprintf(`{"error":"%s", "hint":"%s"}`, m, m)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       s,
		Headers:    AllowCors(),
	}

}

// Returns a API gateway appropriate error to client
func InvalidParams(m string) events.APIGatewayProxyResponse {
	s := fmt.Sprintf(`{"error":"%s", "hint":"%s"}`, m, m)

	log.Println("InvalidParams reached with :", s)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       s,
		Headers:    AllowCors(),
	}

}

func Success(s string) events.APIGatewayProxyResponse {
	log.Println(s)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       s,
		Headers:    AllowCors(),
	}
}

func SuccessEmptyArray() events.APIGatewayProxyResponse {

	ba, _ := json.Marshal([]string{})

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(ba),
		Headers:    AllowCors(),
	}
}

func SuccessNoContent() events.APIGatewayProxyResponse {
	log.Println("SuccessNoContent Reached")
	// At this point, we have successfully updated the record, so we will return a 204 with an empty body string
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
		Body:       "",
		Headers:    AllowCors(),
	}
}

func SuccessCreated() events.APIGatewayProxyResponse {
	log.Println("SuccessCreated Reached")
	// At this point, we have successfully updated the record, so we will return a 204 with an empty body string
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "",
		Headers:    AllowCors(),
	}
}

func MethodNotAllowed() events.APIGatewayProxyResponse {

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       string(""),
		Headers:    AllowCors(),
	}

}

func SuccessWithAudio(s string) events.APIGatewayProxyResponse {
	log.Println(s)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       s,
		Headers:    AllowCorsWithAudio(),
	}
}

func AllowCorsWithAudio() map[string]string {
	return map[string]string{
		"X-Requested-With":             "*",
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,x-requested-with",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST,GET,DELETE,PUT,OPTIONS",
		"Content-Type":                 "audio/x-wav",
	}
}
