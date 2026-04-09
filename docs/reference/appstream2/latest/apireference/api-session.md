# Session

Describes a streaming session.

## Contents

**FleetName**

The name of the fleet for the streaming session.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**Id**

The identifier of the streaming session.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**StackName**

The name of the stack for the streaming session.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**State**

The current state of the streaming session.

Type: String

Valid Values: `ACTIVE | PENDING | EXPIRED`

Required: Yes

**UserId**

The identifier of the user for whom the session was created.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 128.

Required: Yes

**AuthenticationType**

The authentication method. The user is authenticated using a streaming URL
( `API`) or SAML 2.0 federation ( `SAML`).

Type: String

Valid Values: `API | SAML | USERPOOL | AWS_AD`

Required: No

**ConnectionState**

Specifies whether a user is connected to the streaming session.

Type: String

Valid Values: `CONNECTED | NOT_CONNECTED`

Required: No

**InstanceDrainStatus**

The drain status of the instance hosting the streaming session. This only applies to multi-session fleets.

Type: String

Valid Values: `ACTIVE | DRAINING | NOT_APPLICABLE`

Required: No

**InstanceId**

The identifier for the instance hosting the session.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**MaxExpirationTime**

The time when the streaming session is set to expire. This time is based on the `MaxUserDurationinSeconds` value, which determines the maximum length of time that a streaming session can run. A streaming session might end earlier than the time specified in `SessionMaxExpirationTime`, when the `DisconnectTimeOutInSeconds` elapses or the user chooses to end his or her session. If the `DisconnectTimeOutInSeconds` elapses, or the user chooses to end his or her session, the streaming instance is terminated and the streaming session ends.

Type: Timestamp

Required: No

**NetworkAccessConfiguration**

The network details for the streaming session.

Type: [NetworkAccessConfiguration](api-networkaccessconfiguration.md) object

Required: No

**StartTime**

The time when a streaming instance is dedicated for the user.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/session.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/session.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/session.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceAccountCredentials

SharedImagePermissions

All content copied from https://docs.aws.amazon.com/.
