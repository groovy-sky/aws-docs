This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::AccessPoint

The AWS::S3::AccessPoint resource is an Amazon S3 resource type that you can use to access
buckets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3::AccessPoint",
  "Properties" : {
      "Bucket" : String,
      "BucketAccountId" : String,
      "Name" : String,
      "Policy" : Json,
      "PublicAccessBlockConfiguration" : PublicAccessBlockConfiguration,
      "Tags" : [ Tag, ... ],
      "VpcConfiguration" : VpcConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::S3::AccessPoint
Properties:
  Bucket: String
  BucketAccountId: String
  Name: String
  Policy: Json
  PublicAccessBlockConfiguration:
    PublicAccessBlockConfiguration
  Tags:
    - Tag
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`Bucket`

The name of the bucket associated with this access point.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketAccountId`

The AWS account ID associated with the S3 bucket associated with this
access point.

_Required_: No

_Type_: String

_Pattern_: `^\d{12}$`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of this access point. If you don't specify a name, AWS CloudFormation
generates a unique ID and uses that ID for the access point name.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

_Minimum_: `3`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The access point policy associated with this access point.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicAccessBlockConfiguration`

The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You
can enable the configuration options in any combination. For more information about when
Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: [PublicAccessBlockConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-accesspoint-publicaccessblockconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of tags that you can apply to access points. Tags are key-value pairs of metadata used to categorize your access points and control access. For more information, see [Using tags for attribute-based access control (ABAC)](../../../s3/latest/userguide/tagging.md#using-tags-for-abac).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-accesspoint-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

The Virtual Private Cloud (VPC) configuration for this access point, if one exists.

_Required_: No

_Type_: [VpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-accesspoint-vpcconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the access point name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Alias`

The alias for this access point.

`Arn`

This property contains the details of the ARN for the access point.

`Name`

The name of this access point.

`NetworkOrigin`

Indicates whether this access point allows access from the internet. If
`VpcConfiguration` is specified for this access point, then
`NetworkOrigin` is `VPC`, and the access point doesn't allow access
from the internet. Otherwise, `NetworkOrigin` is `Internet`, and the
access point allows access from the internet, subject to the access point and bucket access
policies.

_Allowed values_: `VPC` \| `Internet`

## Examples

- [Create an S3 Access Point](#aws-resource-s3-accesspoint--examples--Create_an_S3_Access_Point)

- [Create an S3 Access Point restricted to a VPC](#aws-resource-s3-accesspoint--examples--Create_an_S3_Access_Point_restricted_to_a_VPC)

### Create an S3 Access Point

The following example creates an Amazon S3 access point for the given S3 bucket. This
access point allows user `JaneDoe` to make GetObject and PutObject operations
only for bucket objects prefixed with `/janedoe`. You must include
`/object` in the resource ARN path.

For more information, see [Configuring IAM policies for\
using access points](../../../s3/latest/userguide/access-points-policies.md) and [Managing and using access\
points](../../../s3/latest/userguide/using-access-points.md) in the _Amazon S3 User Guide_.

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
        "S3AccessPoint": {
            "Type": "AWS::S3::AccessPoint",
            "Properties": {
                "Bucket": {
                    "Ref": "S3Bucket"
                },
                "Name": "my-access-point",
                "Policy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Action": [
                                "s3:GetObject",
                                "s3:PutObject"
                            ],
                            "Effect": "Allow",
                            "Resource": [
                                {
                                    "Fn::Sub": "arn:${AWS::Partition}:s3:${AWS::Region}:${AWS::AccountId}:accesspoint/my-access-point/object/janedoe/*"
                                }
                            ],
                            "Principal": {
                                "AWS": {
                                    "Fn::Sub": "arn:${AWS::Partition}:iam::${AWS::AccountId}:user/JaneDoe"
                                }
                            }
                        }
                    ]
                }
            }
        }
    },
   "Outputs": {
        "S3AccessPointArn": {
            "Value": {
                "Fn::GetAtt": ["S3AccessPoint", "Arn"]
            },
            "Description": "ARN of the sample Amazon S3 access point."
        },
        "S3AccessPointName": {
            "Value": {
                "Fn::GetAtt": ["S3AccessPoint", "Name"]
            },
            "Description": "Name of the sample Amazon S3 access point."
        },
        "S3AccessPointAlias": {
            "Value": {
                "Fn::GetAtt": ["S3AccessPoint", "Alias"]
            },
            "Description": "Alias of the sample Amazon S3 access point."
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
  S3BucketPolicy:
    Type: 'AWS::S3::BucketPolicy'
    Properties:
      Bucket: !Ref S3Bucket
      PolicyDocument:
        Version: 2012-10-17
        Statement:
          - Action: '*'
            Effect: Allow
            Resource:
              - !GetAtt
                - S3Bucket
                - Arn
              - !Join
                - ''
                - - !GetAtt
                    - S3Bucket
                    - Arn
                  - /*
            Principal:
              AWS: '*'
            Condition:
              StringEquals:
                's3:DataAccessPointAccount': !Ref 'AWS::AccountId'
  S3AccessPoint:
    Type: 'AWS::S3::AccessPoint'
    Properties:
      Bucket: !Ref S3Bucket
      Name: my-access-point
      Policy:
        Version: 2012-10-17
        Statement:
          - Action:
              - 's3:GetObject'
              - 's3:PutObject'
            Effect: Allow
            Resource:
               - !Sub  'arn:${AWS::Partition}:s3:${AWS::Region}:${AWS::AccountId}:accesspoint/my-access-point/object/janedoe/*'
            Principal:
              AWS: !Sub 'arn:${AWS::Partition}:iam::${AWS::AccountId}:user/JaneDoe'
Outputs:
  S3AccessPointArn:
    Value:
      Fn::GetAtt:
      - S3AccessPoint
      - Arn
    Description: ARN of the sample Amazon S3 access point.
  S3AccessPointName:
    Value:
      Fn::GetAtt:
      - S3AccessPoint
      - Name
    Description: Name of the sample Amazon S3 access point.
  S3AccessPointAlias:
    Value:
      Fn::GetAtt:
      - S3AccessPoint
      - Alias
    Description: Alias of the sample Amazon S3 access point.
```

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

PublicAccessBlockConfiguration
