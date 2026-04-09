# DeploymentSummary

Information about the deployment.

## Contents

**CompletedAt**

Time the deployment completed.

Type: Timestamp

Required: No

**ConfigurationName**

The name of the configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ConfigurationVersion**

The version of the configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**DeploymentDurationInMinutes**

Total amount of time the deployment lasted.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1440.

Required: No

**DeploymentNumber**

The sequence number of the deployment.

Type: Integer

Required: No

**FinalBakeTimeInMinutes**

The amount of time that AWS AppConfig monitors for alarms before considering the
deployment to be complete and no longer eligible for automatic rollback.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1440.

Required: No

**GrowthFactor**

The percentage of targets to receive a deployed configuration during each
interval.

Type: Float

Valid Range: Minimum value of 1.0. Maximum value of 100.0.

Required: No

**GrowthType**

The algorithm used to define how percentage grows over time.

Type: String

Valid Values: `LINEAR | EXPONENTIAL`

Required: No

**PercentageComplete**

The percentage of targets for which the deployment is available.

Type: Float

Valid Range: Minimum value of 1.0. Maximum value of 100.0.

Required: No

**StartedAt**

Time the deployment started.

Type: Timestamp

Required: No

**State**

The state of the deployment.

Type: String

Valid Values: `BAKING | VALIDATING | DEPLOYING | COMPLETE | ROLLING_BACK | ROLLED_BACK | REVERTED`

Required: No

**VersionLabel**

A user-defined label for an AWS AppConfig hosted configuration version.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `.*[^0-9].*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/deploymentsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/deploymentsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/deploymentsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentStrategy

Environment

All content copied from https://docs.aws.amazon.com/.
