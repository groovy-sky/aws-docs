This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy ResponseHeadersPolicyConfig

A response headers policy configuration.

A response headers policy configuration contains metadata about the response headers policy,
and configurations for sets of HTTP response headers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comment" : String,
  "CorsConfig" : CorsConfig,
  "CustomHeadersConfig" : CustomHeadersConfig,
  "Name" : String,
  "RemoveHeadersConfig" : RemoveHeadersConfig,
  "SecurityHeadersConfig" : SecurityHeadersConfig,
  "ServerTimingHeadersConfig" : ServerTimingHeadersConfig
}

```

### YAML

```yaml

  Comment: String
  CorsConfig:
    CorsConfig
  CustomHeadersConfig:
    CustomHeadersConfig
  Name: String
  RemoveHeadersConfig:
    RemoveHeadersConfig
  SecurityHeadersConfig:
    SecurityHeadersConfig
  ServerTimingHeadersConfig:
    ServerTimingHeadersConfig

```

## Properties

`Comment`

A comment to describe the response headers policy.

The comment cannot be longer than 128 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CorsConfig`

A configuration for a set of HTTP response headers that are used for cross-origin
resource sharing (CORS).

_Required_: No

_Type_: [CorsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-responseheaderspolicy-corsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomHeadersConfig`

A configuration for a set of custom HTTP response headers.

_Required_: No

_Type_: [CustomHeadersConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-responseheaderspolicy-customheadersconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name to identify the response headers policy.

The name must be unique for response headers policies in this AWS account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveHeadersConfig`

A configuration for a set of HTTP headers to remove from the HTTP response.

_Required_: No

_Type_: [RemoveHeadersConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-responseheaderspolicy-removeheadersconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityHeadersConfig`

A configuration for a set of security-related HTTP response headers.

_Required_: No

_Type_: [SecurityHeadersConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-responseheaderspolicy-securityheadersconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerTimingHeadersConfig`

A configuration for enabling the `Server-Timing` header in HTTP responses
sent from CloudFront.

_Required_: No

_Type_: [ServerTimingHeadersConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-responseheaderspolicy-servertimingheadersconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RemoveHeadersConfig

SecurityHeadersConfig
