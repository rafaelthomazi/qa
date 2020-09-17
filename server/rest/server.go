package rest

import (
	"github.com/rafaelthomazi/qa/qa/service"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	basePath = "/api/qa"
)

// NewServer creates a REST server instance
func NewServer(svc service.Service, port string, logger *zap.Logger) *http.Server {
	endpoints := makeEndpoints(svc, logger)
	r := makeHandlers(endpoints)

	server := &http.Server{
		Addr: port,
		Handler: handlers.CORS(
			handlers.AllowedHeaders(
				[]string{"Content-Type"},
			),
			handlers.AllowedMethods(
				[]string{"GET", "POST", "PUT", "DELETE"},
			),
		)(r),
	}

	return server
}

func makeHandlers(endpoints Endpoints) *mux.Router {
	getQuestionHandler := kithttp.NewServer(
		endpoints.GetQuestionEndpoint,
		decodeIDParamRequest,
		encodeResponse,
	)

	getQuestionsHandler := kithttp.NewServer(
		endpoints.GetQuestionsEndpoint,
		decodeBlankRequest,
		encodeResponse,
	)

	createQuestionHandler := kithttp.NewServer(
		endpoints.CreateQuestionEndpoint,
		decodeQuestionRequest,
		encodeResponse,
	)

	updateQuestionHandler := kithttp.NewServer(
		endpoints.UpdateQuestionEndpoint,
		decodeUpdateQuestionRequest,
		encodeResponse,
	)

	deleteQuestionHandler := kithttp.NewServer(
		endpoints.DeleteQuestionEndpoint,
		decodeIDParamRequest,
		encodeResponse,
	)

	r := mux.NewRouter()

	// Questions Routes
	r.Methods("GET").Path(basePath + "/questions/{id}").Handler(getQuestionHandler)
	r.Methods("GET").Path(basePath + "/questions").Handler(getQuestionsHandler)
	r.Methods("POST").Path(basePath + "/questions").Handler(createQuestionHandler)
	r.Methods("PUT").Path(basePath + "/questions/{id}").Handler(updateQuestionHandler)
	r.Methods("DELETE").Path(basePath + "/questions/{id}").Handler(deleteQuestionHandler)

	return r
}
