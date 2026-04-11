This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket LifecycleConfiguration

Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see
[Object Lifecycle\
Management](../../../s3/latest/dev/object-lifecycle-mgmt.md) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ Rule, ... ],
  "TransitionDefaultMinimumObjectSize" : String
}

```

### YAML

```yaml

  Rules:
    - Rule
  TransitionDefaultMinimumObjectSize: String

```

## Properties

`Rules`

A lifecycle rule for individual objects in an Amazon S3 bucket.

_Required_: Yes

_Type_: Array of [Rule](aws-properties-s3-bucket-rule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitionDefaultMinimumObjectSize`

Indicates which default minimum object size behavior is applied to the lifecycle
configuration.

###### Note

This parameter applies to general purpose buckets only. It isn't supported for
directory bucket lifecycle configurations.

- `all_storage_classes_128K` \- Objects smaller than 128 KB will not transition to any storage class by default.

- `varies_by_storage_class` \- Objects smaller than 128 KB will
transition to Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By
default, all other storage classes will prevent transitions smaller than 128 KB.

To customize the minimum object size for any transition you can add a filter that
specifies a custom `ObjectSizeGreaterThan` or `ObjectSizeLessThan` in
the body of your transition rule. Custom filters always take precedence over the default
transition behavior.

_Required_: No

_Type_: String

_Allowed values_: `varies_by_storage_class | all_storage_classes_128K`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Manage the lifecycle for S3 objects

The following example template shows an S3 bucket with a lifecycle configuration rule.
The rule applies to all objects with the `glacier` key prefix. The objects are
transitioned to Glacier after one day, and deleted after one year.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "Private",
                "LifecycleConfiguration": {
                    "Rules": [
                        {
                            "Id": "GlacierRule",
                            "Prefix": "glacier",
                            "Status": "Enabled",
                            "ExpirationInDays": 365,
                            "Transitions": [
                                {
                                    "TransitionInDays": 1,
                                    "StorageClass": "GLACIER"
                                }
                            ]
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a lifecycle configuration."
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      AccessControl: Private
      LifecycleConfiguration:
        Rules:
          - Id: GlacierRule
            Prefix: glacier
            Status: Enabled
            ExpirationInDays: 365
            Transitions:
              - TransitionInDays: 1
                StorageClass: GLACIER
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with a lifecycle configuration.
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaConfiguration

LoggingConfiguration

All content copied from https://docs.aws.amazon.com/.
