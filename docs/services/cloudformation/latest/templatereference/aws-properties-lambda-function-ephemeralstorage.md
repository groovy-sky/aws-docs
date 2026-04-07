This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function EphemeralStorage

The size of the function's `/tmp` directory in MB. The default value is 512,
but it can be any whole number between 512 and 10,240 MB.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Size" : Integer
}

```

### YAML

```yaml

  Size: Integer

```

## Properties

`Size`

The size of the function's `/tmp` directory.

_Required_: Yes

_Type_: Integer

_Minimum_: `512`

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Environment

FileSystemConfig
