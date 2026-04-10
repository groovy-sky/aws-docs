This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayConnect

Creates a Connect attachment from a specified transit gateway attachment. A Connect attachment is a GRE-based tunnel attachment that you can use to establish a connection between a transit gateway and an appliance.

A Connect attachment uses an existing VPC or AWS Direct Connect attachment as the underlying transport mechanism.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayConnect",
  "Properties" : {
      "Options" : TransitGatewayConnectOptions,
      "Tags" : [ Tag, ... ],
      "TransportTransitGatewayAttachmentId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayConnect
Properties:
  Options:
    TransitGatewayConnectOptions
  Tags:
    - Tag
  TransportTransitGatewayAttachmentId: String

```

## Properties

`Options`

The Connect attachment options.

- protocol (gre)

_Required_: Yes

_Type_: [TransitGatewayConnectOptions](aws-properties-ec2-transitgatewayconnect-transitgatewayconnectoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the attachment.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-transitgatewayconnect-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransportTransitGatewayAttachmentId`

The ID of the attachment from which the Connect attachment was created.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the transit gateway attachment.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The creation time.

`State`

The state of the attachment.

`TransitGatewayAttachmentId`

The ID of the transit gateway attachment.

`TransitGatewayId`

The ID of the transit gateway.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
