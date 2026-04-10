This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::DHCPOptions

Specifies a set of DHCP options for your VPC.

You must specify at least one of the following properties:
`DomainNameServers`, `NetbiosNameServers`,
`NtpServers`. If you specify `NetbiosNameServers`, you must specify
`NetbiosNodeType`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::DHCPOptions",
  "Properties" : {
      "DomainName" : String,
      "DomainNameServers" : [ String, ... ],
      "Ipv6AddressPreferredLeaseTime" : Integer,
      "NetbiosNameServers" : [ String, ... ],
      "NetbiosNodeType" : Integer,
      "NtpServers" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::DHCPOptions
Properties:
  DomainName: String
  DomainNameServers:
    - String
  Ipv6AddressPreferredLeaseTime: Integer
  NetbiosNameServers:
    - String
  NetbiosNodeType: Integer
  NtpServers:
    - String
  Tags:
    - Tag

```

## Properties

`DomainName`

This value is used to complete unqualified DNS hostnames. If you're using
AmazonProvidedDNS in `us-east-1`, specify `ec2.internal`. If you're
using AmazonProvidedDNS in another Region, specify
_region_. `compute.internal` (for example,
`ap-northeast-1.compute.internal`). Otherwise, specify a domain name (for
example, _MyCompany.com_).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainNameServers`

The IPv4 addresses of up to four domain name servers, or `AmazonProvidedDNS`.
The default is `AmazonProvidedDNS`. To have your instance receive a custom
DNS hostname as specified in `DomainName`, you must set this property to a
custom DNS server.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6AddressPreferredLeaseTime`

A value (in seconds, minutes, hours, or years) for how frequently a running instance with an IPv6 assigned to it goes through DHCPv6 lease renewal. Acceptable values are between 140 and 2147483647 seconds (approximately 68 years). If no value is entered, the default lease time is 140 seconds. If you use long-term addressing for EC2 instances, you can increase the lease time and avoid frequent lease renewal requests. Lease renewal typically occurs when half of the lease time has elapsed.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetbiosNameServers`

The IPv4 addresses of up to four NetBIOS name servers.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetbiosNodeType`

The NetBIOS node type (1, 2, 4, or 8). We recommend that you specify 2 (broadcast and
multicast are not currently supported).

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NtpServers`

The IPv4 addresses of up to four Network Time Protocol (NTP) servers.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Any tags assigned to the DHCP options set.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-dhcpoptions-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DhcpOptionsId`

The ID of the DHCP options set.

## Examples

### Create a DHCP options set

The following example creates a DHCP options set.

#### YAML

```yaml

myDhcpOptions:
    Type: AWS::EC2::DHCPOptions
    Properties:
        DomainName: example.com
        DomainNameServers:
          - AmazonProvidedDNS
        NtpServers:
          - 10.2.5.1
        NetbiosNameServers:
          - 10.2.5.1
        NetbiosNodeType: 2
        Tags:
        - Key: project
          Value: 123
```

#### JSON

```json

{
    "myDhcpOptions" : {
        "Type" : "AWS::EC2::DHCPOptions",
        "Properties" : {
            "DomainName" : "example.com",
            "DomainNameServers" : [ "AmazonProvidedDNS" ],
            "NtpServers" : [ "10.2.5.1" ],
            "NetbiosNameServers" : [ "10.2.5.1" ],
            "NetbiosNodeType" : 2,
            "Tags" : [ { "Key" : "project", "Value" : "123" } ]
        }
    }
}
```

## See also

- [CreateDhcpOptions](../../../../reference/awsec2/latest/apireference/apireference-query-createdhcpoptions.md) in the _Amazon EC2 API Reference_

- [DHCP options sets](../../../vpc/latest/userguide/vpc-dhcp-options.md) in the _Amazon VPC User_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
