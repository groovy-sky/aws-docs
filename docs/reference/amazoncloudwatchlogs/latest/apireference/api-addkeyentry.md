# AddKeyEntry

This object defines one key that will be added with the [addKeys](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-addKey) processor.

## Contents

**key**

The key of the new entry to be added to the log event

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**value**

The value of the new entry to be added to the log event

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**overwriteIfExists**

Specifies whether to overwrite the value if the key already exists in the log event. If
you omit this, the default is `false`.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/addkeyentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/addkeyentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/addkeyentry.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccountPolicy

AddKeys

All content copied from https://docs.aws.amazon.com/.
