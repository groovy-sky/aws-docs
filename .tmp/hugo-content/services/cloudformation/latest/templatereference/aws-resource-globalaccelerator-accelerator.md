This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GlobalAccelerator::Accelerator

The `AWS::GlobalAccelerator::Accelerator` resource is a Global Accelerator resource type that contains information about
how you create an accelerator. An accelerator includes one or more listeners that process inbound connections and direct traffic
to one or more endpoint groups, each of which includes endpoints, such as Application Load Balancers, Network Load Balancers,
and Amazon EC2 instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GlobalAccelerator::Accelerator",
  "Properties" : {
      "Enabled" : Boolean,
      "IpAddresses" : [ String, ... ],
      "IpAddressType" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GlobalAccelerator::Accelerator
Properties:
  Enabled: Boolean
  IpAddresses:
    - String
  IpAddressType: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Enabled`

Indicates whether the accelerator is enabled. The value is true or false. The default value is true.

If the value is set to true, the accelerator cannot be deleted. If set to false, accelerator can be deleted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddresses`

Optionally, if you've added your own IP address pool to Global Accelerator (BYOIP), you can choose IP addresses
from your own pool to use for the accelerator's static IP addresses when you create an accelerator. You can
specify one or two addresses, separated by a comma. Do not include the /32 suffix.

Only one IP address from each of your IP address ranges can be used for each accelerator. If you specify only
one IP address from your IP address range, Global Accelerator assigns a second static IP address for the
accelerator from the AWS IP address pool.

Note that you can't update IP addresses for an existing accelerator. To change them, you must create a new
accelerator with the new addresses.

For more information, see [Bring Your Own \
IP Addresses (BYOIP)](../../../global-accelerator/latest/dg/using-byoip.md) in the _AWS Global Accelerator Developer Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The IP address type that an accelerator supports. For a standard accelerator, the value can be IPV4 or DUAL\_STACK.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUAL_STACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the accelerator. The name must contain only alphanumeric characters or
hyphens (-), and must not begin or end with a hyphen.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{0,64}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Create tags for an accelerator.

For more information, see [Tagging](../../../global-accelerator/latest/dg/tagging-in-global-accelerator.md) in the _AWS Global Accelerator Developer Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-globalaccelerator-accelerator-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the accelerator, such as
`arn:aws:globalaccelerator::012345678901:accelerator/1234abcd-abcd-1234-abcd-1234abcdefgh`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AcceleratorArn`

The ARN of the accelerator, such as
`arn:aws:globalaccelerator::012345678901:accelerator/1234abcd-abcd-1234-abcd-1234abcdefgh`.

`DnsName`

The Domain Name System (DNS) name that Global Accelerator creates that points to an accelerator's static IPv4 addresses.

`DualStackDnsName`

The DNS name that Global Accelerator creates that points to a dual-stack accelerator's four
static IP addresses: two IPv4 addresses and two IPv6 addresses.

`Ipv4Addresses`

The array of IPv4 addresses in the IP address set. An IP address set can have a maximum of two IP addresses.

`Ipv6Addresses`

The array of IPv6 addresses in the IP address set. An IP address set can have a maximum of two IP addresses.

## Examples

### Add an accelerator

These are examples to specify an accelerator.

#### JSON

```json

"Resources": {
    "Accelerator": {
        "Type": "AWS::GlobalAccelerator::Accelerator",
        "Properties": {
            "Name": "SampleAccelerator",
            "Enabled": true
        }
    },
    "Outputs": {
        "AcceleratorDnsName": {
            "Description": "Accelerator DNS Name",
            "Value": {
                "Fn::GetAtt": [
                    "Accelerator",
                    "DnsName"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Accelerator:
  Type: AWS::GlobalAccelerator::Accelerator
  Properties:
    Name: SampleAccelerator
    Enabled: true
Outputs:
  AcceleratorDnsName:
    Description: Accelerator DNS Name
    Value:
      Fn::GetAtt:
      - Accelerator
      - DnsName
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Global Accelerator

Tag

All content copied from https://docs.aws.amazon.com/.
