This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsGuru::NotificationChannel NotificationChannelConfig

Information about notification channels you have configured with DevOps Guru.
The one
supported notification channel is Amazon Simple Notification Service (Amazon SNS).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filters" : NotificationFilterConfig,
  "Sns" : SnsChannelConfig
}

```

### YAML

```yaml

  Filters:
    NotificationFilterConfig
  Sns:
    SnsChannelConfig

```

## Properties

`Filters`

The filter configurations for the Amazon SNS notification topic you use with DevOps Guru.
If you do not provide filter configurations, the default configurations are to receive notifications for all message types of `High` or `Medium` severity.

_Required_: No

_Type_: [NotificationFilterConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsguru-notificationchannel-notificationfilterconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Sns`

Information about a notification channel configured in DevOps Guru to send notifications
when insights are created.

If you use an Amazon SNS topic in another account, you must attach a policy to it that grants DevOps Guru permission
to send it notifications. DevOps Guru adds the required policy on your behalf to send notifications using Amazon SNS in your account. DevOps Guru only supports standard SNS topics.
For more information, see [Permissions \
for Amazon SNS topics](https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-required-permissions.html).

If you use an Amazon SNS topic that is encrypted by an AWS Key Management Service customer-managed key (CMK), then you must add permissions
to the CMK. For more information, see [Permissions for \
AWS KMS–encrypted Amazon SNS topics](https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-kms-permissions.html).

_Required_: No

_Type_: [SnsChannelConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsguru-notificationchannel-snschannelconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DevOpsGuru::NotificationChannel

NotificationFilterConfig
