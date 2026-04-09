# Azure AD

Azure AD is a SAML-based authentication plugin that works with Azure AD identity provider.
This plugin does not support multifactor authentication (MFA). If you require MFA support,
consider using the `BrowserAzureAD` plugin instead.

## Authentication Type

**Connection string name****Parameter type****Default value****Connection string example**AuthenticationTypeRequired`IAM Credentials``AuthenticationType=AzureAD;`

## User ID

Your user name for connecting to Azure AD.

**Connection string name****Parameter type****Default value****Connection string example**UIDRequired`none``UID=jane.doe@example.com;`

## Password

Your password for connecting to Azure AD.

**Connection string name****Parameter type****Default value****Connection string example**PWDRequired`none``PWD=password_3EXAMPLE;`

## Preferred role

The Amazon Resource Name (ARN) of the role to assume. For information about ARN roles,
see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the _AWS Security Token Service API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**preferred\_roleOptional`none``preferred_role=arn:aws:iam::123456789012:id/user1;`

## Session duration

The duration, in seconds, of the role session. For more information, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the _AWS Security Token Service API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**durationOptional`900``duration=900;`

## Tenant ID

Specifies your application tenant ID.

**Connection string name****Parameter type****Default value****Connection string example**idp\_tenantRequired`none``idp_tenant=123zz112z-z12d-1z1f-11zz-f111aa111234;`

## Client ID

Specifies your application client ID.

**Connection string name****Parameter type****Default value****Connection string example**client\_idRequired`none``client_id=9178ac27-a1bc-1a2b-1a2b-a123abcd1234;`

## Client secret

Specifies your client secret.

**Connection string name****Parameter type****Default value****Connection string example**client\_secretRequired`none``client_secret=zG12q~.xzG1xxxZ1wX1.~ZzXXX1XxkHZizeT1zzZ;`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AD FS

Browser Azure AD

All content copied from https://docs.aws.amazon.com/.
