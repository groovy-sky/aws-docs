# Common authentication parameters

The parameters in this section are common to the authentication types as noted.

## Use Proxy for IdP

Enables communication between the driver and the IdP through the proxy. This
option is available for the following authentication plugins:

- AD FS

- Azure AD

- Browser Azure AD

- Browser SSO OIDC

- JWT trusted identity propagation

- JWT

- JWT trusted identity propagation

- Browser trusted identity propagation

- Okta

- Ping

**Connection string name****Parameter type****Default value****Connection string example**UseProxyForIdPOptional`0``UseProxyForIdP=1;`

## Use Lake Formation

Uses the [AssumeDecoratedRoleWithSAML](https://docs.aws.amazon.com/lake-formation/latest/APIReference/API_AssumeDecoratedRoleWithSAML.html) Lake Formation API action to retrieve temporary IAM
credentials instead of the [AssumeRoleWithSAML](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRoleWithSAML.html) AWS STS API action. This option is available for the
Azure AD, Browser Azure AD, Browser SAML, Okta, Ping, and AD FS authentication
plugins.

**Connection string name****Parameter type****Default value****Connection string example**LakeformationEnabledOptional`0``LakeformationEnabled=1;`

## SSL insecure (IdP)

Disables SSL when communicating with the IdP. This option is available for the
Azure AD, Browser Azure AD, Okta, Ping, and AD FS authentication plugins.

###### Important

Breaking change in v2.1.0.0: The default behavior
for SSL certificate validation when connecting to identity providers has changed.
In versions before 2.1.0.0, SSL validation was disabled by default. Starting in
v2.1.0.0, SSL validation is enabled by default for all IdP connections. The driver
also enforces TLS 1.2 as the minimum TLS version. If you use a local identity
provider without a valid SSL certificate (for testing purposes only), set
`SSL_Insecure=1` in your connection string.

**Connection string name****Parameter type****Default value****Connection string example**SSL\_InsecureOptional`0``SSL_Insecure=1;`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Ping

Endpoints
