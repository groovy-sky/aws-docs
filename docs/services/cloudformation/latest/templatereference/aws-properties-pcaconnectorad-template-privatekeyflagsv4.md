This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template PrivateKeyFlagsV4

Private key flags for v4 templates specify the client compatibility, if the private key
can be exported, if user input is required when using a private key, if an alternate
signature algorithm should be used, and if certificates are renewed using the same private
key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientVersion" : String,
  "ExportableKey" : Boolean,
  "RequireAlternateSignatureAlgorithm" : Boolean,
  "RequireSameKeyRenewal" : Boolean,
  "StrongKeyProtectionRequired" : Boolean,
  "UseLegacyProvider" : Boolean
}

```

### YAML

```yaml

  ClientVersion: String
  ExportableKey: Boolean
  RequireAlternateSignatureAlgorithm: Boolean
  RequireSameKeyRenewal: Boolean
  StrongKeyProtectionRequired: Boolean
  UseLegacyProvider: Boolean

```

## Properties

`ClientVersion`

Defines the minimum client compatibility.

_Required_: Yes

_Type_: String

_Allowed values_: `WINDOWS_SERVER_2012 | WINDOWS_SERVER_2012_R2 | WINDOWS_SERVER_2016`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportableKey`

Allows the private key to be exported.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireAlternateSignatureAlgorithm`

Requires the PKCS #1 v2.1 signature format for certificates. You should verify that your
CA, objects, and applications can accept this signature format.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireSameKeyRenewal`

Renew certificate using the same private key.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrongKeyProtectionRequired`

Require user input when using the private key for enrollment.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseLegacyProvider`

Specifies the cryptographic service provider category used to generate private keys. Set
to TRUE to use Legacy Cryptographic Service Providers and FALSE to use Key Storage
Providers.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrivateKeyFlagsV3

SubjectNameFlagsV2

All content copied from https://docs.aws.amazon.com/.
