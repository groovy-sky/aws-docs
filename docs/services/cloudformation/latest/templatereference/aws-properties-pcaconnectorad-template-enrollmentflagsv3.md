This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template EnrollmentFlagsV3

Template configurations for v3 template schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableKeyReuseOnNtTokenKeysetStorageFull" : Boolean,
  "IncludeSymmetricAlgorithms" : Boolean,
  "NoSecurityExtension" : Boolean,
  "RemoveInvalidCertificateFromPersonalStore" : Boolean,
  "UserInteractionRequired" : Boolean
}

```

### YAML

```yaml

  EnableKeyReuseOnNtTokenKeysetStorageFull: Boolean
  IncludeSymmetricAlgorithms: Boolean
  NoSecurityExtension: Boolean
  RemoveInvalidCertificateFromPersonalStore: Boolean
  UserInteractionRequired: Boolean

```

## Properties

`EnableKeyReuseOnNtTokenKeysetStorageFull`

Allow renewal using the same key.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeSymmetricAlgorithms`

Include symmetric algorithms allowed by the subject.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoSecurityExtension`

This flag instructs the CA to not include the security extension
szOID\_NTDS\_CA\_SECURITY\_EXT (OID:1.3.6.1.4.1.311.25.2), as specified in \[MS-WCCE\] sections
2.2.2.7.7.4 and 3.2.2.6.2.1.4.5.9, in the issued certificate. This addresses a Windows
Kerberos elevation-of-privilege vulnerability.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveInvalidCertificateFromPersonalStore`

Delete expired or revoked certificates instead of archiving them.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserInteractionRequired`

Require user interaction when the subject is enrolled and the private key associated
with the certificate is used.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnrollmentFlagsV2

EnrollmentFlagsV4

All content copied from https://docs.aws.amazon.com/.
