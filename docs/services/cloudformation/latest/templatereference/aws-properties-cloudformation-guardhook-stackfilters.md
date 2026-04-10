This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::GuardHook StackFilters

The `StackFilters` property type specifies stack level filters for a
Hook.

The `StackNames` or `StackRoles` properties are optional.
However, you must specify at least one of these properties.

For more information, see [CloudFormation Hooks stack level filters](../../../cloudformation-cli/latest/hooks-userguide/hooks-stack-level-filtering.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilteringCriteria" : String,
  "StackNames" : StackNames,
  "StackRoles" : StackRoles
}

```

### YAML

```yaml

  FilteringCriteria: String
  StackNames:
    StackNames
  StackRoles:
    StackRoles

```

## Properties

`FilteringCriteria`

The filtering criteria.

- All stack names and stack roles ( `All`): The Hook will only be
invoked when all specified filters match.

- Any stack names and stack roles ( `Any`): The Hook will be invoked
if at least one of the specified filters match.

_Required_: Yes

_Type_: String

_Allowed values_: `ALL | ANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackNames`

Includes or excludes specific stacks from Hook invocations.

_Required_: No

_Type_: [StackNames](aws-properties-cloudformation-guardhook-stacknames.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackRoles`

Includes or excludes specific stacks from Hook invocations based on their associated
IAM roles.

_Required_: No

_Type_: [StackRoles](aws-properties-cloudformation-guardhook-stackroles.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Location

StackNames

All content copied from https://docs.aws.amazon.com/.
