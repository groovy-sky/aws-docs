This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::Filter PortRangeFilter

An object that describes the details of a port range filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BeginInclusive" : Integer,
  "EndInclusive" : Integer
}

```

### YAML

```yaml

  BeginInclusive: Integer
  EndInclusive: Integer

```

## Properties

`BeginInclusive`

The port number the port range begins at.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndInclusive`

The port number the port range ends at.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PackageFilter

StringFilter

All content copied from https://docs.aws.amazon.com/.
