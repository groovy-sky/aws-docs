This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentConditionProperty

The `ComponentConditionProperty` property specifies a conditional expression
for setting a component property. Use `ComponentConditionProperty` to set a
property to different values conditionally, based on the value of another property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Else" : ComponentProperty,
  "Field" : String,
  "Operand" : String,
  "OperandType" : String,
  "Operator" : String,
  "Property" : String,
  "Then" : ComponentProperty
}

```

### YAML

```yaml

  Else:
    ComponentProperty
  Field: String
  Operand: String
  OperandType: String
  Operator: String
  Property: String
  Then:
    ComponentProperty

```

## Properties

`Else`

The value to assign to the property if the condition is not met.

_Required_: No

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Field`

The name of a field. Specify this when the property is a data model.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operand`

The value of the property to evaluate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperandType`

The type of the property to evaluate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The operator to use to perform the evaluation, such as `eq` to represent
equals.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Property`

The name of the conditional property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Then`

The value to assign to the property if the condition is met.

_Required_: No

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentChild

ComponentDataConfiguration

All content copied from https://docs.aws.amazon.com/.
