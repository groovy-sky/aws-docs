This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ExcludedRule

Specifies a single rule in a rule group whose action you want to override to `Count`.

###### Note

Instead of this option, use `RuleActionOverrides`. It accepts any valid action setting, including `Count`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

The name of the rule whose action you want to override to `Count`.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultAction

FieldIdentifier

All content copied from https://docs.aws.amazon.com/.
