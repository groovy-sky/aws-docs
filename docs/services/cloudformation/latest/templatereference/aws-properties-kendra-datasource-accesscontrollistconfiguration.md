This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource AccessControlListConfiguration

Specifies access control list files for the documents in a data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyPath" : String
}

```

### YAML

```yaml

  KeyPath: String

```

## Properties

`KeyPath`

Path to the AWS S3 bucket that contains the access control list
files.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Kendra::DataSource

AclConfiguration

All content copied from https://docs.aws.amazon.com/.
