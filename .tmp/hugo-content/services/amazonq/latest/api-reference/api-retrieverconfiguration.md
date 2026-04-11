# RetrieverConfiguration

Provides information on how the retriever used for your Amazon Q Business application is
configured.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**kendraIndexConfiguration**

Provides information on how the Amazon Kendra index used as a retriever for your
Amazon Q Business application is configured.

Type: [KendraIndexConfiguration](api-kendraindexconfiguration.md) object

Required: No

**nativeIndexConfiguration**

Provides information on how a Amazon Q Business index used as a retriever for your
Amazon Q Business application is configured.

Type: [NativeIndexConfiguration](api-nativeindexconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/retrieverconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/retrieverconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/retrieverconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Retriever

RetrieverContentSource

All content copied from https://docs.aws.amazon.com/.
