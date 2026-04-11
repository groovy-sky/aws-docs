This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket TopicConfiguration

A container for specifying the configuration for publication of messages to an Amazon Simple
Notification Service (Amazon SNS) topic when Amazon S3 detects specified events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Event" : String,
  "Filter" : NotificationFilter,
  "Topic" : String
}

```

### YAML

```yaml

  Event: String
  Filter:
    NotificationFilter
  Topic: String

```

## Properties

`Event`

The Amazon S3 bucket event about which to send notifications. For more information, see [Supported Event Types](../../../s3/latest/dev/notificationhowto.md) in
the _Amazon S3 User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

The filtering rules that determine for which objects to send notifications. For example,
you can create a filter so that Amazon S3 sends notifications only when image files with a
`.jpg` extension are added to the bucket.

_Required_: No

_Type_: [NotificationFilter](aws-properties-s3-bucket-notificationfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Topic`

The Amazon Resource Name (ARN) of the Amazon SNS topic to which Amazon S3 publishes a message when it
detects events of the specified type.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Receive S3 bucket notifications to an SNS topic

The following example template shows an Amazon S3 bucket with a notification
configuration that sends an event to the specified SNS topic when S3 has lost all replicas
of an object.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "Private",
                "NotificationConfiguration": {
                    "TopicConfigurations": [
                        {
                            "Topic": "arn:aws:sns:us-east-1:123456789012:TestTopic",
                            "Event": "s3:ReducedRedundancyLostObject"
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a notification configuration."
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      AccessControl: Private
      NotificationConfiguration:
        TopicConfigurations:
          - Topic: 'arn:aws:sns:us-east-1:123456789012:TestTopic'
            Event: 's3:ReducedRedundancyLostObject'
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with a notification configuration.
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tiering

Transition

All content copied from https://docs.aws.amazon.com/.
