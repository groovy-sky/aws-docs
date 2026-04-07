# Editing an Amazon SQS queue using the console

You can use the Amazon SQS console to edit queue configuration parameters (except the queue
type) and modify or remove features as needed.

###### To edit an Amazon SQS queue (console)

1. Open the [Queues\
    page](https://console.aws.amazon.com/sqs) of the Amazon SQS console.

2. Select a queue, and then choose **Edit**.

3. (Optional) Under **Configuration**, update the queue's [configuration\
    parameters](sqs-configure-queue-parameters.md).

4. (Optional) To update the [access\
    policy](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-add-permissions.html), under **Access policy**, modify the
    **JSON policy**.

5. (Optional) To update a dead-letter queue [redrive allow\
    policy](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-dead-letter-queue-redrive.html), expand **Redrive allow policy**.

6. (Optional) To update or remove [encryption](sqs-configure-sse-existing-queue.md), expand
    **Encryption**.

7. (Optional) To add, update, or remove a [dead-letter queue](sqs-configure-dead-letter-queue.md) (which
    allows you to receive undeliverable messages), expand **Dead-letter**
**queue**.

8. (Optional) To add, update, or remove the [tags](sqs-configure-tag-queue.md) for the queue, expand
    **Tags**.

9. Choose **Save**.

- The console displays the **Details** page for the
queue.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Managing a queue

Receiving and deleting a message
