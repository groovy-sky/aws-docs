This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::DBClusterParameterGroup Tag

Metadata assigned to an Amazon Neptune resource consisting of a key-value pair.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

A key is the required name of the tag. The string value can be from 1 to 128 Unicode
characters in length and can't be prefixed with `aws:` or `rds:`.
The string can only contain the set of Unicode letters, digits, white-space,
'\_', '.', '/', '=', '+', '-' (Java regex: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-\]\*)$").

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A value is the optional value of the tag. The string value can be from 1 to 256 Unicode
characters in length and can't be prefixed with `aws:` or `rds:`.
The string can only contain the set of Unicode letters, digits, white-space,
'\_', '.', '/', '=', '+', '-' (Java regex: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-\]\*)$").

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Neptune::DBClusterParameterGroup

AWS::Neptune::DBInstance
