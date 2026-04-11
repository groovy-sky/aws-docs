# CopyValue

This processor copies values within a log event. You can also use this processor to add
metadata to log events by copying the values of the following metadata keys into the log
events: `@logGroupName`, `@logGroupStream`, `@accountId`,
`@regionName`.

For more information about this processor including examples, see [copyValue](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-copyValue) in the _CloudWatch Logs User Guide_.

## Contents

**entries**

An array of `CopyValueEntry` objects, where each object contains the
information about one field value to copy.

Type: Array of [CopyValueEntry](api-copyvalueentry.md) objects

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/copyvalue.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/copyvalue.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/copyvalue.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigurationTemplateDeliveryConfigValues

CopyValueEntry

All content copied from https://docs.aws.amazon.com/.
