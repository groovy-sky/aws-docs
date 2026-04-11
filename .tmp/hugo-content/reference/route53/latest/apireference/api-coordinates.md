# Coordinates

A complex type that lists the coordinates for a geoproximity resource record.

## Contents

**Latitude**

Specifies a coordinate of the north–south position of a geographic point on the surface of
the Earth (-90 - 90).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 6.

Pattern: `[-+]?[0-9]{1,2}(\.[0-9]{0,2})?`

Required: Yes

**Longitude**

Specifies a coordinate of the east–west position of a geographic point on the surface of
the Earth (-180 - 180).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 7.

Pattern: `[-+]?[0-9]{1,3}(\.[0-9]{0,2})?`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/coordinates.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/coordinates.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/coordinates.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CollectionSummary

DelegationSet
