This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::LambdaHook StackFilters

The `StackFilters` property type specifies stack level filters for a
Hook.

The `StackNames` or `StackRoles` properties are optional.
However, you must specify at least one of these properties.

For more information, see [CloudFormation Hooks stack level filters](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-stack-level-filtering.html).

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

_Type_: [StackNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-lambdahook-stacknames.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackRoles`

Includes or excludes specific stacks from Hook invocations based on their associated
IAM roles.

_Required_: No

_Type_: [StackRoles](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-lambdahook-stackroles.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudFormation::LambdaHook

StackNames
