This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow InputConfiguration

The transport parameters that are associated with an incoming media stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputPort" : Integer,
  "Interface" : Interface
}

```

### YAML

```yaml

  InputPort: Integer
  Interface:
    Interface

```

## Properties

`InputPort`

The port that the flow listens on for an incoming media stream.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interface`

The VPC interface where the media stream comes in from.

_Required_: Yes

_Type_: [Interface](aws-properties-mediaconnect-flow-interface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatewayBridgeSource

Interface

All content copied from https://docs.aws.amazon.com/.
