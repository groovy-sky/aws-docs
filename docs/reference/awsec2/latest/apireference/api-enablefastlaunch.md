# EnableFastLaunch

When you enable Windows fast launch for a Windows AMI, images are pre-provisioned, using
snapshots to launch instances up to 65% faster. To create the optimized Windows image, Amazon EC2
launches an instance and runs through Sysprep steps, rebooting as required. Then it creates a
set of reserved snapshots that are used for subsequent launches. The reserved snapshots are
automatically replenished as they are used, depending on your settings for launch
frequency.

###### Note

You can only change these settings for Windows AMIs that you own or that have been
shared with you.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ImageId**

Specify the ID of the image for which to enable Windows fast launch.

Type: String

Required: Yes

**LaunchTemplate**

The launch template to use when launching Windows instances from pre-provisioned
snapshots. Launch template parameters can include either the name or ID of the launch
template, but not both.

Type: [FastLaunchLaunchTemplateSpecificationRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FastLaunchLaunchTemplateSpecificationRequest.html) object

Required: No

**MaxParallelLaunches**

The maximum number of instances that Amazon EC2 can launch at the same time to create
pre-provisioned snapshots for Windows fast launch. Value must be `6` or
greater.

Type: Integer

Required: No

**ResourceType**

The type of resource to use for pre-provisioning the AMI for Windows fast launch.
Supported values include: `snapshot`, which is the default value.

Type: String

Required: No

**SnapshotConfiguration**

Configuration settings for creating and managing the snapshots that are used for
pre-provisioning the AMI for Windows fast launch. The associated `ResourceType`
must be `snapshot`.

Type: [FastLaunchSnapshotConfigurationRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FastLaunchSnapshotConfigurationRequest.html) object

Required: No

## Response Elements

The following elements are returned by the service.

**imageId**

The image ID that identifies the AMI for which Windows fast launch was enabled.

Type: String

**launchTemplate**

The launch template that is used when launching Windows instances from pre-provisioned
snapshots.

Type: [FastLaunchLaunchTemplateSpecificationResponse](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FastLaunchLaunchTemplateSpecificationResponse.html) object

**maxParallelLaunches**

The maximum number of instances that Amazon EC2 can launch at the same time to create
pre-provisioned snapshots for Windows fast launch.

Type: Integer

**ownerId**

The owner ID for the AMI for which Windows fast launch was enabled.

Type: String

**requestId**

The ID of the request.

Type: String

**resourceType**

The type of resource that was defined for pre-provisioning the AMI for Windows fast
launch.

Type: String

Valid Values: `snapshot`

**snapshotConfiguration**

Settings to create and manage the pre-provisioned snapshots that Amazon EC2 uses for faster
launches from the Windows AMI. This property is returned when the associated
`resourceType` is `snapshot`.

Type: [FastLaunchSnapshotConfigurationResponse](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FastLaunchSnapshotConfigurationResponse.html) object

**state**

The current state of Windows fast launch for the specified AMI.

Type: String

Valid Values: `enabling | enabling-failed | enabled | enabled-failed | disabling | disabling-failed`

**stateTransitionReason**

The reason that the state changed for Windows fast launch for the AMI.

Type: String

**stateTransitionTime**

The time that the state changed for Windows fast launch for the AMI.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/EnableFastLaunch)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/EnableFastLaunch)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EnableEbsEncryptionByDefault

EnableFastSnapshotRestores
