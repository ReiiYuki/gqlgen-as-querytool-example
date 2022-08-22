package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"gqlgen-as-querytool/graph"
	"gqlgen-as-querytool/graph/generated"
	"gqlgen-as-querytool/graph/model"
)

type GraphExecutor struct {
	Exec *executor.Executor
}

func (graphExecutor GraphExecutor) Execute(query string, variables map[string]interface{}) *graphql.Response {
	ctx := graphql.StartOperationTrace(context.Background())
	now := graphql.Now()

	rc, err := graphExecutor.Exec.CreateOperationContext(ctx, &graphql.RawParams{
		Query:         query,
		OperationName: "",
		ReadTime: graphql.TraceTiming{
			Start: now,
			End:   now,
		},
		Variables: variables,
	})
	if err != nil {
		return graphExecutor.Exec.DispatchError(ctx, err)
	}

	resp, ctx2 := graphExecutor.Exec.DispatchOperation(ctx, rc)

	return resp(ctx2)
}

type QueryResponse struct {
	Todos []model.Todo
}

type MutationResponse struct {
	CreateTodo model.Todo
}

func convert[T interface{}](value json.RawMessage) T {
	var out T

	json.Unmarshal(value, &out)

	return out
}

func main() {
	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})

	graphExecutor := GraphExecutor{
		executor.New(schema),
	}

	resps1 := graphExecutor.Execute(`
		query Query {
			todos {
				text
			}
		}
	`, nil)

	fmt.Println(convert[QueryResponse](resps1.Data).Todos) // should result as {[]}

	resps2 := graphExecutor.Execute(`
		mutation Mutation ($input: NewTodo!) {
			createTodo(input: $input) {
				text
			}
		}
	`, map[string]interface{}{
		"input": map[string]interface{}{
			"userId": "test",
			"text":   "Test Todo",
		},
	})

	fmt.Println(convert[MutationResponse](resps2.Data).CreateTodo.Text) // should result as Test Todo

	resps3 := graphExecutor.Execute(`
		query Query {
			todos {
				text
			}
		}
	`, nil)

	fmt.Println(convert[QueryResponse](resps3.Data).Todos) // should result as <[{ Test Todo false <nil>}]
}
