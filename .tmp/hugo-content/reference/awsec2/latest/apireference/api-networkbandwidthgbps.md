# NetworkBandwidthGbps

The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).

###### Note

Setting the minimum bandwidth does not guarantee that your instance will achieve the
minimum bandwidth. Amazon EC2 will identify instance types that support the specified minimum
bandwidth, but the actual bandwidth of your instance might go below the specified minimum
at times. For more information, see [Available instance bandwidth](../../../../services/ec2/latest/userguide/ec2-instance-network-bandwidth.md#available-instance-bandwidth) in the
_Amazon EC2 User Guide_.

## Contents

**Max** (request), **max** (response)

The maximum amount of network bandwidth, in Gbps. If this parameter is not specified, there is no
maximum limit.

Type: Double

Required: No

**Min** (request), **min** (response)

The minimum amount of network bandwidth, in Gbps. If this parameter is not specified, there is no minimum
limit.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkbandwidthgbps.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkbandwidthgbps.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkbandwidthgbps.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkAclEntry

NetworkBandwidthGbpsRequest
