This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler HudiTarget

Specifies an Apache Hudi data source.

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

The name of the connection to use to connect to the Hudi target. If your Hudi files are stored in buckets that require VPC authorization, you can set their connection properties here.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exclusions`

A list of glob patterns used to exclude from the crawl.
For more information, see [Catalog Tables with a Crawler](../../../glue/latest/dg/add-crawler.md).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumTraversalDepth`

The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Hudi metadata folder in your Amazon S3 path. Used to limit the crawler run time.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Paths`

An array of Amazon S3 location strings for Hudi, each indicating the root folder with which the metadata files for a Hudi table resides. The Hudi folder may be located in a child folder of the root folder.

The crawler will scan all folders underneath a path for a Hudi folder.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDBTarget

IcebergTarget

All content copied from https://docs.aws.amazon.com/.
