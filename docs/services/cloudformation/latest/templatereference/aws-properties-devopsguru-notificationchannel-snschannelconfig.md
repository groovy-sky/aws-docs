This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsGuru::NotificationChannel SnsChannelConfig

Contains the Amazon Resource Name (ARN) of an Amazon Simple Notification Service topic.

If you use an Amazon SNS topic in another account, you must attach a policy to it that grants DevOps Guru permission
to send it notifications. DevOps Guru adds the required policy on your behalf to send notifications using Amazon SNS in your account. DevOps Guru only supports standard SNS topics.
For more information, see [Permissions \
for Amazon SNS topics](../../../devops-guru/latest/userguide/sns-required-permissions.md).

If you use an Amazon SNS topic that is encrypted by an AWS Key Management Service customer-managed key (CMK), then you must add permissions
to the CMK. For more information, see [Permissions for \
AWS KMS–encrypted Amazon SNS topics](../../../devops-guru/latest/userguide/sns-kms-permissions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TopicArn" : String
}

```

### YAML

```yaml

  TopicArn: String

```

## Properties

`TopicArn`

The Amazon Resource Name (ARN) of an Amazon Simple Notification Service topic.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\d{12}:[^:]+$`

_Minimum_: `36`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationFilterConfig

AWS::DevOpsGuru::ResourceCollection

All content copied from https://docs.aws.amazon.com/.
