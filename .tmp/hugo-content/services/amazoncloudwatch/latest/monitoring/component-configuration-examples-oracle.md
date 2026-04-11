# RDS Oracle

The following example shows a component configuration in JSON format for RDS
Oracle.

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
      "logType": "ORACLE_ALERT",
      "monitor": true,
    },
    {
      "logType": "ORACLE_LISTENER",
      "monitor": false
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RDS MariaDB and RDS MySQL

RDS PostgreSQL

All content copied from https://docs.aws.amazon.com/.
