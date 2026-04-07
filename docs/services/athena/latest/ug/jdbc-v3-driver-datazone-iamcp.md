# DataZone IAM Credentials Provider

An authentication mechanism that uses IAM credentials to connect to
DataZone-governed data in Athena.

## DataZone domain identifier

Identifier of the DataZone domain to use.

Parameter nameAliasParameter typeDefault valueDataZoneDomainIdnoneRequirednone

## DataZone environment identifier

Identifier of the DataZone environment to use.

Parameter nameAliasParameter typeDefault valueDataZoneEnvironmentIdnoneRequirednone

## DataZone domain region

The AWS Region where your DataZone domain is provisioned.

Parameter nameAliasParameter typeDefault valueDataZoneDomainRegionnoneRequirednone

## DataZone endpoint override

The DataZone API endpoint to use instead of the endpoint default for the provided
AWS Region.

Parameter nameAliasParameter typeDefault valueDataZoneEndpointOverridenoneOptionalnone

## User

Your AWS access key ID. For more information about access keys, see [AWS\
security credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/security-creds.html) in the _IAM User_
_Guide_.

Parameter nameAliasParameter typeDefault valueUserAccessKeyIdOptionalnone

## Password

Your AWS secret key ID. For more information about access keys, see [AWS\
security credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/security-creds.html) in the _IAM User_
_Guide_.

Parameter nameAliasParameter typeDefault valuePasswordSecretAccessKeyOptionalnone

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataZone IdC

Other JDBC 3.x configuration
