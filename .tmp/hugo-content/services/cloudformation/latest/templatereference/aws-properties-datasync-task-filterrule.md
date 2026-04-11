This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Task FilterRule

Specifies which files, folders, and objects to include or exclude when transferring files
from source to destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilterType" : String,
  "Value" : String
}

```

### YAML

```yaml

  FilterType: String
  Value: String

```

## Properties

`FilterType`

The type of filter rule to apply. AWS DataSync only supports the SIMPLE\_PATTERN
rule type.

_Required_: No

_Type_: String

_Allowed values_: `SIMPLE_PATTERN`

_Pattern_: `^[A-Z0-9_]+$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A single filter string that consists of the patterns to include or exclude. The patterns
are delimited by "\|" (that is, a pipe), for example: `/folder1|/folder2`

_Required_: No

_Type_: String

_Pattern_: `^[^\x00]+$`

_Maximum_: `409600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Destination

ManifestConfig

All content copied from https://docs.aws.amazon.com/.
