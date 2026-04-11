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

Type: [ImportInstanceTaskDetails](api-importinstancetaskdetails.md) object

Required: No

**importVolume**

If the task is for importing a volume, this contains information about the import volume task.

Type: [ImportVolumeTaskDetails](api-importvolumetaskdetails.md) object

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/conversiontask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/conversiontask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/conversiontask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ConnectionTrackingSpecificationResponse

CpuOptions
