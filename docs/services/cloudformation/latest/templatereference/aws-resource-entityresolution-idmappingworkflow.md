This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdMappingWorkflow

Creates an `IdMappingWorkflow` object which stores the configuration of the
data processing job to be run. Each `IdMappingWorkflow` must have a unique
workflow name. To modify an existing workflow, use the UpdateIdMappingWorkflow
API.

###### Important

Incremental processing is not supported for ID mapping workflows.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EntityResolution::IdMappingWorkflow",
  "Properties" : {
      "Description" : String,
      "IdMappingIncrementalRunConfig" : IdMappingIncrementalRunConfig,
      "IdMappingTechniques" : IdMappingTechniques,
      "InputSourceConfig" : [ IdMappingWorkflowInputSource, ... ],
      "OutputSourceConfig" : [ IdMappingWorkflowOutputSource, ... ],
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "WorkflowName" : String
    }
}

```

### YAML

```yaml

Type: AWS::EntityResolution::IdMappingWorkflow
Properties:
  Description: String
  IdMappingIncrementalRunConfig:
    IdMappingIncrementalRunConfig
  IdMappingTechniques:
    IdMappingTechniques
  InputSourceConfig:
    - IdMappingWorkflowInputSource
  OutputSourceConfig:
    - IdMappingWorkflowOutputSource
  RoleArn: String
  Tags:
    - Tag
  WorkflowName: String

```

## Properties

`Description`

A description of the workflow.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdMappingIncrementalRunConfig`

Property description not available.

_Required_: No

_Type_: [IdMappingIncrementalRunConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-entityresolution-idmappingworkflow-idmappingincrementalrunconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdMappingTechniques`

An object which defines the ID mapping technique and any additional
configurations.

_Required_: Yes

_Type_: [IdMappingTechniques](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSourceConfig`

A list of `InputSource` objects, which have the fields
`InputSourceARN` and `SchemaName`.

_Required_: Yes

_Type_: Array of [IdMappingWorkflowInputSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowinputsource.html)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputSourceConfig`

A list of `IdMappingWorkflowOutputSource` objects, each of which contains
fields `outputS3Path` and `KMSArn`.

_Required_: No

_Type_: Array of [IdMappingWorkflowOutputSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowoutputsource.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role. AWS Entity Resolution assumes
this role to create resources on your behalf as part of workflow execution.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-entityresolution-idmappingworkflow-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowName`

The name of the workflow. There can't be multiple `IdMappingWorkflows` with
the same name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_0-9-]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Entity Resolution

IdMappingIncrementalRunConfig
