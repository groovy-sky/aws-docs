# PredictQAppInputOptions

The input options for generating an Q App definition.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**conversation**

A conversation to use as input for generating the Q App definition.

Type: Array of [ConversationMessage](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_ConversationMessage.html) objects

Array Members: Minimum number of 1 item. Maximum number of 25 items.

Required: No

**problemStatement**

A problem statement to use as input for generating the Q App definition.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10000.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/PredictQAppInputOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/PredictQAppInputOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/PredictQAppInputOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PredictAppDefinition

PrincipalOutput
