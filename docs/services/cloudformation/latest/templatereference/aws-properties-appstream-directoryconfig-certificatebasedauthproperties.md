This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::DirectoryConfig CertificateBasedAuthProperties

The certificate-based authentication properties used to authenticate SAML 2.0 Identity
Provider (IdP) user identities to Active Directory domain-joined streaming instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateAuthorityArn" : String,
  "Status" : String
}

```

### YAML

```yaml

  CertificateAuthorityArn: String
  Status: String

```

## Properties

`CertificateAuthorityArn`

The ARN of the AWS Certificate Manager Private CA resource.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the certificate-based authentication properties. Fallback is turned on by default when certificate-based authentication is **Enabled**. Fallback allows users to log in using their AD
domain password if certificate-based authentication is unsuccessful, or to unlock a
desktop lock screen. **Enabled\_no\_directory\_login\_fallback** enables certificate-based
authentication, but does not allow users to log in using their AD domain password. Users
will be disconnected to re-authenticate using certificates.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED | ENABLED_NO_DIRECTORY_LOGIN_FALLBACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppStream::DirectoryConfig

ServiceAccountCredentials
