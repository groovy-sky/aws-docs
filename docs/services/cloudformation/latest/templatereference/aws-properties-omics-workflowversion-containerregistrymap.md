This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::WorkflowVersion ContainerRegistryMap

Use a container registry map to specify mappings between the ECR
private repository and one or more upstream registries.
For more information, see [Container images](../../../omics/latest/dev/workflows-ecr.md) in the _AWS HealthOmics User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageMappings" : [ ImageMapping, ... ],
  "RegistryMappings" : [ RegistryMapping, ... ]
}

```

### YAML

```yaml

  ImageMappings:
    - ImageMapping
  RegistryMappings:
    - RegistryMapping

```

## Properties

`ImageMappings`

Image mappings specify path mappings between the ECR
private repository and their corresponding external repositories.

_Required_: No

_Type_: Array of [ImageMapping](aws-properties-omics-workflowversion-imagemapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RegistryMappings`

Mapping that provides the ECR repository path where upstream container images are pulled and synchronized.

_Required_: No

_Type_: Array of [RegistryMapping](aws-properties-omics-workflowversion-registrymapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Omics::WorkflowVersion

DefinitionRepository

All content copied from https://docs.aws.amazon.com/.
