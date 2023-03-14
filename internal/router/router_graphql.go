package router

import "github.com/ch3nnn/webstack-go/internal/graph/handler"

func setGraphQLRouter(r *resource) {
	// graphQL 控制器
	gqlHandler := handler.New(r.logger, r.db, r.cache)

	gql := r.mux.Group("/graphql")
	{
		gql.GET("", gqlHandler.Playground())
		gql.POST("/query", gqlHandler.Query())
	}
}
