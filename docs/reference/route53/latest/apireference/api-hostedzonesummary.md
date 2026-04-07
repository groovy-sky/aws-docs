# HostedZoneSummary

In the response to a `ListHostedZonesByVPC` request, the
`HostedZoneSummaries` element contains one `HostedZoneSummary`
element for each hosted zone that the specified Amazon VPC is associated with. Each
`HostedZoneSummary` element contains the hosted zone name and ID, and
information about who owns the hosted zone.

## Contents

**HostedZoneId**

The Route 53 hosted zone ID of a private hosted zone that the specified VPC is
associated with.

Type: String

Length Constraints: Maximum length of 32.

Required: Yes

**Name**

The name of the private hosted zone, such as `example.com`.

Type: String

Length Constraints: Maximum length of 1024.

Required: Yes

**Owner**

The owner of a private hosted zone that the specified VPC is associated with. The
owner can be either an AWS account or an AWS
service.

Type: [HostedZoneOwner](https://docs.aws.amazon.com/Route53/latest/APIReference/API_HostedZoneOwner.html) object

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/HostedZoneSummary)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/HostedZoneSummary)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/HostedZoneSummary)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HostedZoneOwner

KeySigningKey
