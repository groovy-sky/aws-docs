# RenameKeyEntry

This object defines one key that will be renamed with the [renameKey](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-renameKey) processor.

## Contents

**key**

The key to rename

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**renameTo**

The string to use for the new key name

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**overwriteIfExists**

Specifies whether to overwrite the existing value if the destination key already exists.
The default is `false`

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/renamekeyentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/renamekeyentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/renamekeyentry.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RejectedLogEventsInfo

RenameKeys

All content copied from https://docs.aws.amazon.com/.
