This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::CarrierGateway

Creates a carrier gateway. For more information about carrier gateways, see [Carrier gateways](../../../wavelength/latest/developerguide/how-wavelengths-work.md#wavelength-carrier-gateway) in the _AWS Wavelength Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::CarrierGateway",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::CarrierGateway
Properties:
  Tags:
    - Tag
  VpcId: String

```

## Properties

`Tags`

The tags assigned to the carrier gateway.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-carriergateway-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC associated with the carrier gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the carrier gateway ID. For example:
`cagw-05a8da9a199afb1c7`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CarrierGatewayId`

The ID of the carrier gateway.

`OwnerId`

The AWS account ID of the owner of the carrier gateway.

`State`

The state of the carrier gateway.

## See also

- [Carrier gateways](../../../vpc/latest/userguide/carrier-gateway.md) in _Amazon VPC User Guide_

- [CreateCarrierGateway](../../../../reference/awsec2/latest/apireference/api-createcarriergateway.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagSpecification

Tag

All content copied from https://docs.aws.amazon.com/.
