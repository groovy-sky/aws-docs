This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::AnycastIpList IpamCidrConfig

Configuration for an IPAM CIDR that defines a specific IP address range, IPAM pool, and associated Anycast IP address.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidr" : String,
  "IpamPoolArn" : String
}

```

### YAML

```yaml

  Cidr: String
  IpamPoolArn: String

```

## Properties

`Cidr`

The CIDR that specifies the IP address range for this IPAM configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamPoolArn`

The Amazon Resource Name (ARN) of the IPAM pool that the CIDR block is assigned to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnycastIpList

IpamCidrConfigResult

All content copied from https://docs.aws.amazon.com/.
