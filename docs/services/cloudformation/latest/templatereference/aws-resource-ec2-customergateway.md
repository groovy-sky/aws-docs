This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::CustomerGateway

Specifies a customer gateway.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::CustomerGateway",
  "Properties" : {
      "BgpAsn" : Integer,
      "BgpAsnExtended" : Number,
      "CertificateArn" : String,
      "DeviceName" : String,
      "IpAddress" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::CustomerGateway
Properties:
  BgpAsn: Integer
  BgpAsnExtended: Number
  CertificateArn: String
  DeviceName: String
  IpAddress: String
  Tags:
    - Tag
  Type: String

```

## Properties

`BgpAsn`

For customer gateway devices that support BGP, specify the device's ASN. You must specify either `BgpAsn` or `BgpAsnExtended` when creating the customer gateway. If the ASN is larger than `2,147,483,647`, you must use `BgpAsnExtended`.

Default: 65000

Valid values: `1` to `2,147,483,647`

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BgpAsnExtended`

For customer gateway devices that support BGP, specify the device's ASN. You must specify either `BgpAsn` or `BgpAsnExtended` when creating the customer gateway. If the ASN is larger than `2,147,483,647`, you must use `BgpAsnExtended`.

Valid values: `2,147,483,648` to `4,294,967,295`

_Required_: No

_Type_: Number

_Minimum_: `2147483648`

_Maximum_: `4294967294`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateArn`

The Amazon Resource Name (ARN) for the customer gateway certificate.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:acm:[a-z]{2}((-gov)|(-iso([a-z]{1})?))?-[a-z]+-\d{1}:\d{12}:certificate\/[a-zA-Z0-9-_]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeviceName`

The name of customer gateway device.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddress`

The IP address for the customer gateway device's outside interface. The address must be
static. If `OutsideIpAddressType` in your VPN connection options is set to
`PrivateIpv4`, you can use an RFC6598 or RFC1918 private IPv4 address. If
`OutsideIpAddressType` is set to `Ipv6`, you can use an IPv6 address.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

One or more tags for the customer gateway.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-customergateway-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of VPN connection that this customer gateway supports
( `ipsec.1`).

_Required_: Yes

_Type_: String

_Allowed values_: `ipsec.1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the customer gateway.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CustomerGatewayId`

The ID of the customer gateway.

## Examples

### Create a customer gateway

The following example creates a customer gateway.

#### YAML

```yaml

myCustomerGateway:
    Type: AWS::EC2::CustomerGateway
    Properties:
        Type: ipsec.1
        BgpAsn: 65534
        IpAddress: 12.1.2.3
```

#### JSON

```json

{
    "myCustomerGateway" : {
        "Type" : "AWS::EC2::CustomerGateway",
        "Properties" : {
            "Type" : "ipsec.1",
            "BgpAsn" : "65534",
            "IpAddress" : "12.1.2.3"
        }
    }
}
```

## See also

- [CreateCustomerGateway](../../../../reference/awsec2/latest/apireference/apireference-query-createcustomergateway.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::ClientVpnTargetNetworkAssociation

Tag

All content copied from https://docs.aws.amazon.com/.
