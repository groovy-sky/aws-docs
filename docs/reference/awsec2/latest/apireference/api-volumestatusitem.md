# VolumeStatusItem

Describes the volume status.

## Contents

**ActionsSet.N**

The details of the operation.

Type: Array of [VolumeStatusAction](api-volumestatusaction.md) objects

Required: No

**AttachmentStatuses.N**

Information about the instances to which the volume is attached.

Type: Array of [VolumeStatusAttachmentStatus](api-volumestatusattachmentstatus.md) objects

Required: No

**availabilityZone**

The Availability Zone of the volume.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone.

Type: String

Required: No

**EventsSet.N**

A list of events associated with the volume.

Type: Array of [VolumeStatusEvent](api-volumestatusevent.md) objects

Required: No

**initializationStatusDetails**

Information about the volume initialization. It can take up to 5 minutes
for the volume initialization information to be updated.

Only available for volumes created from snapshots. Not available for empty
volumes created without a snapshot.

For more information, see
[Initialize Amazon EBS volumes](../../../../services/ebs/latest/userguide/initalize-volume.md).

Type: [InitializationStatusDetails](api-initializationstatusdetails.md) object

Required: No

**outpostArn**

The Amazon Resource Name (ARN) of the Outpost.

Type: String

Required: No

**volumeId**

The volume ID.

Type: String

Required: No

**volumeStatus**

The volume status.

Type: [VolumeStatusInfo](api-volumestatusinfo.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/volumestatusitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/volumestatusitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/volumestatusitem.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VolumeStatusInfo

Vpc
