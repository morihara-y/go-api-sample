package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/morihara-y/go-api-sample/application/usecase"
	"github.com/morihara-y/go-api-sample/domain/service"
	"github.com/morihara-y/go-api-sample/infrastracuture/dao"
	"github.com/morihara-y/go-api-sample/interfaces/api/server/handler"
	"github.com/urfave/cli"
)

const title = "Comment Response"
const discription = "You can fetch comments"
const version = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = title
	app.Usage = discription
	app.Version = version

	app.Action = func(context *cli.Context) {
		execute(context.String("port"))
	}
	app.Flags = getAppFlags()

	app.Run(os.Args)
}

func execute(port string) {
	userCommentRepo := dao.NewUserCommentRepository()
	cmntService := service.NewCommentService(userCommentRepo)
	apiUsecase := usecase.NewAPIUsecase(cmntService)

	router := gin.Default()
	var cmntHandler handler.CommentHandler
	v1 := router.Group("/api/v1")
	{
		comments := v1.Group("/comments")
		{
			cmntHandler = handler.NewCmntHandler(apiUsecase)
			comments.GET("/:id", cmntHandler.FetchComment)
		}
	}
	router.Run(":" + port)
}

func getAppFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "8080",
			Usage: "input port number",
		},
	}
}
