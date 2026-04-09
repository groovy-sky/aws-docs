# ImportDataRule

Specifies a rule that controls how data is imported from S3 into the file
system.

## Contents

**prefix**

The S3 key prefix that scopes this import rule. Only objects with keys beginning with
this prefix are subject to the rule.

Type: String

Pattern: `(|.*/)`

Required: Yes

**sizeLessThan**

The upper size limit in bytes for this import rule. Only objects with a size strictly
less than this value will have data imported into the file system.

Type: Long

Valid Range: Minimum value of 0. Maximum value of 52673613135872.

Required: Yes

**trigger**

The event that triggers data import. Valid values are
`ON_DIRECTORY_FIRST_ACCESS` (import when a directory is first accessed) and
`ON_FILE_ACCESS` (import when a file is accessed).

Type: String

Valid Values: `ON_DIRECTORY_FIRST_ACCESS | ON_FILE_ACCESS`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/importdatarule.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/importdatarule.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/importdatarule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpirationDataRule

ListAccessPointsDescription

All content copied from https://docs.aws.amazon.com/.
