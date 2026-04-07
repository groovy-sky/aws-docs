# Configuring cost allocation tags for a queue using the Amazon SQS console

To organize and identify your Amazon SQS queues, you can add cost allocation tags. For more
information, see [Amazon SQS cost allocation tags](sqs-queue-tags.md).

- The Tagging tab on the Details page displays the queue's tags.

- You can add or modify tags when [creating](creating-sqs-standard-queues.md#step-create-standard-queue) or [editing](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-edit-queue.html) a queue.

###### To configure tags for an existing queue (console)

1. Open the Amazon SQS console at
    [https://console.aws.amazon.com/sqs/](https://console.aws.amazon.com/sqs).

2. In the navigation pane, choose **Queues**.

3. Choose a queue and choose **Edit**.

4. Scroll to the **Tags** section.

5. Add, modify, or remove the queue tags:
1. To add a tag, choose **Add new tag**, enter a
       **Key** and **Value**, and then choose
       **Add new tag**.

2. To update a tag, change its **Key** and
       **Value**.

3. To remove a tag, choose **Remove** next to its key-value
       pair.
6. When you finish configuring the tags, choose **Save**.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Configuring SSE-KMS for a queue

Subscribing a queue to a topic
