# API Gateway REST API stages

The following example shows a component configuration in JSON format for API Gateway
REST API stages.

```json

{
     "alarmMetrics" : [
         {
             "alarmMetricName" : "4XXError",
             "monitor" : true
         },
         {
             "alarmMetricName" : "5XXError",
             "monitor" : true
         }
     ],
    "logs" : [
        {
            "logType" : "API_GATEWAY_EXECUTION",
            "monitor" : true
        },
        {
            "logType" : "API_GATEWAY_ACCESS",
            "monitor" : true
        }
    ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon VPC Network Address Translation (NAT) gateways

Application Elastic Load Balancing

All content copied from https://docs.aws.amazon.com/.
