# Amazon Relational Database Service (RDS) instance

The following example shows a component configuration in JSON format for an
Amazon RDS instance.

```json

{
  "alarmMetrics" : [
    {
      "alarmMetricName" : "BurstBalance",
      "monitor" : true
    }, {
      "alarmMetricName" : "WriteThroughput",
      "monitor" : false
    }
  ],

  "alarms" : [
    {
      "alarmName" : "my_rds_instance_alarm",
      "severity" : "MEDIUM"
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Relational Database Service (RDS) Aurora MySQL

Amazon Route 53 health check

All content copied from https://docs.aws.amazon.com/.
