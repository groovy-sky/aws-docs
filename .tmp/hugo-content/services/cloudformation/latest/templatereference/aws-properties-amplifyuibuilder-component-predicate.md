This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component Predicate

The `Predicate` property specifies information for generating Amplify DataStore queries. Use `Predicate` to retrieve a subset of the
data in a collection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "And" : [ Predicate, ... ],
  "Field" : String,
  "Operand" : String,
  "OperandType" : String,
  "Operator" : String,
  "Or" : [ Predicate, ... ]
}

```

### YAML

```yaml

  And:
    - Predicate
  Field: String
  Operand: String
  OperandType: String
  Operator: String
  Or:
    - Predicate

```

## Properties

`And`

A list of predicates to combine logically.

_Required_: No

_Type_: Array of [Predicate](aws-properties-amplifyuibuilder-component-predicate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Field`

The field to query.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operand`

The value to use when performing the evaluation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperandType`

The type of value to use when performing the evaluation.

_Required_: No

_Type_: String

_Pattern_: `^boolean|string|number$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The operator to use to perform the evaluation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Or`

A list of predicates to combine logically.

_Required_: No

_Type_: Array of [Predicate](aws-properties-amplifyuibuilder-component-predicate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MutationActionSetStateParameter

SortProperty

All content copied from https://docs.aws.amazon.com/.
