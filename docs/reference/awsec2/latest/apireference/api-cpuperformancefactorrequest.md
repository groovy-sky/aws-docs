# CpuPerformanceFactorRequest

The CPU performance to consider, using an instance family as the baseline reference.

## Contents

**Reference.N**

Specify an instance family to use as the baseline reference for CPU performance. All
instance types that match your specified attributes will be compared against the CPU
performance of the referenced instance family, regardless of CPU manufacturer or
architecture differences.

###### Note

Currently, only one instance family can be specified in the list.

Type: Array of [PerformanceFactorReferenceRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PerformanceFactorReferenceRequest.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CpuPerformanceFactorRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CpuPerformanceFactorRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CpuPerformanceFactorRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CpuPerformanceFactor

CreateFleetError
