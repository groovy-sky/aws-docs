# Azure AD credentials

A SAML-based authentication mechanism that enables authentication to Athena using the
Azure AD identity provider. This method assumes that a federation has already been set
up between Athena and Azure AD.

###### Note

Some of the parameter names in this section have aliases. The aliases are
functional equivalents of the parameter names and have been provided for backward
compatibility with the JDBC 2.x driver. Because the parameter names have been
improved to follow a clearer, more consistent naming convention, we recommend that
you use them instead of the aliases, which have been deprecated.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `AzureAD`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProviderAWSCredentialsProviderClass (deprecated)Requirednone`AzureAD`

## User

The email address of the Azure AD user to use for authentication with Azure
AD.

Parameter nameAliasParameter typeDefault valueUserUID (deprecated)Requirednone

## Password

The password for the Azure AD user.

Parameter nameAliasParameter typeDefault valuePasswordPWD (deprecated)Requirednone

## Azure AD tenant ID

The tenant ID of your Azure AD application.

Parameter nameAliasParameter typeDefault valueAzureAdTenantIdtenant\_id (deprecated)Requirednone

## Azure AD client ID

The client ID of your Azure AD application.

Parameter nameAliasParameter typeDefault valueAzureAdClientIdclient\_id (deprecated)Requirednone

## Azure AD client secret

The client secret of your Azure AD application.

Parameter nameAliasParameter typeDefault valueAzureAdClientSecretclient\_secret (deprecated)Requirednone

## Preferred role

The Amazon Resource Name (ARN) of the role to assume. For information about ARN
roles, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) in the
_AWS Security Token Service API Reference_.

Parameter nameAliasParameter typeDefault valuePreferredRolepreferred\_role (deprecated)Optionalnone

## Role session duration

The duration, in seconds, of the role session. For more information, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) in the _AWS Security Token Service API Reference_.

Parameter nameAliasParameter typeDefault valueRoleSessionDurationDuration (deprecated)Optional3600

## Lake Formation enabled

Specifies whether to use the [AssumeDecoratedRoleWithSAML](https://docs.aws.amazon.com/lake-formation/latest/APIReference/API_AssumeDecoratedRoleWithSAML.html) Lake Formation API action to retrieve temporary IAM
credentials instead of the [AssumeRoleWithSAML](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRoleWithSAML.html) AWS STS API action.

Parameter nameAliasParameter typeDefault valueLakeFormationEnablednoneOptionalFALSE

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Browser trusted identity propagation

Okta
