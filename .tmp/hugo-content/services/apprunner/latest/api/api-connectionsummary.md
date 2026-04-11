# ConnectionSummary

Provides summary information about an AWS App Runner connection resource.

## Contents

**ConnectionArn**

The Amazon Resource Name (ARN) of this connection.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

**ConnectionName**

The customer-provided connection name.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 32.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,31}`

Required: No

**CreatedAt**

The App Runner connection creation time, expressed as a Unix time stamp.

Type: Timestamp

Required: No

**ProviderType**

The source repository provider.

Type: String

Valid Values: `GITHUB | BITBUCKET`

Required: No

**Status**

The current state of the App Runner connection. When the state is `AVAILABLE`, you can use the connection to create an App Runner service.

Type: String

Valid Values: `PENDING_HANDSHAKE | AVAILABLE | ERROR | DELETED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/connectionsummary.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/connectionsummary.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/connectionsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Connection

CustomDomain

All content copied from https://docs.aws.amazon.com/.
