This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate PrivateIpAdd

Specifies a secondary private IPv4 address for a network interface.

`PrivateIpAdd` is a property of [AWS::EC2::LaunchTemplate NetworkInterface](../userguide/aws-properties-ec2-launchtemplate-networkinterface.md).

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

Indicates whether the private IPv4 address is the primary private IPv4 address. Only
one IPv4 address can be designated as primary.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateIpAddress`

The private IPv4 address.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [PrivateIpAddressSpecification](../../../../reference/awsec2/latest/apireference/api-privateipaddressspecification.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrivateDnsNameOptions

Reference

All content copied from https://docs.aws.amazon.com/.
