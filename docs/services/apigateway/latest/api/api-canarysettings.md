# CanarySettings

Configuration settings of a canary deployment.

## Contents

**deploymentId**

The ID of the canary deployment.

Type: String

Required: No

**percentTraffic**

The percent (0-100) of traffic diverted to a canary deployment.

Type: Double

Required: No

**stageVariableOverrides**

Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.

Type: String to string map

Required: No

**useStageCache**

A Boolean flag to indicate whether the canary deployment uses the stage cache or not.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/canarysettings.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/canarysettings.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/canarysettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasePathMapping

ClientCertificate

All content copied from https://docs.aws.amazon.com/.
