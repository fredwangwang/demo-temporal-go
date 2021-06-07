# Demo Temporal using Go SDK

## Run

1. `docker-compose up`
2. `go run ./workflow/main.go` to "execute" (submit) a new workflow
3. `go run ./worker/main.go` to spin up the worker to actually execute the workflow and its actions
4. http://localhost:8088 to check out the execution status

## Take

The most basic understanding: an enhenced RPC framework.

As the workflow (client) only says what workflow to execute with what activities, it is up to the workers (servers) to actually execute those workflows.

Due to the fact it is so generic but handles so many common "workflow" cases, it is a power tool to boost productivity to focus
on the actual business logic instead of redo the plumbing everytime a similar use case arises.

## [Value Props](https://docs.temporal.io/docs/go/run-your-first-app-tutorial/#lore-check)

1. Temporal gives you full visibility in the state of your Workflow and code execution.
1. Temporal maintains the state of your Workflow, even through server outages and errors.
1. Temporal makes it easy to timeout and retry Activity code using options that exist outside of your business logic.
1. Temporal enables you to perform "live debugging" of your business logic while the Workflow is running.

## Docs & Refs

https://docs.temporal.io/docs/go/introduction
https://docs.temporal.io/docs/go/run-your-first-app-tutorial/
https://medium.com/@ijayakantha/microservices-the-saga-pattern-for-distributed-transactions-c489d0ac0247
