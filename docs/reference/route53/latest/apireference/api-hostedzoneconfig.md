# HostedZoneConfig

A complex type that contains an optional comment about your hosted zone. If you don't
want to specify a comment, omit both the `HostedZoneConfig` and
`Comment` elements.

## Contents

**Comment**

Any comments that you want to include about the hosted zone.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**PrivateZone**

A value that indicates whether this is a private hosted zone.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/hostedzoneconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/hostedzoneconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/hostedzoneconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

HostedZone

HostedZoneFailureReasons
