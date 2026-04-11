This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::StreamingDistribution

This resource is deprecated.
Amazon CloudFront is deprecating real-time messaging protocol (RTMP) distributions on December 31, 2020.
For more information, [read the announcement](https://forums.aws.amazon.com/ann.jspa?annID=7356) on the Amazon CloudFront discussion forum.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::StreamingDistribution",
  "Properties" : {
      "StreamingDistributionConfig" : StreamingDistributionConfig,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::StreamingDistribution
Properties:
  StreamingDistributionConfig:
    StreamingDistributionConfig
  Tags:
    - Tag

```

## Properties

`StreamingDistributionConfig`

The current configuration information for the RTMP distribution.

_Required_: Yes

_Type_: [StreamingDistributionConfig](aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A complex type that contains zero or more `Tag` elements.

_Required_: Yes

_Type_: Array of [Tag](aws-properties-cloudfront-streamingdistribution-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the streaming distribution ID, such as
`E1E7FEN9T35R9W`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainName`

The domain name of the resource, such as `d111111abcdef8.cloudfront.net`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

XSSProtection

Logging

All content copied from https://docs.aws.amazon.com/.
