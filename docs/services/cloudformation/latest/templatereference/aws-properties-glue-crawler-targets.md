This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler Targets

Specifies data stores to crawl.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogTargets" : [ CatalogTarget, ... ],
  "DeltaTargets" : [ DeltaTarget, ... ],
  "DynamoDBTargets" : [ DynamoDBTarget, ... ],
  "HudiTargets" : [ HudiTarget, ... ],
  "IcebergTargets" : [ IcebergTarget, ... ],
  "JdbcTargets" : [ JdbcTarget, ... ],
  "MongoDBTargets" : [ MongoDBTarget, ... ],
  "S3Targets" : [ S3Target, ... ]
}

```

### YAML

```yaml

  CatalogTargets:
    - CatalogTarget
  DeltaTargets:
    - DeltaTarget
  DynamoDBTargets:
    - DynamoDBTarget
  HudiTargets:
    - HudiTarget
  IcebergTargets:
    - IcebergTarget
  JdbcTargets:
    - JdbcTarget
  MongoDBTargets:
    - MongoDBTarget
  S3Targets:
    - S3Target

```

## Properties

`CatalogTargets`

Specifies AWS Glue Data Catalog targets.

_Required_: No

_Type_: Array of [CatalogTarget](aws-properties-glue-crawler-catalogtarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeltaTargets`

Specifies an array of Delta data store targets.

_Required_: No

_Type_: Array of [DeltaTarget](aws-properties-glue-crawler-deltatarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDBTargets`

Specifies Amazon DynamoDB targets.

_Required_: No

_Type_: Array of [DynamoDBTarget](aws-properties-glue-crawler-dynamodbtarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HudiTargets`

Property description not available.

_Required_: No

_Type_: Array of [HudiTarget](aws-properties-glue-crawler-huditarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IcebergTargets`

Specifies Apache Iceberg data store targets.

_Required_: No

_Type_: Array of [IcebergTarget](aws-properties-glue-crawler-icebergtarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JdbcTargets`

Specifies JDBC targets.

_Required_: No

_Type_: Array of [JdbcTarget](aws-properties-glue-crawler-jdbctarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MongoDBTargets`

A list of Mongo DB targets.

_Required_: No

_Type_: Array of [MongoDBTarget](aws-properties-glue-crawler-mongodbtarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Targets`

Specifies Amazon Simple Storage Service (Amazon S3) targets.

_Required_: No

_Type_: Array of [S3Target](aws-properties-glue-crawler-s3target.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaChangePolicy

AWS::Glue::CustomEntityType

All content copied from https://docs.aws.amazon.com/.
