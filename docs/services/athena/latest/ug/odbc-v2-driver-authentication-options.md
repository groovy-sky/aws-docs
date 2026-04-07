# Authentication options

You can connect to Amazon Athena using the following authentication types. For all types, the
connection string name is `AuthenticationType`, the parameter type is
`Required`, and the default value is `IAM Credentials`. For
information about the parameters for each authentication type, visit the corresponding link.
For common authentication parameters, see [Common authentication parameters](odbc-v2-driver-common-authentication-parameters.md).

Authentication typeConnection string example[IAM credentials](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-iam-credentials.html)`AuthenticationType=IAM Credentials;`[IAM profile](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-iam-profile.html)`AuthenticationType=IAM Profile;`[AD FS](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-ad-fs.html)`AuthenticationType=ADFS;`[Azure AD](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-azure-ad.html)`AuthenticationType=AzureAD;`[Browser Azure AD](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-browser-azure-ad.html)`AuthenticationType=BrowserAzureAD;`[Browser SAML](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-browser-saml.html)`AuthenticationType=BrowserSAML;`[Browser SSO OIDC](odbc-v2-driver-browser-sso-oidc.md)`AuthenticationType=BrowserSSOOIDC;`[Default credentials](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-default-credentials.html)`AuthenticationType=Default Credentials;`[External credentials](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-external-credentials.html)`AuthenticationType=External Credentials;`[Instance profile](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-instance-profile.html)`AuthenticationType=Instance Profile;`[JWT](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-jwt.html)`AuthenticationType=JWT;`[JWT Trusted identity propagation credentials provider](odbc-v2-driver-jwt-tip.md)`AuthenticationType=JWT_TIP;`[Browser trusted identity propagation credentials](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-browser-oidc-tip.html)`AuthenticationType=BrowserOidcTip;`[Okta](https://docs.aws.amazon.com/athena/latest/ug/odbc-v2-driver-okta.html)`AuthenticationType=Okta;`[Ping](odbc-v2-driver-ping.md)`AuthenticationType=Ping;`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Main parameters

IAM credentials
