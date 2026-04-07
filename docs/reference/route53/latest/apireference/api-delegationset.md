# DelegationSet

A complex type that lists the name servers in a delegation set, as well as the
`CallerReference` and the `ID` for the delegation set.

## Contents

**NameServers**

A complex type that contains a list of the authoritative name servers for a hosted
zone or for a reusable delegation set.

Type: Array of strings

Array Members: Minimum number of 1 item.

Length Constraints: Maximum length of 1024.

Required: Yes

**CallerReference**

The value that you specified for `CallerReference` when you created the
reusable delegation set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**Id**

The ID that Amazon Route 53 assigns to a reusable delegation set.

Type: String

Length Constraints: Maximum length of 32.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/DelegationSet)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/DelegationSet)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/DelegationSet)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Coordinates

Dimension
