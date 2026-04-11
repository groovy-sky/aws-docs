This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::GuardHook StackRoles

Specifies the stack roles for the `StackFilters` property type to include
or exclude specific stacks from Hook invocations based on their associated IAM
roles.

For more information, see [CloudFormation Hooks stack level filters](../../../cloudformation-cli/latest/hooks-userguide/hooks-stack-level-filtering.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exclude" : [ String, ... ],
  "Include" : [ String, ... ]
}

```

### YAML

```yaml

  Exclude:
    - String
  Include:
    - String

```

## Properties

`Exclude`

The IAM role ARNs for stacks you want to exclude. The Hook will be invoked on all
stacks except those initiated by the specified roles.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Include`

The IAM role ARNs to target stacks associated with these roles. Only stack operations
initiated by these roles will invoke the Hook.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StackNames

TargetFilters

All content copied from https://docs.aws.amazon.com/.
