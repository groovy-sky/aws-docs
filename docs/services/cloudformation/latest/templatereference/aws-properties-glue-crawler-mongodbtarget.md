This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler MongoDBTarget

Specifies an Amazon DocumentDB or MongoDB data store to crawl.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionName" : String,
  "Path" : String
}

```

### YAML

```yaml

  ConnectionName: String
  Path: String

```

## Properties

`ConnectionName`

The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path of the Amazon DocumentDB or MongoDB target (database/collection).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LakeFormationConfiguration

RecrawlPolicy

All content copied from https://docs.aws.amazon.com/.
