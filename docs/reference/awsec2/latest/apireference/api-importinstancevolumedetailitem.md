# ImportInstanceVolumeDetailItem

Describes an import volume task.

## Contents

**availabilityZone**

The Availability Zone where the resulting instance will reside.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone where the resulting instance will reside.

Type: String

Required: No

**bytesConverted**

The number of bytes converted so far.

Type: Long

Required: No

**description**

A description of the task.

Type: String

Required: No

**image**

The image.

Type: [DiskImageDescription](api-diskimagedescription.md) object

Required: No

**status**

The status of the import of this particular disk image.

Type: String

Required: No

**statusMessage**

The status information or errors related to the disk image.

Type: String

Required: No

**volume**

The volume.

Type: [DiskImageVolumeDescription](api-diskimagevolumedescription.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/importinstancevolumedetailitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/importinstancevolumedetailitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/importinstancevolumedetailitem.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImportInstanceTaskDetails

ImportSnapshotTask
