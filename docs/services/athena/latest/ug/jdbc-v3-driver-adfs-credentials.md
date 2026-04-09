# AD FS credentials

A SAML-based authentication mechanism that enables authentication to Athena using
Microsoft Active Directory Federation Services (AD FS). This method assumes that the
user has already set up a federation between Athena and AD FS.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `ADFS`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProviderAWSCredentialsProviderClass (deprecated)Requirednone`ADFS`

## User

The email address of the AD FS user to use for authentication with AD FS.

Parameter nameAliasParameter typeDefault valueUserUID (deprecated)Required for form-based authentication. Optional for Windows
Integrated Authentication.none

## Password

The password for the AD FS user.

Parameter nameAliasParameter typeDefault valuePasswordPWD (deprecated)Required for form-based authentication. Optional for Windows
Integrated Authentication.none

## ADFS host name

The address for your AD FS server.

Parameter nameAliasParameter typeDefault valueAdfsHostNameIdP\_Host (deprecated)Requirednone

## ADFS port number

The port number to use to connect to your AD FS server.

Parameter nameAliasParameter typeDefault valueAdfsPortNumberIdP\_Port (deprecated)Requirednone

## ADFS relying party

The trusted relying party. Use this parameter to override the AD FS relying party
endpoint URL.

Parameter nameAliasParameter typeDefault valueAdfsRelyingPartyLoginToRP (deprecated)Optional`urn:amazon:webservices`

## ADFS WIA enabled

Boolean. Use this parameter to enable Windows Integrated Authentication (WIA) with
AD FS.

Parameter nameAliasParameter typeDefault valueAdfsWiaEnabled`none`Optional`FALSE`

## Preferred role

The Amazon Resource Name (ARN) of the role to assume. For information about ARN
roles, see [`AssumeRole`](../../../../reference/sts/latest/apireference/api-assumerole.md) in the _AWS Security Token Service API_
_Reference_.

Parameter nameAliasParameter typeDefault valuePreferredRolepreferred\_role (deprecated)Optionalnone

## Role session duration

The duration, in seconds, of the role session. For more information, see [`AssumeRole`](../../../../reference/sts/latest/apireference/api-assumerole.md) in the _AWS Security Token Service API_
_Reference_.

Parameter nameAliasParameter typeDefault valueRoleSessionDurationDuration (deprecated)Optional`3600`

## Lake Formation enabled

Specifies whether to use the [`AssumeDecoratedRoleWithSAML`](../../../../reference/lake-formation/latest/apireference/api-assumedecoratedrolewithsaml.md) Lake Formation API action to
retrieve temporary IAM credentials instead of the [`AssumeRoleWithSAML`](../../../../reference/sts/latest/apireference/api-assumerolewithsaml.md) AWS STS API action.

Parameter nameAliasParameter typeDefault valueLakeFormationEnabled`none`OptionalFALSE

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ping

Browser Azure
AD

All content copied from https://docs.aws.amazon.com/.
