## Equivalent code sample in Java

```
  private static void getActivitiesWithRetriesOver (int retryCount) {

    ListOpenWorkflowExecutionsRequest listOpenWorkflowExecutionsRequest =
      ListOpenWorkflowExecutionsRequest.newBuilder()
          .setNamespace(client.getOptions().getNamespace())
          .build();

    ListOpenWorkflowExecutionsResponse listOpenWorkflowExecutionsResponse =
      service.blockingStub().listOpenWorkflowExecutions(listOpenWorkflowExecutionsRequest);

    for (WorkflowExecutionInfo info : listOpenWorkflowExecutionsResponse.getExecutionsList()) {

      DescribeWorkflowExecutionRequest describeWorkflowExecutionRequest =
        DescribeWorkflowExecutionRequest.newBuilder()
            .setNamespace(client.getOptions().getNamespace())
            .setExecution(info.getExecution()).build();

        DescribeWorkflowExecutionResponse describeWorkflowExecutionResponse =
          service.blockingStub().describeWorkflowExecution(describeWorkflowExecutionRequest);

        for (PendingActivityInfo activityInfo : describeWorkflowExecutionResponse.getPendingActivitiesList()) {

          if (activityInfo.getAttempt() > retryCount) {
            System.out.println("Activity Type: " + activityInfo.getActivityType());
            System.out.println("Activity attempt: " + activityInfo.getAttempt());
            System.out.println("Last failure message : " + activityInfo.getLastFailure().getMessage());
            // ...
          }
        }
      }
    }
```
