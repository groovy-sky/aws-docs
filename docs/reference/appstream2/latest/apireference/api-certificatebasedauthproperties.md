# CertificateBasedAuthProperties

The certificate-based authentication properties used to authenticate SAML 2.0 Identity
Provider (IdP) user identities to Active Directory domain-joined streaming instances.
Fallback is turned on by default when certificate-based authentication is **Enabled** . Fallback allows users to log in using their AD
domain password if certificate-based authentication is unsuccessful, or to unlock a
desktop lock screen. **Enabled\_no\_directory\_login\_fallback** enables certificate-based
authentication, but does not allow users to log in using their AD domain password. Users
will be disconnected to re-authenticate using certificates.

## Contents

**CertificateAuthorityArn**

The ARN of the AWS Certificate Manager Private CA resource.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**Status**

The status of the certificate-based authentication properties.

Type: String

Valid Values: `DISABLED | ENABLED | ENABLED_NO_DIRECTORY_LOGIN_FALLBACK`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/certificatebasedauthproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/certificatebasedauthproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/certificatebasedauthproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationSettingsResponse

ComputeCapacity

All content copied from https://docs.aws.amazon.com/.
