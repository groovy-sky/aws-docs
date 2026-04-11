This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayMulticastDomain

Creates a multicast domain using the specified transit gateway.

The transit gateway must be in the available state before you create a domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayMulticastDomain",
  "Properties" : {
      "Options" : Options,
      "Tags" : [ Tag, ... ],
      "TransitGatewayId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayMulticastDomain
Properties:
  Options:
    Options
  Tags:
    - Tag
  TransitGatewayId: String

```

## Properties

`Options`

The options for the transit gateway multicast domain.

- AutoAcceptSharedAssociations (enable \| disable)

- Igmpv2Support (enable \| disable)

- StaticSourcesSupport (enable \| disable)

_Required_: No

_Type_: [Options](aws-properties-ec2-transitgatewaymulticastdomain-options.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the transit gateway multicast domain.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-transitgatewaymulticastdomain-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayId`

The ID of the transit gateway.

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

`CreationTime`

The time the multicast domain was created.

`State`

The state of the multicast domain.

`TransitGatewayMulticastDomainArn`

The Amazon Resource Name (ARN) of the multicast domain.

`TransitGatewayMulticastDomainId`

The ID of the multicast domain.

## See also

- [CreateTransitGatewayMulticastDomain](../../../../reference/awsec2/latest/apireference/api-createtransitgatewaymulticastdomain.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::TransitGatewayMeteringPolicyEntry

Options

All content copied from https://docs.aws.amazon.com/.
