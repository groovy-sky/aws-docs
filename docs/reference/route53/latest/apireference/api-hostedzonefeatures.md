# HostedZoneFeatures

Represents the features configuration for a hosted zone, including the status of various features and any associated failure reasons.

## Contents

**AcceleratedRecoveryStatus**

The current status of accelerated recovery for the hosted zone.

Type: String

Valid Values: `ENABLING | ENABLE_FAILED | ENABLING_HOSTED_ZONE_LOCKED | ENABLED | DISABLING | DISABLE_FAILED | DISABLED | DISABLING_HOSTED_ZONE_LOCKED`

Required: No

**FailureReasons**

Information about any failures that occurred when attempting to enable or configure features for the hosted zone.

Type: [HostedZoneFailureReasons](api-hostedzonefailurereasons.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/hostedzonefeatures.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/hostedzonefeatures.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/hostedzonefeatures.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

HostedZoneFailureReasons

HostedZoneLimit
