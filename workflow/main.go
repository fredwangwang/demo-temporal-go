package main

import (
	"context"
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/fredwangwang/demo-temporal-go/shared"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "transfer money wf",
		TaskQueue: shared.TransferMoneyTaskQueue,
	}

	transferDetails := shared.TransferDetails{
		Amount:      123.4,
		FromAccount: "somebdy-01",
		ToAccount:   "huan-01",
		ReferenceID: uuid.NewString(),
	}

	wr, err := c.ExecuteWorkflow(context.Background(), options, getFunctionName(shared.ITrasnferMoney.TransferMoney), transferDetails)
	if err != nil {
		log.Fatalln("unable to start TransferMoney workflow", err)
	}
	printResults(transferDetails, wr.GetID(), wr.GetRunID())
}

func printResults(transferDetails shared.TransferDetails, workflowID, runID string) {
	log.Printf(
		"\nTransfer of $%f from account %s to account %s is processing. ReferenceID: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)
}

func getFunctionName(i interface{}) string {
	if fullName, ok := i.(string); ok {
		return fullName
	}
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	elements := strings.Split(fullName, ".")
	shortName := elements[len(elements)-1]
	// This allows to call activities by method pointer
	// Compiler adds -fm suffix to a function name which has a receiver
	// Note that this works even if struct pointer used to get the function is nil
	// It is possible because nil receivers are allowed.
	// For example:
	// var a *Activities
	// ExecuteActivity(ctx, a.Foo)
	// will call this function which is going to return "Foo"
	return strings.TrimSuffix(shortName, "-fm")
}
