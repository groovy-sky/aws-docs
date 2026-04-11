This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution Restrictions

A complex type that identifies ways in which you want to restrict distribution of your
content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GeoRestriction" : GeoRestriction
}

```

### YAML

```yaml

  GeoRestriction:
    GeoRestriction

```

## Properties

`GeoRestriction`

A complex type that controls the countries in which your content is distributed. CloudFront
determines the location of your users using `MaxMind` GeoIP databases. To disable geo restriction, remove the [Restrictions](../userguide/aws-properties-cloudfront-distribution-distributionconfig.md#cfn-cloudfront-distribution-distributionconfig-restrictions) property from your stack template.

_Required_: Yes

_Type_: [GeoRestriction](aws-properties-cloudfront-distribution-georestriction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Restrictions](../../../../reference/cloudfront/latest/apireference/api-restrictions.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterDefinition

S3OriginConfig

All content copied from https://docs.aws.amazon.com/.
