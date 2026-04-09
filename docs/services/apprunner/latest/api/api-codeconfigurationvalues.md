# CodeConfigurationValues

Describes the basic configuration needed for building and running an AWS App Runner service. This type doesn't support the full set of possible
configuration options. Fur full configuration capabilities, use a `apprunner.yaml` file in the source code repository.

## Contents

**Runtime**

A runtime environment type for building and running an App Runner service.
It represents a
programming language runtime.

Type: String

Valid Values: `PYTHON_3 | NODEJS_12 | NODEJS_14 | CORRETTO_8 | CORRETTO_11 | NODEJS_16 | GO_1 | DOTNET_6 | PHP_81 | RUBY_31 | PYTHON_311 | NODEJS_18 | NODEJS_22`

Required: Yes

**BuildCommand**

The command App Runner runs to build your
application.

Type: String

Pattern: `[^\x0a\x0d]+`

Required: No

**Port**

The port that your application listens to in the container.

Default: `8080`

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**RuntimeEnvironmentSecrets**

An array of key-value pairs representing the secrets and parameters that get referenced to your service as an environment variable. The supported
values are either the full Amazon Resource Name (ARN) of the AWS Secrets Manager secret or the full ARN of the parameter in the AWS Systems Manager
Parameter Store.

###### Note

- If the AWS Systems Manager Parameter Store parameter exists in the same AWS Region as the service that you're launching, you can use
either the full ARN or name of the secret. If the parameter exists in a different Region, then the full ARN must be specified.

- Currently, cross account referencing of AWS Systems Manager Parameter Store parameter is not supported.

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 2048.

Value Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**RuntimeEnvironmentVariables**

The environment variables that are available to your running AWS App Runner service. An array of key-value pairs.

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 51200.

Key Pattern: `.*`

Value Length Constraints: Minimum length of 0. Maximum length of 51200.

Value Pattern: `.*`

Required: No

**StartCommand**

The command App Runner runs to start your
application.

Type: String

Pattern: `[^\x0a\x0d]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/codeconfigurationvalues.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/codeconfigurationvalues.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/codeconfigurationvalues.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeConfiguration

CodeRepository

All content copied from https://docs.aws.amazon.com/.
