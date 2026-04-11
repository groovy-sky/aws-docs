This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::LambdaHook TargetFiltersItems

Specifies the resource types, actions, and invocation points to target for the
`TargetFilters` property type.

For more information, see [CloudFormation Hook target filters](../../../cloudformation-cli/latest/hooks-userguide/hooks-target-filtering.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ String, ... ],
  "InvocationPoints" : [ String, ... ],
  "TargetNames" : [ String, ... ]
}

```

### YAML

```yaml

  Actions:
    - String
  InvocationPoints:
    - String
  TargetNames:
    - String

```

## Properties

`Actions`

The actions to target.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationPoints`

The invocation points to target.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetNames`

The resource types to target, such as `AWS::S3::Bucket` or
`AWS::DynamoDB::Table`.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetFilters

AWS::CloudFormation::Macro

All content copied from https://docs.aws.amazon.com/.
