This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayMulticastDomainAssociation

Associates the specified subnets and transit gateway attachments with the specified
transit gateway multicast domain.

The transit gateway attachment must be in the available state before you can add a
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayMulticastDomainAssociation",
  "Properties" : {
      "SubnetId" : String,
      "TransitGatewayAttachmentId" : String,
      "TransitGatewayMulticastDomainId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayMulticastDomainAssociation
Properties:
  SubnetId: String
  TransitGatewayAttachmentId: String
  TransitGatewayMulticastDomainId: String

```

## Properties

`SubnetId`

The IDs of the subnets to associate with the transit gateway multicast domain.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayAttachmentId`

The ID of the transit gateway attachment.

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

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the transit gateway multicast domain ID. For example:
`tgw-mcast-domain-000fb24d04EXAMPLE`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourceId`

The ID of the resource.

`ResourceType`

The type of resource, for example a VPC attachment.

`State`

The state of the resource.

## See also

- [AssociateTransitGatewayMulticastDomain](../../../../reference/awsec2/latest/apireference/api-associatetransitgatewaymulticastdomain.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::TransitGatewayMulticastGroupMember

All content copied from https://docs.aws.amazon.com/.
