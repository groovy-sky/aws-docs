This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConcentrator

Describes a VPN concentrator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPNConcentrator",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "TransitGatewayId" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPNConcentrator
Properties:
  Tags:
    - Tag
  TransitGatewayId: String
  Type: String

```

## Properties

`Tags`

Any tags assigned to the VPN concentrator.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-vpnconcentrator-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayId`

The ID of the transit gateway associated with the VPN concentrator.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of VPN concentrator.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`TransitGatewayAttachmentId`

The ID of the transit gateway attachment for the VPN concentrator.

`VpnConcentratorId`

The ID of the VPN concentrator to associate with the VPN connection.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
