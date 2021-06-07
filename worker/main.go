package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fredwangwang/demo-temporal-go/shared"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

var counter = 0

func TransferMoney(ctx workflow.Context, transferDetails shared.TransferDetails) error {
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    500,
	}
	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Actvitivy functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failures by default, this is just an example.
		RetryPolicy: retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// An interesting impl... The second argument can be either a string to the function name (RPC) or a reference to a function ptr.
	// If a function ptr is given, essentially it is still translate into strings, and get called on they worker side.
	err := workflow.ExecuteActivity(ctx, Withdraw, transferDetails).Get(ctx, nil)
	if err != nil {
		return err
	}
	log.Println("finished withdraw, sleep 20s")

	// workflow.sleep is recorded as well.
	// saying deposit fails, then only deposit is going to be rerun.
	// workflow.Sleep(ctx, 5*time.Second)
	log.Println("starting deposit")
	err = workflow.ExecuteActivity(ctx, Deposit, transferDetails).Get(ctx, nil)
	if err != nil {
		return err
	}
	log.Println("finished deposit")

	return nil
}

func Withdraw(ctx context.Context, transferDetails shared.TransferDetails) error {
	fmt.Printf(
		"\nWithdrawing $%f from account %s. ReferenceId: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ReferenceID,
	)
	return nil
}

func Deposit(ctx context.Context, transferDetails shared.TransferDetails) error {
	fmt.Printf(
		"\nDepositing $%f into account %s. ReferenceId: %s\n",
		transferDetails.Amount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	// Switch out comments on the return statements to simulate an error
	counter += 1
	if counter%2 == 0 {
		return fmt.Errorf("deposit did not occur due to an issue")
	}
	return nil
}

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Worker and Activity functions
	w := worker.New(c, shared.TransferMoneyTaskQueue, worker.Options{})
	w.RegisterWorkflow(TransferMoney)
	w.RegisterActivity(Withdraw)
	w.RegisterActivity(Deposit)
	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
