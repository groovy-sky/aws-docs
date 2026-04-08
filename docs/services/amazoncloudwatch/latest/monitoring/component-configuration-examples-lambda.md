# AWS Lambda Function

The following example shows a component configuration in JSON format for
AWS Lambda Function.

```json

{
  "alarmMetrics": [
    {
      "alarmMetricName": "Errors",
      "monitor": true
    },
    {
      "alarmMetricName": "Throttles",
      "monitor": true
    },
    {
      "alarmMetricName": "IteratorAge",
      "monitor": true
    },
    {
      "alarmMetricName": "Duration",
      "monitor": true
    }
  ],
  "logs": [
    {
      "logType": "DEFAULT",
      "monitor": true
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Application Elastic Load Balancing

AWS Network Firewall rule group

All content copied from https://docs.aws.amazon.com/.
