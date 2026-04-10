This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::Workflow

Creates a private workflow. Before you create a private workflow, you must create and configure these required resources:

- _Workflow definition file:_ A workflow definition file written in WDL,
Nextflow, or CWL. The workflow definition specifies the inputs and outputs
for runs that use the workflow. It also includes specifications for the runs and run tasks for your workflow,
including compute and memory requirements. The workflow definition file must be in `.zip` format.
For more information, see [Workflow definition files](../../../omics/latest/dev/workflow-definition-files.md) in AWS HealthOmics.

- You can use Amazon Q CLI to build and validate your workflow definition files in WDL, Nextflow, and CWL. For more information, see
[Example prompts for Amazon Q CLI](../../../omics/latest/dev/getting-started.md#omics-q-prompts)
and the [AWS HealthOmics Agentic generative AI tutorial](https://github.com/aws-samples/aws-healthomics-tutorials/tree/main/generative-ai)
on GitHub.

- _(Optional) Parameter template file:_ A parameter template file written in JSON.
Create the file to define the run parameters, or AWS HealthOmics generates the parameter template for you. For more information,
see [Parameter template files for HealthOmics workflows](../../../omics/latest/dev/parameter-templates.md).

- _ECR container images:_ Create container images for
the workflow in a private ECR repository, or synchronize images from a supported upstream
registry with your Amazon ECR private repository.

- _(Optional) Sentieon licenses:_ Request a Sentieon license to use the
Sentieon software in private workflows.

For more information, see [Creating or updating a private workflow in AWS HealthOmics](../../../omics/latest/dev/creating-private-workflows.md) in the _AWS HealthOmics User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Omics::Workflow",
  "Properties" : {
      "Accelerators" : String,
      "ContainerRegistryMap" : ContainerRegistryMap,
      "ContainerRegistryMapUri" : String,
      "DefinitionRepository" : DefinitionRepository,
      "DefinitionUri" : String,
      "Description" : String,
      "Engine" : String,
      "Main" : String,
      "Name" : String,
      "ParameterTemplate" : {Key: Value, ...},
      "ParameterTemplatePath" : String,
      "readmeMarkdown" : String,
      "readmePath" : String,
      "readmeUri" : String,
      "StorageCapacity" : Number,
      "StorageType" : String,
      "Tags" : {Key: Value, ...},
      "WorkflowBucketOwnerId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Omics::Workflow
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
  Name: String
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
  WorkflowBucketOwnerId: String

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

_Type_: [ContainerRegistryMap](aws-properties-omics-workflow-containerregistrymap.md)

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

_Type_: [DefinitionRepository](aws-properties-omics-workflow-definitionrepository.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefinitionUri`

The URI of a definition for the workflow.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The parameter's description.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

An engine for the workflow.

_Required_: No

_Type_: String

_Allowed values_: `WDL | NEXTFLOW | CWL`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Main`

The path of the main definition file for the workflow.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The workflow's name.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterTemplate`

The workflow's parameter template.

_Required_: No

_Type_: Object of [WorkflowParameter](aws-properties-omics-workflow-workflowparameter.md)

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

The default static storage capacity (in gibibytes) for runs that use this workflow or workflow version.
The `storageCapacity` can be overwritten at run time. The storage capacity is not required for runs with a `DYNAMIC` storage type.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `STATIC | DYNAMIC`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags for the workflow.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowBucketOwnerId`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the details of this resource. For example:

`{ "Ref": "Workflow.Type" }` `Ref` returns the type of workflow.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN for the workflow.

`CreationTime`

When the workflow was created.

`Id`

The workflow's ID.

`Status`

The workflow's status.

`Type`

The workflow's type.

`Uuid`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SseConfig

ContainerRegistryMap

All content copied from https://docs.aws.amazon.com/.
