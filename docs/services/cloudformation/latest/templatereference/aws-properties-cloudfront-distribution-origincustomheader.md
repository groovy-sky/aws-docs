This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution OriginCustomHeader

A complex type that contains `HeaderName` and `HeaderValue`
elements, if any, for this distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderName" : String,
  "HeaderValue" : String
}

```

### YAML

```yaml

  HeaderName: String
  HeaderValue: String

```

## Properties

`HeaderName`

The name of a header that you want CloudFront to send to your origin. For more information,
see [Adding\
Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/forward-custom-headers.html) in the _Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderValue`

The value for the header that you specified in the `HeaderName`
field.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [OriginCustomHeader](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_OriginCustomHeader.html) in the _Amazon CloudFront API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Origin

OriginGroup
