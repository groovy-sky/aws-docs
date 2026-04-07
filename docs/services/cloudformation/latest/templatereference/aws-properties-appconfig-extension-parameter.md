This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::Extension Parameter

A value such as an Amazon Resource Name (ARN) or an Amazon Simple Notification Service topic entered
in an extension when invoked. Parameter values are specified in an extension association.
For more information about extensions, see [Extending\
workflows](../../../appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Dynamic" : Boolean,
  "Required" : Boolean
}

```

### YAML

```yaml

  Description: String
  Dynamic: Boolean
  Required: Boolean

```

## Properties

`Description`

Information about the parameter.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dynamic`

Indicates whether this parameter's value can be supplied at the extension's action point
instead of during extension association. Dynamic parameters can't be marked
`Required`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Required`

A parameter value must be specified in the extension association.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Action

Tag
