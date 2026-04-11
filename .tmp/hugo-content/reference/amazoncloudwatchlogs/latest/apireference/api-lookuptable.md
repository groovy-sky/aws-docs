# LookupTable

Contains metadata about a lookup table returned by `DescribeLookupTables`.

## Contents

**description**

The description of the lookup table.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**kmsKeyId**

The ARN of the AWS KMS key used to encrypt the lookup table data, if
applicable.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**lastUpdatedTime**

The time when the lookup table was last updated, expressed as the number of
milliseconds after `Jan 1, 1970 00:00:00 UTC`.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**lookupTableArn**

The ARN of the lookup table.

Type: String

Required: No

**lookupTableName**

The name of the lookup table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `^[a-zA-Z0-9_]+$`

Required: No

**recordsCount**

The number of data rows in the lookup table, excluding the header row.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**sizeBytes**

The size of the lookup table in bytes.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**tableFields**

The column headers from the first row of the CSV file.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/lookuptable.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/lookuptable.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/lookuptable.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogStream

LowerCaseString

All content copied from https://docs.aws.amazon.com/.
