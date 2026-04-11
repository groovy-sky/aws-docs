This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign SetDimension

Specifies the dimension type and values for a segment dimension.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DimensionType" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  DimensionType: String
  Values:
    - String

```

## Properties

`DimensionType`

The type of segment dimension to use. Valid values are:
`INCLUSIVE`, endpoints that match the criteria are included in the
segment; and, `EXCLUSIVE`, endpoints that match the criteria are
excluded from the segment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The criteria values to use for the segment dimension. Depending on the value
of the `DimensionType` property, endpoints are included or excluded
from the segment if their values match the criteria values.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Schedule

Template

All content copied from https://docs.aws.amazon.com/.
