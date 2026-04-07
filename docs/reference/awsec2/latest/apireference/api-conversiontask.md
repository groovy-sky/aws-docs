# ConversionTask

Describes a conversion task.

## Contents

**conversionTaskId**

The ID of the conversion task.

Type: String

Required: No

**expirationTime**

The time when the task expires. If the upload isn't complete before the expiration time, we automatically cancel
the task.

Type: String

Required: No

**importInstance**

If the task is for importing an instance, this contains information about the import instance task.

Type: [ImportInstanceTaskDetails](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportInstanceTaskDetails.html) object

Required: No

**importVolume**

If the task is for importing a volume, this contains information about the import volume task.

Type: [ImportVolumeTaskDetails](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportVolumeTaskDetails.html) object

Required: No

**state**

The state of the conversion task.

Type: String

Valid Values: `active | cancelling | cancelled | completed`

Required: No

**statusMessage**

The status message related to the conversion task.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the task.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ConversionTask)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ConversionTask)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ConversionTask)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectionTrackingSpecificationResponse

CpuOptions
