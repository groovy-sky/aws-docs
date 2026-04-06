# Receiving S3 on Outposts event notifications by using Amazon CloudWatch Events

You can use CloudWatch Events to create a rule for any Amazon S3 on Outposts API event. When you
create a rule, you can choose to get notified through all supported CloudWatch targets,
including Amazon Simple Queue Service (Amazon SQS), Amazon Simple Notification Service (Amazon SNS), and AWS Lambda. For more information, see
the list of [AWS services that\
can be targets for CloudWatch Events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/WhatIsCloudWatchEvents.html) in the _Amazon CloudWatch Events User Guide_. To choose a target service to work with your
S3 on Outposts, see [Creating a CloudWatch Events rule that triggers on an AWS API call using\
AWS CloudTrail](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/Create-CloudWatch-Events-CloudTrail-Rule.html) in the _Amazon CloudWatch Events User Guide_.

###### Note

For S3 on Outposts object operations, AWS API call events sent by CloudTrail will match
your rules only if you have trails (optionally with event selectors) configured to
receive those events. For more information, see [Working with CloudTrail log files](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create-event-selectors-for-a-trail.html) in the _AWS CloudTrail_
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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatch metrics

CloudTrail logs
