# ResourceIdentifierSummary

Describes the target resources of a specific type in your import template (for example,
all `AWS::S3::Bucket` resources) and the properties you can provide during the
import to identify resources of that type.

## Contents

**LogicalResourceIds.member.N**

The logical IDs of the target resources of the specified `ResourceType`, as
defined in the import template.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 200 items.

Required: No

**ResourceIdentifiers.member.N**

The resource properties you can provide during the import to identify your target
resources. For example, `BucketName` is a possible identifier property for
`AWS::S3::Bucket` resources.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**ResourceType**

The template resource type of the target resources, such as
`AWS::S3::Bucket`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/resourceidentifiersummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/resourceidentifiersummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/resourceidentifiersummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceDriftIgnoredAttribute

ResourceLocation

All content copied from https://docs.aws.amazon.com/.
