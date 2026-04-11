This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::RecordSetGroup GeoLocation

A complex type that contains information about a geographic location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContinentCode" : String,
  "CountryCode" : String,
  "SubdivisionCode" : String
}

```

### YAML

```yaml

  ContinentCode: String
  CountryCode: String
  SubdivisionCode: String

```

## Properties

`ContinentCode`

For geolocation resource record sets, a two-letter abbreviation that identifies a continent. Route 53 supports the following continent codes:

- **AF**: Africa

- **AN**: Antarctica

- **AS**: Asia

- **EU**: Europe

- **OC**: Oceania

- **NA**: North America

- **SA**: South America

Constraint: Specifying `ContinentCode` with either `CountryCode` or `SubdivisionCode` returns an
`InvalidInput` error.

_Required_: No

_Type_: String

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CountryCode`

For geolocation resource record sets, the two-letter code for a country.

Route 53 uses the two-letter country codes that are specified in
[ISO standard 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubdivisionCode`

For geolocation resource record sets, the two-letter code for a state of the United States.
Route 53 doesn't support any other values for `SubdivisionCode`. For a list of state abbreviations, see
[Appendix B: Two–Letter State and Possession Abbreviations](https://pe.usps.com/text/pub28/28apb.htm)
on the United States Postal Service website.

If you specify `subdivisioncode`, you must also specify `US` for `CountryCode`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Return values](../userguide/aws-resource-route53-recordsetgroup.md#aws-resource-route53-recordsetgroup-return-values)
in the topic
[AWS::Route53::RecordSetGroup](../userguide/aws-resource-route53-recordsetgroup.md)

- [GeoLocation](../../../../reference/route53/latest/apireference/api-geolocation.md) in the _Amazon Route 53 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Coordinates

GeoProximityLocation

All content copied from https://docs.aws.amazon.com/.
