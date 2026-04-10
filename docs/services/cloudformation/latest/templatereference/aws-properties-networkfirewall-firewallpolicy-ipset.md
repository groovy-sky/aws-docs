This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy IPSet

A list of IP addresses and address ranges, in CIDR notation. This is part of a rule variable.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Definition" : [ String, ... ]
}

```

### YAML

```yaml

  Definition:
    - String

```

## Properties

`Definition`

The list of IP addresses and address ranges, in CIDR notation.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowTimeouts

PolicyVariables

All content copied from https://docs.aws.amazon.com/.
