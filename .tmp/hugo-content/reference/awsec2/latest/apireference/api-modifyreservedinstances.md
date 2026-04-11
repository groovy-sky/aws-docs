# ModifyReservedInstances

Modifies the configuration of your Reserved Instances, such as the Availability Zone,
instance count, or instance type. The Reserved Instances to be modified must be identical,
except for Availability Zone, network platform, and instance type.

For more information, see [Modify Reserved Instances](../../../../services/ec2/latest/userguide/ri-modifying.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

A unique, case-sensitive token you provide to ensure idempotency of your modification
request. For more information, see [Ensuring\
Idempotency](run-instance-idempotency.md).

Type: String

Required: No

**ReservedInstancesConfigurationSetItemType.N**

The configuration settings for the Reserved Instances to modify.

Type: Array of [ReservedInstancesConfiguration](api-reservedinstancesconfiguration.md) objects

Required: Yes

**ReservedInstancesId.N**

The IDs of the Reserved Instances to modify.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**reservedInstancesModificationId**

The ID for the modification.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example illustrates one usage of ModifyReservedInstances.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyReservedInstances
&ClientToken=myClientToken
&ReservedInstancesConfigurationSetItemType.1.AvailabilityZone=us-east-1a
&ReservedInstancesConfigurationSetItemType.1.InstanceCount=1
&ReservedInstancesConfigurationSetItemType.1.Platform=EC2-VPC
&ReservedInstancesConfigurationSetItemType.1.InstanceType=m1.small
&ReservedInstancesId.1=d16f7a91-4d0f-4f19-9d7f-a74d26b1ccfa
&AUTHPARAMS
```

#### Sample Response

```

<ModifyReservedInstancesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
<requestId>bef729b6-0731-4489-8881-2258746ae163</requestId>
<reservedInstancesModificationId>rimod-3aae219d-3d63-47a9-a7e9-e764example</reservedInstancesModificationId>
</ModifyReservedInstancesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyreservedinstances.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyreservedinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyPublicIpDnsNameOptions

ModifyRouteServer
