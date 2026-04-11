This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration ContainerDistributionConfiguration

Container distribution settings for encryption, licensing, and sharing in a specific
Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerTags" : [ String, ... ],
  "Description" : String,
  "TargetRepository" : TargetContainerRepository
}

```

### YAML

```yaml

  ContainerTags:
    - String
  Description: String
  TargetRepository:
    TargetContainerRepository

```

## Properties

`ContainerTags`

Tags that are attached to the container distribution configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the container distribution configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRepository`

The destination repository for the container distribution configuration.

_Required_: No

_Type_: [TargetContainerRepository](aws-properties-imagebuilder-distributionconfiguration-targetcontainerrepository.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmiDistributionConfiguration

Distribution

All content copied from https://docs.aws.amazon.com/.
