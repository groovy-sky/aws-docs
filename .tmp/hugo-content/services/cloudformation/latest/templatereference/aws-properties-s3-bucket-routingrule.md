This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket RoutingRule

Specifies the redirect behavior and when a redirect is applied. For more information about routing
rules, see [Configuring\
advanced conditional redirects](../../../s3/latest/dev/how-to-page-redirect.md#advanced-conditional-redirects) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RedirectRule" : RedirectRule,
  "RoutingRuleCondition" : RoutingRuleCondition
}

```

### YAML

```yaml

  RedirectRule:
    RedirectRule
  RoutingRuleCondition:
    RoutingRuleCondition

```

## Properties

`RedirectRule`

Container for redirect information. You can redirect requests to another host, to another page, or
with another protocol. In the event of an error, you can specify a different error code to
return.

_Required_: Yes

_Type_: [RedirectRule](aws-properties-s3-bucket-redirectrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingRuleCondition`

A container for describing a condition that must be met for the specified redirect to apply. For
example, 1. If request is for pages in the `/docs` folder, redirect to the
`/documents` folder. 2. If request results in HTTP error 4xx, redirect request to another
host where you might process the error.

_Required_: No

_Type_: [RoutingRuleCondition](aws-properties-s3-bucket-routingrulecondition.md)

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

ReplicationTimeValue

RoutingRuleCondition

All content copied from https://docs.aws.amazon.com/.
