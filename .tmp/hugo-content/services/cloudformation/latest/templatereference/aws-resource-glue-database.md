This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Database

The `AWS::Glue::Database` resource specifies a logical grouping of tables
in AWS Glue. For more information, see [Defining a Database in Your Data\
Catalog](../../../glue/latest/dg/define-database.md) and [Database Structure](../../../glue/latest/dg/aws-glue-api-catalog-databases.md#aws-glue-api-catalog-databases-Database) in the _AWS Glue Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Database",
  "Properties" : {
      "CatalogId" : String,
      "DatabaseInput" : DatabaseInput,
      "DatabaseName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Database
Properties:
  CatalogId: String
  DatabaseInput:
    DatabaseInput
  DatabaseName: String

```

## Properties

`CatalogId`

The AWS account ID for the account in which to create the catalog object.

###### Note

To specify the account ID, you can use the `Ref` intrinsic function
with the `AWS::AccountId` pseudo parameter. For example: `!Ref
                    AWS::AccountId`

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseInput`

The metadata for the database.

_Required_: Yes

_Type_: [DatabaseInput](aws-properties-glue-database-databaseinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of the catalog database.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the database name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::CustomEntityType

DatabaseIdentifier

All content copied from https://docs.aws.amazon.com/.
