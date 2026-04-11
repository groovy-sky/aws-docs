# Validator

A validator provides a syntactic or semantic check to ensure the configuration that you
want to deploy functions as intended. To validate your application configuration data, you
provide a schema or an AWS Lambda function that runs against the configuration. The
configuration deployment or update can only proceed when the configuration data is valid.
For more information, see [About validators](../../../../services/appconfig/latest/userguide/appconfig-creating-configuration-profile.md#appconfig-creating-configuration-and-profile-validators) in the _AWS AppConfig User Guide_.

## Contents

**Content**

Either the JSON Schema content or the Amazon Resource Name (ARN) of an Lambda
function.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 32768.

Required: Yes

**Type**

AWS AppConfig supports validators of type `JSON_SCHEMA` and
`LAMBDA`

Type: String

Valid Values: `JSON_SCHEMA | LAMBDA`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/validator.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/validator.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/validator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Parameter

AWS AppConfig Data

All content copied from https://docs.aws.amazon.com/.
