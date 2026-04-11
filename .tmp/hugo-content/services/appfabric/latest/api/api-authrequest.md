# AuthRequest

Contains authorization request information, which is required for AWS AppFabric to get the
OAuth2 access token for an application.

## Contents

**code**

The authorization code returned by the application after permission is granted in the
application OAuth page (after clicking on the AuthURL).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**redirectUri**

The redirect URL that is specified in the AuthURL and the application client.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `https://[-a-zA-Z0-9-._~:/?#@!$&'()*+,;=]+`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/authrequest.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/authrequest.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/authrequest.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuditLogProcessingConfiguration

Credential

All content copied from https://docs.aws.amazon.com/.
