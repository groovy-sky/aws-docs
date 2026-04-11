This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow

Creates a matching workflow that defines the configuration for a data processing job.
The workflow name must be unique. To modify an existing workflow, use
`UpdateMatchingWorkflow`.

###### Important

For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER`, incremental
processing is not supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EntityResolution::MatchingWorkflow",
  "Properties" : {
      "Description" : String,
      "IncrementalRunConfig" : IncrementalRunConfig,
      "InputSourceConfig" : [ InputSource, ... ],
      "OutputSourceConfig" : [ OutputSource, ... ],
      "ResolutionTechniques" : ResolutionTechniques,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "WorkflowName" : String
    }
}

```

### YAML

```yaml

Type: AWS::EntityResolution::MatchingWorkflow
Properties:
  Description: String
  IncrementalRunConfig:
    IncrementalRunConfig
  InputSourceConfig:
    - InputSource
  OutputSourceConfig:
    - OutputSource
  ResolutionTechniques:
    ResolutionTechniques
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

`IncrementalRunConfig`

Optional. An object that defines the incremental run type. This object contains only the
`incrementalRunType` field, which appears as "Automatic" in the console.

###### Important

For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER`,
incremental processing is not supported.

_Required_: No

_Type_: [IncrementalRunConfig](aws-properties-entityresolution-matchingworkflow-incrementalrunconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSourceConfig`

A list of `InputSource` objects, which have the fields
`InputSourceARN` and `SchemaName`.

_Required_: Yes

_Type_: Array of [InputSource](aws-properties-entityresolution-matchingworkflow-inputsource.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputSourceConfig`

A list of `OutputSource` objects, each of which contains fields
`outputS3Path`, `applyNormalization`, `KMSArn`, and
`output`.

_Required_: Yes

_Type_: Array of [OutputSource](aws-properties-entityresolution-matchingworkflow-outputsource.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolutionTechniques`

An object which defines the `resolutionType` and the
`ruleBasedProperties`.

_Required_: Yes

_Type_: [ResolutionTechniques](aws-properties-entityresolution-matchingworkflow-resolutiontechniques.md)

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

_Type_: Array of [Tag](aws-properties-entityresolution-matchingworkflow-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowName`

The name of the workflow. There can't be multiple `MatchingWorkflows` with
the same name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_0-9-]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CustomerProfilesIntegrationConfig

All content copied from https://docs.aws.amazon.com/.
