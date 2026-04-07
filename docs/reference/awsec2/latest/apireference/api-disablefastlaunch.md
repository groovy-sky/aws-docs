# DisableFastLaunch

Discontinue Windows fast launch for a Windows AMI, and clean up existing pre-provisioned
snapshots. After you disable Windows fast launch, the AMI uses the standard launch process for
each new instance. Amazon EC2 must remove all pre-provisioned snapshots before you can enable
Windows fast launch again.

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

**Force**

Forces the image settings to turn off Windows fast launch for your Windows AMI. This
parameter overrides any errors that are encountered while cleaning up resources in your
account.

Type: Boolean

Required: No

**ImageId**

Specify the ID of the image for which to disable Windows fast launch.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**imageId**

The ID of the image for which Windows fast launch was disabled.

Type: String

**launchTemplate**

The launch template that was used to launch Windows instances from pre-provisioned
snapshots.

Type: [FastLaunchLaunchTemplateSpecificationResponse](api-fastlaunchlaunchtemplatespecificationresponse.md) object

**maxParallelLaunches**

The maximum number of instances that Amazon EC2 can launch at the same time to create
pre-provisioned snapshots for Windows fast launch.

Type: Integer

**ownerId**

The owner of the Windows AMI for which Windows fast launch was disabled.

Type: String

**requestId**

The ID of the request.

Type: String

**resourceType**

The pre-provisioning resource type that must be cleaned after turning off Windows fast
launch for the Windows AMI. Supported values include: `snapshot`.

Type: String

Valid Values: `snapshot`

**snapshotConfiguration**

Parameters that were used for Windows fast launch for the Windows AMI before Windows fast
launch was disabled. This informs the clean-up process.

Type: [FastLaunchSnapshotConfigurationResponse](api-fastlaunchsnapshotconfigurationresponse.md) object

**state**

The current state of Windows fast launch for the specified Windows AMI.

Type: String

Valid Values: `enabling | enabling-failed | enabled | enabled-failed | disabling | disabling-failed`

**stateTransitionReason**

The reason that the state changed for Windows fast launch for the Windows AMI.

Type: String

**stateTransitionTime**

The time that the state changed for Windows fast launch for the Windows AMI.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisableFastLaunch)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisableFastLaunch)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disablefastlaunch.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disablefastlaunch.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disablefastlaunch.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disablefastlaunch.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disablefastlaunch.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disablefastlaunch.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisableFastLaunch)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disablefastlaunch.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisableEbsEncryptionByDefault

DisableFastSnapshotRestores
