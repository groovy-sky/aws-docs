# BlockedPhrasesConfiguration

Provides information about the phrases blocked from chat by your chat control
configuration.

## Contents

**blockedPhrases**

A list of phrases blocked from a Amazon Q Business web experience chat.

###### Note

Each phrase can contain a maximum of 36 characters. The list can contain a maximum of 20 phrases.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 36.

Pattern: `\P{C}*`

Required: No

**systemMessageOverride**

The configured custom message displayed to an end user informing them that they've
used a blocked phrase during chat.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 350.

Pattern: `\P{C}*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/blockedphrasesconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/blockedphrasesconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/blockedphrasesconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicAuthConfiguration

BlockedPhrasesConfigurationUpdate

All content copied from https://docs.aws.amazon.com/.
