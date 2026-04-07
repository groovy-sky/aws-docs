# VolumeStatusItem

Describes the volume status.

## Contents

**ActionsSet.N**

The details of the operation.

Type: Array of [VolumeStatusAction](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VolumeStatusAction.html) objects

Required: No

**AttachmentStatuses.N**

Information about the instances to which the volume is attached.

Type: Array of [VolumeStatusAttachmentStatus](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VolumeStatusAttachmentStatus.html) objects

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

Type: Array of [VolumeStatusEvent](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VolumeStatusEvent.html) objects

Required: No

**initializationStatusDetails**

Information about the volume initialization. It can take up to 5 minutes
for the volume initialization information to be updated.

Only available for volumes created from snapshots. Not available for empty
volumes created without a snapshot.

For more information, see
[Initialize Amazon EBS volumes](https://docs.aws.amazon.com/ebs/latest/userguide/initalize-volume.html).

Type: [InitializationStatusDetails](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InitializationStatusDetails.html) object

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

Type: [VolumeStatusInfo](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VolumeStatusInfo.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/VolumeStatusItem)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/VolumeStatusItem)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/VolumeStatusItem)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VolumeStatusInfo

Vpc
