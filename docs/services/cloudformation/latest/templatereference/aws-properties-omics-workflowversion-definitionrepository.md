This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::WorkflowVersion DefinitionRepository

Contains information about a source code repository that hosts the workflow definition files.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "connectionArn" : String,
  "excludeFilePatterns" : [ String, ... ],
  "fullRepositoryId" : String,
  "sourceReference" : SourceReference
}

```

### YAML

```yaml

  connectionArn: String
  excludeFilePatterns:
    - String
  fullRepositoryId: String
  sourceReference:
    SourceReference

```

## Properties

`connectionArn`

The Amazon Resource Name (ARN) of the connection to the source code repository.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[\\w]+)*:.+:.+:[0-9]{12}:.+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`excludeFilePatterns`

A list of file patterns to exclude when retrieving the workflow definition from the repository.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`fullRepositoryId`

The full repository identifier, including the repository owner and name. For example, 'repository-owner/repository-name'.

_Required_: No

_Type_: String

_Pattern_: `.+/.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`sourceReference`

The source reference for the repository, such as a branch name, tag, or commit ID.

_Required_: No

_Type_: [SourceReference](aws-properties-omics-workflowversion-sourcereference.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerRegistryMap

ImageMapping

All content copied from https://docs.aws.amazon.com/.
