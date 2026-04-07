This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentTemplateTarget

Specifies a target for an experiment. You must specify at least one Amazon Resource Name (ARN) or
at least one resource tag. You cannot specify both ARNs and tags.

For more information, see [Targets](https://docs.aws.amazon.com/fis/latest/userguide/targets.html)
in the _AWS Fault Injection Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filters" : [ ExperimentTemplateTargetFilter, ... ],
  "Parameters" : {Key: Value, ...},
  "ResourceArns" : [ String, ... ],
  "ResourceTags" : {Key: Value, ...},
  "ResourceType" : String,
  "SelectionMode" : String
}

```

### YAML

```yaml

  Filters:
    - ExperimentTemplateTargetFilter
  Parameters:
    Key: Value
  ResourceArns:
    - String
  ResourceTags:
    Key: Value
  ResourceType: String
  SelectionMode: String

```

## Properties

`Filters`

The filters to apply to identify target resources using specific attributes.

_Required_: No

_Type_: Array of [ExperimentTemplateTargetFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fis-experimenttemplate-experimenttemplatetargetfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameters for the resource type.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,64}`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArns`

The Amazon Resource Names (ARNs) of the targets.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

The tags for the target resources.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,128}`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The resource type.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectionMode`

Scopes the identified resources to a specific count or percentage.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExperimentTemplateStopCondition

ExperimentTemplateTargetFilter
