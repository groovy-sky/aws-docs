# CodeConfiguration

Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.

## Contents

**ConfigurationSource**

The source of the App Runner configuration. Values are interpreted as follows:

- `REPOSITORY` – App Runner reads configuration values from the `apprunner.yaml` file in the source code repository and
ignores `CodeConfigurationValues`.

- `API` – App Runner uses configuration values provided in `CodeConfigurationValues` and ignores the
`apprunner.yaml` file in the source code repository.

Type: String

Valid Values: `REPOSITORY | API`

Required: Yes

**CodeConfigurationValues**

The basic configuration for building and running the App Runner service. Use it to quickly launch an App Runner service without providing a
`apprunner.yaml` file in the source code repository (or ignoring the file if it exists).

Type: [CodeConfigurationValues](api-codeconfigurationvalues.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/codeconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/codeconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/codeconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateValidationRecord

CodeConfigurationValues

All content copied from https://docs.aws.amazon.com/.
