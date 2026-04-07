# NetworkBandwidthGbpsRequest

The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).

###### Note

Setting the minimum bandwidth does not guarantee that your instance will achieve the
minimum bandwidth. Amazon EC2 will identify instance types that support the specified minimum
bandwidth, but the actual bandwidth of your instance might go below the specified minimum
at times. For more information, see [Available instance bandwidth](../../../../services/ec2/latest/userguide/ec2-instance-network-bandwidth.md#available-instance-bandwidth) in the
_Amazon EC2 User Guide_.

## Contents

**Max**

The maximum amount of network bandwidth, in Gbps. To specify no maximum limit, omit this
parameter.

Type: Double

Required: No

**Min**

The minimum amount of network bandwidth, in Gbps. To specify no minimum limit, omit this
parameter.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/NetworkBandwidthGbpsRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/NetworkBandwidthGbpsRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/NetworkBandwidthGbpsRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkBandwidthGbps

NetworkCardInfo
