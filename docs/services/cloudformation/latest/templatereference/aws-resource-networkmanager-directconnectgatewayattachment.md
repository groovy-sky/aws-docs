This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::DirectConnectGatewayAttachment

Creates an AWS Direct Connect gateway attachment

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::DirectConnectGatewayAttachment",
  "Properties" : {
      "CoreNetworkId" : String,
      "DirectConnectGatewayArn" : String,
      "EdgeLocations" : [ String, ... ],
      "ProposedNetworkFunctionGroupChange" : ProposedNetworkFunctionGroupChange,
      "ProposedSegmentChange" : ProposedSegmentChange,
      "RoutingPolicyLabel" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::DirectConnectGatewayAttachment
Properties:
  CoreNetworkId: String
  DirectConnectGatewayArn: String
  EdgeLocations:
    - String
  ProposedNetworkFunctionGroupChange:
    ProposedNetworkFunctionGroupChange
  ProposedSegmentChange:
    ProposedSegmentChange
  RoutingPolicyLabel: String
  Tags:
    - Tag

```

## Properties

`CoreNetworkId`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DirectConnectGatewayArn`

The Direct Connect gateway attachment ARN.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[^:]{1,63}:directconnect::[^:]{0,63}:dx-gateway\/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EdgeLocations`

Property description not available.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProposedNetworkFunctionGroupChange`

Describes proposed changes to a network function group.

_Required_: No

_Type_: [ProposedNetworkFunctionGroupChange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-directconnectgatewayattachment-proposednetworkfunctiongroupchange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProposedSegmentChange`

Describes a proposed segment change. In some cases, the segment change must first be evaluated and accepted.

_Required_: No

_Type_: [ProposedSegmentChange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-directconnectgatewayattachment-proposedsegmentchange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingPolicyLabel`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-directconnectgatewayattachment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AttachmentId`

Property description not available.

`AttachmentPolicyRuleNumber`

Property description not available.

`AttachmentType`

Property description not available.

`CoreNetworkArn`

Property description not available.

`CreatedAt`

Property description not available.

`LastModificationErrors`

Property description not available.

`NetworkFunctionGroupName`

The name of the network function group.

`OwnerAccountId`

Property description not available.

`ResourceArn`

Property description not available.

`SegmentName`

Property description not available.

`State`

Property description not available.

`UpdatedAt`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ProposedNetworkFunctionGroupChange
