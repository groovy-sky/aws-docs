This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::Document DocumentRequires

An SSM document required by the current document.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Version" : String
}

```

### YAML

```yaml

  Name: String
  Version: String

```

## Properties

`Name`

The name of the required SSM document. The name can be an Amazon Resource Name (ARN).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.:/]{3,200}$`

_Maximum_: `200`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Version`

The document version required by the current document.

_Required_: No

_Type_: String

_Pattern_: `([$]LATEST|[$]DEFAULT|^[1-9][0-9]*$)`

_Maximum_: `8`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttachmentsSource

Tag
