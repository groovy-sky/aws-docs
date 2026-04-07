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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/HostedZoneConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/HostedZoneConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/HostedZoneConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HostedZone

HostedZoneFailureReasons
