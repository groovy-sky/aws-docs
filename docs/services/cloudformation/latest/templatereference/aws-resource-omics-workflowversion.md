This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::WorkflowVersion

Creates a new workflow version for the workflow that you specify with the `workflowId` parameter.

When you create a new version of a workflow, you need to specify the configuration
for the new version. It doesn't inherit any configuration values from the workflow.

Provide a version name that is unique for
this workflow. You cannot change the name after HealthOmics creates the version.

###### Note

Don't include any personally identifiable information (PII) in the version name. Version names appear in
the workflow version ARN.

For more information, see [Workflow versioning in AWS HealthOmics](../../../omics/latest/dev/workflow-versions.md) in the _AWS HealthOmics User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Omics::WorkflowVersion",
  "Properties" : {
      "Accelerators" : String,
      "ContainerRegistryMap" : ContainerRegistryMap,
      "ContainerRegistryMapUri" : String,
      "DefinitionRepository" : DefinitionRepository,
      "DefinitionUri" : String,
      "Description" : String,
      "Engine" : String,
      "Main" : String,
      "ParameterTemplate" : {Key: Value, ...},
      "ParameterTemplatePath" : String,
      "readmeMarkdown" : String,
      "readmePath" : String,
      "readmeUri" : String,
      "StorageCapacity" : Number,
      "StorageType" : String,
      "Tags" : {Key: Value, ...},
      "VersionName" : String,
      "WorkflowBucketOwnerId" : String,
      "WorkflowId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Omics::WorkflowVersion
Properties:
  Accelerators: String
  ContainerRegistryMap:
    ContainerRegistryMap
  ContainerRegistryMapUri: String
  DefinitionRepository:
    DefinitionRepository
  DefinitionUri: String
  Description: String
  Engine: String
  Main: String
  ParameterTemplate:
    Key: Value
  ParameterTemplatePath: String
  readmeMarkdown: String
  readmePath: String
  readmeUri: String
  StorageCapacity: Number
  StorageType: String
  Tags:
    Key: Value
  VersionName: String
  WorkflowBucketOwnerId: String
  WorkflowId: String

```

## Properties

`Accelerators`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `GPU`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContainerRegistryMap`

Use a container registry map to specify mappings between the ECR
private repository and one or more upstream registries.
For more information, see [Container images](../../../omics/latest/dev/workflows-ecr.md) in the _AWS HealthOmics User Guide_.

_Required_: No

_Type_: [ContainerRegistryMap](aws-properties-omics-workflowversion-containerregistrymap.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContainerRegistryMapUri`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `750`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefinitionRepository`

Contains information about a source code repository that hosts the workflow definition files.

_Required_: No

_Type_: [DefinitionRepository](aws-properties-omics-workflowversion-definitionrepository.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefinitionUri`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the workflow version.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `WDL | NEXTFLOW | CWL`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Main`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterTemplate`

Property description not available.

_Required_: No

_Type_: Object of [WorkflowParameter](aws-properties-omics-workflowversion-workflowparameter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterTemplatePath`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[\S]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`readmeMarkdown`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`readmePath`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`readmeUri`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^s3://([a-z0-9][a-z0-9-.]{1,61}[a-z0-9])/((.{1,1024}))$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageCapacity`

Property description not available.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `STATIC | DYNAMIC`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionName`

The name of the workflow version.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9\-\._]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkflowBucketOwnerId`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `1`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkflowId`

The workflow's ID.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+$`

_Minimum_: `1`

_Maximum_: `18`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

ARN of the workflow version.

`CreationTime`

The creation time of the workflow version.

`Status`

The status of the workflow version.

`Type`

The type of the workflow version.

`Uuid`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowParameter

ContainerRegistryMap

All content copied from https://docs.aws.amazon.com/.
