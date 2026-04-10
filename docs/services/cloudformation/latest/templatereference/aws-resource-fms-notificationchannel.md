This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::NotificationChannel

Designates the IAM role and Amazon Simple Notification Service (SNS) topic to use to record SNS logs.

To perform this action outside of the console, you must configure the SNS topic to allow the
role `AWSServiceRoleForFMS` to publish SNS logs. For more information, see
[Firewall Manager required permissions for API actions](../../../waf/latest/developerguide/fms-api-permissions-ref.md) in the _AWS Firewall Manager Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FMS::NotificationChannel",
  "Properties" : {
      "SnsRoleName" : String,
      "SnsTopicArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::FMS::NotificationChannel
Properties:
  SnsRoleName: String
  SnsTopicArn: String

```

## Properties

`SnsRoleName`

The Amazon Resource Name (ARN) of the IAM role that allows Amazon SNS to record AWS Firewall Manager activity.

_Required_: Yes

_Type_: String

_Pattern_: `^([^\s]+)$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsTopicArn`

The Amazon Resource Name (ARN) of the SNS topic that collects notifications from AWS Firewall Manager.

_Required_: Yes

_Type_: String

_Pattern_: `^([^\s]+)$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

The `Ref` for this resource returns the `SnsTopicArn`. This is
the Amazon Resource Name (ARN) that uniquely identifies the Amazon Simple Notification Service (Amazon SNS) topic. For
example, `arn:aws:sns:us-west-2:111122223333:MyTopic`. For more information about SNS, see [Amazon Simple Notification Service Resource Type Reference](../userguide/aws-sns.md).

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource physical ID, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create a Firewall Manager notification channel

The following shows an example SNS notification channel for Firewall Manager.

#### YAML

```yaml

NotificationChannel:
    Type: AWS::FMS::NotificationChannel
    Properties:
      SnsRoleName: !Sub arn:aws:iam::${AWS::AccountId}:role/aws-service-role/fms.amazonaws.com/AWSServiceRoleForFMS
      SnsTopicArn: !Ref SnsTopic
```

#### JSON

```json

"NotificationChannel": {
    "Type": "AWS::FMS::NotificationChannel",
    "Properties": {
        "SnsRoleName": {
            "Fn::Sub": "arn:aws:iam::${AWS::AccountId}:role/aws-service-role/fms.amazonaws.com/AWSServiceRoleForFMS"
        },
        "SnsTopicArn": {
            "Ref": "SnsTopic"
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Firewall Manager

AWS::FMS::Policy

All content copied from https://docs.aws.amazon.com/.
