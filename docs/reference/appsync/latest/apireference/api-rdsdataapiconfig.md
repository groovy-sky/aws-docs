# RdsDataApiConfig

Contains the metadata required to introspect the RDS cluster.

## Contents

**databaseName**

The name of the database in the cluster.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**resourceArn**

The resource ARN of the RDS cluster.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:[a-z-]*:rds:[a-z0-9-]*:\d{12}:cluster:[0-9A-Za-z_/-]*$`

Required: Yes

**secretArn**

The secret's ARN that was obtained from Secrets Manager. A secret consists of secret
information, the secret value, plus metadata about the secret. A secret value can be a
string or binary. It typically includes the ARN, secret name and description, policies,
tags, encryption key from the Key Management Service, and key rotation data.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:[a-z-]*:secretsmanager:[a-z0-9-]*:\d{12}:secret:[0-9A-Za-z_/+=.@!-]*$`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/rdsdataapiconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/rdsdataapiconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/rdsdataapiconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipelineConfig

RdsHttpEndpointConfig

All content copied from https://docs.aws.amazon.com/.
