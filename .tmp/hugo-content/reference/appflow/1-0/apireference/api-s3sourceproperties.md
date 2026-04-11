# S3SourceProperties

The properties that are applied when Amazon S3 is being used as the flow source.

## Contents

**bucketName**

The Amazon S3 bucket name where the source files are stored.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `\S+`

Required: Yes

**bucketPrefix**

The object key for the Amazon S3 bucket in which the source files are stored.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

**s3InputFormatConfig**

When you use Amazon S3 as the source, the configuration format that you provide
the flow input data.

Type: [S3InputFormatConfig](api-s3inputformatconfig.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/s3sourceproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/s3sourceproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/s3sourceproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3OutputFormatConfig

SalesforceConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
