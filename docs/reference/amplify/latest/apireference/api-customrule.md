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

- [AWS SDK for C++](../../../goto/sdkforcpp/amplify-2017-07-25/customrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/amplify-2017-07-25/customrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/amplify-2017-07-25/customrule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateSettings

DomainAssociation

All content copied from https://docs.aws.amazon.com/.
