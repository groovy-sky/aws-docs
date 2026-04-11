# LaunchTemplateCpuOptions

The CPU options for the instance.

## Contents

**amdSevSnp**

Indicates whether the instance is enabled for AMD SEV-SNP. For more information, see
[AMD SEV-SNP\
for Amazon EC2 instances](../../../../services/ec2/latest/userguide/sev-snp.md).

Type: String

Valid Values: `enabled | disabled`

Required: No

**coreCount**

The number of CPU cores for the instance.

Type: Integer

Required: No

**nestedVirtualization**

Indicates whether the instance is enabled for nested virtualization.

Type: String

Valid Values: `enabled | disabled`

Required: No

**threadsPerCore**

The number of threads per CPU core.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplatecpuoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplatecpuoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplatecpuoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateConfig

LaunchTemplateCpuOptionsRequest
