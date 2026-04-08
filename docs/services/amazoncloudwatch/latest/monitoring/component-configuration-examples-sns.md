# Amazon SNS topic

The following example shows a component configuration in JSON format for Amazon SNS
topic.

```json

{
  "alarmMetrics": [
    {
      "alarmMetricName": "NumberOfNotificationsFailed",
      "monitor": true
    },
    {
      "alarmMetricName": "NumberOfNotificationsFilteredOut-InvalidAttributes",
      "monitor": true
    },
    {
      "alarmMetricName": "NumberOfNotificationsFilteredOut-NoMessageAttributes",
      "monitor": true
    },
    {
      "alarmMetricName": "NumberOfNotificationsFailedToRedriveToDlq",
      "monitor": true
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Simple Queue Service (SQS)

Amazon Virtual Private Cloud (Amazon VPC)

All content copied from https://docs.aws.amazon.com/.
