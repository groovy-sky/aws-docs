# BaselinePerformanceFactorsRequest

The baseline performance to consider, using an instance family as a baseline reference.
The instance family establishes the lowest acceptable level of performance. Amazon EC2 uses this
baseline to guide instance type selection, but there is no guarantee that the selected
instance types will always exceed the baseline for every application.

Currently, this parameter only supports CPU performance as a baseline performance
factor. For example, specifying `c6i` would use the CPU performance of the
`c6i` family as the baseline reference.

## Contents

**Cpu**

The CPU performance to consider, using an instance family as the baseline reference.

Type: [CpuPerformanceFactorRequest](api-cpuperformancefactorrequest.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/baselineperformancefactorsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/baselineperformancefactorsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/baselineperformancefactorsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BaselinePerformanceFactors

BlobAttributeValue
