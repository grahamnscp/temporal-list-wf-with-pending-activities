package main

import (
  "log"
  "context"
  "os"

  "go.temporal.io/sdk/client"
  "go.temporal.io/api/workflowservice/v1"

  u "listwf/utils"
)

func main() {

  clientOptions, err := u.LoadClientOptions()
  if err != nil {
    log.Fatalln("Failed to load Temporal Cloud environment:", err)
  }

  client, err := client.Dial(clientOptions)
  if err != nil {
    log.Fatalln("Unable to create client", err)
  }
  defer client.Close()

  namespace := os.Getenv("TEMPORAL_NAMESPACE")

  //
  log.Println("ListOpenWorkflowExecutions on " + namespace + " that have PendingActivities:\n")
  openWorkflows, err := client.ListOpenWorkflow(context.Background(), &workflowservice.ListOpenWorkflowExecutionsRequest{
    Namespace: namespace,
  })
  if err != nil {
    log.Fatalln("fail to list open workflows", err)
  }

  for _, openWorkflow := range openWorkflows.GetExecutions() {
    describe, err := client.DescribeWorkflowExecution(context.Background(), openWorkflow.Execution.WorkflowId, openWorkflow.Execution.RunId)
    if err != nil {
      log.Fatalln("fail to descibe workflow", err)
    }

    pendingActivity := describe.GetPendingActivities()
    if (pendingActivity != nil) {
      log.Println("GetPendingActivities for workflow execution " + openWorkflow.Execution.WorkflowId)

      for _, pendingActivity := range describe.GetPendingActivities() {
        log.Printf("  Activity: '%s' has '%d' Retries, Error: '%v'",  pendingActivity.GetActivityType().Name, pendingActivity.GetAttempt(), pendingActivity.GetLastFailure().Message)
      }
    }
  }
}
