This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket CorsRule

Specifies a cross-origin access rule for an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedHeaders" : [ String, ... ],
  "AllowedMethods" : [ String, ... ],
  "AllowedOrigins" : [ String, ... ],
  "ExposedHeaders" : [ String, ... ],
  "Id" : String,
  "MaxAge" : Integer
}

```

### YAML

```yaml

  AllowedHeaders:
    - String
  AllowedMethods:
    - String
  AllowedOrigins:
    - String
  ExposedHeaders:
    - String
  Id: String
  MaxAge: Integer

```

## Properties

`AllowedHeaders`

Headers that are specified in the `Access-Control-Request-Headers` header. These headers
are allowed in a preflight OPTIONS request. In response to any preflight OPTIONS request, Amazon S3 returns
any requested headers that are allowed.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedMethods`

An HTTP method that you allow the origin to run.

_Allowed values_: `GET` \| `PUT` \|
`HEAD` \| `POST` \| `DELETE`

_Required_: Yes

_Type_: Array of String

_Allowed values_: `GET | PUT | HEAD | POST | DELETE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedOrigins`

One or more origins you want customers to be able to access the bucket from.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExposedHeaders`

One or more headers in the response that you want customers to be able to access from
their applications (for example, from a JavaScript `XMLHttpRequest` object).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

A unique identifier for this rule. The value must be no more than 255 characters.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAge`

The time in seconds that your browser is to cache the preflight response for the specified
resource.

_Required_: No

_Type_: Integer

_Minimum_: `0`

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CorsConfiguration

DataExport

All content copied from https://docs.aws.amazon.com/.
