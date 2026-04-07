# InstanceCreditSpecificationRequest

Describes the credit option for CPU usage of a burstable performance instance.

## Contents

**InstanceId**

The ID of the instance.

Type: String

Required: Yes

**CpuCredits**

The credit option for CPU usage of the instance.

Valid values: `standard` \| `unlimited`

T3 instances with `host` tenancy do not support the `unlimited`
CPU credit option.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/InstanceCreditSpecificationRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/InstanceCreditSpecificationRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/InstanceCreditSpecificationRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InstanceCreditSpecification

InstanceEventWindow
