# NativeIndexConfiguration

Configuration information for an Amazon Q Business index.

## Contents

**indexId**

The identifier for the Amazon Q Business index.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**boostingOverride**

Overrides the default boosts applied by Amazon Q Business to supported document attribute
data types.

Type: String to [DocumentAttributeBoostingConfiguration](api-documentattributeboostingconfiguration.md) object map

Map Entries: Maximum number of items.

Key Length Constraints: Minimum length of 1. Maximum length of 200.

Key Pattern: `[a-zA-Z0-9_][a-zA-Z0-9_-]*`

Required: No

**version**

A read-only field that specifies the version of the `NativeIndexConfiguration`.

Amazon Q Business introduces enhanced document retrieval capabilities in version 2 of `NativeIndexConfiguration`, focusing on streamlined metadata boosting that prioritizes recency and source relevance to deliver more accurate responses to your queries. Version 2 has the following differences from version 1:

- Version 2 supports a single Date field (created\_at OR last\_updated\_at) for recency boosting

- Version 2 supports a single String field with an ordered list of up to 5 values

- Version 2 introduces number-based boost levels (ONE, TWO) alongside the text-based levels

- Version 2 allows specifying prioritization between Date and String fields

- Version 2 maintains backward compatibility with existing configurations

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/nativeindexconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/nativeindexconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/nativeindexconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataEvent

NoAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
