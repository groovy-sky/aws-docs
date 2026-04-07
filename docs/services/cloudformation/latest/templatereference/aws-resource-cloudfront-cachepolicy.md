This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::CachePolicy

A cache policy.

When it's attached to a cache behavior, the cache policy determines the
following:

- The values that CloudFront includes in the cache key. These values can include HTTP
headers, cookies, and URL query strings. CloudFront uses the cache key to find an
object in its cache that it can return to the viewer.

- The default, minimum, and maximum time to live (TTL) values that you want
objects to stay in the CloudFront cache.

The headers, cookies, and query strings that are included in the cache key are also included
in requests that CloudFront sends to the origin. CloudFront sends a request when it can't find a
valid object in its cache that matches the request's cache key. If you want to send
values to the origin but _not_ include them in the cache key, use
`OriginRequestPolicy`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::CachePolicy",
  "Properties" : {
      "CachePolicyConfig" : CachePolicyConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::CachePolicy
Properties:
  CachePolicyConfig:
    CachePolicyConfig

```

## Properties

`CachePolicyConfig`

The cache policy configuration.

_Required_: Yes

_Type_: [CachePolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-cachepolicy-cachepolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the cache policy ID. For example:
`2766f7b2-75c5-41c6-8f06-bf4303a2f2f5`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The unique identifier for the cache policy. For example:
`2766f7b2-75c5-41c6-8f06-bf4303a2f2f5`.

`LastModifiedTime`

The date and time when the cache policy was last modified.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tags

CachePolicyConfig
