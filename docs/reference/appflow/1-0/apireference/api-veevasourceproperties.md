# VeevaSourceProperties

The properties that are applied when using Veeva as a flow source.

## Contents

**object**

The object specified in the Veeva flow source.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

**documentType**

The document type specified in the Veeva document extract flow.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `[\s\w_-]+`

Required: No

**includeAllVersions**

Boolean value to include All Versions of files in Veeva document extract flow.

Type: Boolean

Required: No

**includeRenditions**

Boolean value to include file renditions in Veeva document extract flow.

Type: Boolean

Required: No

**includeSourceFiles**

Boolean value to include source files in Veeva document extract flow.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/VeevaSourceProperties)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/VeevaSourceProperties)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/VeevaSourceProperties)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VeevaMetadata

ZendeskConnectorProfileCredentials
