This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::OriginRequestPolicy OriginRequestPolicyConfig

An origin request policy configuration.

This configuration determines the values that CloudFront includes in requests that it sends
to the origin. Each request that CloudFront sends to the origin includes the following:

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
  "Comment" : String,
  "CookiesConfig" : CookiesConfig,
  "HeadersConfig" : HeadersConfig,
  "Name" : String,
  "QueryStringsConfig" : QueryStringsConfig
}

```

### YAML

```yaml

  Comment: String
  CookiesConfig:
    CookiesConfig
  HeadersConfig:
    HeadersConfig
  Name: String
  QueryStringsConfig:
    QueryStringsConfig

```

## Properties

`Comment`

A comment to describe the origin request policy. The comment cannot be longer than 128
characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CookiesConfig`

The cookies from viewer requests to include in origin requests.

_Required_: Yes

_Type_: [CookiesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-originrequestpolicy-cookiesconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeadersConfig`

The HTTP headers to include in origin requests. These can include headers from viewer
requests and additional headers added by CloudFront.

_Required_: Yes

_Type_: [HeadersConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-originrequestpolicy-headersconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A unique name to identify the origin request policy.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStringsConfig`

The URL query strings from viewer requests to include in origin requests.

_Required_: Yes

_Type_: [QueryStringsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-originrequestpolicy-querystringsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HeadersConfig

QueryStringsConfig
