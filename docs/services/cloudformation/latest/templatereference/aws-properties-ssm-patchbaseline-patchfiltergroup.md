This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::PatchBaseline PatchFilterGroup

The `PatchFilterGroup` property type specifies a set of patch filters for
an AWS Systems Manager patch baseline, typically used for approval rules for a
Systems Manager patch baseline.

`PatchFilterGroup` is the property type for the `GlobalFilters`
property of the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource and the `PatchFilterGroup`
property of the [Rule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PatchFilters" : [ PatchFilter, ... ]
}

```

### YAML

```yaml

  PatchFilters:
    - PatchFilter

```

## Properties

`PatchFilters`

The set of patch filters that make up the group.

_Required_: No

_Type_: Array of [PatchFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-patchbaseline-patchfilter.html)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PatchFilter

PatchSource
