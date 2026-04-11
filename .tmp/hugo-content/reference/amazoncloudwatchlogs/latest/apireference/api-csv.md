# CSV

The `CSV` processor parses comma-separated values (CSV) from the log events
into columns.

For more information about this processor including examples, see [csv](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-csv) in the _CloudWatch Logs User Guide_.

## Contents

**columns**

An array of names to use for the columns in the transformed log event.

If you omit this, default column names ( `[column_1, column_2 ...]`) are
used.

Type: Array of strings

Array Members: Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**delimiter**

The character used to separate each column in the original comma-separated value log
event. If you omit this, the processor looks for the comma `,` character as the
delimiter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2.

Required: No

**destination**

The path to the parent field to put transformed key value pairs under.
If you omit this value, the key value pairs will be placed under the root node.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**quoteCharacter**

The character used used as a text qualifier for a single column of data. If you omit this,
the double quotation mark `"` character is used.

Type: String

Length Constraints: Fixed length of 1.

Required: No

**source**

The path to the field in the log event that has the comma separated values to be parsed.
If you omit this value, the whole log message is processed.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/csv.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/csv.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/csv.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyValueEntry

DataSource

All content copied from https://docs.aws.amazon.com/.
