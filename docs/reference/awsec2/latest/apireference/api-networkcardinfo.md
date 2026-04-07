# NetworkCardInfo

Describes the network card support of the instance type.

## Contents

**additionalFlexibleNetworkInterfaces**

The number of additional network interfaces that can be attached to an instance when using
flexible Elastic Network Adapter (ENA) queues. This number is in addition to the base number
specified by `maximumNetworkInterfaces`.

Type: Integer

Required: No

**baselineBandwidthInGbps**

The baseline network performance of the network card, in Gbps.

Type: Double

Required: No

**defaultEnaQueueCountPerInterface**

The default number of the ENA queues for each interface.

Type: Integer

Required: No

**maximumEnaQueueCount**

The maximum number of the ENA queues.

Type: Integer

Required: No

**maximumEnaQueueCountPerInterface**

The maximum number of the ENA queues for each interface.

Type: Integer

Required: No

**maximumNetworkInterfaces**

The maximum number of network interfaces for the network card.

Type: Integer

Required: No

**networkCardIndex**

The index of the network card.

Type: Integer

Required: No

**networkPerformance**

The network performance of the network card.

Type: String

Required: No

**peakBandwidthInGbps**

The peak (burst) network performance of the network card, in Gbps.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkcardinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkcardinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkcardinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkBandwidthGbpsRequest

NetworkInfo
