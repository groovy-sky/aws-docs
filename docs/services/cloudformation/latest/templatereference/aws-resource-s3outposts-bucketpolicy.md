This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Outposts::BucketPolicy

This resource applies a bucket policy to an Amazon S3 on Outposts bucket.

If you are using an identity other than the root user of the AWS account
that owns the S3 on Outposts bucket, the calling identity must have
the `s3-outposts:PutBucketPolicy` permissions on the specified
Outposts bucket and belong to the bucket owner's account in order to use
this resource.

If you don't have `s3-outposts:PutBucketPolicy` permissions,
S3 on Outposts returns a `403 Access Denied` error.

###### Important

The root user of the AWS account that owns an Outposts bucket can
_always_ use this resource, even if the policy explicitly denies the
root user the ability to perform actions on this resource.

For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and
[Access Policy Language Overview](../../../s3/latest/userguide/access-policy-language-overview.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Outposts::BucketPolicy",
  "Properties" : {
      "Bucket" : String,
      "PolicyDocument" : Json
    }
}

```

### YAML

```yaml

Type: AWS::S3Outposts::BucketPolicy
Properties:
  Bucket: String
  PolicyDocument: Json

```

## Properties

`Bucket`

The name of the Amazon S3 Outposts bucket to which the policy applies.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[^:]+:s3-outposts:[a-zA-Z0-9\-]+:\d{12}:outpost\/[^:]+\/bucket\/[^:]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyDocument`

A policy document containing permissions to add to the specified bucket. In IAM, you must
provide policy documents in JSON format. However, in CloudFormation, you can provide the
policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting
it to IAM. For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language\
Overview](../../../s3/latest/userguide/access-policy-language-overview.md).

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the S3 on Outposts bucket Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create an Amazon S3 on Outposts bucket with a bucket policy

The following example creates an Amazon S3 on Outposts bucket and adds a bucket policy to that bucket.

###### Note

To add a bucket policy to a bucket, you must create your Outposts bucket before or
at the same time as you add your bucket policy.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Bucket with no tags + Bucket Policy",
    "Resources": {
        "ExampleS3OutpostsBucket": {
            "Type": "AWS::S3Outposts::Bucket",
            "Properties": {
                "BucketName": "DOC-EXAMPLE-BUCKET",
                "OutpostId": "op-01ac5d28a6a232904"
            }
        },
        "ExampleS3OutpostsBucketPolicy": {
            "Type": "AWS::S3Outposts::BucketPolicy",
            "Properties": {
                "Bucket": {
                    "Ref": "ExampleS3OutpostsBucket"
                },
                "PolicyDocument": {
                    "Version": "2012-10-17",
                    "ID": "BucketPolicy",
                    "Statement": [
                        {
                            "Sid": "st1",
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": "arn:aws:iam::123456789012:root"
                            },
                            "Action": "s3-outposts:*",
                            "Resource": "arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-01ac5d28a6a232904/bucket/DOC-EXAMPLE-BUCKET"
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
        "ExampleS3OutpostsBucketPolicyARN": {
            "Description": "The ARN of the BucketPolicy",
            "Value": {
                "Ref": "ExampleS3OutpostsBucketPolicy"
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

AWSTemplateFormatVersion: 2010-09-09
Description: Bucket with no tags + Bucket Policy
Resources:
  ExampleS3OutpostsBucket:
    Type: 'AWS::S3Outposts::Bucket'
    Properties:
      BucketName: DOC-EXAMPLE-BUCKET
      OutpostId: op-01ac5d28a6a232904
  ExampleS3OutpostsBucketPolicy:
    Type: 'AWS::S3Outposts::BucketPolicy'
    Properties:
      Bucket: !Ref ExampleS3OutpostsBucket
      PolicyDocument:
        Version: 2012-10-17
        ID: BucketPolicy
        Statement:
          - Sid: st1
            Effect: Allow
            Principal:
              AWS: 'arn:aws:iam::123456789012:root'
            Action: 's3-outposts:*'
            Resource: >-
              arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-01ac5d28a6a232904/bucket/DOC-EXAMPLE-BUCKET
Outputs:
  ExampleS3OutpostsBucketARN:
    Description: The ARN of ExampleS3OutpostsBucket
    Value: !Ref ExampleS3OutpostsBucket
  ExampleS3OutpostsBucketPolicyARN:
    Description: The ARN of the BucketPolicy
    Value: !Ref ExampleS3OutpostsBucketPolicy
  ExampleS3OutpostsStackID:
    Description: The stack ID
    Value: !Ref 'AWS::StackID'
    Export:
      Name: !Sub '${AWS::StackName}-StackID'

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::S3Outposts::Endpoint
