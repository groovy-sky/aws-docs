This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElementalInference::Feed ClippingConfig

A type of OutputConfig, used when the output in a feed is for the clip feature.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CallbackMetadata" : String
}

```

### YAML

```yaml

  CallbackMetadata: String

```

## Properties

`CallbackMetadata`

Metadata that you want to include in the event that Elemental Inference sends to Amazon EventBridge.
The metadata can help you distinguish between different events.

_Required_: No

_Type_: String

_Pattern_: `^[\w \-\.',@:;]*$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ElementalInference::Feed

GetOutput

All content copied from https://docs.aws.amazon.com/.
