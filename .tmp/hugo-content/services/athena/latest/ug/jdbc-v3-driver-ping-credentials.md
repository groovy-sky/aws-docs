# Ping credentials

A SAML-based authentication mechanism that enables authentication to Athena using the
Ping Federate identity provider. This method assumes that a federation has already been
set up between Athena and Ping Federate.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `Ping`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProviderAWSCredentialsProviderClass (deprecated)Requirednone`Ping`

## User

The email address of the Ping Federate user to use for authentication with Ping
Federate.

Parameter nameAliasParameter typeDefault valueUserUID (deprecated)Requirednone

## Password

The password for the Ping Federate user.

Parameter nameAliasParameter typeDefault valuePasswordPWD (deprecated)Requirednone

## PingHostName

The address for your Ping server. To find your address, visit the following URL
and view the **SSO Application Endpoint** field.

```nohighlight

https://your-pf-host-#:9999/pingfederate/your-pf-app#/spConnections
```

Parameter nameAliasParameter typeDefault valuePingHostNameIdP\_Host
(deprecated)Requirednone

## PingPortNumber

The port number to use to connect to your IdP host.

Parameter nameAliasParameter typeDefault valuePingPortNumberIdP\_Port (deprecated)Requirednone

## PingPartnerSpId

The service provider address. To find the service provider address, visit the
following URL and view the **SSO Application Endpoint**
field.

```nohighlight

https://your-pf-host-#:9999/pingfederate/your-pf-app#/spConnections
```

Parameter nameAliasParameter typeDefault value

PingPartnerSpId

Partner\_SPID (deprecated)Requirednone

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

Okta

AD FS

All content copied from https://docs.aws.amazon.com/.
