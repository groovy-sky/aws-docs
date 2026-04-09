# StorageLensGroupOrOperator

A container element for specifying `Or` rule conditions. The rule conditions
determine the subset of objects to which the `Or` rule applies. Objects can
match any of the listed filter conditions, which are joined by the `Or` logical
operator. Only one of each filter condition is allowed.

## Contents

**MatchAnyPrefix**

Filters objects that match any of the specified prefixes.

Type: Array of strings

Required: No

**MatchAnySuffix**

Filters objects that match any of the specified suffixes.

Type: Array of strings

Required: No

**MatchAnyTag**

Filters objects that match any of the specified S3 object tags.

Type: Array of [S3Tag](api-control-s3tag.md) data types

Required: No

**MatchObjectAge**

Filters objects that match the specified object age range.

Type: [MatchObjectAge](api-control-matchobjectage.md) data type

Required: No

**MatchObjectSize**

Filters objects that match the specified object size range.

Type: [MatchObjectSize](api-control-matchobjectsize.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/storagelensgrouporoperator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/storagelensgrouporoperator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/storagelensgrouporoperator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensGroupLevelSelectionCriteria

StorageLensTableDestination

All content copied from https://docs.aws.amazon.com/.
