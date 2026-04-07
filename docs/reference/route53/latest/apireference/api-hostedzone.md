# HostedZone

A complex type that contains general information about the hosted zone.

## Contents

**CallerReference**

The value that you specified for `CallerReference` when you created the
hosted zone.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**Id**

The ID that Amazon Route 53 assigned to the hosted zone when you created it.

Type: String

Length Constraints: Maximum length of 32.

Required: Yes

**Name**

The name of the domain. For public hosted zones, this is the name that you have
registered with your DNS registrar.

For information about how to specify characters other than `a-z`,
`0-9`, and `-` (hyphen) and how to specify internationalized
domain names, see [CreateHostedZone](api-createhostedzone.md).

Type: String

Length Constraints: Maximum length of 1024.

Required: Yes

**Config**

A complex type that includes the `Comment` and `PrivateZone`
elements. If you omitted the `HostedZoneConfig` and `Comment`
elements from the request, the `Config` and `Comment` elements
don't appear in the response.

Type: [HostedZoneConfig](api-hostedzoneconfig.md) object

Required: No

**Features**

The features configuration for the hosted zone, including accelerated recovery settings and status information.

Type: [HostedZoneFeatures](https://docs.aws.amazon.com/Route53/latest/APIReference/API_HostedZoneFeatures.html) object

Required: No

**LinkedService**

If the hosted zone was created by another service, the service that created the hosted
zone. When a hosted zone is created by another service, you can't edit or delete it
using Route 53.

Type: [LinkedService](https://docs.aws.amazon.com/Route53/latest/APIReference/API_LinkedService.html) object

Required: No

**ResourceRecordSetCount**

The number of resource record sets in the hosted zone.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/HostedZone)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/HostedZone)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/HostedZone)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HealthCheckObservation

HostedZoneConfig
