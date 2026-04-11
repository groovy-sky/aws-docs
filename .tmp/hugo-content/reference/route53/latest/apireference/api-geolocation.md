# GeoLocation

A complex type that contains information about a geographic location.

## Contents

**ContinentCode**

The two-letter code for the continent.

Amazon Route 53 supports the following continent codes:

- **AF**: Africa

- **AN**: Antarctica

- **AS**: Asia

- **EU**: Europe

- **OC**: Oceania

- **NA**: North America

- **SA**: South America

Constraint: Specifying `ContinentCode` with either `CountryCode`
or `SubdivisionCode` returns an `InvalidInput` error.

Type: String

Length Constraints: Fixed length of 2.

Required: No

**CountryCode**

For geolocation resource record sets, the two-letter code for a country.

Amazon Route 53 uses the two-letter country codes that are specified in [ISO standard 3166-1\
alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).

Route 53 also supports the country code **UA** for
Ukraine.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2.

Required: No

**SubdivisionCode**

For geolocation resource record sets, the two-letter code for a state of the United
States. Route 53 doesn't support any other values for `SubdivisionCode`. For
a list of state abbreviations, see [Appendix B: Two–Letter State and Possession Abbreviations](https://pe.usps.com/text/pub28/28apb.htm) on the United
States Postal Service website.

If you specify `subdivisioncode`, you must also specify `US` for
`CountryCode`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 3.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/geolocation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/geolocation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/geolocation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DNSSECStatus

GeoLocationDetails
