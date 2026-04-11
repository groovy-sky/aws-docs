# Amazon Simple Queue Service (SQS)

The following example shows a component configuration in JSON format for
Amazon Simple Queue Service.

```json

{
  "alarmMetrics" : [
    {
      "alarmMetricName" : "ApproximateAgeOfOldestMessage"
    }, {
      "alarmMetricName" : "NumberOfEmptyReceives"
    }
  ],
  "alarms" : [
    {
      "alarmName" : "my_sqs_alarm",
      "severity" : "MEDIUM"
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3 bucket

Amazon SNS topic

All content copied from https://docs.aws.amazon.com/.
