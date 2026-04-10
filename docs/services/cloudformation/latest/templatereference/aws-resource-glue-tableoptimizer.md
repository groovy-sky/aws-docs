This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::TableOptimizer

An AWS Glue resource for enabling table optimizers to improve read performance for open table formats.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::TableOptimizer",
  "Properties" : {
      "CatalogId" : String,
      "DatabaseName" : String,
      "TableName" : String,
      "TableOptimizerConfiguration" : TableOptimizerConfiguration,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::TableOptimizer
Properties:
  CatalogId: String
  DatabaseName: String
  TableName: String
  TableOptimizerConfiguration:
    TableOptimizerConfiguration
  Type: String

```

## Properties

`CatalogId`

The catalog ID of the table.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The name of the database. For Hive compatibility, this is folded to lowercase when it is
stored.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableName`

The table name. For Hive compatibility, this must be entirely
lowercase.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableOptimizerConfiguration`

Specifies configuration details of a table optimizer.

_Required_: Yes

_Type_: [TableOptimizerConfiguration](aws-properties-glue-tableoptimizer-tableoptimizerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of table optimizer. The valid values are:

- compaction - for managing compaction with a table optimizer.

- retention - for managing the retention of snapshot with a table optimizer.

- orphan\_file\_deletion - for managing the deletion of orphan files with a table optimizer.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

## Examples

- [Compaction table optimizer example](#aws-resource-glue-tableoptimizer--examples--Compaction_table_optimizer_example)

- [Snapshot retention table optimizer example](#aws-resource-glue-tableoptimizer--examples--Snapshot_retention_table_optimizer_example)

- [Orphan file deletion table optimizer example](#aws-resource-glue-tableoptimizer--examples--Orphan_file_deletion_table_optimizer_example)

### Compaction table optimizer example

#### JSON

```json

{
  "GlueTableOptimizer": {
    "Type": "AWS::Glue::TableOptimizer",
    "Properties": {
      "CatalogId": {"Ref": "AWS::AccountId"},
      "DatabaseName": {"Ref": "GlueDatabase"},
      "TableName": {"Ref": "GlueTable"},
      "Type": "compaction",
      "TableOptimizerConfiguration": {
        "RoleArn": {"Fn:: GetAtt": ["CompactionRole", "Arn"]},
        "Enabled": true
      }
    }
  }
}

```

#### YAML

```yaml

GlueTableOptimizer:
  Type: AWS::Glue::TableOptimizer
  Properties:
    CatalogId: !Ref AWS::AccountId
    DatabaseName: !Ref GlueDatabase
    TableName: !Ref GlueTable
    Type: "compaction"
    TableOptimizerConfiguration:
      RoleArn: !GetAtt CompactionRole.Arn
      Enabled: True

```

### Snapshot retention table optimizer example

#### JSON

```json

{
  "GlueTableOptimizer": {
    "Type": "AWS::Glue::TableOptimizer",
    "Properties": {
      "CatalogId": {"Ref": "AWS::AccountId"},
      "DatabaseName": {"Ref": "GlueDatabase"},
      "TableName": {"Ref": "GlueTable"},
      "Type": "retention",
      "TableOptimizerConfiguration": {
        "RoleArn": {"Fn::GetAtt": ["RetentionRole", "Arn"]},
        "Enabled": true,
        "RetentionConfiguration": {
          "icebergConfiguration": {
            "snapshotRetentionPeriodInDays": 7,
            "numberOfSnapshotsToRetain": 2,
            "cleanExpiredFiles": true
          }
        }
      }
    }
  }
}

```

#### YAML

```yaml

GlueTableOptimizer:
    Type: AWS::Glue::TableOptimizer
    Properties:
        CatalogId: !Ref AWS::AccountId
        DatabaseName: !Ref GlueDatabase
        TableName: !Ref GlueTable
        Type: "retention"
        TableOptimizerConfiguration:
            RoleArn: !GetAtt TableOptimizerRole.Arn
            Enabled: False
            RetentionConfiguration:
              IcebergConfiguration:
                SnapshotRetentionPeriodInDays: 7
                NumberOfSnapshotsToRetain: 5
                CleanExpiredFiles: True

```

### Orphan file deletion table optimizer example

#### JSON

```json

{
  "OrphanFileDeletionOptimizer": {
    "Type": "AWS::Glue::TableOptimizer",
    "Properties": {
      "CatalogId": {"Ref": "AWS::AccountId"},
      "DatabaseName": {"Ref": "GlueDatabase"},
      "TableName": {"Ref": "GlueTable"},
      "Type": "orphan_file_deletion",
      "TableOptimizerConfiguration": {
        "RoleArn": {"Fn::GetAtt": ["OrphanFileDeletionRole", "Arn"]},
        "Enabled": true,
        "OrphanFileDeletionConfiguration": {
          "icebergConfiguration": {
            "orphanFileRetentionPeriodInDays": 5,
            "location": "s3://my-bucket/table-location/"
          }
        }
      }
    }
  }
}

```

#### YAML

```yaml

GlueTableOptimizer:
    Type: AWS::Glue::TableOptimizer
    Properties:
      CatalogId: !Ref AWS::AccountId
      DatabaseName: !Ref GlueDatabase
      TableName: !Ref GlueTable
      Type: "orphan_file_deletion"
      TableOptimizerConfiguration:
        RoleArn: !GetAtt TableOptimizerRole.Arn
        Enabled: False
        OrphanFileDeletionConfiguration:
          IcebergConfiguration:
            OrphanFileRetentionPeriodInDays: 10
            Location: !Join ['', ['s3://', !ImportValue TestIcebergS3Bucket, '/orphan-files/']]

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableInput

IcebergConfiguration

All content copied from https://docs.aws.amazon.com/.
