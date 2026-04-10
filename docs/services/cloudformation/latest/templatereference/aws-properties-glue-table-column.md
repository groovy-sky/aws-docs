This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table Column

A column in a `Table`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comment" : String,
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Comment: String
  Name: String
  Type: String

```

## Properties

`Comment`

A free-form text comment.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the `Column`.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The data type of the `Column`.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Column Structure](../../../glue/latest/dg/aws-glue-api-catalog-tables.md#aws-glue-api-catalog-tables-Column) in the _AWS Glue Developer_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::Table

IcebergInput

All content copied from https://docs.aws.amazon.com/.
