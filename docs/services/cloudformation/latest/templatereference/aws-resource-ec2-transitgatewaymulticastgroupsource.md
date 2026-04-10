This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayMulticastGroupSource

Registers sources (network interfaces) with the specified transit gateway multicast
domain.

A multicast source is a network interface attached to a supported instance that sends
multicast traffic. For information about supported instances, see [Multicast Considerations](../../../vpc/latest/tgw/transit-gateway-limits.md#multicast-limits) in _Amazon VPC Transit_
_Gateways_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayMulticastGroupSource",
  "Properties" : {
      "GroupIpAddress" : String,
      "NetworkInterfaceId" : String,
      "TransitGatewayMulticastDomainId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayMulticastGroupSource
Properties:
  GroupIpAddress: String
  NetworkInterfaceId: String
  TransitGatewayMulticastDomainId: String

```

## Properties

`GroupIpAddress`

The IP address assigned to the transit gateway multicast group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaceId`

The group sources' network interface IDs to register with the transit gateway multicast group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayMulticastDomainId`

The ID of the transit gateway multicast domain.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the transit gateway multicast domain group source.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GroupMember`

Information about the registered transit gateway multicast domain group members.

`GroupSource`

Indicates that the resource is a transit gateway group member.

`ResourceId`

The ID of the resource.

`ResourceType`

The type of resource, for example a VPC attachment.

`SourceType`

The type of source.

`SubnetId`

The ID of the subnet.

`TransitGatewayAttachmentId`

The ID of the transit gateway attachment.

## See also

- [RegisterTransitGatewayMulticastGroupSources](../../../../reference/awsec2/latest/apireference/api-registertransitgatewaymulticastgroupsources.md) in the _Amazon_
_EC2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::TransitGatewayMulticastGroupMember

AWS::EC2::TransitGatewayPeeringAttachment

All content copied from https://docs.aws.amazon.com/.
