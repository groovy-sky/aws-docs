This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInterface PrivateIpAddressSpecification

Describes a secondary private IPv4 address for a network interface.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Primary" : Boolean,
  "PrivateIpAddress" : String
}

```

### YAML

```yaml

  Primary: Boolean
  PrivateIpAddress: String

```

## Properties

`Primary`

Sets the private IP address as the primary private address. You can set only one primary
private IP address. If you don't specify a primary private IP address, Amazon EC2
automatically assigns a primary private IP address.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PrivateIpAddress`

The private IP address of the network interface.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ipv6PrefixSpecification

PublicIpDnsNameOptions

All content copied from https://docs.aws.amazon.com/.
