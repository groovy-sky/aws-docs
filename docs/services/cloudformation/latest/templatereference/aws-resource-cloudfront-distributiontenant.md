This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::DistributionTenant

The distribution tenant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::DistributionTenant",
  "Properties" : {
      "ConnectionGroupId" : String,
      "Customizations" : Customizations,
      "DistributionId" : String,
      "Domains" : [ String, ... ],
      "Enabled" : Boolean,
      "ManagedCertificateRequest" : ManagedCertificateRequest,
      "Name" : String,
      "Parameters" : [ Parameter, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::DistributionTenant
Properties:
  ConnectionGroupId: String
  Customizations:
    Customizations
  DistributionId: String
  Domains:
    - String
  Enabled: Boolean
  ManagedCertificateRequest:
    ManagedCertificateRequest
  Name: String
  Parameters:
    - Parameter
  Tags:
    - Tag

```

## Properties

`ConnectionGroupId`

The ID of the connection group for the distribution tenant. If you don't specify a connection group, CloudFront uses the default connection group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Customizations`

Customizations for the distribution tenant. For each distribution tenant, you can specify the geographic restrictions, and the Amazon Resource Names (ARNs) for the ACM certificate and AWS WAF web ACL. These are specific values that you can override or disable from the multi-tenant distribution that was used to create the distribution tenant.

_Required_: No

_Type_: [Customizations](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distributiontenant-customizations.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DistributionId`

The ID of the multi-tenant distribution.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domains`

The domains associated with the distribution tenant.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether the distribution tenant is in an enabled state. If disabled, the distribution tenant won't serve traffic.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedCertificateRequest`

An object that represents the request for the Amazon CloudFront managed ACM certificate.

_Required_: No

_Type_: [ManagedCertificateRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the distribution tenant.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

A list of parameter values to add to the resource. A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution.

_Required_: No

_Type_: Array of [Parameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distributiontenant-parameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A complex type that contains zero or more `Tag` elements.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distributiontenant-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the distribution tenant.

`CreatedTime`

The date and time when the distribution tenant was created.

`DomainResults`

Property description not available.

`ETag`

The current version of the distribution tenant.

`Id`

The ID of the distribution tenant.

`LastModifiedTime`

The date and time when the distribution tenant was updated.

`Status`

The status of the distribution tenant.

## Examples

- [Create a distribution tenant that inherits its certificate](#aws-resource-cloudfront-distributiontenant--examples--Create_a_distribution_tenant_that_inherits_its_certificate)

- [Create a distribution tenant with its own certificate](#aws-resource-cloudfront-distributiontenant--examples--Create_a_distribution_tenant_with_its_own_certificate)

- [Create a CloudFront hosted distribution tenant](#aws-resource-cloudfront-distributiontenant--examples--Create_a_hosted_distribution_tenant)

- [Create a self hosted distribution tenant](#aws-resource-cloudfront-distributiontenant--examples--Create_a_self_hosted_distribution_tenant)

### Create a distribution tenant that inherits its certificate

The following example specifies a distribution tenant that inherits its certificate from its parent multi-tenant distribution.

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
    },
    "MyDistributionTenant": {
      "Type": "AWS::CloudFront::DistributionTenant",
      "Properties": {
        "Domains": [
          "my-distribution-tenant.example.com"
        ],
        "DistributionId": "MyMultiTenantDistribution.Id",
        "Name": "MyDistributionTenant",
        "Enabled": true,
        "Parameters": [
          {
            "Name": "tenantName",
            "Value": "first-user"
          }
        ]
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

  MyDistributionTenant:
    Type: AWS::CloudFront::DistributionTenant
    Properties:
      Domains:
        - my-distribution-tenant.example.com
      DistributionId: !GetAtt MyMultiTenantDistribution.Id
      Name: MyDistributionTenant
      Enabled: true
      Parameters:
        - Name: tenantName
          Value: first-user
```

### Create a distribution tenant with its own certificate

The following example specifies a distribution tenant with its own certificate.

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
    },
    "MyDistributionTenant": {
      "Type": "AWS::CloudFront::DistributionTenant",
      "Properties": {
        "Domains": [
          "my-distribution-tenant.example.com"
        ],
        "DistributionId": "MyMultiTenantDistribution.Id",
        "Name": "MyDistributionTenant",
        "Enabled": true,
        "Parameters": [
          {
            "Name": "tenantName",
            "Value": "first-user"
          }
        ]
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

  MyDistributionTenant:
    Type: AWS::CloudFront::DistributionTenant
    Properties:
      Domains:
        - my-distribution-tenant.example.com
      DistributionId: !GetAtt MyMultiTenantDistribution.Id
      Name: MyDistributionTenant
      Enabled: true
      Customizations:
        Certificate:
          Arn: arn:aws:acm:us-east-1:123456789012:certificate/1954f095-11b6-4daf-9952-0c308a00944d
      Parameters:
        - Name: tenantName
          Value: first-user
```

### Create a CloudFront hosted distribution tenant

The following example specifies a CloudFront hosted distribution tenant.

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
    },
    "MyConnectionGroup": {
      "Type": "AWS::CloudFront::ConnectionGroup",
      "Properties": {
        "Name": "cf-hosted-connection-group-cfn"
      }
    },
    "RecordSet": {
      "Type": "AWS::Route53::RecordSet",
      "Properties": {
        "Name": "my-distribution-tenant.example.com",
        "Type": "CNAME",
        "HostedZoneId": "Z06559422OQIFCZO0EORK",
        "TTL": 300,
        "ResourceRecords": [
          "MyConnectionGroup.RoutingEndpoint"
        ]
      }
    },
    "MyDistributionTenant": {
      "Type": "AWS::CloudFront::DistributionTenant",
      "Properties": {
        "ConnectionGroupId": "MyConnectionGroup.Id",
        "Domains": [
          "my-distribution-tenant.example.com"
        ],
        "DistributionId": "MyMultiTenantDistribution.Id",
        "Name": "MyDistributionTenant",
        "Enabled": true,
        "ManagedCertificateRequest": {
          "ValidationTokenHost": "cloudfront"
        },
        "Parameters": [
          {
            "Name": "tenantName",
            "Value": "first-user"
          }
        ]
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

  MyConnectionGroup:
    Type: AWS::CloudFront::ConnectionGroup
    Properties:
      Name: cf-hosted-connection-group-cfn

  RecordSet:
    Type: AWS::Route53::RecordSet
    Properties:
      Name: my-distribution-tenant.example.com
      Type: CNAME
      HostedZoneId: Z06559422OQIFCZO0EORK
      TTL: 300
      ResourceRecords:
        - !GetAtt MyConnectionGroup.RoutingEndpoint

  MyDistributionTenant:
    Type: AWS::CloudFront::DistributionTenant
    Properties:
      ConnectionGroupId: !GetAtt MyConnectionGroup.Id
      Domains:
        - my-distribution-tenant.example.com
      DistributionId: !GetAtt MyMultiTenantDistribution.Id
      Name: MyDistributionTenant
      Enabled: true
      ManagedCertificateRequest:
        ValidationTokenHost: cloudfront
      Parameters:
        - Name: tenantName
          Value: first-user
```

### Create a self hosted distribution tenant

The following example specifies a self hosted distribution tenant.

###### Important

You must set up token validation for the distribution tenant when using this option. For more information, see [Request certificates for your CloudFront distribution tenant](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/managed-cloudfront-certificates.html) in the _Amazon CloudFront Developer Guide_.

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
    },
    "MyConnectionGroup": {
      "Type": "AWS::CloudFront::ConnectionGroup",
      "Properties": {
        "Name": "cf-hosted-connection-group-cfn"
      }
    },
    "RecordSet": {
      "Type": "AWS::Route53::RecordSet",
      "Properties": {
        "Name": "my-distribution-tenant.example.com",
        "Type": "CNAME",
        "HostedZoneId": "Z06559422OQIFCZO0EORK",
        "TTL": 300,
        "ResourceRecords": [
          "MyConnectionGroup.RoutingEndpoint"
        ]
      }
    },
    "MyDistributionTenant": {
      "Type": "AWS::CloudFront::DistributionTenant",
      "Properties": {
        "ConnectionGroupId": "MyConnectionGroup.Id",
        "Domains": [
          "my-distribution-tenant.example.com"
        ],
        "DistributionId": "MyMultiTenantDistribution.Id",
        "Name": "MyDistributionTenant",
        "Enabled": true,
        "ManagedCertificateRequest": {
          "ValidationTokenHost": "self-hosted",
          "PrimaryDomainName": "my-distribution-tenant.example.com"
        },
        "Parameters": [
          {
            "Name": "tenantName",
            "Value": "first-user"
          }
        ]
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

  MyConnectionGroup:
    Type: AWS::CloudFront::ConnectionGroup
    Properties:
      Name: cf-hosted-connection-group-cfn

  RecordSet:
    Type: AWS::Route53::RecordSet
    Properties:
      Name: my-distribution-tenant.example.com
      Type: CNAME
      HostedZoneId: Z06559422OQIFCZO0EORK
      TTL: 300
      ResourceRecords:
        - !GetAtt MyConnectionGroup.RoutingEndpoint

  MyDistributionTenant:
    Type: AWS::CloudFront::DistributionTenant
    Properties:
      ConnectionGroupId: !GetAtt MyConnectionGroup.Id
      Domains:
        - my-distribution-tenant.example.com
      DistributionId: !GetAtt MyMultiTenantDistribution.Id
      Name: MyDistributionTenant
      Enabled: true
      ManagedCertificateRequest:
        ValidationTokenHost: self-hosted
        PrimaryDomainName: my-distribution-tenant.example.com
      Parameters:
        - Name: tenantName
          Value: first-user
```

## See also

- [Understand how multi-tenant distributions work](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md) in the _Amazon CloudFront Developer Guide_

- [Request certificates for your CloudFront distribution tenant](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/managed-cloudfront-certificates.html) in the _Amazon CloudFront Developer Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcOriginConfig

Certificate
