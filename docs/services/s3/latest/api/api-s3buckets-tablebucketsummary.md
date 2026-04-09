# TableBucketSummary

Contains details about a table bucket.

## Contents

**arn**

The Amazon Resource Name (ARN) of the table bucket.

Type: String

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

**createdAt**

The date and time the table bucket was created at.

Type: Timestamp

Required: Yes

**name**

The name of the table bucket.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `[0-9a-z-]*`

Required: Yes

**ownerAccountId**

The ID of the account that owns the table bucket.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `[0-9].*`

Required: Yes

**tableBucketId**

The system-assigned unique identifier for the table bucket.

Type: String

Required: No

**type**

The type of the table bucket.

Type: String

Valid Values: `customer | aws`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/tablebucketsummary.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/tablebucketsummary.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/tablebucketsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableBucketReplicationRule

TableMaintenanceConfigurationValue

All content copied from https://docs.aws.amazon.com/.
