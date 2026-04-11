# AvailableCapacity

The capacity information for instances that can be launched onto the Dedicated Host.

## Contents

**AvailableInstanceCapacity.N**

The number of instances that can be launched onto the Dedicated Host depending on the
host's available capacity. For Dedicated Hosts that support multiple instance types,
this parameter represents the number of instances for each instance size that is
supported on the host.

Type: Array of [InstanceCapacity](api-instancecapacity.md) objects

Required: No

**availableVCpus**

The number of vCPUs available for launching instances onto the Dedicated Host.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/availablecapacity.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/availablecapacity.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/availablecapacity.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AvailabilityZoneSubGeography

BaselineEbsBandwidthMbps
