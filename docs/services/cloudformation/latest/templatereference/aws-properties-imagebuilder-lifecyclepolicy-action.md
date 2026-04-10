This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy Action

Contains selection criteria for the lifecycle policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncludeResources" : IncludeResources,
  "Type" : String
}

```

### YAML

```yaml

  IncludeResources:
    IncludeResources
  Type: String

```

## Properties

`IncludeResources`

Specifies the resources that the lifecycle policy applies to.

_Required_: No

_Type_: [IncludeResources](aws-properties-imagebuilder-lifecyclepolicy-includeresources.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the lifecycle action to take.

_Required_: Yes

_Type_: String

_Allowed values_: `DELETE | DEPRECATE | DISABLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ImageBuilder::LifecyclePolicy

AmiExclusionRules

All content copied from https://docs.aws.amazon.com/.
