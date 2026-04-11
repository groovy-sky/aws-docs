This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent Function

Defines parameters that the agent needs to invoke from the user to complete the function. Corresponds to an action in an action group.

This data type is used in the following API operations:

- [CreateAgentActionGroup request](../../../../reference/bedrock/latest/apireference/api-agent-createagentactiongroup.md#API_agent_CreateAgentActionGroup_RequestSyntax)

- [CreateAgentActionGroup response](../../../../reference/bedrock/latest/apireference/api-agent-createagentactiongroup.md#API_agent_CreateAgentActionGroup_ResponseSyntax)

- [UpdateAgentActionGroup request](../../../../reference/bedrock/latest/apireference/api-agent-updateagentactiongroup.md#API_agent_UpdateAgentActionGroup_RequestSyntax)

- [UpdateAgentActionGroup response](../../../../reference/bedrock/latest/apireference/api-agent-updateagentactiongroup.md#API_agent_UpdateAgentActionGroup_ResponseSyntax)

- [GetAgentActionGroup response](../../../../reference/bedrock/latest/apireference/api-agent-getagentactiongroup.md#API_agent_GetAgentActionGroup_ResponseSyntax)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Name" : String,
  "Parameters" : {Key: Value, ...},
  "RequireConfirmation" : String
}

```

### YAML

```yaml

  Description: String
  Name: String
  Parameters:
    Key: Value
  RequireConfirmation: String

```

## Properties

`Description`

A description of the function and its purpose.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the function.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameters that the agent elicits from the user to fulfill the function.

_Required_: No

_Type_: Object of [ParameterDetail](aws-properties-bedrock-agent-parameterdetail.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireConfirmation`

Contains information if user confirmation is required to invoke the function.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomOrchestration

FunctionSchema

All content copied from https://docs.aws.amazon.com/.
