This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::TransitGatewayRouteTableAttachment

Creates a transit gateway route table attachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::TransitGatewayRouteTableAttachment",
  "Properties" : {
      "NetworkFunctionGroupName" : String,
      "PeeringId" : String,
      "ProposedNetworkFunctionGroupChange" : ProposedNetworkFunctionGroupChange,
      "ProposedSegmentChange" : ProposedSegmentChange,
      "RoutingPolicyLabel" : String,
      "Tags" : [ Tag, ... ],
      "TransitGatewayRouteTableArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::TransitGatewayRouteTableAttachment
Properties:
  NetworkFunctionGroupName: String
  PeeringId: String
  ProposedNetworkFunctionGroupChange:
    ProposedNetworkFunctionGroupChange
  ProposedSegmentChange:
    ProposedSegmentChange
  RoutingPolicyLabel: String
  Tags:
    - Tag
  TransitGatewayRouteTableArn: String

```

## Properties

`NetworkFunctionGroupName`

The name of the network function group.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeeringId`

The ID of the transit gateway peering.

_Required_: Yes

_Type_: String

_Pattern_: `^peering-([0-9a-f]{8,17})$`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProposedNetworkFunctionGroupChange`

Describes proposed changes to a network function group.

_Required_: No

_Type_: [ProposedNetworkFunctionGroupChange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-transitgatewayroutetableattachment-proposednetworkfunctiongroupchange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProposedSegmentChange`

This property is read-only. Values can't be assigned to it.

_Required_: No

_Type_: [ProposedSegmentChange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-transitgatewayroutetableattachment-proposedsegmentchange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingPolicyLabel`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The list of key-value pairs associated with the transit gateway route table attachment.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-transitgatewayroutetableattachment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayRouteTableArn`

The ARN of the transit gateway attachment route table. For example, `"TransitGatewayRouteTableArn": "arn:aws:ec2:us-west-2:123456789012:transit-gateway-route-table/tgw-rtb-9876543210123456"`.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the attachment ID of the transit gateway route table. For example: `attachment-12367e74104d31234`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttachmentId`

The ID of the transit gateway route table attachment.

`AttachmentPolicyRuleNumber`

The policy rule number associated with the attachment.

`AttachmentType`

The type of attachment. This will be `TRANSIT_GATEWAY_ROUTE_TABLE`.

`CoreNetworkArn`

The ARN of the core network.

`CoreNetworkId`

The ID of the core network.

`CreatedAt`

The timestamp when the transit gateway route table attachment was created.

`EdgeLocation`

The Region where the core network edge is located.

`LastModificationErrors`

Property description not available.

`OwnerAccountId`

The ID of the transit gateway route table attachment owner.

`ResourceArn`

The resource ARN for the transit gateway route table attachment.

`SegmentName`

The name of the attachment's segment.

`State`

The state of the attachment. This can be: `REJECTED` \| `PENDING_ATTACHMENT_ACCEPTANCE` \| `CREATING` \| `FAILED` \| `AVAILABLE` \| `UPDATING` \| ` PENDING_NETWORK_UPDATE` \| `PENDING_TAG_ACCEPTANCE` \| `DELETING`.

`UpdatedAt`

The timestamp when the transit gateway route table attachment was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::NetworkManager::TransitGatewayRegistration

ProposedNetworkFunctionGroupChange
