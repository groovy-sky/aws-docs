This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket MetadataTableConfiguration

###### Important

We recommend that you create your S3 Metadata configurations by using the V2
[MetadataConfiguration](aws-properties-s3-bucket-metadataconfiguration.md) resource type. We no longer recommend using the V1
`MetadataTableConfiguration` resource type.

If you created your S3 Metadata configuration before July 15, 2025, we recommend that you
delete and re-create your configuration by using the
[MetadataConfiguration](aws-properties-s3-bucket-metadataconfiguration.md) resource type so that you can expire journal table records and
create a live inventory table.

Creates a V1 S3 Metadata configuration for a general purpose bucket. For more information, see
[Accelerating\
data discovery with S3 Metadata](../../../s3/latest/userguide/metadata-tables-overview.md) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3TablesDestination" : S3TablesDestination
}

```

### YAML

```yaml

  S3TablesDestination:
    S3TablesDestination

```

## Properties

`S3TablesDestination`

The destination information for the metadata table configuration. The destination table bucket must
be in the same Region and AWS account as the general purpose bucket. The specified metadata table name
must be unique within the `aws_s3_metadata` namespace in the destination table bucket.

_Required_: Yes

_Type_: [S3TablesDestination](aws-properties-s3-bucket-s3tablesdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create metadata table configuration

The following example creates an S3 Metadata configuration for the specified general
purpose bucket. To use this example, replace `
                            amzn-s3-demo-table-bucket
                        ` with
the name of the table bucket where you want to store your metadata table, replace
`
                            amzn-s3-demo-bucket1
                        ` with the name of your general purpose bucket,
and replace `my_metadata_table_name` with the name that you want to use for your
metadata table.

#### JSON

```json

{
  "Resources": {
    "S3TableBucket": {
      "Type": "AWS::S3Tables::TableBucket",
      "Properties": {
        "TableBucketName": "amzn-s3-demo-table-bucket"
      }
    },
    "S3Bucket": {
      "Type": "AWS::S3::Bucket",
      "DeletionPolicy": "Retain",
      "Properties": {
        "BucketName": "amzn-s3-demo-bucket1",
        "MetadataTableConfiguration": {
          "S3TablesDestination": {
            "TableBucketArn": {
              "Fn::GetAtt": [
                "S3TableBucket",
                "TableBucketARN"
              ]
            },
            "TableName": "my_metadata_table_name",
            "TableNamespace": "aws_s3_metadata"
          }
        }
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  S3TableBucket:
    Type: AWS::S3Tables::TableBucket
    Properties:
      TableBucketName: amzn-s3-demo-table-bucket

  S3Bucket:
    Type: AWS::S3::Bucket
    DeletionPolicy: Retain
    Properties:
      BucketName: amzn-s3-demo-bucket1
      MetadataTableConfiguration:
        S3TablesDestination:
          TableBucketArn: !GetAtt S3TableBucket.TableBucketARN
          TableName: my_metadata_table_name
          TableNamespace: aws_s3_metadata

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataDestination

MetadataTableEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
