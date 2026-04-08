# LoggingConfig

Contains logging configuration information for an extension.

## Contents

**LogGroupName**

The Amazon CloudWatch Logs group to which CloudFormation sends error logging information when invoking the
extension's handlers.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: Yes

**LogRoleArn**

The Amazon Resource Name (ARN) of the role that CloudFormation should assume when sending log
entries to CloudWatch Logs.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `arn:.+:iam::[0-9]{12}:role/.+`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/loggingconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/loggingconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/loggingconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LiveResourceDrift

ManagedExecution
