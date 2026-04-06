This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket AnalyticsConfiguration

Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Prefix" : String,
  "StorageClassAnalysis" : StorageClassAnalysis,
  "TagFilters" : [ TagFilter, ... ]
}

```

### YAML

```yaml

  Id: String
  Prefix: String
  StorageClassAnalysis:
    StorageClassAnalysis
  TagFilters:
    - TagFilter

```

## Properties

`Id`

The ID that identifies the analytics configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix that an object must have to be included in the analytics results.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageClassAnalysis`

Contains data related to access patterns to be collected and made available to analyze the
tradeoffs between different storage classes.

_Required_: Yes

_Type_: [StorageClassAnalysis](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-storageclassanalysis.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagFilters`

The tags to use when evaluating an analytics filter.

The analytics only includes objects that meet the filter's criteria. If no filter is
specified, all of the contents of the bucket are included in the analysis.

_Required_: No

_Type_: Array of [TagFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-tagfilter.html)

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

AccessControlTranslation

BlockedEncryptionTypes
