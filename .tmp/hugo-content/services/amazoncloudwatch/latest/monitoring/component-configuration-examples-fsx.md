# Amazon FSx

The following example shows a component configuration in JSON format for
Amazon FSx.

```json

{
  "alarmMetrics": [
    {
      "alarmMetricName": "DataReadBytes",
      "monitor": true
    },
    {
      "alarmMetricName": "DataWriteBytes",
      "monitor": true
    },
    {
      "alarmMetricName": "DataReadOperations",
      "monitor": true
    },
    {
      "alarmMetricName": "DataWriteOperations",
      "monitor": true
    },
    {
      "alarmMetricName": "MetadataOperations",
      "monitor": true
    },
    {
      "alarmMetricName": "FreeStorageCapacity",
      "monitor": true
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Elastic File System (Amazon EFS)

Amazon Relational Database Service (RDS) Aurora MySQL

All content copied from https://docs.aws.amazon.com/.
