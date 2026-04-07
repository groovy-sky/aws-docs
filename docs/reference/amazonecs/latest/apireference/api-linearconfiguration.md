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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/LinearConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/LinearConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/LinearConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KeyValuePair

LinuxParameters
