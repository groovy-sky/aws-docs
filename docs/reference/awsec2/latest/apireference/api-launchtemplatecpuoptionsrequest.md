# LaunchTemplateCpuOptionsRequest

The CPU options for the instance. Both the core count and threads per core must be
specified in the request.

## Contents

**AmdSevSnp**

Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported
with M6a, R6a, and C6a instance types only. For more information, see [AMD SEV-SNP for\
Amazon EC2 instances](../../../../services/ec2/latest/userguide/sev-snp.md).

Type: String

Valid Values: `enabled | disabled`

Required: No

**CoreCount**

The number of CPU cores for the instance.

Type: Integer

Required: No

**NestedVirtualization**

Indicates whether to enable the instance for nested virtualization.
Nested virtualization is supported only on 8th generation Intel-based instance types (c8i, m8i, r8i, and their flex variants).
When nested virtualization is enabled, Virtual Secure Mode (VSM) is automatically disabled for the instance.

Type: String

Valid Values: `enabled | disabled`

Required: No

**ThreadsPerCore**

The number of threads per CPU core. To disable multithreading for the instance,
specify a value of `1`. Otherwise, specify the default value of
`2`.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplatecpuoptionsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplatecpuoptionsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplatecpuoptionsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateCpuOptions

LaunchTemplateEbsBlockDevice
