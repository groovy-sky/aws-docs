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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancecreditspecificationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancecreditspecificationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancecreditspecificationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceCreditSpecification

InstanceEventWindow
