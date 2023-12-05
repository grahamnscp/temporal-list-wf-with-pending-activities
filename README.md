## Lists workflow pending activities - util to see activities stuck looping

### Example run
```
$ go run list-workflows.go
2023/12/05 16:56:31 ListOpenWorkflowExecutions on grahamnscp-ns.abcde that have PendingActivities:

2023/12/05 16:56:31 GetPendingActivities for workflow execution go-txfr-sorder-payment-46004
2023/12/05 16:56:31   Activity: 'Withdraw' has '24' Retries, Error: 'Banking Service currently unavailable'
2023/12/05 16:56:31 GetPendingActivities for workflow execution go-txfr-webtask-wkfl-28908
2023/12/05 16:56:31   Activity: 'Withdraw' has '25' Retries, Error: 'Banking Service currently unavailable'
```

Ref: [Temporal Community post](https://community.temporal.io/t/activity-retries-and-alerting/3484)
