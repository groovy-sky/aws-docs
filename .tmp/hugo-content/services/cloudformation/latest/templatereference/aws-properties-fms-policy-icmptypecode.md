This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy IcmpTypeCode

ICMP protocol: The ICMP type and code.

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

ICMP code.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

ICMP type.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::FMS::Policy

IEMap

All content copied from https://docs.aws.amazon.com/.
