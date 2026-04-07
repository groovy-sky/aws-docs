# GpuDeviceInfo

Describes the GPU accelerators for the instance type.

## Contents

**count**

The number of GPUs for the instance type.

Type: Integer

Required: No

**gpuPartitionSize**

The size of each GPU as a fraction of a full GPU, between 0 (excluded) and 1 (included).

Type: Double

Required: No

**logicalGpuCount**

Total number of GPU devices of this type.

Type: Integer

Required: No

**manufacturer**

The manufacturer of the GPU accelerator.

Type: String

Required: No

**memoryInfo**

Describes the memory available to the GPU accelerator.

Type: [GpuDeviceMemoryInfo](api-gpudevicememoryinfo.md) object

Required: No

**name**

The name of the GPU accelerator.

Type: String

Required: No

**WorkloadSet.N**

A list of workload types this GPU supports.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/gpudeviceinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/gpudeviceinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/gpudeviceinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FpgaInfo

GpuDeviceMemoryInfo
