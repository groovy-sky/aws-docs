This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::DirectoryConfig ServiceAccountCredentials

The credentials for the service account used by the streaming instance to connect to the directory.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountName" : String,
  "AccountPassword" : String
}

```

### YAML

```yaml

  AccountName: String
  AccountPassword: String

```

## Properties

`AccountName`

The user name of the account. This account must have the following privileges: create computer objects,
join computers to the domain, and change/reset the password on descendant computer objects for the
organizational units specified.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccountPassword`

The password for the account.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateBasedAuthProperties

AWS::AppStream::Entitlement

All content copied from https://docs.aws.amazon.com/.
