This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentVariant

The `ComponentVariant` property specifies the style configuration of a unique
variation of a main component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Overrides" : {Key: Value, ...},
  "VariantValues" : {Key: Value, ...}
}

```

### YAML

```yaml

  Overrides:
    Key: Value
  VariantValues:
    Key: Value

```

## Properties

`Overrides`

The properties of the component variant that can be overriden when customizing an instance
of the component. You can't specify `tags` as a valid property for
`overrides`.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariantValues`

The combination of variants that comprise this variant.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentPropertyBindingProperties

FormBindingElement

All content copied from https://docs.aws.amazon.com/.
