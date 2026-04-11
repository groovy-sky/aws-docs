# CsvOptions

Processing options for the CSV file being imported.

## Contents

###### Note

In the following list, the required parameters are described first.

**Delimiter**

The delimiter used for separating items in the CSV file being imported.

Type: String

Length Constraints: Fixed length of 1.

Pattern: `[,;:|\t ]`

Required: No

**HeaderList**

List of the headers used to specify a common header for all source CSV files being
imported. If this field is specified then the first line of each CSV file is treated as
data instead of the header. If this field is not specified the the first line of each
CSV file is treated as the header.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 255 items.

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `[\x20-\x21\x23-\x2B\x2D-\x7E]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/csvoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/csvoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/csvoptions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateReplicationGroupMemberAction

Delete

All content copied from https://docs.aws.amazon.com/.
