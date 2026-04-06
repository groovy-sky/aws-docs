# Logging IP traffic using VPC Flow Logs

VPC Flow Logs is a feature that enables you to capture information about the IP traffic
going to and from network interfaces in your VPC. Flow log data can be published to the
following locations: Amazon CloudWatch Logs, Amazon S3, or Amazon Data Firehose. The configured delivery path and
permissions that enable network traffic logs to be sent to a destination like CloudWatch
Logs or S3 are referred to as _subscriptions_. After you
create a flow log, you can retrieve and view the flow log records in the log group, bucket,
or delivery stream that you configured.

Flow logs can help you with a number of tasks, such as:

- Diagnosing overly restrictive security group rules

- Monitoring the traffic that is reaching your instance

- Determining the direction of the traffic to and from the network interfaces

Flow log data is collected outside of the path of your network traffic, and therefore does
not affect network throughput or latency. You can create or delete flow logs without any
risk of impact to network performance.

###### Note

This section only talks about flow logs for VPCs. For information about flow logs for transit gateways introduced in version 6, see [Logging network traffic using Transit Gateway Flow Logs](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-flow-logs.html) in the _Amazon VPC Transit Gateways User Guide_.

###### Contents

- [Flow logs basics](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-basics.html)

- [Flow log records](https://docs.aws.amazon.com/vpc/latest/userguide/flow-log-records.html)

- [Flow log record examples](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-records-examples.html)

- [Flow log limitations](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-limitations.html)

- [Pricing](#flow-logs-pricing)

- [Work with flow logs](https://docs.aws.amazon.com/vpc/latest/userguide/working-with-flow-logs.html)

- [Publish flow logs to CloudWatch Logs](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-cwl.html)

- [Publish flow logs to Amazon S3](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-s3.html)

- [Publish flow logs to Amazon Data Firehose](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-firehose.html)

- [Query flow logs using Amazon Athena](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-athena.html)

- [Troubleshoot VPC Flow Logs](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-troubleshooting.html)

## Pricing

Data ingestion and archival charges for vended logs apply when you publish flow logs.
For more information about pricing when publishing vended logs, open [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing), select
**Logs** and find **Vended Logs**.

To track charges from publishing flow logs, you can apply cost allocation tags
to your destination resource. Thereafter, your AWS cost allocation report includes
usage and costs aggregated by these tags. You can apply tags that represent business
categories (such as cost centers, application names, or owners) to organize your costs.
For more information, see the following:

- [Using Cost Allocation Tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md) in the
_AWS Billing User Guide_

- [Tag log groups in Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html#log-group-tagging) in the _Amazon CloudWatch Logs User Guide_

- [Using cost allocation S3 bucket tags](https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html) in the _Amazon Simple Storage Service User Guide_

- [Tagging Your Delivery Streams](https://docs.aws.amazon.com/firehose/latest/dev/firehose-tagging.html) in the _Amazon Data Firehose Developer Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Monitoring

Flow logs basics
