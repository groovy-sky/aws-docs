This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket InventoryConfiguration

Specifies the S3 Inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket\
inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the _Amazon S3 API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : Destination,
  "Enabled" : Boolean,
  "Id" : String,
  "IncludedObjectVersions" : String,
  "OptionalFields" : [ String, ... ],
  "Prefix" : String,
  "ScheduleFrequency" : String
}

```

### YAML

```yaml

  Destination:
    Destination
  Enabled: Boolean
  Id: String
  IncludedObjectVersions: String
  OptionalFields:
    - String
  Prefix: String
  ScheduleFrequency: String

```

## Properties

`Destination`

Contains information about where to publish the inventory results.

_Required_: Yes

_Type_: [Destination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-destination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether the inventory is enabled or disabled. If set to `True`, an inventory
list is generated. If set to `False`, no inventory list is generated.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID used to identify the inventory configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedObjectVersions`

Object versions to include in the inventory list. If set to `All`, the list includes all
the object versions, which adds the version-related fields `VersionId`,
`IsLatest`, and `DeleteMarker` to the list. If set to `Current`, the
list does not contain these version-related fields.

_Required_: Yes

_Type_: String

_Allowed values_: `All | Current`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionalFields`

Contains the optional fields that are included in the inventory results.

_Required_: No

_Type_: Array of String

_Allowed values_: `Size | LastModifiedDate | StorageClass | ETag | IsMultipartUploaded | ReplicationStatus | EncryptionStatus | ObjectLockRetainUntilDate | ObjectLockMode | ObjectLockLegalHoldStatus | IntelligentTieringAccessTier | BucketKeyStatus | ChecksumAlgorithm | ObjectAccessControlList | ObjectOwner | LifecycleExpirationDate`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

Specifies the inventory filter prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleFrequency`

Specifies the schedule for generating inventory results.

_Required_: Yes

_Type_: String

_Allowed values_: `Daily | Weekly`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Specify analytics and inventory configurations for an S3 bucket

The following example specifies analytics and inventory results to be generated for an
S3 bucket, including the format of the results and the destination bucket. The inventory
list generates reports weekly and includes the current version of each object.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "S3 Bucket with Inventory and Analytics Configurations",
    "Resources": {
        "Helper": {
            "Type": "AWS::S3::Bucket"
        },
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AnalyticsConfigurations": [
                    {
                        "Id": "AnalyticsConfigurationId",
                        "StorageClassAnalysis": {
                            "DataExport": {
                                "Destination": {
                                    "BucketArn": {
                                        "Fn::GetAtt": [
                                            "Helper",
                                            "Arn"
                                        ]
                                    },
                                    "Format": "CSV",
                                    "Prefix": "AnalyticsDestinationPrefix"
                                },
                                "OutputSchemaVersion": "V_1"
                            }
                        },
                        "Prefix": "AnalyticsConfigurationPrefix",
                        "TagFilters": [
                            {
                                "Key": "AnalyticsTagKey",
                                "Value": "AnalyticsTagValue"
                            }
                        ]
                    }
                ],
                "InventoryConfigurations": [
                    {
                        "Id": "InventoryConfigurationId",
                        "Destination": {
                            "BucketArn": {
                                "Fn::GetAtt": [
                                    "Helper",
                                    "Arn"
                                ]
                            },
                            "Format": "CSV",
                            "Prefix": "InventoryDestinationPrefix"
                        },
                        "Enabled": true,
                        "IncludedObjectVersions": "Current",
                        "Prefix": "InventoryConfigurationPrefix",
                        "ScheduleFrequency": "Weekly"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: S3 Bucket with Inventory and Analytics Configurations
Resources:
  Helper:
    Type: 'AWS::S3::Bucket'
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      AnalyticsConfigurations:
        - Id: AnalyticsConfigurationId
          StorageClassAnalysis:
            DataExport:
              Destination:
                BucketArn: !GetAtt
                  - Helper
                  - Arn
                Format: CSV
                Prefix: AnalyticsDestinationPrefix
              OutputSchemaVersion: V_1
          Prefix: AnalyticsConfigurationPrefix
          TagFilters:
            - Key: AnalyticsTagKey
              Value: AnalyticsTagValue
      InventoryConfigurations:
        - Id: InventoryConfigurationId
          Destination:
            BucketArn: !GetAtt
              - Helper
              - Arn
            Format: CSV
            Prefix: InventoryDestinationPrefix
          Enabled: true
          IncludedObjectVersions: Current
          Prefix: InventoryConfigurationPrefix
          ScheduleFrequency: Weekly
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IntelligentTieringConfiguration

InventoryTableConfiguration
