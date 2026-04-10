This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::Flywheel EntityRecognitionConfig

Configuration required for an entity recognition model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EntityTypes" : [ EntityTypesListItem, ... ]
}

```

### YAML

```yaml

  EntityTypes:
    - EntityTypesListItem

```

## Properties

`EntityTypes`

Up to 25 entity types that the model is trained to recognize.

_Required_: No

_Type_: Array of [EntityTypesListItem](aws-properties-comprehend-flywheel-entitytypeslistitem.md)

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentClassificationConfig

EntityTypesListItem

All content copied from https://docs.aws.amazon.com/.
