# TargetCapacitySpecificationRequest

The number of units to request. You can choose to set the target capacity as the number of
instances. Or you can set the target capacity to a performance characteristic that is important to your application workload,
such as vCPUs, memory, or I/O. If the request type is `maintain`, you can
specify a target capacity of 0 and add capacity later.

You can use the On-Demand Instance `MaxTotalPrice` parameter, the Spot Instance
`MaxTotalPrice` parameter, or both parameters to ensure that your fleet cost
does not exceed your budget. If you set a maximum price per hour for the On-Demand Instances and Spot Instances
in your request, EC2 Fleet will launch instances until it reaches the maximum amount that you're
willing to pay. When the maximum amount you're willing to pay is reached, the fleet stops
launching instances even if it hasn't met the target capacity. The
`MaxTotalPrice` parameters are located in [OnDemandOptionsRequest](api-ondemandoptionsrequest.md)
and [SpotOptionsRequest](api-spotoptionsrequest.md).

## Contents

**TotalTargetCapacity**

The number of units to request, filled using the default target capacity type.

Type: Integer

Required: Yes

**DefaultTargetCapacityType**

The default target capacity type.

Type: String

Valid Values: `spot | on-demand | capacity-block | reserved-capacity`

Required: No

**OnDemandTargetCapacity**

The number of On-Demand units to request.

Type: Integer

Required: No

**SpotTargetCapacity**

The number of Spot units to request.

Type: Integer

Required: No

**TargetCapacityUnitType**

The unit for the target capacity. You can specify this parameter only when using
attributed-based instance type selection.

Default: `units` (the number of instances)

Type: String

Valid Values: `vcpu | memory-mib | units`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/targetcapacityspecificationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/targetcapacityspecificationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/targetcapacityspecificationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TargetCapacitySpecification

TargetConfiguration
