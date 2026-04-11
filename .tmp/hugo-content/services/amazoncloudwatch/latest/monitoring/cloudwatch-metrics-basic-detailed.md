# Basic monitoring and detailed monitoring in CloudWatch

CloudWatch provides two categories of monitoring: _basic monitoring_ and
_detailed monitoring_.

Many AWS services offer basic monitoring by publishing a default set of metrics to CloudWatch
with no charge to customers. By default, when you start using one of these AWS services,
basic monitoring is automatically enabled. For a list of services that offer basic monitoring,
see [AWS services that publish CloudWatch metrics](aws-services-cloudwatch-metrics.md).

Detailed monitoring is offered by only some services. It also incurs charges. To use it
for an AWS service, you must choose to activate it. For more information about pricing, see
[Amazon CloudWatch pricing](http://aws.amazon.com/cloudwatch/pricing).

Detailed monitoring options differ based on the services that offer it. For example, Amazon EC2
detailed monitoring provides more frequent metrics, published at one-minute intervals, instead
of the five-minute intervals used in Amazon EC2 basic monitoring. Detailed monitoring for Amazon S3 and
Amazon Managed Streaming for Apache Kafka means more fine-grained metrics.

In different AWS services, detailed monitoring also has different names. For example, in
Amazon EC2 it is called detailed monitoring, in AWS Elastic Beanstalk it is called enhanced monitoring, and in
Amazon S3 it is called request metrics.

Using detailed monitoring for Amazon EC2 helps you better manage your Amazon EC2 resources, so that
you can find trends and take action faster. For Amazon S3 request metrics are available at
one-minute intervals to help you quickly identify and act on operational issues. On Amazon MSK,
when you enable the `PER_BROKER`, `PER_TOPIC_PER_BROKER`, or
`PER_TOPIC_PER_PARTITION` level monitoring, you get additional metrics that
provide more visibility.

The following table lists the services that offer detailed monitoring. It also includes
links to the documentation for those services that explain more about the detailed monitoring
and provide instructions for how to activate it.

ServiceDocumentation

Amazon API Gateway

[Dimensions for API Gateway metrics](../../../apigateway/latest/developerguide/api-gateway-metrics-and-dimensions.md#api-gateway-metricdimensions)

AWS AppSync

[CloudWatch\
metrics](../../../appsync/latest/devguide/monitoring.md#cw-metrics)

Amazon CloudFront

[Viewing additional CloudFront distribution metrics](../../../amazoncloudfront/latest/developerguide/viewing-cloudfront-metrics.md#monitoring-console.distributions-additional)

Amazon EC2

[Manage detailed\
monitoring for your EC2 instances](../../../ec2/latest/userguide/manage-detailed-monitoring.md)

Elastic Beanstalk

[Enhanced health reporting\
and monitoring](../../../elasticbeanstalk/latest/dg/health-enhanced.md)

Amazon Kinesis Data Streams

[Enhanced Shard-level Metrics](../../../streams/latest/dev/monitoring-with-cloudwatch.md#kinesis-metrics-shard)

AWS Lambda

[Event source mapping metrics](../../../lambda/latest/dg/monitoring-metrics-types.md#event-source-mapping-metrics)

Amazon MSK

[Amazon MSK Metrics for Monitoring with CloudWatch](../../../msk/latest/developerguide/metrics-details.md)

Amazon S3

[Amazon S3 request metrics in CloudWatch](../../../s3/latest/userguide/metrics-dimensions.md#s3-request-cloudwatch-metrics)

Amazon SES

[Collect CloudWatch detailed monitoring metrics using Amazon SES event\
publishing](../../../ses/latest/dg/event-publishing-add-event-destination-cloudwatch.md#event-publishing-add-event-destination-cloudwatch-dimensions).

Additionally, CloudWatch offers out-of-the-box monitoring solutions with more detailed metrics
and pre-created dashboards for some AWS services, as shown in the following table.

ServiceFeature documentation

Lambda

[Lambda\
Insights](lambda-insights.md)

Amazon ECS

[Container Insights for Amazon ECS](deploy-container-insights-ecs.md)

Amazon EKS

[Container Insights for Amazon EKS and Kubernetes](deploy-container-insights-eks.md)

Amazon RDS

[CloudWatch Database Insights](database-insights.md)

Amazon Aurora

[CloudWatch Database Insights](database-insights.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Concepts

Query metrics with CloudWatch Metrics Insights

All content copied from https://docs.aws.amazon.com/.
