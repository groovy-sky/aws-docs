# AWS Step Functions

The following example shows a component configurations in JSON format for
AWS Step Functions.

```json

{
  "alarmMetrics": [
    {
      "alarmMetricName": "ExecutionsFailed",
      "monitor": true
    },
    {
      "alarmMetricName": "LambdaFunctionsFailed",
      "monitor": true
    },
    {
      "alarmMetricName": "ProvisionedRefillRate",
      "monitor": true
    }
  ],
  "logs": [
    {
     "logGroupName": "/aws/states/HelloWorld-Logs",
      "logType": "STEP_FUNCTION",
      "monitor": true,
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Network Firewall rule group association

Customer-grouped Amazon EC2 instances

All content copied from https://docs.aws.amazon.com/.
