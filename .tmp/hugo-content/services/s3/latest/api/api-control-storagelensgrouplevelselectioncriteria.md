# StorageLensGroupLevelSelectionCriteria

Indicates which Storage Lens group ARNs to include or exclude in the Storage Lens group
aggregation. You can only attach Storage Lens groups to your Storage Lens dashboard if
they're included in your Storage Lens group aggregation. If this value is left null, then
all Storage Lens groups are selected.

## Contents

**Exclude**

Indicates which Storage Lens group ARNs to exclude from the Storage Lens group
aggregation.

Type: Array of strings

Length Constraints: Minimum length of 4. Maximum length of 1024.

Pattern: `arn:[a-z\-]+:s3:[a-z0-9\-]+:\d{12}:storage\-lens\-group\/.*`

Required: No

**Include**

Indicates which Storage Lens group ARNs to include in the Storage Lens group
aggregation.

Type: Array of strings

Length Constraints: Minimum length of 4. Maximum length of 1024.

Pattern: `arn:[a-z\-]+:s3:[a-z0-9\-]+:\d{12}:storage\-lens\-group\/.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/storagelensgrouplevelselectioncriteria.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/storagelensgrouplevelselectioncriteria.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/storagelensgrouplevelselectioncriteria.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensGroupLevel

StorageLensGroupOrOperator

All content copied from https://docs.aws.amazon.com/.
