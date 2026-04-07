This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkAclEntry Icmp

Describes the ICMP type and code.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Code" : Integer,
  "Type" : Integer
}

```

### YAML

```yaml

  Code: Integer
  Type: Integer

```

## Properties

`Code`

The Internet Control Message Protocol (ICMP) code. You can use -1 to specify all ICMP
codes for the given ICMP type. Required if you specify 1 (ICMP)
for the protocol parameter.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The Internet Control Message Protocol (ICMP) type. You can use -1 to specify all ICMP
types. Conditional requirement: Required if you specify 1 (ICMP) for the
`CreateNetworkAclEntry` protocol parameter.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::NetworkAclEntry

PortRange
