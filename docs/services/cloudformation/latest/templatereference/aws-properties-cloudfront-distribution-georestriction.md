This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution GeoRestriction

A complex type that controls the countries in which your content is distributed. CloudFront
determines the location of your users using `MaxMind` GeoIP databases. To disable geo restriction, remove the [Restrictions](../userguide/aws-properties-cloudfront-distribution-distributionconfig.md#cfn-cloudfront-distribution-distributionconfig-restrictions) property from your stack template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Locations" : [ String, ... ],
  "RestrictionType" : String
}

```

### YAML

```yaml

  Locations:
    - String
  RestrictionType: String

```

## Properties

`Locations`

A complex type that contains a `Location` element for each country in
which you want CloudFront either to distribute your content ( `whitelist`) or not
distribute your content ( `blacklist`).

The `Location` element is a two-letter, uppercase country code for a
country that you want to include in your `blacklist` or
`whitelist`. Include one `Location` element for each
country.

CloudFront and `MaxMind` both use `ISO 3166` country codes. For the
current list of countries and the corresponding codes, see `ISO
				3166-1-alpha-2` code on the _International Organization for_
_Standardization_ website. You can also refer to the country list on the
CloudFront console, which includes both country names and codes.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestrictionType`

The method that you want to use to restrict distribution of your content by country:

- `none`: No geo restriction is enabled, meaning access to content is
not restricted by client geo location.

- `blacklist`: The `Location` elements specify the
countries in which you don't want CloudFront to distribute your content.

- `whitelist`: The `Location` elements specify the
countries in which you want CloudFront to distribute your content.

_Required_: Yes

_Type_: String

_Allowed values_: `blacklist | whitelist | none`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Add georestrictions

The following examples show how to add georestrictions.

#### JSON

```json

{
  "Restrictions": {
    "GeoRestriction": {
      "RestrictionType": "whitelist",
      "Locations": [
        "ES",
        "GB"
      ]
    }
  }
}
```

#### YAML

```yaml

        Restrictions:
          GeoRestriction:
            RestrictionType: whitelist
            Locations:
              - ES
              - GB
```

## See also

- [GeoRestriction](../../../../reference/cloudfront/latest/apireference/api-georestriction.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunctionAssociation

GrpcConfig

All content copied from https://docs.aws.amazon.com/.
