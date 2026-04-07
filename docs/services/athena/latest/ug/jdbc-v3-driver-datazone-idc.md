# DataZone IdC Credentials Provider

An authentication mechanism that enables connecting to DataZone-governed data in Athena
using IAM Identity Center.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `DataZoneIdc`. Note that the
`AWSCredentialsProviderClass` alias is deprecated; use the
`CredentialsProvider` parameter name instead.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProviderAWSCredentialsProviderClass (deprecated)RequirednoneDataZoneIdc

## DataZone domain identifier

Identifier of the DataZone domain to use.

Parameter nameAliasParameter typeDefault valueDataZoneDomainIdnoneRequirednone

## DataZone environment identifier

Identifier of the DataZone environment to use.

Parameter nameAliasParameter typeDefault valueDataZoneEnvironmentIdnoneRequirednone

## DataZone domain region

The AWS Region where your DataZone domain is provisioned.

Parameter nameAliasParameter typeDefault valueDataZoneDomainRegionnoneRequirednone

## Region

The AWS Region where your DataZone environment and Athena workgroup are
provisioned.

Parameter nameAliasParameter typeDefault valueRegionnoneRequirednone

## IAM Identity Center issuer URL

The issuer URL of the IAM Identity Center instance that the DataZone domain uses.

Parameter nameAliasParameter typeDefault valueIdentityCenterIssuerUrlnoneRequirednone

## DataZone endpoint override

The DataZone API endpoint to be used instead of the default for the provided
AWS Region.

Parameter nameAliasParameter typeDefault valueDataZoneEndpointOverridenoneOptionalnone

## Enable token caching

When enabled, allows the same IAM Identity Center access token to be used across driver
connections. This prevents SQL tools that create multiple driver connections from
launching multiple browser windows. If you enable this parameter, we recommend that
you close the SQL tool immediately after using it to clear the token cache and
require re-authentication.

Parameter nameAliasParameter typeDefault valueEnableTokenCachingnoneOptionalFALSE

## Listen port

The port number that listens for the IAM Identity Center response.

Parameter nameAliasParameter typeDefault valueListenPortnoneOptional8000

## Identity provider response time out

The duration, in seconds, before the driver stops waiting for the response from
IAM Identity Center.

Parameter nameAliasParameter typeDefault valueIdpResponseTimeoutnoneOptional120

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Browser
SAML

DataZone IAM
