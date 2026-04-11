This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Contact ChannelTargetInfo

Information about the contact channel that Incident Manager uses to engage the
contact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChannelId" : String,
  "RetryIntervalInMinutes" : Integer
}

```

### YAML

```yaml

  ChannelId: String
  RetryIntervalInMinutes: Integer

```

## Properties

`ChannelId`

The Amazon Resource Name (ARN) of the contact channel.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-cn|aws-us-gov):ssm-contacts:[-\w+=\/,.@]*:[0-9]+:([\w+=\/,.@:-])*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryIntervalInMinutes`

The number of minutes to wait before retrying to send engagement if the engagement
initially failed.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSMContacts::Contact

ContactTargetInfo

All content copied from https://docs.aws.amazon.com/.
