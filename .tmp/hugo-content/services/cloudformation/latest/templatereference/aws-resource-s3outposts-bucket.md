This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Outposts::Bucket

The AWS::S3Outposts::Bucket resource specifies a new Amazon S3 on Outposts bucket.
To create an S3 on Outposts bucket, you must have S3 on Outposts capacity provisioned on your Outpost.
For more information, see [Using Amazon S3 on Outposts](../../../s3/latest/userguide/s3onoutposts.md).

S3 on Outposts buckets support the following:

- Tags

- Lifecycle configuration rules for deleting expired objects

For a complete list of restrictions and Amazon S3 feature limitations on S3 on Outposts,
see [Amazon S3 on Outposts Restrictions and Limitations](../../../s3/latest/userguide/s3onoutpostsrestrictionslimitations.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Outposts::Bucket",
  "Properties" : {
      "BucketName" : String,
      "LifecycleConfiguration" : LifecycleConfiguration,
      "OutpostId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::S3Outposts::Bucket
Properties:
  BucketName: String
  LifecycleConfiguration:
    LifecycleConfiguration
  OutpostId: String
  Tags:
    - Tag

```

## Properties

`BucketName`

A name for the S3 on Outposts bucket. If you don't specify a name, AWS CloudFormation generates a
unique ID and uses that ID for the bucket name.
The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-)
and must follow [Amazon S3 bucket restrictions and limitations](../../../s3/latest/userguide/bucketrestrictions.md).
For more information, see [Bucket \
naming rules](../../../s3/latest/userguide/bucketrestrictions.md#bucketnamingrules).

###### Important

If you specify a name, you can't perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you need to
replace the resource, specify a new name.

_Required_: Yes

_Type_: String

_Pattern_: `(?=^.{3,63}$)(?!^(\d+\.)+\d+$)(^(([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])\.)*([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])$)`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LifecycleConfiguration`

Creates a new lifecycle configuration for the S3 on Outposts bucket or replaces an existing
lifecycle configuration. Outposts buckets only support lifecycle configurations that delete/expire objects
after a certain period of time and abort incomplete multipart uploads.

_Required_: No

_Type_: [LifecycleConfiguration](aws-properties-s3outposts-bucket-lifecycleconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutpostId`

The ID of the Outpost of the specified bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^(op-[a-f0-9]{17}|\d{12}|ec2)$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Sets the tags for an S3 on Outposts bucket. For more information, see [Using Amazon S3 on Outposts](../../../s3/latest/userguide/s3onoutposts.md).

Use tags to organize your AWS bill to reflect your own cost structure. To do this, sign up to get your
AWS account bill with tag key values included. Then, to see the cost of combined resources, organize your
billing information according to resources with the same tag key values. For example, you can tag several
resources with a specific application name, and then organize your billing information to see the total cost
of that application across several services. For more information, see
[Cost allocation and \
tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md).

###### Note

Within a bucket, if you add a tag that has the same key as an existing tag, the new value overwrites
the old value. For more information, see [Using cost allocation and bucket tags](../../../s3/latest/userguide/costalloctagging.md).

To use this resource, you must have permissions to perform the
`s3-outposts:PutBucketTagging`. The S3 on Outposts bucket owner has this
permission by default and can grant this permission to others. For more information about
permissions, see [Permissions \
Related to Bucket Subresource Operations](../../../s3/latest/userguide/using-with-s3-actions.md#using-with-s3-actions-related-to-bucket-subresources) and [Managing access permissions to your Amazon S3 resources](../../../s3/latest/userguide/s3-access-control.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-s3outposts-bucket-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the S3 on Outposts bucket Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the ARN of the specified bucket.

Example: `arn:aws:s3Outposts:::DOC-EXAMPLE-BUCKET`

## Examples

The following examples create an Amazon S3 on Outposts bucket using AWS
CloudFormation.

- [Create an S3 on Outposts bucket](#aws-resource-s3outposts-bucket--examples--Create_an_S3_on_Outposts_bucket)

- [Creates an S3 on Outposts bucket with tags](#aws-resource-s3outposts-bucket--examples--Creates_an_S3_on_Outposts_bucket_with_tags)

- [Creates an S3 on Outposts bucket with tags and lifecycle configuration](#aws-resource-s3outposts-bucket--examples--Creates_an_S3_on_Outposts_bucket_with_tags_and_lifecycle_configuration)

### Create an S3 on Outposts bucket

The following example creates an S3 on Outposts bucket without tags or lifecycle configuration.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Bucket, no tags, no lifecycle configuration",
    "Resources": {
        "ExampleS3OutpostsBucket": {
            "Type": "AWS::S3Outposts::Bucket",
            "Properties": {
                "BucketName": "DOC-EXAMPLE-BUCKET",
                "OutpostId": "op-01ac5d28a6a232904"
            }
        }
    },
    "Outputs": {
        "ExampleS3OutpostsBucketARN": {
            "Description": "The ARN of ExampleS3OutpostsBucket",
            "Value": {
                "Ref": "ExampleS3OutpostsBucket"
            }
        },
        "ExampleS3OutpostsStackID": {
            "Description": "The stack ID",
            "Value": {
                "Ref": "AWS::StackID"
            },
            "Export": {
                "Name": {
                    "Fn::Sub": "${AWS::StackName}-StackID"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Bucket, no tags, no lifecycle configuration
Resources:
  ExampleS3OutpostsBucket:
    Type: AWS::S3Outposts::Bucket
    Properties:
      BucketName: DOC-EXAMPLE-BUCKET
      OutpostId: op-01ac5d28a6a232904
Outputs:
  ExampleS3OutpostsBucketARN:
    Description: The ARN of ExampleS3OutpostsBucket
    Value:
      Ref: ExampleS3OutpostsBucket
  ExampleS3OutpostsStackID:
    Description: The stack ID
    Value:
      Ref: AWS::StackID
    Export:
      Name:
        Fn::Sub: "${AWS::StackName}-StackID"

```

### Creates an S3 on Outposts bucket with tags

The following example creates an S3 on Outposts bucket with tags and no lifecycle configuration.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Bucket, tags, no lifecycle configuration",
    "Resources": {
        "ExampleS3OutpostsBucket": {
            "Type": "AWS::S3Outposts::Bucket",
            "Properties": {
                "BucketName": "DOC-EXAMPLE-BUCKET",
                "OutpostId": "op-01ac5d28a6a232904",
                "Tags": [
                    {
                        "Key": "stage",
                        "Value": "beta"
                    },
                    {
                        "Key": "purpose",
                        "Value": "testing"
                    }
                ]
            }
        }
    },
    "Outputs": {
        "ExampleS3OutpostsBucketARN": {
            "Description": "The ARN of ExampleS3OutpostsBucket",
            "Value": {
                "Ref": "ExampleS3OutpostsBucket"
            }
        },
        "ExampleS3OutpostsStackID": {
            "Description": "The stack ID",
            "Value": {
                "Ref": "AWS::StackID"
            },
            "Export": {
                "Name": {
                    "Fn::Sub": "${AWS::StackName}-StackID"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Bucket, tags, no lifecycle configuration
Resources:
  ExampleS3OutpostsBucket:
    Type: AWS::S3Outposts::Bucket
    Properties:
      BucketName: DOC-EXAMPLE-BUCKET
      OutpostId: op-01ac5d28a6a232904
      Tags:
      - Key: stage
        Value: beta
      - Key: purpose
        Value: testing
Outputs:
  ExampleS3OutpostsBucketARN:
    Description: The ARN of ExampleS3OutpostsBucket
    Value:
      Ref: ExampleS3OutpostsBucket
  ExampleS3OutpostsStackID:
    Description: The stack ID
    Value:
      Ref: AWS::StackID
    Export:
      Name:
        Fn::Sub: "${AWS::StackName}-StackID"
```

### Creates an S3 on Outposts bucket with tags and lifecycle configuration

The following example creates an S3 on Outposts bucket with tags and four lifecycle configuration rules.
Three of the four lifecycle rules are disabled.

###### Note

All lifecycle rules must have values for either `ExpirationInDays`,
`ExpirationDate`, or `DaysAfterInitiation` for `AbortIncompleteMultipartUpload`
to be valid.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Bucket, tags, lifecycle configuration",
    "Resources": {
        "ExampleS3OutpostsBucket": {
            "Type": "AWS::S3Outposts::Bucket",
            "Properties": {
                "BucketName": "DOC-EXAMPLE-BUCKET",
                "OutpostId": "op-01ac5d28a6a232904",
                "Tags": [
                    {
                        "Key": "stage",
                        "Value": "beta"
                    },
                    {
                        "Key": "purpose",
                        "Value": "testing"
                    }
                ],
                "LifecycleConfiguration": {
                    "Rules": [
                        {
                            "ExpirationInDays": 2,
                            "ID": "rule1",
                            "Status": "Enabled"
                        },
                        {
                            "AbortIncompleteMultipartUpload": {
                                "DaysAfterInitiation": 2
                            },
                            "ID": "rule2",
                            "Status": "Disabled",
                            "Filter": {
                                "AndOperator": {
                                    "Prefix": "st",
                                    "Tags": [
                                        {
                                            "Key": "purpose",
                                            "Value": "testing"
                                        }
                                    ]
                                }
                            }
                        },
                        {
                            "ExpirationDate": "2020-02-25T00:00:00Z",
                            "ID": "rule3",
                            "Status": "Disabled",
                            "Filter": {
                                "Tag": {
                                    "Key": "stage",
                                    "Value": "beta"
                                }
                            }
                        },
                        {
                            "ExpirationInDays": 4,
                            "ID": "rule4",
                            "Status": "Disabled",
                            "Filter": {
                                "Prefix": "st"
                            }
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "ExampleS3OutpostsBucketARN": {
            "Description": "The ARN of ExampleS3OutpostsBucket",
            "Value": {
                "Ref": "ExampleS3OutpostsBucket"
            }
        },
        "ExampleStackID": {
            "Description": "The stack ID",
            "Value": {
                "Ref": "AWS::StackID"
            },
            "Export": {
                "Name": {
                    "Fn::Sub": "${AWS::StackName}-StackID"
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
AWSTemplateFormatVersion: '2010-09-09'
Description: Bucket, tags, lifecycle configuration
Resources:
  ExampleS3OutpostsBucket:
    Type: AWS::S3Outposts::Bucket
    Properties:
      BucketName: DOC-EXAMPLE-BUCKET
      OutpostId: op-01ac5d28a6a232904
      Tags:
      - Key: stage
        Value: beta
      - Key: purpose
        Value: testing
      LifecycleConfiguration:
        Rules:
        - ExpirationInDays: 2
          ID: rule1
          Status: Enabled
        - AbortIncompleteMultipartUpload:
            DaysAfterInitiation: 2
          ID: rule2
          Status: Disabled
          Filter:
            AndOperator:
              Prefix: st
              Tags:
              - Key: purpose
                Value: testing
        - ExpirationDate: '2020-02-25T00:00:00Z'
          ID: rule3
          Status: Disabled
          Filter:
            Tag:
              Key: stage
              Value: beta
        - ExpirationInDays: 4
          ID: rule4
          Status: Disabled
          Filter:
            Prefix: st
Outputs:
  ExampleS3OutpostsBucketARN:
    Description: The ARN of ExampleS3OutpostsBucket
    Value:
      Ref: ExampleS3OutpostsBucket
  ExampleStackID:
    Description: The stack ID
    Value:
      Ref: AWS::StackID
    Export:
      Name:
        Fn::Sub: "${AWS::StackName}-StackID"

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfiguration

AbortIncompleteMultipartUpload

All content copied from https://docs.aws.amazon.com/.
