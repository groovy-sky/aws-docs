This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler DeltaTarget

Specifies a Delta data store to crawl one or more Delta tables.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionName" : String,
  "CreateNativeDeltaTable" : Boolean,
  "DeltaTables" : [ String, ... ],
  "WriteManifest" : Boolean
}

```

### YAML

```yaml

  ConnectionName: String
  CreateNativeDeltaTable: Boolean
  DeltaTables:
    - String
  WriteManifest: Boolean

```

## Properties

`ConnectionName`

The name of the connection to use to connect to the Delta table target.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateNativeDeltaTable`

Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeltaTables`

A list of the Amazon S3 paths to the Delta tables.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteManifest`

Specifies whether to write the manifest files to the Delta table path.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CatalogTarget

DynamoDBTarget

All content copied from https://docs.aws.amazon.com/.
