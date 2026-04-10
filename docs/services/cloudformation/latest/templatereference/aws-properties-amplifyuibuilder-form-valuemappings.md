This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form ValueMappings

The `ValueMappings` property specifies the data binding configuration for a value map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BindingProperties" : {Key: Value, ...},
  "Values" : [ ValueMapping, ... ]
}

```

### YAML

```yaml

  BindingProperties:
    Key: Value
  Values:
    - ValueMapping

```

## Properties

`BindingProperties`

The information to bind fields to data at runtime.

_Required_: No

_Type_: Object of [FormInputBindingPropertiesValue](aws-properties-amplifyuibuilder-form-forminputbindingpropertiesvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value and display value pairs.

_Required_: Yes

_Type_: Array of [ValueMapping](aws-properties-amplifyuibuilder-form-valuemapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValueMapping

AWS::AmplifyUIBuilder::Theme

All content copied from https://docs.aws.amazon.com/.
