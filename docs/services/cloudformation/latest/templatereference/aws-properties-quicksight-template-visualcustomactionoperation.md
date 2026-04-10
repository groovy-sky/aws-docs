This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template VisualCustomActionOperation

The operation that is defined by the custom action.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilterOperation" : CustomActionFilterOperation,
  "NavigationOperation" : CustomActionNavigationOperation,
  "SetParametersOperation" : CustomActionSetParametersOperation,
  "URLOperation" : CustomActionURLOperation
}

```

### YAML

```yaml

  FilterOperation:
    CustomActionFilterOperation
  NavigationOperation:
    CustomActionNavigationOperation
  SetParametersOperation:
    CustomActionSetParametersOperation
  URLOperation:
    CustomActionURLOperation

```

## Properties

`FilterOperation`

The filter operation that filters data included in a visual or in an entire sheet.

_Required_: No

_Type_: [CustomActionFilterOperation](aws-properties-quicksight-template-customactionfilteroperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NavigationOperation`

The navigation operation that navigates between different sheets in the same analysis.

_Required_: No

_Type_: [CustomActionNavigationOperation](aws-properties-quicksight-template-customactionnavigationoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SetParametersOperation`

The set parameter operation that sets parameters in custom action.

_Required_: No

_Type_: [CustomActionSetParametersOperation](aws-properties-quicksight-template-customactionsetparametersoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`URLOperation`

The URL operation that opens a link to another webpage.

_Required_: No

_Type_: [CustomActionURLOperation](aws-properties-quicksight-template-customactionurloperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisualCustomAction

VisualInteractionOptions

All content copied from https://docs.aws.amazon.com/.
