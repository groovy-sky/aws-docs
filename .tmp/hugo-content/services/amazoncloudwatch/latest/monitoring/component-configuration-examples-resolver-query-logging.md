# Amazon Route 53 Resolver query logging configuration

The following example shows a component configuration in JSON format for Amazon Route 53 Resolver
query logging configuration.

```json

{
  "logs": [
    {
      "logGroupName": "/resolver-query-log-config/logs",
      "logType": "ROUTE53_RESOLVER_QUERY_LOGS",
      "monitor": true
    }
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Route 53 Resolver endpoint

Amazon S3 bucket

All content copied from https://docs.aws.amazon.com/.
