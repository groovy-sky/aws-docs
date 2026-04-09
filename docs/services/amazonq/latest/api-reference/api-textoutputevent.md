# TextOutputEvent

An output event for an AI-generated response in an Amazon Q Business web
experience.

## Contents

**conversationId**

The identifier of the conversation with which the text output event is
associated.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**systemMessage**

An AI-generated message in a `TextOutputEvent`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**systemMessageId**

The identifier of an AI-generated message in a `TextOutputEvent`.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**systemMessageType**

The type of AI-generated message in a `TextOutputEvent`. Amazon Q Business
currently supports two types of messages:

- `RESPONSE` \- The Amazon Q Business system response.

- `GROUNDED_RESPONSE` \- The corrected, hallucination-reduced,
response returned by Amazon Q Business. Available only if hallucination reduction is
supported and configured for the application and detected in the end user chat
query by Amazon Q Business.

Type: String

Valid Values: `RESPONSE | GROUNDED_RESPONSE`

Required: No

**userMessageId**

The identifier of an end user message in a `TextOutputEvent`.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/textoutputevent.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/textoutputevent.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/textoutputevent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextInputEvent

TextSegment

All content copied from https://docs.aws.amazon.com/.
