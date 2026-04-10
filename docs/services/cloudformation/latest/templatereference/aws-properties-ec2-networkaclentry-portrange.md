This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkAclEntry PortRange

Describes a range of ports.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "From" : Integer,
  "To" : Integer
}

```

### YAML

```yaml

  From: Integer
  To: Integer

```

## Properties

`From`

The first port in the range. Required if you specify 6 (TCP) or 17 (UDP) for the
protocol parameter.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`To`

The last port in the range. Required if you specify 6 (TCP) or 17 (UDP) for the protocol
parameter.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Icmp

AWS::EC2::NetworkInsightsAccessScope

All content copied from https://docs.aws.amazon.com/.
