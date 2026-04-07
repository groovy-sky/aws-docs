This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::Pool TwoWay

The pool's two-way SMS configuration object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChannelArn" : String,
  "ChannelRole" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  ChannelArn: String
  ChannelRole: String
  Enabled: Boolean

```

## Properties

`ChannelArn`

The Amazon Resource Name (ARN) of the two way channel.

_Required_: No

_Type_: String

_Pattern_: `^[ \S]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelRole`

An optional IAM Role Arn for a service to assume, to be able to post inbound SMS messages.

_Required_: No

_Type_: String

_Pattern_: `^arn:\S+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

By default this is set to false. When set to true you can receive incoming text
messages from your end recipients using the TwoWayChannelArn.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::SMSVOICE::ProtectConfiguration
