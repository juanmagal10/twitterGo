package handlers

import (
	"context"
	"fmt"

	"twitterGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi{
	fmt.Println("voy a procesar "+ ctx.Value(models.Key("path")).(string)+">"+ctx.Value(models.Key("method")).(string))

	var r models.RespApi
	r.Status=400

	switch ctx.Value(models.Key("method")).(string){
	case "POST":
		switch ctx.Value(models.Key("path")).(string){

		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string){
			
		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string){
			
		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string){
			
		}
		//
	}

	r.Message="method invalid"
	return r
}