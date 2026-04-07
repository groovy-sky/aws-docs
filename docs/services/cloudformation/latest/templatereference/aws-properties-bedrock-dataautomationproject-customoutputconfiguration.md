This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject CustomOutputConfiguration

Blueprints to apply to objects processed by the project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Blueprints" : [ BlueprintItem, ... ]
}

```

### YAML

```yaml

  Blueprints:
    - BlueprintItem

```

## Properties

`Blueprints`

A list of blueprints.

_Required_: No

_Type_: Array of [BlueprintItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-dataautomationproject-blueprintitem.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ChannelLabelingConfiguration

DocumentBoundingBox
