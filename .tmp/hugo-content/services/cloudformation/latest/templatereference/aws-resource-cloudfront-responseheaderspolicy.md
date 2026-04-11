This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy

A response headers policy.

A response headers policy contains information about a set of HTTP response headers.

After you create a response headers policy, you can use its ID to attach it to one or more
cache behaviors in a CloudFront distribution. When it's attached to a cache behavior, the
response headers policy affects the HTTP headers that CloudFront includes in HTTP responses to
requests that match the cache behavior. CloudFront adds or removes response headers according
to the configuration of the response headers policy.

For more information, see [Adding or removing HTTP headers in CloudFront responses](../../../amazoncloudfront/latest/developerguide/modifying-response-headers.md) in the
_Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::ResponseHeadersPolicy",
  "Properties" : {
      "ResponseHeadersPolicyConfig" : ResponseHeadersPolicyConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::ResponseHeadersPolicy
Properties:
  ResponseHeadersPolicyConfig:
    ResponseHeadersPolicyConfig

```

## Properties

`ResponseHeadersPolicyConfig`

A response headers policy configuration.

_Required_: Yes

_Type_: [ResponseHeadersPolicyConfig](aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the response headers policy ID. For example:
`57f99797-3b20-4e1b-a728-27972a74082a`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The unique identifier for the response headers policy. For example:
`57f99797-3b20-4e1b-a728-27972a74082a`.

`LastModifiedTime`

The date and time when the response headers policy was last modified.

## Examples

### Create a response headers policy

The following example creates a response headers policy.

#### JSON

```json

{
  "Resources": {
    "MyResponseHeadersPolicy": {
      "Type": "AWS::CloudFront::ResponseHeadersPolicy",
      "Properties": {
        "ResponseHeadersPolicyConfig": {
          "Name": {
            "Fn::Join": [
              "-",
              [
                "canary-response-header-policy",
                {
                  "Fn::Select": [
                    0,
                    {
                      "Fn::Split": [
                        "-",
                        {
                          "Fn::Select": [
                            2,
                            {
                              "Fn::Split": [
                                "/",
                                {
                                  "Ref": "AWS::StackId"
                                }
                              ]
                            }
                          ]
                        }
                      ]
                    }
                  ]
                }
              ]
            ]
          },
          "Comment": "Sample Comment",
          "CorsConfig": {
            "AccessControlAllowCredentials": true,
            "AccessControlAllowHeaders": {
              "Items": ["Header1"]
            },
            "AccessControlAllowMethods": {
              "Items": ["GET", "POST"]
            },
            "AccessControlAllowOrigins": {
              "Items": ["example1.com", "example2.com", "example3.com"]
            },
            "AccessControlExposeHeaders": {
              "Items": [
                "ExposeHeader1",
                "ExposeHeader2",
                "ExposeHeader3",
                "ExposeHeader4"
              ]
            },
            "AccessControlMaxAgeSec": 1200,
            "OriginOverride": true
          },
          "CustomHeadersConfig": {
            "Items": [
              {
                "Header": "HeaderX",
                "Override": true,
                "Value": "ValueX"
              },
              {
                "Header": "HeaderY",
                "Override": true,
                "Value": "ValueY"
              }
            ]
          },
          "RemoveHeadersConfig": {
            "Items": [
              {
                "Header": "RemoveHeaderX"
              },
              {
                "Header": "RemoveHeaderY"
              }
            ]
          },
          "SecurityHeadersConfig": {
            "ContentSecurityPolicy": {
              "ContentSecurityPolicy": "MyPolicy",
              "Override": true
            },
            "ContentTypeOptions": {
              "Override": true
            },
            "FrameOptions": {
              "FrameOption": "SAMEORIGIN",
              "Override": true
            },
            "ReferrerPolicy": {
              "ReferrerPolicy": "origin-when-cross-origin",
              "Override": true
            },
            "StrictTransportSecurity": {
              "AccessControlMaxAgeSec": 2400,
              "IncludeSubdomains": true,
              "Override": true,
              "Preload": true
            },
            "XSSProtection": {
              "ModeBlock": true,
              "Override": true,
              "Protection": true
            }
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  ABCDE1FGHI:
  MyWaitConditionHandle:
    Type: AWS::CloudFormation::WaitConditionHandle
  MyResponseHeadersPolicy:
    Type: AWS::CloudFront::ResponseHeadersPolicy
    Properties:
      ResponseHeadersPolicyConfig:
        Name: !Join
          - "-"
          - - "canary-response-header-policy"
            - !Select
              - 0
              - !Split
                - "-"
                - !Select
                  - 2
                  - !Split
                    - "/"
                    - !Ref "AWS::StackId"
        Comment: "Sample Comment"
        CorsConfig:
          AccessControlAllowCredentials: true
          AccessControlAllowHeaders:
            Items:
              - "Header1"
          AccessControlAllowMethods:
            Items:
              - "GET"
              - "POST"
          AccessControlAllowOrigins:
            Items:
              - "example1.com"
              - "example2.com"
              - "example3.com"
          AccessControlExposeHeaders:
            Items:
              - "ExposeHeader1"
              - "ExposeHeader2"
              - "ExposeHeader3"
              - "ExposeHeader4"
          AccessControlMaxAgeSec: 1200
          OriginOverride: true
        CustomHeadersConfig:
          Items:
            - Header: "HeaderX"
              Override: true
              Value: "ValueX"
            - Header: "HeaderY"
              Override: true
              Value: "ValueY"
        RemoveHeadersConfig:
          Items:
            - Header: "RemoveHeaderX"
            - Header: "RemoveHeaderY"
        SecurityHeadersConfig:
          ContentSecurityPolicy:
            ContentSecurityPolicy: "MyPolicy"
            Override: true
          ContentTypeOptions:
            Override: true
          FrameOptions:
            FrameOption: "SAMEORIGIN"
            Override: true
          ReferrerPolicy:
            ReferrerPolicy: "origin-when-cross-origin"
            Override: true
          StrictTransportSecurity:
            AccessControlMaxAgeSec: 2400
            IncludeSubdomains: true
            Override: true
            Preload: true
          XSSProtection:
            ModeBlock: true
            Override: true
            Protection: true
  RHPolicyWaiterCustomResource:
    Type: AWS::CloudFormation::CustomResource
    DeletionPolicy: Retain
    Properties:
      ServiceToken: !ImportValue RHPolicyLambdaWaiter
      WaitSeconds: 300
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisStreamConfig

AccessControlAllowHeaders

All content copied from https://docs.aws.amazon.com/.
