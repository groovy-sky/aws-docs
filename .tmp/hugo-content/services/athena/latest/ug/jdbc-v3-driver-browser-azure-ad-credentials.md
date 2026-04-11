# Browser Azure AD credentials

Browser Azure AD is a SAML-based authentication mechanism that works with the Azure AD
identity provider and supports multi-factor authentication. Unlike the standard Azure AD
authentication mechanism, this mechanism does not require a user name, password, or
client secret in the connection parameters. Like the standard Azure AD authentication
mechanism, Browser Azure AD also assumes the user has already set up federation between
Athena and Azure AD.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `BrowserAzureAD`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProviderAWSCredentialsProviderClass (deprecated)Requirednone`BrowserAzureAD`

## Azure AD tenant ID

The tenant ID of your Azure AD application

Parameter nameAliasParameter typeDefault valueAzureAdTenantIdtenant\_id (deprecated)Requirednone

## Azure AD client ID

The client ID of your Azure AD application

Parameter nameAliasParameter typeDefault valueAzureAdClientIdclient\_id (deprecated)Requirednone

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

AD FS

Browser
SAML

All content copied from https://docs.aws.amazon.com/.
