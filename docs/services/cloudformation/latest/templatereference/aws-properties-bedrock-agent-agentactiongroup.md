This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent AgentActionGroup

Contains details of the inline agent's action group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionGroupExecutor" : ActionGroupExecutor,
  "ActionGroupName" : String,
  "ActionGroupState" : String,
  "ApiSchema" : APISchema,
  "Description" : String,
  "FunctionSchema" : FunctionSchema,
  "ParentActionGroupSignature" : String,
  "SkipResourceInUseCheckOnDelete" : Boolean
}

```

### YAML

```yaml

  ActionGroupExecutor:
    ActionGroupExecutor
  ActionGroupName: String
  ActionGroupState: String
  ApiSchema:
    APISchema
  Description: String
  FunctionSchema:
    FunctionSchema
  ParentActionGroupSignature: String
  SkipResourceInUseCheckOnDelete: Boolean

```

## Properties

`ActionGroupExecutor`

The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking
the action or the custom control method for handling the information elicited from the user.

_Required_: No

_Type_: [ActionGroupExecutor](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-actiongroupexecutor.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ActionGroupName`

The name of the action group.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ActionGroupState`

Specifies whether the action group is available for the agent to invoke or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApiSchema`

Contains either details about the S3 object containing the OpenAPI schema for the action group or the JSON or YAML-formatted
payload defining the schema. For more information, see [Action group OpenAPI schemas](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-api-schema.html).

_Required_: No

_Type_: [APISchema](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-apischema.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the action group.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionSchema`

Contains details about the function schema for the action group or the JSON or YAML-formatted payload defining the schema.

_Required_: No

_Type_: [FunctionSchema](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-functionschema.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentActionGroupSignature`

If this field is set as `AMAZON.UserInput`, the agent can request the user for additional information when trying to complete a task. The `description`, `apiSchema`, and `actionGroupExecutor` fields must be blank for this action group.

During orchestration, if the agent determines that it needs to invoke an API in an action group, but doesn't have enough information to complete the API request, it will invoke this action group instead and return an [Observation](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_Observation.html) reprompting the user for more information.

_Required_: No

_Type_: String

_Allowed values_: `AMAZON.UserInput | AMAZON.CodeInterpreter`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkipResourceInUseCheckOnDelete`

Specifies whether to delete the resource even if it's in use. By default, this value
is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ActionGroupExecutor

AgentCollaborator
