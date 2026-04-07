# Creating alarms for dead-letter queues using Amazon CloudWatch

Set up a CloudWatch alarm to monitor messages in a dead-letter queue using the [ApproximateNumberOfMessagesVisible](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-available-cloudwatch-metrics.html) metric. For detailed
instructions, see [Creating CloudWatch alarms for Amazon SQS metrics](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/set-cloudwatch-alarms-for-metrics.html). When the alarm triggers, indicating
messages have been moved to the dead-letter queue, you can [poll](sqs-short-and-long-polling.md) the queue to review and retrieve
them.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudTrail update and permission
requirements

Message metadata for Amazon SQS
