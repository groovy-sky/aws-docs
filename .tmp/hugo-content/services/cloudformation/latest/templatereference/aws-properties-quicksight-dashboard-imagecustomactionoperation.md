This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ImageCustomActionOperation

The operation that is defined by the custom action.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NavigationOperation" : CustomActionNavigationOperation,
  "SetParametersOperation" : CustomActionSetParametersOperation,
  "URLOperation" : CustomActionURLOperation
}

```

### YAML

```yaml

  NavigationOperation:
    CustomActionNavigationOperation
  SetParametersOperation:
    CustomActionSetParametersOperation
  URLOperation:
    CustomActionURLOperation

```

## Properties

`NavigationOperation`

Property description not available.

_Required_: No

_Type_: [CustomActionNavigationOperation](aws-properties-quicksight-dashboard-customactionnavigationoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SetParametersOperation`

Property description not available.

_Required_: No

_Type_: [CustomActionSetParametersOperation](aws-properties-quicksight-dashboard-customactionsetparametersoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`URLOperation`

Property description not available.

_Required_: No

_Type_: [CustomActionURLOperation](aws-properties-quicksight-dashboard-customactionurloperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageCustomAction

ImageInteractionOptions

All content copied from https://docs.aws.amazon.com/.
