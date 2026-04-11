This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket NotificationConfiguration

Describes the notification configuration for an Amazon S3 bucket.

###### Note

If you create the target resource and related permissions in the same template, you
might have a circular dependency.

For example, you might use the `AWS::Lambda::Permission` resource to grant
the bucket permission to invoke an AWS Lambda function. However, AWS CloudFormation can't create the bucket until the bucket has permission to
invoke the function (AWS CloudFormation checks whether the bucket can
invoke the function). If you're using Refs to pass the bucket name, this leads to a circular
dependency.

To avoid this dependency, you can create all resources without specifying the
notification configuration. Then, update the stack with a notification configuration.

For more information on permissions, see [AWS::Lambda::Permission](../userguide/aws-resource-lambda-permission.md) and [Granting Permissions to Publish Event Notification Messages to a\
Destination](../../../s3/latest/dev/notificationhowto.md#grant-destinations-permissions-to-s3).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventBridgeConfiguration" : EventBridgeConfiguration,
  "LambdaConfigurations" : [ LambdaConfiguration, ... ],
  "QueueConfigurations" : [ QueueConfiguration, ... ],
  "TopicConfigurations" : [ TopicConfiguration, ... ]
}

```

### YAML

```yaml

  EventBridgeConfiguration:
    EventBridgeConfiguration
  LambdaConfigurations:
    - LambdaConfiguration
  QueueConfigurations:
    - QueueConfiguration
  TopicConfigurations:
    - TopicConfiguration

```

## Properties

`EventBridgeConfiguration`

Enables delivery of events to Amazon EventBridge.

_Required_: No

_Type_: [EventBridgeConfiguration](aws-properties-s3-bucket-eventbridgeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaConfigurations`

Describes the AWS Lambda functions to invoke and the events for which to invoke them.

_Required_: No

_Type_: Array of [LambdaConfiguration](aws-properties-s3-bucket-lambdaconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueConfigurations`

The Amazon Simple Queue Service queues to publish messages to and the events for which to publish
messages.

_Required_: No

_Type_: Array of [QueueConfiguration](aws-properties-s3-bucket-queueconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicConfigurations`

The topic to which notifications are sent and the events for which notifications are
generated.

_Required_: No

_Type_: Array of [TopicConfiguration](aws-properties-s3-bucket-topicconfiguration.md)

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

NoncurrentVersionTransition

NotificationFilter

All content copied from https://docs.aws.amazon.com/.
