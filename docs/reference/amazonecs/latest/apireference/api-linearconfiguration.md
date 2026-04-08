# LinearConfiguration

Configuration for linear deployment strategy that shifts production traffic in equal
percentage increments with configurable wait times between each step until 100% of
traffic is shifted to the new service revision. This is only valid when you run
`CreateService` or `UpdateService` with
`deploymentController` set to `ECS` and a
`deploymentConfiguration` with a strategy set to `LINEAR`.

## Contents

**stepBakeTimeInMinutes**

The amount of time in minutes to wait between each traffic shifting step during a linear deployment. Valid values are 0 to 1440 minutes (24 hours). The default value is 6. This bake time is not applied after reaching 100 percent traffic.

Type: Integer

Required: No

**stepPercent**

The percentage of production traffic to shift in each step during a linear deployment. Valid
values are multiples of 0.1 from 3.0 to 100.0. The default value is 10.0.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/linearconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/linearconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/linearconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

KeyValuePair

LinuxParameters
