This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket MetadataConfiguration

Creates a V2 Amazon S3 Metadata configuration of a general purpose bucket.
For more information, see
[Accelerating data discovery with S3 Metadata](../../../s3/latest/userguide/metadata-tables-overview.md) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : MetadataDestination,
  "InventoryTableConfiguration" : InventoryTableConfiguration,
  "JournalTableConfiguration" : JournalTableConfiguration
}

```

### YAML

```yaml

  Destination:
    MetadataDestination
  InventoryTableConfiguration:
    InventoryTableConfiguration
  JournalTableConfiguration:
    JournalTableConfiguration

```

## Properties

`Destination`

The destination information for the S3 Metadata configuration.

_Required_: No

_Type_: [MetadataDestination](aws-properties-s3-bucket-metadatadestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InventoryTableConfiguration`

The inventory table configuration for a metadata configuration.

_Required_: No

_Type_: [InventoryTableConfiguration](aws-properties-s3-bucket-inventorytableconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JournalTableConfiguration`

The journal table configuration for a metadata configuration.

_Required_: Yes

_Type_: [JournalTableConfiguration](aws-properties-s3-bucket-journaltableconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a metadata configuration

The following example creates an S3 Metadata configuration for the specified general
purpose bucket. To use this example, replace `
                            amzn-s3-demo-bucket
                        `
with the name of your general purpose bucket. Also make sure to update the
AWS Identity and Access Management (IAM) Amazon Resource Name (ARN) with the name
of the IAM role that you want to use.

#### JSON

```json

{
  "Resources": {
    "S3MetadataKMSKey": {
      "Type": "AWS::KMS::Key",
      "Properties": {
        "Description": "KMS key for S3 metadata encryption",
        "EnableKeyRotation": true,
        "KeyPolicy": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "Enable IAM User Permissions",
              "Effect": "Allow",
              "Principal": {
                "AWS": {
                  "Fn::Sub": "arn:aws:iam::${AWS::AccountId}:role/SpecificRoleName"
                }
              },
              "Action": "kms:*",
              "Resource": "*"
            },
            {
              "Sid": "Allow S3 Metadata Service",
              "Effect": "Allow",
              "Principal": {
                "Service": [
                  "maintenance.s3tables.amazonaws.com",
                  "metadata.s3.amazonaws.com"
                ]
              },
              "Action": [
                "kms:Decrypt",
                "kms:GenerateDataKey"
              ],
              "Resource": "*"
            }
          ]
        }
      }
    },
    "S3MetadataKMSKeyAlias": {
      "Type": "AWS::KMS::Alias",
      "Properties": {
        "AliasName": "alias/s3-metadata-key",
        "TargetKeyId": {
          "Ref": "S3MetadataKMSKey"
        }
      }
    },
    "TestMetadataBucket": {
      "Type": "AWS::S3::Bucket",
      "Properties": {
        "BucketName": "amzn-s3-demo-bucket",
        "MetadataConfiguration": {
          "JournalTableConfiguration": {
            "RecordExpiration": {
              "Expiration": "ENABLED",
              "Days": 10
            },
            "EncryptionConfiguration": {
              "SseAlgorithm": "aws:kms",
              "KmsKeyArn": {
                "Fn::GetAtt": [
                  "S3MetadataKMSKey",
                  "Arn"
                ]
              }
            }
          },
          "InventoryTableConfiguration": {
            "ConfigurationState": "ENABLED",
            "EncryptionConfiguration": {
              "SseAlgorithm": "aws:kms",
              "KmsKeyArn": {
                "Fn::GetAtt": [
                  "S3MetadataKMSKey",
                  "Arn"
                ]
              }
            }
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
  S3MetadataKMSKey:
    Type: 'AWS::KMS::Key'
    Properties:
      Description: 'KMS key for S3 metadata encryption'
      EnableKeyRotation: true
      KeyPolicy:
        Version: '2012-10-17'
        Statement:
          - Sid: 'Enable IAM User Permissions'
            Effect: Allow
            Principal:
              AWS: !Sub 'arn:aws:iam::${AWS::AccountId}:role/SpecificRoleName'
            Action: 'kms:*'
            Resource: '*'
          - Sid: 'Allow S3 Metadata Service'
            Effect: Allow
            Principal:
              Service:
                - 'maintenance.s3tables.amazonaws.com'
                - 'metadata.s3.amazonaws.com'
            Action:
              - 'kms:Decrypt'
              - 'kms:GenerateDataKey'
            Resource: '*'

  S3MetadataKMSKeyAlias:
    Type: 'AWS::KMS::Alias'
    Properties:
      AliasName: 'alias/s3-metadata-key'
      TargetKeyId: !Ref S3MetadataKMSKey

  TestMetadataBucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      BucketName: amzn-s3-demo-bucket
      MetadataConfiguration:
        JournalTableConfiguration:
          RecordExpiration:
            Expiration: ENABLED
            Days: 10
          EncryptionConfiguration:
            SseAlgorithm: aws:kms
            KmsKeyArn: !GetAtt S3MetadataKMSKey.Arn
        InventoryTableConfiguration:
          ConfigurationState: ENABLED
          EncryptionConfiguration:
            SseAlgorithm: aws:kms
            KmsKeyArn: !GetAtt S3MetadataKMSKey.Arn

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingConfiguration

MetadataDestination

All content copied from https://docs.aws.amazon.com/.
