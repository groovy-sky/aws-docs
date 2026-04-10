This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsAccessScope PacketHeaderStatementRequest

Describes a packet header statement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationAddresses" : [ String, ... ],
  "DestinationPorts" : [ String, ... ],
  "DestinationPrefixLists" : [ String, ... ],
  "Protocols" : [ String, ... ],
  "SourceAddresses" : [ String, ... ],
  "SourcePorts" : [ String, ... ],
  "SourcePrefixLists" : [ String, ... ]
}

```

### YAML

```yaml

  DestinationAddresses:
    - String
  DestinationPorts:
    - String
  DestinationPrefixLists:
    - String
  Protocols:
    - String
  SourceAddresses:
    - String
  SourcePorts:
    - String
  SourcePrefixLists:
    - String

```

## Properties

`DestinationAddresses`

The destination addresses.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPorts`

The destination ports.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPrefixLists`

The destination prefix lists.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocols`

The protocols.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceAddresses`

The source addresses.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePorts`

The source ports.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePrefixLists`

The source prefix lists.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessScopePathRequest

PathStatementRequest

All content copied from https://docs.aws.amazon.com/.
