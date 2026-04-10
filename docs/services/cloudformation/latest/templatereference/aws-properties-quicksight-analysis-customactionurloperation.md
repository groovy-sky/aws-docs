This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis CustomActionURLOperation

The URL operation that opens a link to another webpage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "URLTarget" : String,
  "URLTemplate" : String
}

```

### YAML

```yaml

  URLTarget: String
  URLTemplate: String

```

## Properties

`URLTarget`

The target of the `CustomActionURLOperation`.

Valid values are defined as follows:

- `NEW_TAB`: Opens the target URL in a new browser tab.

- `NEW_WINDOW`: Opens the target URL in a new browser window.

- `SAME_TAB`: Opens the target URL in the same browser tab.

_Required_: Yes

_Type_: String

_Allowed values_: `NEW_TAB | NEW_WINDOW | SAME_TAB`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`URLTemplate`

THe URL link of the `CustomActionURLOperation`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomActionSetParametersOperation

CustomColor

All content copied from https://docs.aws.amazon.com/.
