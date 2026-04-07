# Enabling high throughput for FIFO queues in Amazon SQS

You can enable high throughput for any new or existing FIFO queue. The feature includes
three new options when you create and edit FIFO queues:

- **Enable high throughput FIFO** – Makes higher throughput
available for messages in the current FIFO queue.

- **Deduplication scope** – Specifies whether deduplication
occurs at the queue or message group level.

- **FIFO throughput limit** – Specifies whether the throughput
quota on messages in the FIFO queue is set at the queue or message group level.

###### To enable high throughput for a FIFO queue (console)

1. Start [creating](creating-sqs-standard-queues.md#step-create-standard-queue) or [editing](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-edit-queue.html) a FIFO queue.

2. When specifying options for the queue, choose **Enable high throughput**
**FIFO**.

Enabling high throughput for FIFO queues sets the related options as follows:

- **Deduplication scope** is set to **Message**
**group**, the required setting for using high throughput for FIFO
queues.

- **FIFO throughput limit** is set to **Per message group**
**ID**, the required setting for using high throughput for FIFO
queues.

If you change any of the settings required for using high throughput for FIFO queues,
normal throughput is in effect for the queue, and deduplication occurs as
specified.

3. Continue specifying all options for the queue. When you finish, choose
    **Create queue** or **Save**.

After creating or editing the FIFO queue, you can [send\
messages](creating-sqs-standard-queues.md#sqs-send-messages) to it and [receive and delete\
messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/step-receive-delete-message.html), all at a higher TPS. For high throughput quotas, see Message throughput in
[Amazon SQS message quotas](quotas-messages.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

High throughput for FIFO queues

Queue and message identifiers
