# StorageLensGroupAndOperator

A logical operator that allows multiple filter conditions to be joined for more complex
comparisons of Storage Lens group data.

## Contents

**MatchAnyPrefix**

Contains a list of prefixes. At least one prefix must be specified. Up to 10 prefixes
are allowed.

Type: Array of strings

Required: No

**MatchAnySuffix**

Contains a list of suffixes. At least one suffix must be specified. Up to 10 suffixes
are allowed.

Type: Array of strings

Required: No

**MatchAnyTag**

Contains the list of object tags. At least one object tag must be specified. Up to 10
object tags are allowed.

Type: Array of [S3Tag](api-control-s3tag.md) data types

Required: No

**MatchObjectAge**

Contains `DaysGreaterThan` and `DaysLessThan` to define the
object age range (minimum and maximum number of days).

Type: [MatchObjectAge](api-control-matchobjectage.md) data type

Required: No

**MatchObjectSize**

Contains `BytesGreaterThan` and `BytesLessThan` to define the
object size range (minimum and maximum number of Bytes).

Type: [MatchObjectSize](api-control-matchobjectsize.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/storagelensgroupandoperator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/storagelensgroupandoperator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/storagelensgroupandoperator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensGroup

StorageLensGroupFilter

All content copied from https://docs.aws.amazon.com/.
