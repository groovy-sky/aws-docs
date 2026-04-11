# JWT

The JWT (JSON Web Token) plugin provides an interface that uses JSON Web Tokens to assume
an Amazon IAM role. The configuration depends on the identity provider. For information
about configuring federation for Google Cloud and AWS, see [Configure workload identity federation with AWS or Azure](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-clouds) in the Google Cloud
documentation.

## Authentication type

**Connection string name****Parameter type****Default value****Connection string example**AuthenticationTypeRequired`IAM Credentials``AuthenticationType=JWT;`

## Preferred role

The Amazon Resource Name (ARN) of the role to assume. For more information about ARN
roles, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the _AWS Security Token Service API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**preferred\_roleOptional`none``preferred_role=arn:aws:IAM::123456789012:id/user1;`

## Session duration

The duration, in seconds, of the role session. For more information about session
duration, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the
_AWS Security Token Service API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**durationOptional`900``duration=900;`

## JSON web token

The JSON web token that is used to retrieve IAM temporary credentials using the
[AssumeRoleWithWebIdentity](../../../../reference/sts/latest/apireference/api-assumerolewithwebidentity.md) AWS STS API action. For information about
generating JSON web tokens for Google Cloud Platform (GCP) users, see [Using JWT OAuth tokens](https://cloud.google.com/apigee/docs/api-platform/security/oauth/using-jwt-oauth) in the Google Cloud documentation.

**Connection string name****Parameter type****Default value****Connection string example**`web_identity_token`Required`none``web_identity_token=eyJhbGc...<remainder of
                            token>;`

## Role session name

A name for the session. A common technique is to use the name or identifier of the
user of your application as the role session name. This conveniently associates the
temporary security credentials that your application uses with the corresponding
user.

**Connection string name****Parameter type****Default value****Connection string example**role\_session\_nameRequired`none``role_session_name=familiarname;`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Instance profile

JWT trusted identity propagation

All content copied from https://docs.aws.amazon.com/.
