# DirectoryConfig

Describes the configuration information required to join fleets and image builders to Microsoft Active Directory domains.

## Contents

**DirectoryName**

The fully qualified name of the directory (for example, corp.example.com).

Type: String

Required: Yes

**CertificateBasedAuthProperties**

The certificate-based authentication properties used to authenticate SAML 2.0 Identity
Provider (IdP) user identities to Active Directory domain-joined streaming instances.
Fallback is turned on by default when certificate-based authentication is **Enabled** . Fallback allows users to log in using their AD
domain password if certificate-based authentication is unsuccessful, or to unlock a
desktop lock screen. **Enabled\_no\_directory\_login\_fallback** enables certificate-based
authentication, but does not allow users to log in using their AD domain password. Users
will be disconnected to re-authenticate using certificates.

Type: [CertificateBasedAuthProperties](api-certificatebasedauthproperties.md) object

Required: No

**CreatedTime**

The time the directory configuration was created.

Type: Timestamp

Required: No

**OrganizationalUnitDistinguishedNames**

The distinguished names of the organizational units for computer accounts.

Type: Array of strings

Length Constraints: Maximum length of 2000.

Required: No

**ServiceAccountCredentials**

The credentials for the service account used by the fleet or image builder to connect to the directory.

Type: [ServiceAccountCredentials](api-serviceaccountcredentials.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/directoryconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/directoryconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/directoryconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComputeCapacityStatus

DomainJoinInfo

All content copied from https://docs.aws.amazon.com/.
