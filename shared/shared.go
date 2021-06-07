package shared

import "go.temporal.io/sdk/workflow"

const TransferMoneyTaskQueue = "TRANSFER_MONEY_TASK_QUEUE"

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}

type ITrasnferMoney interface {
	TransferMoney(ctx workflow.Context, transferDetails TransferDetails)
}
