# LogConfig

The Amazon CloudWatch Logs configuration.

## Contents

**cloudWatchLogsRoleArn**

The service role that AWS AppSync assumes to publish to CloudWatch
logs in your account.

Type: String

Required: Yes

**fieldLogLevel**

The field logging level. Values can be NONE, ERROR, or ALL.

- **NONE**: No field-level logs are
captured.

- **ERROR**: Logs the following information only for
the fields that are in error:

- The error section in the server response.

- Field-level errors.

- The generated request/response functions that got resolved for error
fields.

- **ALL**: The following information is logged for
all fields in the query:

- Field-level tracing information.

- The generated request/response functions that got resolved for each
field.

Type: String

Valid Values: `NONE | ERROR | ALL | INFO | DEBUG`

Required: Yes

**excludeVerboseContent**

Set to TRUE to exclude sections that contain information such as headers, context, and
evaluated mapping templates, regardless of logging level.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/logconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/logconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/logconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaDataSourceConfig

OpenIDConnectConfig

All content copied from https://docs.aws.amazon.com/.
