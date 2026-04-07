This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Outposts::AccessPoint

The AWS::S3Outposts::AccessPoint resource specifies an access point and associates it with
the specified Amazon S3 on Outposts bucket. For more information, see [Managing data access\
with Amazon S3 access points](../../../s3/latest/userguide/access-points.md).

###### Note

S3 on Outposts supports only VPC-style access points.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Outposts::AccessPoint",
  "Properties" : {
      "Bucket" : String,
      "Name" : String,
      "Policy" : Json,
      "VpcConfiguration" : VpcConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::S3Outposts::AccessPoint
Properties:
  Bucket: String
  Name: String
  Policy: Json
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`Bucket`

The Amazon Resource Name (ARN) of the S3 on Outposts bucket that is associated with this
access point.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[^:]+:s3-outposts:[a-zA-Z0-9\-]+:\d{12}:outpost\/[^:]+\/bucket\/[^:]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of this access point.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$`

_Minimum_: `3`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The access point policy associated with this access point.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

The virtual private cloud (VPC) configuration for this access point, if one exists.

_Required_: Yes

_Type_: [VpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3outposts-accesspoint-vpcconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the access point ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

This resource contains the details of the S3 on Outposts bucket access point ARN. This
resource is read-only.

## Examples

### Creating an access point with an access point policy for your Amazon S3 on Outposts using CloudFormation

The following example shows how you can create an S3 on Outposts bucket and S3 on
Outposts access point in the same CFN stack.

###### Note

To create an access point, you must already have an S3 on Outposts bucket ARN. This
means that you must create your Outposts bucket before or at the same time as you create
the access point.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Bucket, no tags, no lifecycle configuration with access point",
    "Resources": {
        "ExampleS3OutpostsBucket": {
            "Type": "AWS::S3Outposts::Bucket",
            "Properties": {
                "BucketName": "DOC-EXAMPLE-BUCKET",
                "OutpostId": "op-01ac5d28a6a232904"
            }
        },
        "ExampleS3OutpostsAccessPoint": {
            "Type": "AWS::S3Outposts::AccessPoint",
            "Properties": {
                "Bucket": {
                    "Ref": "ExampleS3OutpostsBucket"
                },
                "Name": "ExampleAccessPoint",
                "VpcConfiguration": {
                    "VpcID": "vpc-12345"
                },
                "Policy": {
                    "Version":"2012-10-17",
                    "ID":"AccessPointPolicy",
                    "Statement":[{
                        "Sid":"st1",
                        "Effect":"Allow",
                        "Principal":{"AWS":"arn:aws:iam::123456789012:root"},
                        "Action":"s3-outposts:*",
                        "Resource": "arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-01ac5d28a6a232904/accesspoint/ExampleAccessPoint"
                    }]
                }

            }
        }
    },
    "Outputs": {
        "ExampleS3OutpostsBucketARN": {
            "Description": "The ARN of ExampleS3OutpostsBucket",
            "Value": { "Ref": "ExampleS3OutpostsBucket" }
        },
        "ExampleS3OutpostsAccessPointARN": {
            "Description": "The ARN of ExampleS3OutpostsAccessPoint",
            "Value": {"Ref": "ExampleS3OutpostsAccessPoint" }
        },
        "ExampleS3OutpostsStackID": {
            "Description": "The stack ID",
            "Value": { "Ref": "AWS::StackID" },
            "Export": { "Name": {"Fn::Sub": "${AWS::StackName}-StackID"}}
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Bucket, no tags, no lifecycle configuration with access point
Resources:
  ExampleS3OutpostsBucket:
    Type: AWS::S3Outposts::Bucket
    Properties:
      BucketName: DOC-EXAMPLE-BUCKET
      OutpostId: op-01ac5d28a6a232904
  ExampleS3OutpostsAccessPoint:
    Type: AWS::S3Outposts::AccessPoint
    Properties:
      Bucket:
        Ref: ExampleS3OutpostsBucket
      Name: ExampleAccessPoint
      VpcConfiguration:
        VpcID: vpc-12345
      Policy:
        Version: '2012-10-17'
        ID: AccessPointPolicy
        Statement:
        - Sid: st1
          Effect: Allow
          Principal:
            AWS: arn:aws:iam::123456789012:root
          Action: s3-outposts:*
          Resource: arn:aws:s3-outposts:us-east-1:1234567890:outpost/op-01ac5d28a6a232904/accesspoint/ExampleAccessPoint
Outputs:
  ExampleS3OutpostsBucketARN:
    Description: The ARN of ExampleS3OutpostsBucket
    Value:
      Ref: ExampleS3OutpostsBucket
  ExampleS3OutpostsAccessPointARN:
    Description: The ARN of ExampleS3OutpostsAccessPoint
    Value:
      Ref: ExampleS3OutpostsAccessPoint
  ExampleS3OutpostsStackID:
    Description: The stack ID
    Value:
      Ref: AWS::StackID
    Export:
      Name:
        Fn::Sub: "${AWS::StackName}-StackID"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon S3 on Outposts

VpcConfiguration
