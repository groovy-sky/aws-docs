This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::DirectoryConfig

The `AWS::AppStream::DirectoryConfig` resource specifies the configuration information required to join Amazon WorkSpaces Applications fleets
and image builders to Microsoft Active Directory domains.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::DirectoryConfig",
  "Properties" : {
      "CertificateBasedAuthProperties" : CertificateBasedAuthProperties,
      "DirectoryName" : String,
      "OrganizationalUnitDistinguishedNames" : [ String, ... ],
      "ServiceAccountCredentials" : ServiceAccountCredentials
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::DirectoryConfig
Properties:
  CertificateBasedAuthProperties:
    CertificateBasedAuthProperties
  DirectoryName: String
  OrganizationalUnitDistinguishedNames:
    - String
  ServiceAccountCredentials:
    ServiceAccountCredentials

```

## Properties

`CertificateBasedAuthProperties`

The certificate-based authentication properties used to authenticate SAML 2.0 Identity
Provider (IdP) user identities to Active Directory domain-joined streaming instances.

_Required_: No

_Type_: [CertificateBasedAuthProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-directoryconfig-certificatebasedauthproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DirectoryName`

The fully qualified name of the directory (for example, corp.example.com).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OrganizationalUnitDistinguishedNames`

The distinguished names of the organizational units for computer accounts.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccountCredentials`

The credentials for the service account used by the streaming instance to connect to
the directory. Do not use this parameter directly. Use `ServiceAccountCredentials` as an input parameter with `noEcho` as shown in
the [Parameters](../userguide/parameters-section-structure.md). For best practices information, see [Do Not Embed Credentials in Your Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/best-practices.html#creds).

_Required_: Yes

_Type_: [ServiceAccountCredentials](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CreateDirectoryConfig](https://docs.aws.amazon.com/appstream2/latest/APIReference/API_CreateDirectoryConfig.html) in the _Amazon WorkSpaces Applications API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppStream::ApplicationFleetAssociation

CertificateBasedAuthProperties
