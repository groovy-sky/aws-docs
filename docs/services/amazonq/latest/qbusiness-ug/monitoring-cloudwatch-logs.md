# Monitoring Amazon Q Business user conversations with Amazon CloudWatch Logs

You can use [Amazon CloudWatch Logs](../../../amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md) to
deliver user conversations and response feedback in Amazon Q Business for you to analyze. These logs can be
delivered to multiple destinations, such as Amazon CloudWatch, Amazon S3, or Amazon Data Firehose (standard rates apply).
We recommend that you set up conversation and feedback logging within five minutes of creating
your Amazon Q Business application environment.

The following are examples of tasks you can complete with logs from conversations and
response feedback in Amazon Q Business:

- Identify common user queries and pain points by reviewing the chat message
content.

- Identify number of system-generated messages that have hallucination.

- Monitor the quality of responses by looking at metrics like
`isMessageWithNoAnswer`.

- Understand user sentiment and satisfaction by analyzing the feedback data, including
comments and usefulness ratings.

- Generate custom dashboards and reports to track key metrics and trends over time.

Amazon Q Business supports the `EVENT_LOGS` log type that tracks the specifics of
conversations in an application. You can use `EVENT_LOGS` to monitor Amazon Q Business in all
AWS regions where Amazon Q Business is offered. For more information about the AWS Regions and
endpoints currently supported by Amazon Q Business, see [Supported Regions](quotas-regions.md#regions).

Logs from conversations might include sensitive or personally identifiable data passed in
the chats. You can filter out this information from your logs with the Amazon Q Business console. Or you
can mask this data on your logs using CloudWatch Logs masking policies. For more information, see [Help\
protect sensitive log data with masking](../../../amazoncloudwatch/latest/logs/mask-sensitive-log-data.md).

###### Topics

- [Amazon Q Business chat message and feedback log examples](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cw-log-examples.html)

- [Permissions for monitoring Amazon Q Business with Amazon CloudWatch Logs](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cw-logs-permissions.html)

- [Enabling Amazon Q Business user conversation logging](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cw-logs-enable-logging.html)

- [Amazon Q Business conversation log query examples](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cw-logs-common-queries.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Q Apps metrics

Log examples
