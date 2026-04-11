This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Bridge BridgeFlowSource

The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FlowArn" : String,
  "FlowVpcInterfaceAttachment" : VpcInterfaceAttachment,
  "Name" : String
}

```

### YAML

```yaml

  FlowArn: String
  FlowVpcInterfaceAttachment:
    VpcInterfaceAttachment
  Name: String

```

## Properties

`FlowArn`

The ARN of the cloud flow used as a source of this bridge.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowVpcInterfaceAttachment`

The name of the VPC interface attachment to use for this source.

_Required_: No

_Type_: [VpcInterfaceAttachment](aws-properties-mediaconnect-bridge-vpcinterfaceattachment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the flow source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaConnect::Bridge

BridgeNetworkOutput

All content copied from https://docs.aws.amazon.com/.
