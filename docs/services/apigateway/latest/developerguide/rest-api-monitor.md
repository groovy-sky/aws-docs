# Monitor REST APIs in API Gateway

In this section, you can learn how to monitor your API by using CloudWatch metrics, CloudWatch Logs,
Firehose, and AWS X-Ray. By combining CloudWatch execution logs and CloudWatch metrics, you can log errors
and execution traces, and monitor your API's performance. You might also want to log API
calls to Firehose. You can also use AWS X-Ray to trace calls through the downstream services
that make up your API.

###### Note

API Gateway might not generate logs and metrics in the following cases:

- 413 Request Entity Too Large errors

- 431 Request Header Fields Too Large errors

- Excessive 429 Too Many Requests errors

- 400 series errors from requests sent to a custom domain that has no API mapping

- 500 series errors caused by internal failures

API Gateway will not generate logs and metrics when testing a REST API method. The CloudWatch entries are simulated. For
more information, see [Use the API Gateway console to test a REST API method](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-test-method.html).

###### Topics

- [Monitor REST API execution with Amazon CloudWatch metrics](https://docs.aws.amazon.com/apigateway/latest/developerguide/monitoring-cloudwatch.html)

- [Set up CloudWatch logging for REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)

- [Log REST API calls to Amazon Data Firehose in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-logging-to-kinesis.html)

- [Variables for access logging for API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-variables-for-access-logging.html)

- [Trace user requests to REST APIs using X-Ray in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Invoke a private API

CloudWatch metrics
