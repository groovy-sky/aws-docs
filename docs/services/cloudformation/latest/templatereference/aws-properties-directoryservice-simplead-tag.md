This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectoryService::SimpleAD Tag

Metadata assigned to a directory consisting of a key-value pair.

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

Required name of the tag. The string value can be Unicode characters and cannot be
prefixed with "aws:". The string can contain only the set of Unicode letters, digits,
white-space, '\_', '.', '/', '=', '+', '-', ':', '@'(Java regex:
"^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-\]\*)$").

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The optional value of the tag. The string value can be Unicode characters. The string
can contain only the set of Unicode letters, digits, white-space, '\_', '.', '/', '=', '+', '-', ':', '@'
(Java regex: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-\]\*)$").

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DirectoryService::SimpleAD

VpcSettings
