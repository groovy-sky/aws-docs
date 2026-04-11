# Receiving S3 on Outposts event notifications by using Amazon CloudWatch Events

You can use CloudWatch Events to create a rule for any Amazon S3 on Outposts API event. When you
create a rule, you can choose to get notified through all supported CloudWatch targets,
including Amazon Simple Queue Service (Amazon SQS), Amazon Simple Notification Service (Amazon SNS), and AWS Lambda. For more information, see
the list of [AWS services that\
can be targets for CloudWatch Events](../../../amazoncloudwatch/latest/events/whatiscloudwatchevents.md) in the _Amazon CloudWatch Events User Guide_. To choose a target service to work with your
S3 on Outposts, see [Creating a CloudWatch Events rule that triggers on an AWS API call using\
AWS CloudTrail](../../../amazoncloudwatch/latest/events/create-cloudwatch-events-cloudtrail-rule.md) in the _Amazon CloudWatch Events User Guide_.

###### Note

For S3 on Outposts object operations, AWS API call events sent by CloudTrail will match
your rules only if you have trails (optionally with event selectors) configured to
receive those events. For more information, see [Working with CloudTrail log files](../../../awscloudtrail/latest/userguide/create-event-selectors-for-a-trail.md) in the _AWS CloudTrail_
_User Guide_.

###### Example

The following is a sample rule for the `DeleteObject` operation. To use
this sample rule, replace `amzn-s3-demo-bucket1` with the name of your S3 on Outposts
bucket.

```JSON

{
  "source": [
    "aws.s3-outposts"
  ],
  "detail-type": [
    "AWS API call through CloudTrail"
  ],
  "detail": {
    "eventSource": [
      "s3-outposts.amazonaws.com"
    ],
    "eventName": [
      "DeleteObject"
    ],
    "requestParameters": {
      "bucketName": [
        "amzn-s3-demo-bucket1"
      ]
    }
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatch metrics

CloudTrail logs

All content copied from https://docs.aws.amazon.com/.
