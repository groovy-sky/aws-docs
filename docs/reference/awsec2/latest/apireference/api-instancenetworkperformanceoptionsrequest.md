# InstanceNetworkPerformanceOptionsRequest

Configure network performance options for your instance that are geared towards performance
improvements based on the workload that it runs.

## Contents

**BandwidthWeighting**

Specify the bandwidth weighting option to boost the associated type of baseline bandwidth,
as follows:

default

This option uses the standard bandwidth configuration for your instance type.

vpc-1

This option boosts your networking baseline bandwidth and reduces your EBS baseline
bandwidth.

ebs-1

This option boosts your EBS baseline bandwidth and reduces your networking baseline
bandwidth.

Type: String

Valid Values: `default | vpc-1 | ebs-1`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/InstanceNetworkPerformanceOptionsRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/InstanceNetworkPerformanceOptionsRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/InstanceNetworkPerformanceOptionsRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InstanceNetworkPerformanceOptions

InstancePrivateIpAddress
