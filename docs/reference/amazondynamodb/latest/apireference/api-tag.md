# Tag

Describes a tag. A tag is a key-value pair. You can add up to 50 tags to a single
DynamoDB table.

AWS-assigned tag names and values are automatically assigned the
`aws:` prefix, which the user cannot assign. AWS-assigned
tag names do not count towards the tag limit of 50. User-assigned tag names have the
prefix `user:` in the Cost Allocation Report. You cannot backdate the
application of a tag.

For an overview on tagging DynamoDB resources, see [Tagging\
for DynamoDB](../../../../services/dynamodb/latest/developerguide/tagging.md) in the _Amazon DynamoDB Developer_
_Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**Key**

The key of the tag. Tag keys are case sensitive. Each DynamoDB table can
only have up to one tag with the same key. If you try to add an existing tag (same key),
the existing tag value will be updated to the new value.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**Value**

The value of the tag. Tag values are case-sensitive and can be null.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/tag.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/tag.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/tag.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableWarmThroughputDescription

ThrottlingReason

All content copied from https://docs.aws.amazon.com/.
