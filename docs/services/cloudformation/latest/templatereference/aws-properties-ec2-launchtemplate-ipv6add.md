This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate Ipv6Add

Specifies an IPv6 address in an Amazon EC2 launch template.

`Ipv6Add` is a property of [AWS::EC2::LaunchTemplate NetworkInterface](../userguide/aws-properties-ec2-launchtemplate-networkinterface.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ipv6Address" : String
}

```

### YAML

```yaml

  Ipv6Address: String

```

## Properties

`Ipv6Address`

One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You
can't use this option if you're specifying a number of IPv6 addresses.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [InstanceIpv6AddressRequest](../../../../reference/awsec2/latest/apireference/api-instanceipv6addressrequest.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ipv4PrefixSpecification

Ipv6PrefixSpecification

All content copied from https://docs.aws.amazon.com/.
