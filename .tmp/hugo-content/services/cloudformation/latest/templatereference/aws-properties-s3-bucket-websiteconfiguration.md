This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket WebsiteConfiguration

Specifies website configuration parameters for an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorDocument" : String,
  "IndexDocument" : String,
  "RedirectAllRequestsTo" : RedirectAllRequestsTo,
  "RoutingRules" : [ RoutingRule, ... ]
}

```

### YAML

```yaml

  ErrorDocument: String
  IndexDocument: String
  RedirectAllRequestsTo:
    RedirectAllRequestsTo
  RoutingRules:
    - RoutingRule

```

## Properties

`ErrorDocument`

The name of the error document for the website.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexDocument`

The name of the index document for the website.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedirectAllRequestsTo`

The redirect behavior for every request to this bucket's website endpoint.

###### Important

If you specify this property, you can't specify any other property.

_Required_: No

_Type_: [RedirectAllRequestsTo](aws-properties-s3-bucket-redirectallrequeststo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingRules`

Rules that define when a redirect is applied and the redirect behavior.

_Required_: No

_Type_: Array of [RoutingRule](aws-properties-s3-bucket-routingrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Configure a static website with a routing rule

In this example, `AWS::S3::Bucket's Fn::GetAtt` values are used to provide
outputs. If an HTTP 404 error occurs, the routing rule redirects requests to an EC2
instance and inserts the object key prefix `report-404/` in the redirect. For
example, if you request a page called `out1/ExamplePage.html` and it results in an
HTTP 404 error, the request is routed to a page called
`report-404/ExamplePage.html` on the specified instance. For all other HTTP
error codes, `error.html` is returned.

This example also specifies a metrics configuration called `EntireBucket`
that enables CloudWatch request metrics at the bucket level.

#### JSON

```json

{
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "PublicRead",
                "BucketName": "public-bucket",
                "MetricsConfigurations": [
                    {
                        "Id": "EntireBucket"
                    }
                ],
                "WebsiteConfiguration": {
                    "IndexDocument": "index.html",
                    "ErrorDocument": "error.html",
                    "RoutingRules": [
                        {
                            "RoutingRuleCondition": {
                                "HttpErrorCodeReturnedEquals": "404",
                                "KeyPrefixEquals": "out1/"
                            },
                            "RedirectRule": {
                                "HostName": "ec2-11-22-333-44.compute-1.amazonaws.com",
                                "ReplaceKeyPrefixWith": "report-404/"
                            }
                        }
                    ]
                }
            },
            "DeletionPolicy": "Retain"
        }
    },
    "Outputs": {
        "WebsiteURL": {
            "Value": {
                "Fn::GetAtt": [
                    "S3Bucket",
                    "WebsiteURL"
                ]
            },
            "Description": "URL for website hosted on S3"
        },
        "S3BucketSecureURL": {
            "Value": {
                "Fn::Join": [
                    "",
                    [
                        "https://",
                        {
                            "Fn::GetAtt": [
                                "S3Bucket",
                                "DomainName"
                            ]
                        }
                    ]
                ]
            },
            "Description": "Name of S3 bucket to hold website content"
        }
    }
}
```

#### YAML

```yaml

Resources:
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      AccessControl: PublicRead
      BucketName: public-bucket
      MetricsConfigurations:
        - Id: EntireBucket
      WebsiteConfiguration:
        IndexDocument: index.html
        ErrorDocument: error.html
        RoutingRules:
          - RoutingRuleCondition:
              HttpErrorCodeReturnedEquals: '404'
              KeyPrefixEquals: out1/
            RedirectRule:
              HostName: ec2-11-22-333-44.compute-1.amazonaws.com
              ReplaceKeyPrefixWith: report-404/
    DeletionPolicy: Retain
Outputs:
  WebsiteURL:
    Value: !GetAtt
      - S3Bucket
      - WebsiteURL
    Description: URL for website hosted on S3
  S3BucketSecureURL:
    Value: !Join
      - ''
      - - 'https://'
        - !GetAtt
          - S3Bucket
          - DomainName
    Description: Name of S3 bucket to hold website content
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VersioningConfiguration

AWS::S3::BucketPolicy

All content copied from https://docs.aws.amazon.com/.
