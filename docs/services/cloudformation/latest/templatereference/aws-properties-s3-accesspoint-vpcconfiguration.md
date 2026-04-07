This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::AccessPoint VpcConfiguration

The Virtual Private Cloud (VPC) configuration for this access point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VpcId" : String
}

```

### YAML

```yaml

  VpcId: String

```

## Properties

`VpcId`

If this field is specified, the access point will only allow connections from the
specified VPC ID.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Create an S3 Access Point restricted to a VPC

The following example creates an Amazon S3 access point restricted to a virtual
private cloud (VPC). For more information, see [Configuring IAM policies for\
using access points](../../../s3/latest/userguide/access-points-policies.md) in the _Amazon S3 User Guide_.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket"
        },
        "S3BucketPolicy": {
            "Type": "AWS::S3::BucketPolicy",
            "Properties": {
                "Bucket": {
                    "Ref": "S3Bucket"
                },
                "PolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Action": "*",
                            "Effect": "Allow",
                            "Resource": [
                                {
                                    "Fn::GetAtt": [
                                        "S3Bucket",
                                        "Arn"
                                    ]
                                },
                                {
                                    "Fn::Join": [
                                        "",
                                        [
                                            {
                                                "Fn::GetAtt": [
                                                    "S3Bucket",
                                                    "Arn"
                                                ]
                                            },
                                            "/*"
                                        ]
                                    ]
                                }
                            ],
                            "Principal": {
                                "AWS": "*"
                            },
                            "Condition": {
                                "StringEquals": {
                                    "s3:DataAccessPointAccount": {
                                        "Ref": "AWS::AccountId"
                                    }
                                }
                            }
                        }
                    ]
                }
            }
        },
        "VPC": {
            "Type": "AWS::EC2::VPC",
            "Properties": {
                "CidrBlock": "10.0.0.0/16"
            }
        },
        "S3AccessPoint": {
            "Type": "AWS::S3::AccessPoint",
            "Properties": {
                "Bucket": {
                    "Ref": "S3Bucket"
                },
                "Name": "my-access-point",
                "VpcConfiguration": {
                    "VpcId": {
                        "Ref": "VPC"
                    }
                },
                "PublicAccessBlockConfiguration": {
                    "BlockPublicAcls": true,
                    "IgnorePublicAcls": true,
                    "BlockPublicPolicy": true,
                    "RestrictPublicBuckets": true
                }
            }
        }
    },
    "Outputs": {
        "S3AccessPointArn": {
            "Value": {
                "Ref": "S3AccessPoint"
            },
            "Description": "ARN of the sample Amazon S3 access point."
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  S3Bucket:
    Type: AWS::S3::Bucket
  S3BucketPolicy:
    Type: AWS::S3::BucketPolicy
    Properties:
      Bucket:
        Ref: S3Bucket
      PolicyDocument:
        Version: 2012-10-17
        Statement:
          - Action: "*"
            Effect: Allow
            Resource:
              - Fn::GetAtt:
                  - S3Bucket
                  - Arn
              - Fn::Join:
                  - ""
                  - - Fn::GetAtt:
                        - S3Bucket
                        - Arn
                    - /*
            Principal:
              AWS: "*"
            Condition:
              StringEquals:
                s3:DataAccessPointAccount:
                  Ref: AWS::AccountId
  VPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 10.0.0.0/16
  S3AccessPoint:
    Type: AWS::S3::AccessPoint
    Properties:
      Bucket:
        Ref: S3Bucket
      Name: my-access-point
      VpcConfiguration:
        VpcId:
          Ref: VPC
      PublicAccessBlockConfiguration:
        BlockPublicAcls: true
        IgnorePublicAcls: true
        BlockPublicPolicy: true
        RestrictPublicBuckets: true
Outputs:
  S3AccessPointArn:
    Value:
      Ref: S3AccessPoint
    Description: ARN of the sample Amazon S3 access point.
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::S3::Bucket
