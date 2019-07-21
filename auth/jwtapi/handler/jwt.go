package handler

import (
	"context"
	"encoding/json"

	"github.com/micro/go-micro/util/log"

	"github.com/dmitrymomot/go-micro/auth/jwtapi/client"
	jwt "github.com/dmitrymomot/go-micro/auth/jwtsrv/proto/jwt"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
)

// Jwt service struct
type Jwt struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Call is called by the API as /jwt/call with post body {"name": "foo"}
func (e *Jwt) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Jwt.Call request")

	// extract the client from the context
	jwtClient, ok := client.JwtFromContext(ctx)
	if !ok {
		return errors.InternalServerError("jwtapi.jwtsrv.new_token", "jwt client not found")
	}

	// make request
	response, err := jwtClient.NewToken(ctx, &jwt.NewTokenRequest{
		UserId: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("jwtapi.jwtsrv.new_token", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
