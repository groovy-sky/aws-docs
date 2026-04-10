This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler IcebergTarget

Specifies Apache Iceberg data store targets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionName" : String,
  "Exclusions" : [ String, ... ],
  "MaximumTraversalDepth" : Integer,
  "Paths" : [ String, ... ]
}

```

### YAML

```yaml

  ConnectionName: String
  Exclusions:
    - String
  MaximumTraversalDepth: Integer
  Paths:
    - String

```

## Properties

`ConnectionName`

The name of the connection to use to connect to the Iceberg target.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exclusions`

A list of global patterns used to exclude from the crawl.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumTraversalDepth`

The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Iceberg metadata folder in your Amazon S3 path. Used to limit the crawler run time.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Paths`

One or more Amazon S3 paths that contains Iceberg metadata folders as s3://bucket/prefix .

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HudiTarget

JdbcTarget

All content copied from https://docs.aws.amazon.com/.
