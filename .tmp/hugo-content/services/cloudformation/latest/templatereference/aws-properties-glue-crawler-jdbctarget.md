This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler JdbcTarget

Specifies a JDBC data store to crawl.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionName" : String,
  "EnableAdditionalMetadata" : [ String, ... ],
  "Exclusions" : [ String, ... ],
  "Path" : String
}

```

### YAML

```yaml

  ConnectionName: String
  EnableAdditionalMetadata:
    - String
  Exclusions:
    - String
  Path: String

```

## Properties

`ConnectionName`

The name of the connection to use to connect to the JDBC target.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableAdditionalMetadata`

Specify a value of `RAWTYPES` or `COMMENTS` to enable additional metadata in table responses. `RAWTYPES` provides the native-level datatype. `COMMENTS` provides comments associated with a column or table in the database.

If you do not need additional metadata, keep the field empty.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exclusions`

A list of glob patterns used to exclude from the crawl. For more information, see
[Catalog Tables\
with a Crawler](../../../glue/latest/dg/add-crawler.md).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path of the JDBC target.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergTarget

LakeFormationConfiguration

All content copied from https://docs.aws.amazon.com/.
