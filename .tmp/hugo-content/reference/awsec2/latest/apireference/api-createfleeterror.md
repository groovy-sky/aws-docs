# CreateFleetError

Describes the instances that could not be launched by the fleet.

## Contents

**errorCode**

The error code that indicates why the instance could not be launched. For more
information about error codes, see [Error codes](errors-overview.md).

Type: String

Required: No

**errorMessage**

The error message that describes why the instance could not be launched. For more
information about error messages, see [Error codes](errors-overview.md).

Type: String

Required: No

**launchTemplateAndOverrides**

The launch templates and overrides that were used for launching the instances. The
values that you specify in the Overrides replace the values in the launch template.

Type: [LaunchTemplateAndOverridesResponse](api-launchtemplateandoverridesresponse.md) object

Required: No

**lifecycle**

Indicates if the instance that could not be launched was a Spot, On-Demand, Capacity Block,
or Interruptible Capacity Reservation instance.

Type: String

Valid Values: `spot | on-demand | interruptible-capacity-reservation`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createfleeterror.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createfleeterror.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createfleeterror.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CpuPerformanceFactorRequest

CreateFleetInstance
