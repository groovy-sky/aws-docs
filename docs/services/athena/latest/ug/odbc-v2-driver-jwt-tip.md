# JWT Trusted identity propagation credentials provider

This authentication type allows you to use a JSON web token (JWT) obtained from an
external identity provider as a connection parameter to authenticate with Athena. You can use
this plugin, to enable support for corporate identities via trusted identity propagation.

With trusted identity propagation, identity context is added to an IAM role to identify the user requesting
access to AWS resources. For information on enabling and using trusted identity propagation, see [What is trusted identity propagation?](https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation.html).

## Authentication type

**Connection string name****Parameter type****Default value****Connection string example**AuthenticationTypeRequired`IAM Credentials``AuthenticationType=JWT_TIP;`

## JWT web identity token

The JWT token obtained from an external federated identity provider. This token will
be used to authenticate with Athena. Token caching is enabled by default and allows the
same Identity Center access token to be used across driver connections. We recommend to
provide a fresh JWT token upon "Testing Connection" as the exchanged token is present
only during the duration driver instance is active.

**Connection string name****Parameter type****Default value****Connection string example**web\_identity\_tokenRequired`none``web_identity_token=eyJhbGc...<remainder of token>;`

## Workgroup Arn

The Amazon Resource Name (ARN) of the Amazon Athena workgroup. For more information about
workgroups, see [WorkGroup](../../../../reference/athena/latest/apireference/api-workgroup.md).

**Connection string name****Parameter type****Default value****Connection string example**WorkGroupArnRequired`none``WorkgroupArn=arn:aws:athena:us-west-2:111122223333:workgroup/primary`

## JWT application role ARN

The ARN of the role to assume. This role is used for JWT exchange, getting IAM
Identity Center Customer Managed application ARN through workgroup tags, and getting
Access Role ARN. For more information about assuming roles, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html).

**Connection string name****Parameter type****Default value****Connection string example**ApplicationRoleArnRequired`none``ApplicationRoleArn=arn:aws:iam::111122223333:role/applicationRole;`

## Role session name

A name for the session. It can be anything you like, but typically you pass the name
or identifier that's associated with the user who is using your application. That way,
the temporary security credentials that your application will use are associated with
that user.

**Connection string name****Parameter type****Default value****Connection string example**role\_session\_nameRequired`none``role_session_name=familiarname;`

## Session duration

The duration, in seconds, of the role session. For more information about session
duration, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html).

**Connection string name****Parameter type****Default value****Connection string example**durationOptional`3600``duration=900;`

## JWT access role ARN

The ARN of the role to assume. This is the role that Athena assumes to make calls on
your behalf. For more information about assuming roles, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html).

**Connection string name****Parameter type****Default value****Connection string example**AccessRoleArnOptional`none``AccessRoleArn=arn:aws:iam::111122223333:role/accessRole;`

## IAM Identity Center customer managed application ARN

The ARN of IAM Identity Center customer managed IDC application. For more information about Customer
Managed Applications, see [customer managed\
applications](https://docs.aws.amazon.com/singlesignon/latest/userguide/customermanagedapps.html).

**Connection string name****Parameter type****Default value****Connection string example**CustomerIdcApplicationArnOptional`none``CustomerIdcApplicationArn=arn:aws:sso::111122223333:application/ssoins-111122223333/apl-111122223333`

## Enable file cache

Enables a temporary credentials cache. This connection parameter allows you to cache
temporary credentials and reuse it between multiple processes. Use this option to reduce
the number of web identity tokens when you use BI tools such as Microsoft Power BI. By
default, the driver uses `%USERPROFILE%` in Windows and `HOME`
path to write the file caches. Ensure that you provide read and write access for the
path present in these two environment variables, for a better experience.

###### Note

Starting in v2.1.0.0, cached credentials are stored as plaintext JSON in the
`user-profile/.athena-odbc/` directory with file permissions
restricted to the owning user, consistent with how the AWS CLI
protects locally stored credentials.

**Connection string name****Parameter type****Default value****Connection string example**JwtTipFileCacheOptional`0``JwtTipFileCache=1;`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JWT

Browser trusted identity propagation
