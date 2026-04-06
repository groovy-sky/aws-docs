# PathElement

A single element in a path through the JSON representation of a policy.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**index**

Refers to an index in a JSON array.

Type: Integer

Required: No

**key**

Refers to a key in a JSON object.

Type: String

Required: No

**substring**

Refers to a substring of a literal string in a JSON object.

Type: [Substring](api-substring.md) object

Required: No

**value**

Refers to the value associated with a given key in a JSON object.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/PathElement)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/PathElement)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/PathElement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkOriginConfiguration

PolicyGeneration
