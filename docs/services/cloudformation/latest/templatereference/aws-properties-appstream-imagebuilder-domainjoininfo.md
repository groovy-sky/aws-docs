This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::ImageBuilder DomainJoinInfo

The name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DirectoryName" : String,
  "OrganizationalUnitDistinguishedName" : String
}

```

### YAML

```yaml

  DirectoryName: String
  OrganizationalUnitDistinguishedName: String

```

## Properties

`DirectoryName`

The fully qualified name of the directory (for example, corp.example.com).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationalUnitDistinguishedName`

The distinguished name of the organizational unit for computer accounts.

_Required_: No

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AccessEndpoint

Tag
