This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::DirectoryBucket LifecycleConfiguration

Container for lifecycle rules. You can add as many as 1000 rules.

For more information see, [Creating and managing a lifecycle configuration for directory buckets](../../../s3/latest/userguide/directory-buckets-objects-lifecycle.md) in the
_Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ Rule, ... ]
}

```

### YAML

```yaml

  Rules:
    - Rule

```

## Properties

`Rules`

A lifecycle rule for individual objects in an Amazon S3 Express bucket.

_Required_: Yes

_Type_: Array of [Rule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3express-directorybucket-rule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Manage the lifecycle for S3 objects

The following example template shows an S3 directory bucket with a lifecycle
configuration rule. The rule applies to all objects with the `foo/` key
prefix. The objects are expired after seven days, and incomplete multipart uploads
are deleted 3 days after initiation.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3ExpressBucket": {
            "Type": "AWS::S3Express::DirectoryBucket",
            "Properties": {
                "LocationName": "usw2-az1",
                "DataRedundancy": "SingleAvailabilityZone",
                "LifecycleConfiguration": {
                    "Rules": [
                        {
                            "Id": "ExipiryRule",
                            "Prefix": "foo/",
                            "Status": "Enabled",
                            "ExpirationInDays": 7,
                            "AbortIncompleteMultipartUpload": {
                                "DaysAfterInitiation": 3
                            },
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3ExpressBucket"
            },
            "Description": "Name of the sample Amazon S3 Directory Bucket with a lifecycle configuration."
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  S3ExpressBucket:
    Type: 'AWS::S3Express::DirectoryBucket'
    Properties:
      LocationName: usw2-az1
      DataRedundancy: SingleAvailabilityZone
      LifecycleConfiguration:
        Rules:
          - Id: ExipiryRule
            Prefix: foo/
            Status: Enabled
            ExpirationInDays:7
            AbortIncompleteMultipartUpload:
            DaysAfterInitiation:3
Outputs:
  BucketName:
    Value: !Ref S3ExpressBucket
    Description: Name of the sample Amazon S3 Directory Bucket with a lifecycle configuration.

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BucketEncryption

MetricsConfiguration
