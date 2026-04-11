This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::Workflow ImageMapping

Specifies image mappings that workflow tasks can use.
For example, you can replace all the task references of a public image to use
an equivalent image in your private ECR repository.
You can use image mappings with upstream registries that don't support pull through cache.
You need to manually synchronize the upstream registry with your private repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationImage" : String,
  "SourceImage" : String
}

```

### YAML

```yaml

  DestinationImage: String
  SourceImage: String

```

## Properties

`DestinationImage`

Specifies the URI of the corresponding image in the private ECR registry.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `750`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceImage`

Specifies the URI of the source image in the upstream registry.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `750`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefinitionRepository

RegistryMapping

All content copied from https://docs.aws.amazon.com/.
