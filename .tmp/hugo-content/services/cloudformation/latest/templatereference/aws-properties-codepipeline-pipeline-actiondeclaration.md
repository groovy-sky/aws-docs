This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline ActionDeclaration

Represents information about an action declaration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionTypeId" : ActionTypeId,
  "Commands" : [ String, ... ],
  "Configuration" : Json,
  "EnvironmentVariables" : [ EnvironmentVariable, ... ],
  "InputArtifacts" : [ InputArtifact, ... ],
  "Name" : String,
  "Namespace" : String,
  "OutputArtifacts" : [ OutputArtifact, ... ],
  "OutputVariables" : [ String, ... ],
  "Region" : String,
  "RoleArn" : String,
  "RunOrder" : Integer,
  "TimeoutInMinutes" : Integer
}

```

### YAML

```yaml

  ActionTypeId:
    ActionTypeId
  Commands:
    - String
  Configuration: Json
  EnvironmentVariables:
    - EnvironmentVariable
  InputArtifacts:
    - InputArtifact
  Name: String
  Namespace: String
  OutputArtifacts:
    - OutputArtifact
  OutputVariables:
    - String
  Region: String
  RoleArn: String
  RunOrder: Integer
  TimeoutInMinutes: Integer

```

## Properties

`ActionTypeId`

Specifies the action type and the provider of the action.

_Required_: Yes

_Type_: [ActionTypeId](aws-properties-codepipeline-pipeline-actiontypeid.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Commands`

The shell commands to run with your compute action in CodePipeline. All commands
are supported except multi-line formats. While CodeBuild logs and permissions
are used, you do not need to create any resources in CodeBuild.

###### Note

Using compute time for this action will incur separate charges in AWS CodeBuild.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configuration`

The action's configuration. These are key-value pairs that specify input values for
an action. For more information, see [Action Structure Requirements in CodePipeline](../../../codepipeline/latest/userguide/reference-pipeline-structure.md#action-requirements). For the list of
configuration properties for the AWS CloudFormation action type in CodePipeline, see [Configuration Properties Reference](../userguide/continuous-delivery-codepipeline-action-reference.md) in the _AWS CloudFormation_
_User Guide_. For template snippets with examples, see [Using Parameter Override Functions with CodePipeline Pipelines](../userguide/continuous-delivery-codepipeline-parameter-override-functions.md) in the
_AWS CloudFormation User Guide_.

The values can be represented in either JSON or YAML format. For example, the JSON
configuration item format is as follows:

_JSON:_

`"Configuration" : { Key : Value },`

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentVariables`

The environment variables for the action.

_Required_: No

_Type_: Array of [EnvironmentVariable](aws-properties-codepipeline-pipeline-environmentvariable.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputArtifacts`

The name or ID of the artifact consumed by the action, such as a test or build
artifact. While the field is not a required parameter, most actions have an action
configuration that requires a specified quantity of input artifacts. To refer to the action
configuration specification by action provider, see the [Action structure reference](../../../codepipeline/latest/userguide/action-reference.md)
in the _AWS CodePipeline User Guide_.

###### Note

For a CodeBuild action with multiple input artifacts, one of your input sources must be
designated the PrimarySource. For more information, see the [CodeBuild action\
reference page](../../../codepipeline/latest/userguide/action-reference-codebuild.md) in the _AWS CodePipeline User_
_Guide_.

_Required_: No

_Type_: Array of [InputArtifact](aws-properties-codepipeline-pipeline-inputartifact.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The action declaration's name.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The variable namespace associated with the action. All variables produced as output by
this action fall under this namespace.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputArtifacts`

The name or ID of the result of the action declaration, such as a test or build
artifact. While the field is not a required parameter, most actions have an action
configuration that requires a specified quantity of output artifacts. To refer to the action
configuration specification by action provider, see the [Action structure reference](../../../codepipeline/latest/userguide/action-reference.md)
in the _AWS CodePipeline User Guide_.

_Required_: No

_Type_: Array of [OutputArtifact](aws-properties-codepipeline-pipeline-outputartifact.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputVariables`

The list of variables that are to be exported from the compute action. This is
specifically CodeBuild environment variables as used for that action.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The action declaration's AWS Region, such as us-east-1.

_Required_: No

_Type_: String

_Minimum_: `4`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM service role that performs the declared action. This is assumed
through the roleArn for the pipeline.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunOrder`

The order in which actions are run.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInMinutes`

A timeout duration in minutes that can be applied against the ActionType’s default
timeout value specified in [Quotas for AWS CodePipeline](../../../codepipeline/latest/userguide/limits.md). This attribute is available only to the manual approval ActionType.

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CodePipeline::Pipeline

ActionTypeId

All content copied from https://docs.aws.amazon.com/.
