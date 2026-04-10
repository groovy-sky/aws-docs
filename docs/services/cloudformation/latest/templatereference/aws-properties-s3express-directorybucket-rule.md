This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::DirectoryBucket Rule

Specifies lifecycle rules for an Amazon S3 bucket. For more information, see [Put Bucket\
Lifecycle Configuration](../../../s3/latest/api/restbucketputlifecycle.md) in the _Amazon S3 API Reference_. For
examples, see [Put Bucket Lifecycle Configuration Examples](../../../s3/latest/api/api-putbucketlifecycleconfiguration.md#API_PutBucketLifecycleConfiguration_Examples).

You must specify at least one of the following properties:
`AbortIncompleteMultipartUpload`, or `ExpirationInDays`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AbortIncompleteMultipartUpload" : AbortIncompleteMultipartUpload,
  "ExpirationInDays" : Integer,
  "Id" : String,
  "ObjectSizeGreaterThan" : String,
  "ObjectSizeLessThan" : String,
  "Prefix" : String,
  "Status" : String
}

```

### YAML

```yaml

  AbortIncompleteMultipartUpload:
    AbortIncompleteMultipartUpload
  ExpirationInDays: Integer
  Id: String
  ObjectSizeGreaterThan: String
  ObjectSizeLessThan: String
  Prefix: String
  Status: String

```

## Properties

`AbortIncompleteMultipartUpload`

Specifies the days since the initiation of an incomplete multipart upload that Amazon S3
will wait before permanently removing all parts of the upload.

_Required_: No

_Type_: [AbortIncompleteMultipartUpload](aws-properties-s3express-directorybucket-abortincompletemultipartupload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpirationInDays`

Indicates the number of days after creation when objects are deleted from Amazon S3 and
Amazon S3 Glacier. If you specify an expiration and transition time, you must use the same
time unit for both properties (either in days or by date). The expiration time must also be
later than the transition time.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

Unique identifier for the rule. The value can't be longer than 255 characters.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectSizeGreaterThan`

Specifies the minimum object size in bytes for this rule to apply to. Objects must be
larger than this value in bytes. For more information about size based rules, see [Lifecycle configuration using size-based rules](../../../s3/latest/userguide/lifecycle-configuration-examples.md#lc-size-rules) in the _Amazon S3 User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `[0-9]+`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectSizeLessThan`

Specifies the maximum object size in bytes for this rule to apply to. Objects must be
smaller than this value in bytes. For more information about sized based rules, see [Lifecycle configuration using size-based rules](../../../s3/latest/userguide/lifecycle-configuration-examples.md#lc-size-rules) in the _Amazon S3 User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `[0-9]+`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

Object key prefix that identifies one or more objects to which this rule applies.

###### Important

Replacement must be made for object keys containing special characters (such as carriage returns) when using
XML requests. For more information, see [XML related object key constraints](../../../s3/latest/userguide/object-keys.md#object-key-xml-related-constraints).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

If `Enabled`, the rule is currently being applied. If `Disabled`, the rule is
not currently being applied.

_Required_: Yes

_Type_: String

_Allowed values_: `Enabled | Disabled`

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricsConfiguration

ServerSideEncryptionByDefault

All content copied from https://docs.aws.amazon.com/.
