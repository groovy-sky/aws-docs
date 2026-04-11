# PredictQAppInputOptions

The input options for generating an Q App definition.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**conversation**

A conversation to use as input for generating the Q App definition.

Type: Array of [ConversationMessage](api-qapps-conversationmessage.md) objects

Array Members: Minimum number of 1 item. Maximum number of 25 items.

Required: No

**problemStatement**

A problem statement to use as input for generating the Q App definition.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10000.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/predictqappinputoptions.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/predictqappinputoptions.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/predictqappinputoptions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictAppDefinition

PrincipalOutput

All content copied from https://docs.aws.amazon.com/.
