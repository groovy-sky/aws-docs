This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::FlowSource GatewayBridgeSource

The source configuration for cloud flows receiving a stream from a bridge.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BridgeArn" : String,
  "VpcInterfaceAttachment" : VpcInterfaceAttachment
}

```

### YAML

```yaml

  BridgeArn: String
  VpcInterfaceAttachment:
    VpcInterfaceAttachment

```

## Properties

`BridgeArn`

The ARN of the bridge feeding this flow.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:bridge:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcInterfaceAttachment`

The name of the VPC interface attachment to use for this bridge source.

_Required_: No

_Type_: [VpcInterfaceAttachment](aws-properties-mediaconnect-flowsource-vpcinterfaceattachment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Encryption

Tag

All content copied from https://docs.aws.amazon.com/.
