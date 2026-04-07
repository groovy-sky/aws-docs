# Amazon SQS message quotas

The following table lists quotas related to messages.

QuotaDescriptionBatched message IDA batched message ID can have up to 80 characters. The following characters are accepted: alphanumeric characters, hyphens ( `-`), and underscores ( `_`).Message attributesA message can contain up to 10 metadata attributes.Message batch

A single message batch request can include a maximum of 10 messages. For more
information, see [Configuring AmazonSQSBufferedAsyncClient](sqs-client-side-buffering-request-batching.md#configuring-buffered-async-client) in the [Amazon SQS batch actions](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-batch-api-actions.html)
section.

Message content

A message can include only XML, JSON, and unformatted text. The following Unicode
characters are allowed: `#x9` \| `#xA` \| `#xD` \|
`#x20` to `#xD7FF` \| `#xE000` to
`#xFFFD` \| `#x10000` to `#x10FFFF`

Any characters not included in this list are rejected. For more information, see
the [W3C specification for\
characters](https://www.w3.org/TR/REC-xml).

Message group ID

`MessageGroupId` is required for FIFO queues.
If you don't provide a `MessageGroupId` when sending a message to a FIFO queue, the action fails.
In standard queues, using `MessageGroupId` enables [fair queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-fair-queues.html).
We recommend that you include a `MessageGroupId` in all messages when using fair queues.

The length of
`MessageGroupId` is 128 characters. Valid values: alphanumeric characters
and punctuation ``(!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~)``.

Message retentionBy default, a message is retained for 4 days. The minimum is 60 seconds (1 minute).
The maximum is 1,209,600 seconds (14 days).Message throughput

**[Standard\**
**queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/standard-queues.html)**

Standard queues support a very high, nearly unlimited number of API calls per second, per action
( [`SendMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html), [`ReceiveMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ReceiveMessage.html), or [`DeleteMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessage.html)). This high throughput makes them ideal for use cases that require processing large volumes of messages quickly, such as real-time data streaming or large-scale applications. While standard queues scale automatically with demand, it is essential to monitor usage patterns to ensure optimal performance, especially in regions with higher workloads.

**[FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-fifo-queues.html)**

- Each partition in a FIFO queue is limited to 300 transactions per second, per
API action ( `SendMessage`, `ReceiveMessage`, and
`DeleteMessage`). This limit applies specifically to non-high
throughput mode. By switching to high throughput mode, you can surpass this
default limit. To
enable high-throughput mode, see [Enabling high throughput for FIFO queues in Amazon SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/enable-high-throughput-fifo.html).

- If you use batching, non-high throughput FIFO queues support up to 3,000
messages per second, per API action ( `SendMessage`,
`ReceiveMessage`, and `DeleteMessage`). The 3,000 messages
per second represent 300 API calls, each with a batch of 10 messages.

**[High throughput for\**
**FIFO queues](high-throughput-fifo.md)**

Amazon SQS FIFO limits are based on the number of API requests, not message limits. For
high throughput mode, these API request limits are as follows:

**Transaction throughput limits (Non-batching API**
**calls)**

These limits define how frequently each API operation (such as [SendMessage](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html), [ReceiveMessage](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ReceiveMessage.html), or [DeleteMessage](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessage.html)) can be performed independently, ensuring efficient system
performance within the allowed transactions per second (TPS).

The following limits are based on non-batched API calls:

- US East (N. Virginia), US West (Oregon), and Europe (Ireland): Up to 70,000
transactions per second (TPS).

- US East (Ohio) and Europe (Frankfurt): Up to 19,000 TPS.

- Asia Pacific (Mumbai), Asia Pacific (Singapore), Asia Pacific (Sydney), and
Asia Pacific (Tokyo): Up to 9,000 TPS.

- Europe (London) and South America (São Paulo): Up to 4,500 TPS.

- All other AWS Regions: Default throughput of 2,400 TPS.

**Maximizing throughput with batching**

Processes multiple messages in a single API call, which significantly increasing
efficiency. Instead of handling each message individually, batching allows you to
send, receive, or delete up to 10 messages in a single API request. This reduces the
total number of API calls, allowing you to process more messages per second while
staying within the transaction limits (TPS) for the region, maximizing throughput and
system performance. For more information, see [Increasing throughput using horizontal scaling and action batching with Amazon SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-throughput-horizontal-scaling-and-batching.html).

The following limits are based on batched API calls:

- US East (N. Virginia), US West (Oregon), and Europe (Ireland): Up to 700,000
messages per second (10x the non-batch limit of 70,000 TPS).

- US East (Ohio) and Europe (Frankfurt): Up to 190,000 messages per
second.

- Asia Pacific (Mumbai), Asia Pacific (Singapore), Asia Pacific (Sydney), and
Asia Pacific (Tokyo): Up to 90,000 messages per second.

- Europe (London) and South America (São Paulo): Up to 45,000 messages per
second.

- All other AWS Regions: Up to 24,000 messages per second.

**Optimizing throughput beyond batching**

While batching can greatly increase throughput, it’s important to consider other
strategies for optimizing FIFO performance:

- **Distribute messages across multiple message group**
**IDs** – Since messages within a single group are processed
sequentially, distributing your workload across multiple message groups allows for
better parallelism and higher overall throughput. For more information, see [Partitions and data distribution for high throughput for SQS FIFO queues](high-throughput-fifo.md#partitions-and-data-distribution).

- **Efficient use of API calls** – Minimize
unnecessary API calls, such as frequent visibility changes or repeated message
deletions, to optimize the use of your available TPS and improve
efficiency.

- **Use long poll receives** – Utilize long polling
by setting [`WaitTimeSeconds`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ReceiveMessage.html#SQS-ReceiveMessage-request-WaitTimeSeconds) in your receive requests to reduce empty
responses when no messages are available, lowering unnecessary API calls and
making better use of your TPS quota.

- **Requesting throughput increases** – If
your application requires throughput higher than the default limits, request an increase
using the [Service Quotas](https://console.aws.amazon.com/servicequotas/home/services/sqs/quotas) console. This can be necessary for high-demand workloads
or in regions with lower default limits. To
enable high-throughput mode, see [Enabling high throughput for FIFO queues in Amazon SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/enable-high-throughput-fifo.html).

Message timerThe default (minimum) delay for a message is 0 seconds. The maximum is 15 minutes.Message size

The minimum message size is 1 byte (1 character). The maximum is 1,048,576 bytes
(1 MiB).

To send messages larger than 1 MiB, you can use the [Amazon SQS Extended Client Library for Java](https://github.com/awslabs/amazon-sqs-java-extended-client-lib) and the [Amazon SQS\
Extended Client Library for Python](https://github.com/awslabs/amazon-sqs-python-extended-client-lib). This library allows you to send an Amazon SQS
message that contains a reference to a message payload in Amazon S3. The maximum payload
size is 2 GB.

###### Note

This extended library works only for synchronous clients.

Message visibility timeoutThe default visibility timeout for a message is 30 seconds. The minimum is 0 seconds. The maximum is 12 hours.Policy informationThe maximum quota is 8,192 bytes, 20 statements, 50 principals, or 10 conditions.
For more information, see [Amazon SQS policy quotas](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-policies.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Standard queue quotas

Policy quotas
