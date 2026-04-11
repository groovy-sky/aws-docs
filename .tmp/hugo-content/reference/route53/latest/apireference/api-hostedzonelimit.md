# HostedZoneLimit

A complex type that contains the type of limit that you specified in the request and
the current value for that limit.

## Contents

**Type**

The limit that you requested. Valid values include the following:

- **MAX\_RRSETS\_BY\_ZONE**: The maximum number of
records that you can create in the specified hosted zone.

- **MAX\_VPCS\_ASSOCIATED\_BY\_ZONE**: The maximum
number of Amazon VPCs that you can associate with the specified private hosted
zone.

Type: String

Valid Values: `MAX_RRSETS_BY_ZONE | MAX_VPCS_ASSOCIATED_BY_ZONE`

Required: Yes

**Value**

The current value for the limit that is specified by `Type`.

Type: Long

Valid Range: Minimum value of 1.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/hostedzonelimit.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/hostedzonelimit.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/hostedzonelimit.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

HostedZoneFeatures

HostedZoneOwner
