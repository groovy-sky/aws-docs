# LaunchTemplateNetworkPerformanceOptionsRequest

When you configure network performance options in your launch template, your instance
is geared for performance improvements based on the workload that it runs as soon as
it's available.

## Contents

**BandwidthWeighting**

Specify the bandwidth weighting option to boost the associated type of baseline
bandwidth, as follows:

default

This option uses the standard bandwidth configuration for your instance
type.

vpc-1

This option boosts your networking baseline bandwidth and reduces your EBS
baseline bandwidth.

ebs-1

This option boosts your EBS baseline bandwidth and reduces your networking
baseline bandwidth.

Type: String

Valid Values: `default | vpc-1 | ebs-1`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplatenetworkperformanceoptionsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplatenetworkperformanceoptionsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplatenetworkperformanceoptionsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateNetworkPerformanceOptions

LaunchTemplateOverrides
