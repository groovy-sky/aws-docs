This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::FlowOutput DestinationConfiguration

The transport parameters that you want to associate with an outbound media stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationIp" : String,
  "DestinationPort" : Integer,
  "Interface" : Interface
}

```

### YAML

```yaml

  DestinationIp: String
  DestinationPort: Integer
  Interface:
    Interface

```

## Properties

`DestinationIp`

The IP address where you want MediaConnect to send contents of the media stream.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationPort`

The port that you want MediaConnect to use when it distributes the media stream to the output.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interface`

The VPC interface that you want to use for the media stream associated with the output.

_Required_: Yes

_Type_: [Interface](aws-properties-mediaconnect-flowoutput-interface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaConnect::FlowOutput

EncodingParameters

All content copied from https://docs.aws.amazon.com/.
