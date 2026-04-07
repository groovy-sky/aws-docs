# CustomRule

Describes a custom rewrite or redirect rule.

## Contents

**source**

The source pattern for a URL rewrite or redirect rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(?s).+`

Required: Yes

**target**

The target pattern for a URL rewrite or redirect rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(?s).+`

Required: Yes

**condition**

The condition for a URL rewrite or redirect rule, such as a country code.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `(?s).*`

Required: No

**status**

The status code for a URL rewrite or redirect rule.

200

Represents a 200 rewrite rule.

301

Represents a 301 (moved permanently) redirect rule. This and all future
requests should be directed to the target URL.

302

Represents a 302 temporary redirect rule.

404

Represents a 404 redirect rule.

404-200

Represents a 404 rewrite rule.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 7.

Pattern: `.{3,7}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/amplify-2017-07-25/CustomRule)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/amplify-2017-07-25/CustomRule)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/amplify-2017-07-25/CustomRule)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CertificateSettings

DomainAssociation
