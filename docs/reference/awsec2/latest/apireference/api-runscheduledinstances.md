# RunScheduledInstances

Launches the specified Scheduled Instances.

Before you can launch a Scheduled Instance, you must purchase it and obtain an identifier using [PurchaseScheduledInstances](api-purchasescheduledinstances.md).

You must launch a Scheduled Instance during its scheduled time period. You can't stop or
reboot a Scheduled Instance, but you can terminate it as needed. If you terminate a
Scheduled Instance before the current scheduled time period ends, you can launch it again
after a few minutes.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that ensures the idempotency of the request.
For more information, see [Ensuring Idempotency](run-instance-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceCount**

The number of instances.

Default: 1

Type: Integer

Required: No

**LaunchSpecification**

The launch specification. You must match the instance type, Availability Zone,
network, and platform of the schedule that you purchased.

Type: [ScheduledInstancesLaunchSpecification](api-scheduledinstanceslaunchspecification.md) object

Required: Yes

**ScheduledInstanceId**

The Scheduled Instance ID.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**instanceIdSet**

The IDs of the newly launched instances.

Type: Array of strings

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/RunScheduledInstances)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/RunScheduledInstances)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/runscheduledinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/runscheduledinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/runscheduledinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/runscheduledinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/runscheduledinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/runscheduledinstances.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/RunScheduledInstances)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/runscheduledinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RunInstances

SearchLocalGatewayRoutes
