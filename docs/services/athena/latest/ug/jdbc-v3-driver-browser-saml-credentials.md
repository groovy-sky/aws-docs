# Browser SAML credentials

Browser SAML is a generic authentication plugin that can work with SAML-based identity
providers and supports multi-factor authentication.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `BrowserSaml`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProviderAWSCredentialsProviderClass (deprecated)Requirednone`BrowserSaml`

## Single sign-on login URL

The single sign-on URL for your application on the SAML-based identity
provider.

Parameter nameAliasParameter typeDefault valueSsoLoginUrllogin\_url (deprecated)Requirednone

## Listen port

The port number that is used to listen for the SAML response. This value should
match the URL with which you configured the SAML-based identity provider (for
example, `http://localhost:7890/athena`).

Parameter nameAliasParameter typeDefault valueListenPortlisten\_port (deprecated)Optional`7890`

## Identity provider response timeout

The duration, in seconds, before the driver stops waiting for the SAML response
from Azure AD.

Parameter nameAliasParameter typeDefault valueIdpResponseTimeoutidp\_response\_timeout (deprecated)Optional120

## Preferred role

The Amazon Resource Name (ARN) of the role to assume. For information about ARN
roles, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the
_AWS Security Token Service API Reference_.

Parameter nameAliasParameter typeDefault valuePreferredRolepreferred\_role (deprecated)Optionalnone

## Role session duration

The duration, in seconds, of the role session. For more information, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the _AWS Security Token Service API Reference_.

Parameter nameAliasParameter typeDefault valueRoleSessionDurationDuration (deprecated)Optional3600

## Lake Formation enabled

Specifies whether to use the [AssumeDecoratedRoleWithSAML](../../../../reference/lake-formation/latest/apireference/api-assumedecoratedrolewithsaml.md) Lake Formation API action to retrieve temporary IAM
credentials instead of the [AssumeRoleWithSAML](../../../../reference/sts/latest/apireference/api-assumerolewithsaml.md) AWS STS API action.

Parameter nameAliasParameter typeDefault valueLakeFormationEnablednoneOptionalFALSE

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Browser Azure
AD

DataZone IdC

All content copied from https://docs.aws.amazon.com/.
