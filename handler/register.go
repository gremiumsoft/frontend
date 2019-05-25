package handler

import (
	"context"
	"frontend/gen/models"
	"frontend/gen/restapi/operations"
	"frontend/gen/restapi/operations/quiz"
	"github.com/go-openapi/runtime/middleware"
	pb "github.com/gremiumsoft/api-models-go/quizservice"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
)

const quizServiceAddress = "quizservice-service.default.svc.cluster.local:8000"

// TODO(JN): this should be moved from here and some additional fields
// should be added, request id for ex.
type Params struct {
	Slog *zap.SugaredLogger
}

func Register(p *Params, api *operations.GremiumAPI) {
	api.QuizGetQuizesHandler = quiz.GetQuizesHandler(
		func(p *Params) quiz.GetQuizesHandlerFunc {
			return func(params quiz.GetQuizesParams) middleware.Responder {
				return getQuizzesHandler(p, params)
			}
		}(p))
}

func getQuizzesHandler(p *Params, params quiz.GetQuizesParams) middleware.Responder {
	p.Slog.Info("getQuizzes handler")

	// TODO(JN) we should not create a new connection for every request
	conn, err := grpc.Dial(quizServiceAddress, grpc.WithInsecure())
	if err != nil {
		return retServerError(err)
	}
	defer conn.Close()

	client := pb.NewQuizServiceClient(conn)
	ctx := context.Background()
	qqlist, err := client.GetQuestions(ctx, &pb.QuizRequest{})
	if err != nil {
		return retServerError(err)
	}

	qq := make([]*models.QuizQuestion, 0)
	for _, q := range qqlist.QuizQuestion {
		qq = append(qq, NewQuizQuestion(q))
	}
	return quiz.NewGetQuizesOK().WithPayload(qq)
}

func retServerError(err error) middleware.Responder {
	return quiz.NewGetQuizesDefault(http.StatusInternalServerError).
		WithPayload(&models.Error{
			Message: err.Error(),
		})
}

func NewQuizQuestion(qq *pb.QuizQuestion) *models.QuizQuestion {
	return &models.QuizQuestion{
		ID:               &qq.Id,
		Question:         &qq.Question,
		Answers:          qq.Answers,
		CorrectAnswerIdx: &qq.CorrectAnswerIdx,
	}
}
