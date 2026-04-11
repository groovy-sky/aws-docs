This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::LambdaHook TargetFilters

The `TargetFilters` property type specifies the target filters for the
Hook.

For more information, see [CloudFormation Hook target filters](../../../cloudformation-cli/latest/hooks-userguide/hooks-target-filtering.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetFiltersItems" : TargetFiltersItems
}

```

### YAML

```yaml

  TargetFiltersItems:
    TargetFiltersItems

```

## Properties

`TargetFiltersItems`

The specific resource types, actions, and invocation points to target.

_Required_: No

_Type_: [TargetFiltersItems](aws-properties-cloudformation-lambdahook-targetfiltersitems.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StackRoles

TargetFiltersItems

All content copied from https://docs.aws.amazon.com/.
