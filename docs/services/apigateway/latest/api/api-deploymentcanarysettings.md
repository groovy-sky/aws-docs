# DeploymentCanarySettings

The input configuration for a canary deployment.

## Contents

**percentTraffic**

The percentage (0.0-100.0) of traffic routed to the canary deployment.

Type: Double

Required: No

**stageVariableOverrides**

A stage variable overrides used for the canary release deployment. They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.

Type: String to string map

Required: No

**useStageCache**

A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/deploymentcanarysettings.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/deploymentcanarysettings.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/deploymentcanarysettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deployment

DocumentationPart

All content copied from https://docs.aws.amazon.com/.
