This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet GeoSpatialColumnGroup

Geospatial column group that denotes a hierarchy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : [ String, ... ],
  "CountryCode" : String,
  "Name" : String
}

```

### YAML

```yaml

  Columns:
    - String
  CountryCode: String
  Name: String

```

## Properties

`Columns`

Columns in this hierarchy.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `127 | 16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CountryCode`

Country code.

_Required_: No

_Type_: String

_Allowed values_: `US`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A display name for the hierarchy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FiltersOperation

ImportTableOperation

All content copied from https://docs.aws.amazon.com/.
