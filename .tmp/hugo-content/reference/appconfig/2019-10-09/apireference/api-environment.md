# Environment

## Contents

**ApplicationId**

The application ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**Description**

The description of the environment.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**Id**

The environment ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**Monitors**

Amazon CloudWatch alarms monitored during the deployment.

Type: Array of [Monitor](api-monitor.md) objects

Array Members: Minimum number of 0 items. Maximum number of 5 items.

Required: No

**Name**

The name of the environment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**State**

The state of the environment. An environment can be in one of the following states:
`READY_FOR_DEPLOYMENT`, `DEPLOYING`, `ROLLING_BACK`, or
`ROLLED_BACK`

Type: String

Valid Values: `READY_FOR_DEPLOYMENT | DEPLOYING | ROLLING_BACK | ROLLED_BACK | REVERTED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/environment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/environment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/environment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentSummary

ExtensionAssociationSummary

All content copied from https://docs.aws.amazon.com/.
