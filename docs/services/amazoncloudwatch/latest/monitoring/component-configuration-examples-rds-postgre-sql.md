# RDS PostgreSQL

The following example shows a component configurations in JSON format for RDS
PostgreSQL.

```json

{
  "alarmMetrics": [
    {
      "alarmMetricName": "CPUUtilization",
      "monitor": true
    }
  ],
  "logs": [
    {
      "logType": "POSTGRESQL",
      "monitor": true
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RDS Oracle

SAP ASE on Amazon EC2

All content copied from https://docs.aws.amazon.com/.
