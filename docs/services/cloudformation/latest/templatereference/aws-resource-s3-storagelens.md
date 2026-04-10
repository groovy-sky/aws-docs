This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens

The AWS::S3::StorageLens resource creates an Amazon S3 Storage Lens configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3::StorageLens",
  "Properties" : {
      "StorageLensConfiguration" : StorageLensConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::S3::StorageLens
Properties:
  StorageLensConfiguration:
    StorageLensConfiguration
  Tags:
    - Tag

```

## Properties

`StorageLensConfiguration`

This resource contains the details Amazon S3 Storage Lens configuration.

_Required_: Yes

_Type_: [StorageLensConfiguration](aws-properties-s3-storagelens-storagelensconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A set of tags (key–value pairs) to associate with the Storage Lens configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-s3-storagelens-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the S3 Storage Lens configuration Id, such as
`your-storage-lens-configuration-id`. For more information about using the Ref
function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

Fn::GetAtt returns a value for a specified attribute of this type. For more information,
see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md). The following are the available attributes and sample return
values.

`StorageLensConfiguration.StorageLensArn`

This property contains the details of the ARN of the S3 Storage Lens configuration. This
property is read-only.

## Examples

The following examples create an advanced S3 Storage Lens configuration that enables
advanced metrics, Amazon CloudWatch publishing, and prefix aggregation. This example
also configures a metrics export and adds tags.

### Creating an advanced S3 Storage Lens configuration

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "StorageLens Advanced Configuration Example",
    "Resources": {
        "StorageLensConfigurationExample": {
            "Type": "AWS::S3::StorageLens",
            "Properties": {
                "StorageLensConfiguration": {
                    "Id": "StorageLensAdvancedConfiguration",
                    "AccountLevel": {
                        "ActivityMetrics": {
                            "IsEnabled": true
                        },
                        "AdvancedCostOptimizationMetrics": {
                            "IsEnabled": true
                        },
                        "AdvancedDataProtectionMetrics": {
                            "IsEnabled": true
                        },
                        "DetailedStatusCodesMetrics": {
                            "IsEnabled": true
                        },
                        "BucketLevel": {
                            "ActivityMetrics": {
                                "IsEnabled": true
                            },
                            "AdvancedCostOptimizationMetrics": {
                                "IsEnabled": true
                            },
                            "AdvancedDataProtectionMetrics": {
                                "IsEnabled": true
                            },
                            "DetailedStatusCodesMetrics": {
                                "IsEnabled": true
                            },
                            "PrefixLevel": {
                                "StorageMetrics": {
                                    "IsEnabled": true,
                                    "SelectionCriteria": {
                                        "MaxDepth": 5,
                                        "MinStorageBytesPercentage": 1.23,
                                        "Delimiter": "/"
                                    }
                                }
                            }
                        }
                    },
                    "Exclude": {
                        "Buckets": [
                            {
                                "Fn::Sub": "arn:aws:s3:::source_bucket_1"
                            },
                            {
                                "Fn::Sub": "arn:aws:s3:::source_bucket_2"
                            }
                        ]
                    },
                    "IsEnabled": true,
                    "DataExport": {
                        "S3BucketDestination": {
                            "OutputSchemaVersion": "V_1",
                            "Format": "CSV",
                            "AccountId": "111122223333",
                            "Arn": {
                                "Fn::Sub": "arn:aws:s3:::destination_bucket"
                            },
                            "Prefix": "output-path-prefix",
                            "Encryption": {
                                "SSES3": {}
                            }
                        },
                        "CloudWatchMetrics": {
                            "IsEnabled": true
                        }
                    }
                },
                "Tags": [
                    {
                        "Key": "tag-key-1",
                        "Value": "tag-value-1"
                    },
                    {
                        "Key": "tag-key-2",
                        "Value": "tag-value-2"
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
Description: StorageLens Advanced Configuration Example
Resources:
  StorageLensConfigurationExample:
    Type: AWS::S3::StorageLens
    Properties:
      StorageLensConfiguration:
        Id: "StorageLensAdvancedConfiguration"
        AccountLevel:
          ActivityMetrics:
            IsEnabled: true
          AdvancedCostOptimizationMetrics:
            IsEnabled: true
          AdvancedDataProtectionMetrics:
            IsEnabled: true
          DetailedStatusCodesMetrics:
            IsEnabled: true
          BucketLevel:
            ActivityMetrics:
              IsEnabled: true
            AdvancedCostOptimizationMetrics:
              IsEnabled: true
            AdvancedDataProtectionMetrics:
              IsEnabled: true
            DetailedStatusCodesMetrics:
              IsEnabled: true
            PrefixLevel:
              StorageMetrics:
                IsEnabled: true
                SelectionCriteria:
                  MaxDepth: 5
                  MinStorageBytesPercentage: 1.23
                  Delimiter: "/"
        Exclude:
           Buckets:
            - !Sub 'arn:aws:s3:::source_bucket_1'
            - !Sub 'arn:aws:s3:::source_bucket_2'
        IsEnabled: true
        DataExport:
          S3BucketDestination:
            OutputSchemaVersion: "V_1"
            Format: "CSV"
            AccountId: "111122223333"
            Arn: !Sub 'arn:aws:s3:::destination_bucket'
            Prefix: "output-path-prefix"
            Encryption:
              SSES3: {}
          CloudWatchMetrics:
            IsEnabled: true
      Tags:
        - Key: "tag-key-1"
          Value: "tag-value-1"
        - Key: "tag-key-2"
          Value: "tag-value-2"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyStatus

AccountLevel

All content copied from https://docs.aws.amazon.com/.
