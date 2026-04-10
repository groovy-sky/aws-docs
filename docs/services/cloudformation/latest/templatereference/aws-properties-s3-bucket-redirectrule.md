This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket RedirectRule

Specifies how requests are redirected. In the event of an error, you can specify a different error
code to return.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HostName" : String,
  "HttpRedirectCode" : String,
  "Protocol" : String,
  "ReplaceKeyPrefixWith" : String,
  "ReplaceKeyWith" : String
}

```

### YAML

```yaml

  HostName: String
  HttpRedirectCode: String
  Protocol: String
  ReplaceKeyPrefixWith: String
  ReplaceKeyWith: String

```

## Properties

`HostName`

The host name to use in the redirect request.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpRedirectCode`

The HTTP redirect code to use on the response. Not required if one of the siblings is
present.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

Protocol to use when redirecting requests. The default is the protocol that is used in the original
request.

_Required_: No

_Type_: String

_Allowed values_: `http | https`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplaceKeyPrefixWith`

The object key prefix to use in the redirect request. For example, to redirect requests for all
pages with prefix `docs/` (objects in the `docs/` folder) to
`documents/`, you can set a condition block with `KeyPrefixEquals` set to
`docs/` and in the Redirect set `ReplaceKeyPrefixWith` to
`/documents`. Not required if one of the siblings is present. Can be present only if
`ReplaceKeyWith` is not provided.

###### Important

Replacement must be made for object keys containing special characters (such as carriage returns) when using
XML requests. For more information, see [XML related object key constraints](../../../s3/latest/userguide/object-keys.md#object-key-xml-related-constraints).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplaceKeyWith`

The specific object key to use in the redirect request. For example, redirect request to
`error.html`. Not required if one of the siblings is present. Can be present only if
`ReplaceKeyPrefixWith` is not provided.

###### Important

Replacement must be made for object keys containing special characters (such as carriage returns) when using
XML requests. For more information, see [XML related object key constraints](../../../s3/latest/userguide/object-keys.md#object-key-xml-related-constraints).

_Required_: No

_Type_: String

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

RedirectAllRequestsTo

ReplicaModifications

All content copied from https://docs.aws.amazon.com/.
