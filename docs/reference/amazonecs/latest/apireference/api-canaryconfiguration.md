# CanaryConfiguration

Configuration for a canary deployment strategy that shifts a fixed percentage of traffic to
the new service revision, waits for a specified bake time, then shifts the remaining
traffic.

This is only valid when you run `CreateService` or
`UpdateService` with `deploymentController` set to
`ECS` and a `deploymentConfiguration` with a strategy set to
`CANARY`.

## Contents

**canaryBakeTimeInMinutes**

The amount of time in minutes to wait during the canary phase before shifting the
remaining production traffic to the new service revision. Valid values are 0 to 1440
minutes (24 hours). The default value is 10.

Type: Integer

Required: No

**canaryPercent**

The percentage of production traffic to shift to the new service revision during the canary phase. Valid values are multiples of 0.1 from 0.1 to 100.0. The default value is 5.0.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/canaryconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/canaryconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/canaryconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BaselineEbsBandwidthMbpsRequest

CapacityProvider
