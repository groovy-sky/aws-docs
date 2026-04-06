# RetrieverConfiguration

Provides information on how the retriever used for your Amazon Q Business application is
configured.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**kendraIndexConfiguration**

Provides information on how the Amazon Kendra index used as a retriever for your
Amazon Q Business application is configured.

Type: [KendraIndexConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_KendraIndexConfiguration.html) object

Required: No

**nativeIndexConfiguration**

Provides information on how a Amazon Q Business index used as a retriever for your
Amazon Q Business application is configured.

Type: [NativeIndexConfiguration](api-nativeindexconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/RetrieverConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/RetrieverConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/RetrieverConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Retriever

RetrieverContentSource
