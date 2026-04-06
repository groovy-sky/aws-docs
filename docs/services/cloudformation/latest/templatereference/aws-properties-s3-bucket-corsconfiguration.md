This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket CorsConfiguration

Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information,
see [Enabling Cross-Origin Resource\
Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CorsRules" : [ CorsRule, ... ]
}

```

### YAML

```yaml

  CorsRules:
    - CorsRule

```

## Properties

`CorsRules`

A set of origins and methods (cross-origin access that you want to allow). You can add up to 100
rules to the configuration.

_Required_: Yes

_Type_: Array of [CorsRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-corsrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Enable cross-origin resource sharing

The following example template shows a public S3 bucket with two cross-origin resource
sharing rules.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "PublicRead",
                "CorsConfiguration": {
                    "CorsRules": [
                        {
                            "AllowedHeaders": [
                                "*"
                            ],
                            "AllowedMethods": [
                                "GET"
                            ],
                            "AllowedOrigins": [
                                "*"
                            ],
                            "ExposedHeaders": [
                                "Date"
                            ],
                            "Id": "myCORSRuleId1",
                            "MaxAge": 3600
                        },
                        {
                            "AllowedHeaders": [
                                "x-amz-*"
                            ],
                            "AllowedMethods": [
                                "DELETE"
                            ],
                            "AllowedOrigins": [
                                "http://www.example.com",
                                "http://www.example.net"
                            ],
                            "ExposedHeaders": [
                                "Connection",
                                "Server",
                                "Date"
                            ],
                            "Id": "myCORSRuleId2",
                            "MaxAge": 1800
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
            "Description": "Name of the sample Amazon S3 bucket with CORS enabled."
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
      AccessControl: PublicRead
      CorsConfiguration:
        CorsRules:
          - AllowedHeaders:
              - '*'
            AllowedMethods:
              - GET
            AllowedOrigins:
              - '*'
            ExposedHeaders:
              - Date
            Id: myCORSRuleId1
            MaxAge: 3600
          - AllowedHeaders:
              - x-amz-*
            AllowedMethods:
              - DELETE
            AllowedOrigins:
              - 'http://www.example.com'
              - 'http://www.example.net'
            ExposedHeaders:
              - Connection
              - Server
              - Date
            Id: myCORSRuleId2
            MaxAge: 1800
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with CORS enabled.
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BucketEncryption

CorsRule
