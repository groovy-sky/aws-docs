# ElasticGpus

###### Note

Amazon Elastic Graphics reached end of life on January 8, 2024.

Describes an Elastic Graphics accelerator.

## Contents

**availabilityZone**

The Availability Zone in the which the Elastic Graphics accelerator resides.

Type: String

Required: No

**elasticGpuHealth**

The status of the Elastic Graphics accelerator.

Type: [ElasticGpuHealth](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ElasticGpuHealth.html) object

Required: No

**elasticGpuId**

The ID of the Elastic Graphics accelerator.

Type: String

Required: No

**elasticGpuState**

The state of the Elastic Graphics accelerator.

Type: String

Valid Values: `ATTACHED`

Required: No

**elasticGpuType**

The type of Elastic Graphics accelerator.

Type: String

Required: No

**instanceId**

The ID of the instance to which the Elastic Graphics accelerator is attached.

Type: String

Required: No

**TagSet.N**

The tags assigned to the Elastic Graphics accelerator.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ElasticGpus)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ElasticGpus)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ElasticGpus)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ElasticGpuHealth

ElasticGpuSpecification
