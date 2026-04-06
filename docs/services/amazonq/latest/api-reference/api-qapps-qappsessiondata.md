# QAppSessionData

The response collected for a Amazon Q App session. This container represents a single
response to a Q App session.

## Contents

**cardId**

The card Id associated with the response submitted for a Q App session.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**user**

The user who submitted the response for a Q App session.

Type: [User](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_User.html) object

Required: Yes

**submissionId**

The unique identifier of the submission.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: No

**timestamp**

The date and time when the session data is submitted.

Type: Timestamp

Required: No

**value**

The response submitted for a Q App session.

Type: JSON value

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/QAppSessionData)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/QAppSessionData)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/QAppSessionData)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PrincipalOutput

QPluginCard
