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

Type: [User](api-qapps-user.md) object

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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/qappsessiondata.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/qappsessiondata.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/qappsessiondata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrincipalOutput

QPluginCard

All content copied from https://docs.aws.amazon.com/.
