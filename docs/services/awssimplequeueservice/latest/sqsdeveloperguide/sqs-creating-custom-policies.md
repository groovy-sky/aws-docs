# Using custom policies with the Amazon SQS Access Policy Language

To grant basic permissions (such as [`SendMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html) or [`ReceiveMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ReceiveMessage.html))
based only on an AWS account ID, you don’t need to write a custom policy. Instead,
use the Amazon SQS [`AddPermission`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_AddPermission.html) action.

To allow or deny access based on specific conditions, such as request time or the
requester's IP address, you must create a custom Amazon SQS policy and upload it using
the [SetQueueAttributes](../../../../reference/awssimplequeueservice/latest/apireference/api-setqueueattributes.md) action.

###### Topics

- [Access control architecture](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies-architecture.html)

- [Access control process workflow](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies-process-workflow.html)

- [Access Policy Language key concepts](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies-key-concepts.html)

- [Access Policy Language evaluation logic](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies-evaluation-logic.html)

- [Relationships between explicit and default denials](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies-relationships-between-explicit-default-denials.html)

- [Custom policy limitations](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limitations-of-custom-policies.html)

- [Custom Access Policy Language examples](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies-access-policy-examples.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Basic Amazon SQS policy examples

Access control architecture
