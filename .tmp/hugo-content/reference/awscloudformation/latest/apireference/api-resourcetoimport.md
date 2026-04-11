# ResourceToImport

Describes the target resource of an import operation.

## Contents

**LogicalResourceId**

The logical ID of the target resource as specified in the template.

Type: String

Required: Yes

**ResourceIdentifier**
ResourceIdentifier.entry.N.key (key)

ResourceIdentifier.entry.N.value (value)

A key-value pair that identifies the target resource. The key is an identifier property (for
example, `BucketName` for `AWS::S3::Bucket` resources) and the value is the
actual property value (for example, `MyS3Bucket`).

Type: String to string map

Map Entries: Maximum number of 256 items.

Key Length Constraints: Minimum length of 1. Maximum length of 2048.

Value Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**ResourceType**

The type of resource to import into your stack, such as `AWS::S3::Bucket`. For a
list of supported resource types, see [Resource type\
support for imports and drift detection](../../../../services/cloudformation/latest/userguide/resource-import-supported-resources.md) in the
_AWS CloudFormation User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/resourcetoimport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/resourcetoimport.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/resourcetoimport.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTargetDefinition

RollbackConfiguration

All content copied from https://docs.aws.amazon.com/.
