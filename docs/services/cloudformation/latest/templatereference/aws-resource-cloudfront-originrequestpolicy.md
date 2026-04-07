This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::OriginRequestPolicy

An origin request policy.

When it's attached to a cache behavior, the origin request policy determines the
values that CloudFront includes in requests that it sends to the origin. Each request that
CloudFront sends to the origin includes the following:

- The request body and the URL path (without the domain name) from the viewer
request.

- The headers that CloudFront automatically includes in every origin request,
including `Host`, `User-Agent`, and
`X-Amz-Cf-Id`.

- All HTTP headers, cookies, and URL query strings that are specified in the
cache policy or the origin request policy. These can include items from the
viewer request and, in the case of headers, additional ones that are added by
CloudFront.

CloudFront sends a request when it can't find an object in its cache that matches the
request. If you want to send values to the origin and also include them in the cache
key, use `CachePolicy`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::OriginRequestPolicy",
  "Properties" : {
      "OriginRequestPolicyConfig" : OriginRequestPolicyConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::OriginRequestPolicy
Properties:
  OriginRequestPolicyConfig:
    OriginRequestPolicyConfig

```

## Properties

`OriginRequestPolicyConfig`

The origin request policy configuration.

_Required_: Yes

_Type_: [OriginRequestPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the origin request policy ID. For example:
`befd7079-9bbc-4ebf-8ade-498a3694176c`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The unique identifier for the origin request policy. For example:
`befd7079-9bbc-4ebf-8ade-498a3694176c`.

`LastModifiedTime`

The date and time when the origin request policy was last modified.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OriginAccessControlConfig

CookiesConfig
