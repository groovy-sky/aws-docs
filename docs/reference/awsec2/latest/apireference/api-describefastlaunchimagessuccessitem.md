# DescribeFastLaunchImagesSuccessItem

Describe details about a Windows image with Windows fast launch enabled that meets the
requested criteria. Criteria are defined by the `DescribeFastLaunchImages` action
filters.

## Contents

**imageId**

The image ID that identifies the Windows fast launch enabled image.

Type: String

Required: No

**launchTemplate**

The launch template that the Windows fast launch enabled AMI uses when it launches Windows
instances from pre-provisioned snapshots.

Type: [FastLaunchLaunchTemplateSpecificationResponse](api-fastlaunchlaunchtemplatespecificationresponse.md) object

Required: No

**maxParallelLaunches**

The maximum number of instances that Amazon EC2 can launch at the same time to create
pre-provisioned snapshots for Windows fast launch.

Type: Integer

Required: No

**ownerId**

The owner ID for the Windows fast launch enabled AMI.

Type: String

Required: No

**resourceType**

The resource type that Amazon EC2 uses for pre-provisioning the Windows AMI. Supported values
include: `snapshot`.

Type: String

Valid Values: `snapshot`

Required: No

**snapshotConfiguration**

A group of parameters that are used for pre-provisioning the associated Windows AMI using
snapshots.

Type: [FastLaunchSnapshotConfigurationResponse](api-fastlaunchsnapshotconfigurationresponse.md) object

Required: No

**state**

The current state of Windows fast launch for the specified Windows AMI.

Type: String

Valid Values: `enabling | enabling-failed | enabled | enabled-failed | disabling | disabling-failed`

Required: No

**stateTransitionReason**

The reason that Windows fast launch for the AMI changed to the current state.

Type: String

Required: No

**stateTransitionTime**

The time that Windows fast launch for the AMI changed to the current state.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describefastlaunchimagessuccessitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describefastlaunchimagessuccessitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describefastlaunchimagessuccessitem.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeregisterInstanceTagAttributeRequest

DescribeFastSnapshotRestoreSuccessItem
