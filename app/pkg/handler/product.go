package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/ncostamagna/go-http-utils/response"
	"github.com/ncostamagna/go-monitoring/app/internal/product"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ctxKey string

const (
	ctxParam  ctxKey = "params"
	ctxHeader ctxKey = "header"
)

func NewHTTPServer(_ context.Context, endpoints product.Endpoints) http.Handler {

	r := gin.Default()

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.Use(ginDecode())

	r.GET("/products", gin.WrapH(httptransport.NewServer(endpoint.Endpoint(endpoints.GetAll), decodeGetAllHandler, encodeResponse, opts...)))
	r.POST("/products", gin.WrapH(httptransport.NewServer(endpoint.Endpoint(endpoints.Store), decodeStoreHandler, encodeResponse, opts...)))

	r.GET("/products/:id", gin.WrapH(httptransport.NewServer(endpoint.Endpoint(endpoints.Get), decodeGetHandler, encodeResponse, opts...)))
	r.PATCH("/products/:id", gin.WrapH(httptransport.NewServer(endpoint.Endpoint(endpoints.Update), decodeUpdateHandler, encodeResponse, opts...)))
	r.DELETE("/products/:id", gin.WrapH(httptransport.NewServer(endpoint.Endpoint(endpoints.Delete), decodeDeleteHandler, encodeResponse, opts...)))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}

func ginDecode() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ctxParam, c.Params)
		ctx = context.WithValue(ctx, ctxHeader, c.Request.Header)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func decodeGetHandler(ctx context.Context, _ *http.Request) (interface{}, error) {
	params := ctx.Value(ctxParam).(gin.Params)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}

	return product.GetReq{
		ID: id,
	}, nil
}

func decodeGetAllHandler(_ context.Context, _ *http.Request) (interface{}, error) {

	var req product.GetAllReq

	return req, nil
}

func decodeStoreHandler(_ context.Context, r *http.Request) (interface{}, error) {
	var req product.StoreReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(err.Error())
	}

	return req, nil
}

func decodeUpdateHandler(_ context.Context, _ *http.Request) (interface{}, error) {

	var req product.UpdateReq

	return req, nil
}

func decodeDeleteHandler(_ context.Context, _ *http.Request) (interface{}, error) {

	var req product.DeleteReq

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(response.Response)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := err.(response.Response)
	w.WriteHeader(resp.StatusCode())
	_ = json.NewEncoder(w).Encode(resp)
}
