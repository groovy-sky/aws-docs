This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow MediaStreamAttributes

Attributes that are related to the media stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Fmtp" : Fmtp,
  "Lang" : String
}

```

### YAML

```yaml

  Fmtp:
    Fmtp
  Lang: String

```

## Properties

`Fmtp`

The settings that you want to use to define the media stream.

_Required_: No

_Type_: [Fmtp](aws-properties-mediaconnect-flow-fmtp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lang`

The audio language, in a format that is recognized by the receiver.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaStream

MediaStreamSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
