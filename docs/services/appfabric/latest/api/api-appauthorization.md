# AppAuthorization

Contains information about an app authorization.

## Contents

**app**

The name of the application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**appAuthorizationArn**

The Amazon Resource Name (ARN) of the app authorization.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+`

Required: Yes

**appBundleArn**

The Amazon Resource Name (ARN) of the app bundle for the app authorization.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+`

Required: Yes

**authType**

The authorization type.

Type: String

Valid Values: `oauth2 | apiKey`

Required: Yes

**createdAt**

The timestamp of when the app authorization was created.

Type: Timestamp

Required: Yes

**status**

The state of the app authorization.

The following states are possible:

- `PendingConnect`: The initial state of the app authorization. The app
authorization is created but not yet connected.

- `Connected`: The app authorization is connected to the application, and
is ready to be used.

- `ConnectionValidationFailed`: The app authorization received a
validation exception when trying to connect to the application. If the app
authorization is in this state, you should verify the configured credentials and try
to connect the app authorization again.

- `TokenAutoRotationFailed`: AppFabric failed to refresh the access token. If
the app authorization is in this state, you should try to reconnect the app
authorization.

Type: String

Valid Values: `PendingConnect | Connected | ConnectionValidationFailed | TokenAutoRotationFailed`

Required: Yes

**tenant**

Contains information about an application tenant, such as the application display name
and identifier.

Type: [Tenant](api-tenant.md) object

Required: Yes

**updatedAt**

The timestamp of when the app authorization was last updated.

Type: Timestamp

Required: Yes

**authUrl**

The application URL for the OAuth flow.

Type: String

Required: No

**persona**

The user persona of the app authorization.

This field should always be `admin`.

Type: String

Valid Values: `admin | endUser`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/appauthorization.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/appauthorization.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/appauthorization.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiKeyCredential

AppAuthorizationSummary

All content copied from https://docs.aws.amazon.com/.
