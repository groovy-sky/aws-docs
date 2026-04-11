# ParsePostgres

Use this processor to parse RDS for PostgreSQL vended logs, extract fields, and
and convert them into a JSON format. This processor always processes the entire log event
message. For more information about this processor including examples, see [parsePostGres](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parsePostGres).

For more information about RDS for PostgreSQL log format, see [RDS for PostgreSQL database log filesTCP flag sequence](../../../../services/amazonrds/latest/userguide/user-logaccess-concepts-postgresql.md#USER_LogAccess.Concepts.PostgreSQL.Log_Format.log-line-prefix).

###### Important

If you use this processor, it must be the first processor in your transformer.

## Contents

**source**

Omit this parameter and the whole log message will be processed by this processor. No
other value than `@message` is allowed for `source`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/parsepostgres.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/parsepostgres.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/parsepostgres.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseKeyValue

ParseRoute53

All content copied from https://docs.aws.amazon.com/.
