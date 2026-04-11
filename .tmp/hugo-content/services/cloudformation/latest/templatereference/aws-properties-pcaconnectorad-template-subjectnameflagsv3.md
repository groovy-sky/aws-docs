This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template SubjectNameFlagsV3

Information to include in the subject name and alternate subject name of the
certificate. The subject name can be common name, directory path, DNS as common name, or
left blank. You can optionally include email to the subject name for user templates. If you
leave the subject name blank then you must set a subject alternate name. The subject
alternate name (SAN) can include globally unique identifier (GUID), DNS, domain DNS, email,
service principal name (SPN), and user principal name (UPN). You can leave the SAN blank.
If you leave the SAN blank, then you must set a subject name.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RequireCommonName" : Boolean,
  "RequireDirectoryPath" : Boolean,
  "RequireDnsAsCn" : Boolean,
  "RequireEmail" : Boolean,
  "SanRequireDirectoryGuid" : Boolean,
  "SanRequireDns" : Boolean,
  "SanRequireDomainDns" : Boolean,
  "SanRequireEmail" : Boolean,
  "SanRequireSpn" : Boolean,
  "SanRequireUpn" : Boolean
}

```

### YAML

```yaml

  RequireCommonName: Boolean
  RequireDirectoryPath: Boolean
  RequireDnsAsCn: Boolean
  RequireEmail: Boolean
  SanRequireDirectoryGuid: Boolean
  SanRequireDns: Boolean
  SanRequireDomainDns: Boolean
  SanRequireEmail: Boolean
  SanRequireSpn: Boolean
  SanRequireUpn: Boolean

```

## Properties

`RequireCommonName`

Include the common name in the subject name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireDirectoryPath`

Include the directory path in the subject name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireDnsAsCn`

Include the DNS as common name in the subject name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireEmail`

Include the subject's email in the subject name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SanRequireDirectoryGuid`

Include the globally unique identifier (GUID) in the subject alternate name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SanRequireDns`

Include the DNS in the subject alternate name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SanRequireDomainDns`

Include the domain DNS in the subject alternate name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SanRequireEmail`

Include the subject's email in the subject alternate name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SanRequireSpn`

Include the service principal name (SPN) in the subject alternate name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SanRequireUpn`

Include the user principal name (UPN) in the subject alternate name.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubjectNameFlagsV2

SubjectNameFlagsV4

All content copied from https://docs.aws.amazon.com/.
