This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution

A distribution tells CloudFront where you want content to be delivered from, and the details
about how to track and manage content delivery.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::Distribution",
  "Properties" : {
      "DistributionConfig" : DistributionConfig,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::Distribution
Properties:
  DistributionConfig:
    DistributionConfig
  Tags:
    - Tag

```

## Properties

`DistributionConfig`

The distribution's configuration.

_Required_: Yes

_Type_: [DistributionConfig](aws-properties-cloudfront-distribution-distributionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A complex type that contains zero or more `Tag` elements.

_Required_: No

_Type_: Array of [Tag](aws-properties-cloudfront-distribution-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the CloudFront
distribution ID. For example: `E27LVI50CSW06W`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainName`

The domain name of the resource, such as `d111111abcdef8.cloudfront.net`.

`Id`

The distribution's identifier. For example: `E1U5RQF7T870K0`.

## Examples

- [Create a standard distribution](#aws-resource-cloudfront-distribution--examples--Create_a_standard_distribution)

- [Create a multi-tenant distribution without a certificate](#aws-resource-cloudfront-distribution--examples--Create_a_multi-tenant_distribution_without_a_certificate)

- [Create a multi-tenant distribution with a wildcard certificate](#aws-resource-cloudfront-distribution--examples--Create_a_multi-tenant_distribution_with_a_wildcard_certificate)

### Create a standard distribution

The following example specifies a standard distribution and assigns it a single tag.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "cloudfrontdistribution": {
            "Type": "AWS::CloudFront::Distribution",
            "Properties": {
                "DistributionConfig": {
                    "CacheBehaviors": [
                        {
                            "LambdaFunctionAssociations": [
                                {
                                    "EventType": "string-value",
                                    "LambdaFunctionARN": "string-value"
                                }
                            ]
                        }
                    ],
                    "DefaultCacheBehavior": {
                        "LambdaFunctionAssociations": [
                            {
                                "EventType": "string-value",
                                "LambdaFunctionARN": "string-value"
                            }
                        ]
                    },
                    "IPV6Enabled": "boolean-value",
                    "Origins": [
                        {
                            "CustomOriginConfig": {
                                "OriginKeepaliveTimeout": "integer-value",
                                "OriginReadTimeout": "integer-value"
                            }
                        }
                    ]
                },
                "Tags": [
                    {
                        "Key": "string-value",
                        "Value": "string-value"
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
Resources:
  cloudfrontdistribution:
    Type: AWS::CloudFront::Distribution
    Properties:
      DistributionConfig:
        CacheBehaviors:
          - LambdaFunctionAssociations:
              - EventType: string-value
                LambdaFunctionARN: string-value
        DefaultCacheBehavior:
          LambdaFunctionAssociations:
            - EventType: string-value
              LambdaFunctionARN: string-value
        IPV6Enabled: boolean-value
        Origins:
          - CustomOriginConfig:
              OriginKeepaliveTimeout: integer-value
              OriginReadTimeout: integer-value
      Tags:
        - Key: string-value
          Value: string-value
```

### Create a multi-tenant distribution without a certificate

The following example specifies a basic multi-tenant distribution without a certificate.

#### JSON

```json

{
  "Resources": {
    "MyMultiTenantDistribution": {
      "Type": "AWS::CloudFront::Distribution",
      "Properties": {
        "DistributionConfig": {
          "ConnectionMode": "tenant-only",
          "TenantConfig": {
            "ParameterDefinitions": [
              {
                "Name": "tenantName",
                "Definition": {
                  "StringSchema": {
                    "Comment": "Tenant name",
                    "DefaultValue": "root",
                    "Required": false
                  }
                }
              }
            ]
          },
          "DefaultCacheBehavior": {
            "TargetOriginId": "MyBucket.Arn",
            "ViewerProtocolPolicy": "allow-all",
            "AllowedMethods": [
              "GET",
              "HEAD"
            ],
            "CachePolicyId": "658327ea-f89d-4fab-a63d-7e88639e58f6"
          },
          "Enabled": true,
          "Origins": [
            {
              "DomainName": "MyBucket.RegionalDomainName",
              "Id": "MyBucket.Arn",
              "OriginPath": "/{{tenantName}}",
              "S3OriginConfig": {
                "OriginAccessIdentity": ""
              }
            }
          ]
        }
      }
    },
    "MyBucket": {
      "Type": "AWS::S3::Bucket",
      "Properties": {
        "BucketName": "amzn-s3-demo-bucket",
        "BucketEncryption": {
          "ServerSideEncryptionConfiguration": [
            {
              "ServerSideEncryptionByDefault": {
                "SSEAlgorithm": "aws:kms",
                "KMSMasterKeyID": "alias/aws/s3"
              }
            }
          ]
        },
        "PublicAccessBlockConfiguration": {
          "IgnorePublicAcls": true,
          "RestrictPublicBuckets": true
        }
      }
    },
    "MyBucketBucketPolicy": {
      "Type": "AWS::S3::BucketPolicy",
      "Properties": {
        "Bucket": "MyBucket",
        "PolicyDocument": {
          "Id": "RequireEncryptionInTransit",
          "Version": "2012-10-17",
          "Statement": [
            {
              "Principal": "*",
              "Action": "*",
              "Effect": "Deny",
              "Resource": [
                "MyBucket.Arn",
                "${MyBucket.Arn}/*"
              ],
              "Condition": {
                "Bool": {
                  "aws:SecureTransport": "false"
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyMultiTenantDistribution:
    Type: AWS::CloudFront::Distribution
    Properties:
      DistributionConfig:
        ConnectionMode: tenant-only
        TenantConfig:
          ParameterDefinitions:
            - Name: tenantName
              Definition:
                StringSchema:
                  Comment: "Tenant name"
                  DefaultValue: "root"
                  Required: false
        DefaultCacheBehavior:
          TargetOriginId: !GetAtt MyBucket.Arn
          ViewerProtocolPolicy: allow-all
          AllowedMethods:
            - GET
            - HEAD
          CachePolicyId: 658327ea-f89d-4fab-a63d-7e88639e58f6 # CachingOptimized PolicyId
        Enabled: true
        Origins:
          - DomainName: !GetAtt MyBucket.RegionalDomainName
            Id: !GetAtt MyBucket.Arn
            OriginPath: "/{{tenantName}}"
            S3OriginConfig:
              OriginAccessIdentity: ""

  MyBucket:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: amzn-s3-demo-bucket
      BucketEncryption:
        ServerSideEncryptionConfiguration:
          - ServerSideEncryptionByDefault:
              SSEAlgorithm: aws:kms
              KMSMasterKeyID: alias/aws/s3
      PublicAccessBlockConfiguration:
        IgnorePublicAcls: true
        RestrictPublicBuckets: true
  MyBucketBucketPolicy:
    Type: AWS::S3::BucketPolicy
    Properties:
      Bucket: !Ref MyBucket
      PolicyDocument:
        Id: RequireEncryptionInTransit
        Version: '2012-10-17'
        Statement:
          - Principal: '*'
            Action: '*'
            Effect: Deny
            Resource:
              - !GetAtt MyBucket.Arn
              - !Sub ${MyBucket.Arn}/*
            Condition:
              Bool:
                aws:SecureTransport: 'false'
```

### Create a multi-tenant distribution with a wildcard certificate

The following example specifies a basic multi-tenant distribution with a wildcard certificate.

#### JSON

```json

{
  "Resources": {
    "MyMultiTenantDistribution": {
      "Type": "AWS::CloudFront::Distribution",
      "Properties": {
        "DistributionConfig": {
          "ConnectionMode": "tenant-only",
          "ViewerCertificate": {
            "AcmCertificateArn": "arn:aws:acm:us-east-1:123456789012:certificate/1954f095-11b6-4daf-9952-0c308a00944d",
            "SslSupportMethod": "sni-only"
          },
          "TenantConfig": {
            "ParameterDefinitions": [
              {
                "Name": "tenantName",
                "Definition": {
                  "StringSchema": {
                    "Comment": "Tenant name",
                    "DefaultValue": "root",
                    "Required": false
                  }
                }
              }
            ]
          },
          "DefaultCacheBehavior": {
            "TargetOriginId": "MyBucket.Arn",
            "ViewerProtocolPolicy": "allow-all",
            "AllowedMethods": [
              "GET",
              "HEAD"
            ],
            "CachePolicyId": "658327ea-f89d-4fab-a63d-7e88639e58f6"
          },
          "Enabled": true,
          "Origins": [
            {
              "DomainName": "MyBucket.RegionalDomainName",
              "Id": "MyBucket.Arn",
              "OriginPath": "/{{tenantName}}",
              "S3OriginConfig": {
                "OriginAccessIdentity": ""
              }
            }
          ]
        }
      }
    },
    "MyBucket": {
      "Type": "AWS::S3::Bucket",
      "Properties": {
        "BucketName": "amzn-s3-demo-bucket",
        "BucketEncryption": {
          "ServerSideEncryptionConfiguration": [
            {
              "ServerSideEncryptionByDefault": {
                "SSEAlgorithm": "aws:kms",
                "KMSMasterKeyID": "alias/aws/s3"
              }
            }
          ]
        },
        "PublicAccessBlockConfiguration": {
          "IgnorePublicAcls": true,
          "RestrictPublicBuckets": true
        }
      }
    },
    "MyBucketBucketPolicy": {
      "Type": "AWS::S3::BucketPolicy",
      "Properties": {
        "Bucket": "MyBucket",
        "PolicyDocument": {
          "Id": "RequireEncryptionInTransit",
          "Version": "2012-10-17",
          "Statement": [
            {
              "Principal": "*",
              "Action": "*",
              "Effect": "Deny",
              "Resource": [
                "MyBucket.Arn",
                "${MyBucket.Arn}/*"
              ],
              "Condition": {
                "Bool": {
                  "aws:SecureTransport": "false"
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyMultiTenantDistribution:
    Type: AWS::CloudFront::Distribution
    Properties:
      DistributionConfig:
        ConnectionMode: tenant-only
        ViewerCertificate:
          AcmCertificateArn: arn:aws:acm:us-east-1:123456789012:certificate/1954f095-11b6-4daf-9952-0c308a00944d
          SslSupportMethod: sni-only
        TenantConfig:
          ParameterDefinitions:
            - Name: tenantName
              Definition:
                StringSchema:
                  Comment: "Tenant name"
                  DefaultValue: "root"
                  Required: false
        DefaultCacheBehavior:
          TargetOriginId: !GetAtt MyBucket.Arn
          ViewerProtocolPolicy: allow-all
          AllowedMethods:
            - GET
            - HEAD
          CachePolicyId: 658327ea-f89d-4fab-a63d-7e88639e58f6 # CachingOptimized PolicyId
        Enabled: true
        Origins:
          - DomainName: !GetAtt MyBucket.RegionalDomainName
            Id: !GetAtt MyBucket.Arn
            OriginPath: "/{{tenantName}}"
            S3OriginConfig:
              OriginAccessIdentity: ""

  MyBucket:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: amzn-s3-demo-bucket
      BucketEncryption:
        ServerSideEncryptionConfiguration:
          - ServerSideEncryptionByDefault:
              SSEAlgorithm: aws:kms
              KMSMasterKeyID: alias/aws/s3
      PublicAccessBlockConfiguration:
        IgnorePublicAcls: true
        RestrictPublicBuckets: true
  MyBucketBucketPolicy:
    Type: AWS::S3::BucketPolicy
    Properties:
      Bucket: !Ref MyBucket
      PolicyDocument:
        Id: RequireEncryptionInTransit
        Version: '2012-10-17'
        Statement:
          - Principal: '*'
            Action: '*'
            Effect: Deny
            Resource:
              - !GetAtt MyBucket.Arn
              - !Sub ${MyBucket.Arn}/*
            Condition:
              Bool:
                aws:SecureTransport: 'false'
```

## See also

- [CreateDistribution](../../../../reference/cloudfront/latest/apireference/api-createdistribution.md) in the _Amazon CloudFront API Reference_

- [Understand how multi-tenant distributions work](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md) in the _Amazon CloudFront Developer Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrafficConfig

CacheBehavior

All content copied from https://docs.aws.amazon.com/.
