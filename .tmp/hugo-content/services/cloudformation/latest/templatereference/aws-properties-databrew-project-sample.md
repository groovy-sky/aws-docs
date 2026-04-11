This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Project Sample

Represents the sample size and sampling type for DataBrew to use for interactive data
analysis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Size" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  Size: Integer
  Type: String

```

## Properties

`Size`

The number of rows in the sample.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The way in which DataBrew obtains rows from a dataset.

_Required_: Yes

_Type_: String

_Allowed values_: `FIRST_N | LAST_N | RANDOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataBrew::Project

Tag

All content copied from https://docs.aws.amazon.com/.
