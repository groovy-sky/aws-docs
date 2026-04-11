# ProcessorInfo

Describes the processor used by the instance type.

## Contents

**manufacturer**

The manufacturer of the processor.

Type: String

Required: No

**SupportedArchitectures.N**

The architectures supported by the instance type.

Type: Array of strings

Valid Values: `i386 | x86_64 | arm64 | x86_64_mac | arm64_mac`

Required: No

**SupportedFeatures.N**

Indicates whether the instance type supports AMD SEV-SNP. If the request returns
`amd-sev-snp`, AMD SEV-SNP is supported. Otherwise, it is not supported. For more
information, see [AMD\
SEV-SNP](../../../../services/ec2/latest/userguide/sev-snp.md).

Type: Array of strings

Valid Values: `amd-sev-snp | nested-virtualization`

Required: No

**sustainedClockSpeedInGhz**

The speed of the processor, in GHz.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/processorinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/processorinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/processorinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PrivateIpAddressSpecification

ProductCode
