# DeploymentStrategy

## Contents

**DeploymentDurationInMinutes**

Total amount of time the deployment lasted.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1440.

Required: No

**Description**

The description of the deployment strategy.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**FinalBakeTimeInMinutes**

The amount of time that AWS AppConfig monitored for alarms before considering the
deployment to be complete and no longer eligible for automatic rollback.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1440.

Required: No

**GrowthFactor**

The percentage of targets that received a deployed configuration during each
interval.

Type: Float

Valid Range: Minimum value of 1.0. Maximum value of 100.0.

Required: No

**GrowthType**

The algorithm used to define how percentage grew over time.

Type: String

Valid Values: `LINEAR | EXPONENTIAL`

Required: No

**Id**

The deployment strategy ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**Name**

The name of the deployment strategy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ReplicateTo**

Save the deployment strategy to a Systems Manager (SSM) document.

Type: String

Valid Values: `NONE | SSM_DOCUMENT`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/deploymentstrategy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/deploymentstrategy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/deploymentstrategy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentEvent

DeploymentSummary

All content copied from https://docs.aws.amazon.com/.
