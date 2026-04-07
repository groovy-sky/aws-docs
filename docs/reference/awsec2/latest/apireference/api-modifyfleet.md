# ModifyFleet

Modifies the specified EC2 Fleet.

You can only modify an EC2 Fleet request of type `maintain`.

While the EC2 Fleet is being modified, it is in the `modifying` state.

To scale up your EC2 Fleet, increase its target capacity. The EC2 Fleet launches the additional
Spot Instances according to the allocation strategy for the EC2 Fleet request. If the allocation
strategy is `lowest-price`, the EC2 Fleet launches instances using the Spot Instance
pool with the lowest price. If the allocation strategy is `diversified`, the
EC2 Fleet distributes the instances across the Spot Instance pools. If the allocation strategy
is `capacity-optimized`, EC2 Fleet launches instances from Spot Instance pools with optimal
capacity for the number of instances that are launching.

To scale down your EC2 Fleet, decrease its target capacity. First, the EC2 Fleet cancels any open
requests that exceed the new target capacity. You can request that the EC2 Fleet terminate Spot
Instances until the size of the fleet no longer exceeds the new target capacity. If the
allocation strategy is `lowest-price`, the EC2 Fleet terminates the instances with
the highest price per unit. If the allocation strategy is `capacity-optimized`,
the EC2 Fleet terminates the instances in the Spot Instance pools that have the least available
Spot Instance capacity. If the allocation strategy is `diversified`, the EC2 Fleet terminates
instances across the Spot Instance pools. Alternatively, you can request that the EC2 Fleet keep
the fleet at its current size, but not replace any Spot Instances that are interrupted or
that you terminate manually.

If you are finished with your EC2 Fleet for now, but will use it again later, you can set the
target capacity to 0.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Context**

Reserved.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ExcessCapacityTerminationPolicy**

Indicates whether running instances should be terminated if the total target capacity of
the EC2 Fleet is decreased below the current size of the EC2 Fleet.

Supported only for fleets of type `maintain`.

Type: String

Valid Values: `no-termination | termination`

Required: No

**FleetId**

The ID of the EC2 Fleet.

Type: String

Required: Yes

**LaunchTemplateConfig.N**

The launch template and overrides.

Type: Array of [FleetLaunchTemplateConfigRequest](api-fleetlaunchtemplateconfigrequest.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**TargetCapacitySpecification**

The size of the EC2 Fleet.

Type: [TargetCapacitySpecificationRequest](api-targetcapacityspecificationrequest.md) object

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

If the request succeeds, the response returns `true`. If the request fails,
no response is returned, and instead an error message is returned.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyFleet)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyFleet)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyfleet.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyfleet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyfleet.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyfleet.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyfleet.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyfleet.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyFleet)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyfleet.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyEbsDefaultKmsKeyId

ModifyFpgaImageAttribute
