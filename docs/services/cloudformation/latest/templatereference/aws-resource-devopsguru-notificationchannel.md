This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsGuru::NotificationChannel

Adds a notification channel to DevOps Guru. A notification channel is used to notify you
about important DevOps Guru events, such as when an insight is generated.

If you use an Amazon SNS topic in another account, you must attach a policy to it that grants DevOps Guru permission
to send it notifications. DevOps Guru adds the required policy on your behalf to send notifications using Amazon SNS in your account. DevOps Guru only supports standard SNS topics.
For more information, see [Permissions \
for Amazon SNS topics](https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-required-permissions.html).

If you use an Amazon SNS topic that is encrypted by an AWS Key Management Service customer-managed key (CMK), then you must add permissions
to the CMK. For more information, see [Permissions for \
AWS KMS–encrypted Amazon SNS topics](https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-kms-permissions.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DevOpsGuru::NotificationChannel",
  "Properties" : {
      "Config" : NotificationChannelConfig
    }
}

```

### YAML

```yaml

Type: AWS::DevOpsGuru::NotificationChannel
Properties:
  Config:
    NotificationChannelConfig

```

## Properties

`Config`

A `NotificationChannelConfig` object that contains information about
configured notification channels.

_Required_: Yes

_Type_: [NotificationChannelConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsguru-notificationchannel-notificationchannelconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function,
`Ref` returns Amazon Resource Name (ARN) of the `NotificationChannel`.

For more information about using the `Ref` function, see
[Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following are the available
attributes and sample return values. For more information about using `Fn::GetAtt`, see
[Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`Id`

The ID of the notification channel.

## Examples

- [Create one notification channel with filters](#aws-resource-devopsguru-notificationchannel--examples--Create_one_notification_channel_with_filters)

- [Create two notification channels](#aws-resource-devopsguru-notificationchannel--examples--Create_two_notification_channels)

### Create one notification channel with filters

#### JSON

```json

{
  "Resources": {
    "MyNotificationChannel": {
      "Type": "AWS::DevOpsGuru::NotificationChannel",
      "Properties": {
        "Config": {
	  "Filters": {
	    "MessageTypes": ["NEW_INSIGHT", "CLOSED_INSIGHT", "SEVERITY_UPGRADED"],
	    "Severities": ["MEDIUM", "HIGH"]
	  }
          "Sns": {
            "TopicArn": "arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel"
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyNotificationChannel:
    Type: AWS::DevOpsGuru::NotificationChannel
    Properties:
      Config:
      Filters:
	  MessageTypes:
	    - NEW_INSIGHT
	    - CLOSED_INSIGHT
	    - SEVERITY_UPGRADED
	  Severities:
	    - MEDIUM
	    - HIGH
        Sns:
          TopicArn: arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel
```

### Create two notification channels

#### JSON

```json

{
  "Resources": {
    "MyNotificationChannel1": {
      "Type": "AWS::DevOpsGuru::NotificationChannel",
      "Properties": {
        "Config": {
          "Sns": {
            "TopicArn": "arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel"
          }
        }
      }
    },
    "MyNotificationChannel2": {
      "Type": "AWS::DevOpsGuru::NotificationChannel",
      "Properties": {
        "Config": {
          "Sns": {
            "TopicArn": "arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel2"
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyNotificationChannel1:
    Type: AWS::DevOpsGuru::NotificationChannel
    Properties:
      Config:
        Sns:
          TopicArn: arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel
  MyNotificationChannel2:
    Type: AWS::DevOpsGuru::NotificationChannel
    Properties:
      Config:
        Sns:
          TopicArn: arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel2
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DevOpsGuru::LogAnomalyDetectionIntegration

NotificationChannelConfig
