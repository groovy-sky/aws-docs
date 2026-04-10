This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kinesis::Stream StreamModeDetails

Specifies the capacity mode to which you want to set your data stream. Currently, in
Kinesis Data Streams, you can choose between an **on-demand** capacity mode and a **provisioned** capacity mode for your data streams.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StreamMode" : String
}

```

### YAML

```yaml

  StreamMode: String

```

## Properties

`StreamMode`

Specifies the capacity mode to which you want to set your data stream. Currently, in
Kinesis Data Streams, you can choose between an **on-demand** capacity mode and a **provisioned** capacity mode for your data streams.

_Required_: Yes

_Type_: String

_Allowed values_: `ON_DEMAND | PROVISIONED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StreamEncryption

Tag

All content copied from https://docs.aws.amazon.com/.
