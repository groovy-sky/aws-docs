This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::AnycastIpList

An Anycast static IP list. For more information, see [Request Anycast static IPs to use for allowlisting](../../../amazoncloudfront/latest/developerguide/request-static-ips.md) in the _Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::AnycastIpList",
  "Properties" : {
      "IpAddressType" : String,
      "IpamCidrConfigs" : [ IpamCidrConfig, ... ],
      "IpCount" : Integer,
      "Name" : String,
      "Tags" : Tags
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::AnycastIpList
Properties:
  IpAddressType: String
  IpamCidrConfigs:
    - IpamCidrConfig
  IpCount: Integer
  Name: String
  Tags:
    Tags

```

## Properties

`IpAddressType`

The IP address type for the Anycast static IP list.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | dualstack`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamCidrConfigs`

A list of IPAM CIDR configurations that define the IP address ranges, IPAM pools, and associated Anycast IP addresses.

_Required_: No

_Type_: Array of [IpamCidrConfig](aws-properties-cloudfront-anycastiplist-ipamcidrconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpCount`

The number of IP addresses in the Anycast static IP list.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the Anycast static IP list.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]{1,64}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A complex type that contains zero or more `Tag` elements.

_Required_: No

_Type_: [Tags](aws-properties-cloudfront-anycastiplist-tags.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Anycast IP list. For example:
`aip_7XdPQUqEXAMPLE283z455j`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ETag`

A complex type that contains `Tag` key and `Tag` value.

`Id`

The ID of the Anycast static IP list.

`IpamCidrConfigResults`

The results for the IPAM CIDRs that defines a specific IP address range, IPAM pool, and associated Anycast IP address.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon CloudFront

AnycastIpList

All content copied from https://docs.aws.amazon.com/.
