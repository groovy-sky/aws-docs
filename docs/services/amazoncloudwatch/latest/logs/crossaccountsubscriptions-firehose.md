# Cross-account cross-Region log data sharing using Firehose

To share log data across accounts, you need to establish a log data sender and
receiver:

- **Log data sender**—gets the destination
information from the recipient and lets CloudWatch Logs know that it is ready to send
its log events to the specified destination. In the procedures in the rest
of this section, the log data sender is shown with a fictional AWS account
number of 111111111111.

- **Log data recipient**—sets up a destination that
encapsulates a Amazon Kinesis Data Streams stream and lets CloudWatch Logs know that the recipient wants to
receive log data. The recipient then shares the information about this
destination with the sender. In the procedures in the rest of this section,
the log data recipient is shown with a fictional AWS account number of
222222222222.

The example in this section uses a Firehose delivery stream with Amazon S3 storage. You
can also set up Firehose delivery streams with different settings. For more
information, see [Creating a Firehose Delivery\
Stream](https://docs.aws.amazon.com/firehose/latest/dev/basic-create.html).

###### Note

The log group and the destination must be in the same AWS Region. However,
the AWS resource that the destination points to can be located in a different
Region.

###### Note

Firehose subscription filter for a _**same**_
_**account**_ and
_**cross-Region**_ delivery stream
is supported.

###### Topics

- [Step 1: Create a Firehose delivery stream](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CreateFirehoseStream.html)

- [Step 2: Create a destination](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CreateFirehoseStreamDestination.html)

- [Step 3: Add/validate IAM permissions for the cross-account destination](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Subscription-Filter-CrossAccount-Permissions-Firehose.html)

- [Step 4: Create a subscription filter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CreateSubscriptionFilterFirehose.html)

- [Validating the flow of log events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/ValidateLogEventFlowFirehose.html)

- [Modifying destination membership at runtime](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/ModifyDestinationMembershipFirehose.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Step 2: Update the existing destination access policy

Step 1: Create a Firehose delivery stream
