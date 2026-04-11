This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::AgentAlias

Specifies an agent alias as a resource in a top-level template. Minimally, you must
specify the following properties:

- AgentAliasName – Specify a name for the alias.

For more information about creating aliases for an agent in Amazon Bedrock, see
[Deploy\
an Amazon Bedrock agent](../../../bedrock/latest/userguide/agents-deploy.md).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::AgentAlias",
  "Properties" : {
      "AgentAliasName" : String,
      "AgentId" : String,
      "Description" : String,
      "RoutingConfiguration" : [ AgentAliasRoutingConfigurationListItem, ... ],
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::AgentAlias
Properties:
  AgentAliasName: String
  AgentId: String
  Description: String
  RoutingConfiguration:
    - AgentAliasRoutingConfigurationListItem
  Tags:
    Key: Value

```

## Properties

`AgentAliasName`

The name of the alias of the agent.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AgentId`

The unique identifier of the agent.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the alias of the agent.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingConfiguration`

Contains details about the routing configuration of the alias.

_Required_: No

_Type_: Array of [AgentAliasRoutingConfigurationListItem](aws-properties-bedrock-agentalias-agentaliasroutingconfigurationlistitem.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that you can assign to a resource as key-value pairs. For more information,
see the following resources:

- [Tag naming\
limits and requirements](../../../tag-editor/latest/userguide/tagging.md#tag-conventions)

- [Tagging\
best practices](../../../tag-editor/latest/userguide/tagging.md#tag-best-practices)

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the agent ID and the agent alias ID, separated by a pipe
( `|`).

For example, `{ "Ref": "myAgentAlias" }` could return the value
`"AGENT12345|ALIAS12345"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AgentAliasArn`

The Amazon Resource Name (ARN) of the alias of the agent.

`AgentAliasHistoryEvents`

Contains details about the history of the alias.

`AgentAliasId`

The unique identifier of the alias of the agent.

`AgentAliasStatus`

The status of the alias of the agent and whether it is ready for use. The following statuses are possible:

- CREATING – The agent alias is being created.

- PREPARED – The agent alias is finished being created or updated and is ready to be invoked.

- FAILED – The agent alias API operation failed.

- UPDATING – The agent alias is being updated.

- DELETING – The agent alias is being deleted.

- DISSOCIATED - The agent alias has no version associated with it.

`CreatedAt`

The time at which the alias of the agent was created.

`UpdatedAt`

The time at which the alias was last updated.

## Examples

### Create an alias for an agent

The following example creates an alias that points to version 1 of an
agent.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: "CFN stack for creating an AgentAlias"
Resources:
  ExampleAgentAliasResource:
    Type: AWS::Bedrock::AgentAlias
    Properties:
      AgentId: "1234567890"
      AgentAliasName: "TestAlias"
      Description: "Alias for testing"
      RoutingConfiguration:
        - AgentVersion: "1"
```

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Description": "CFN stack for creating an AgentAlias",
   "Resources": {
      "ExampleAgentAliasResource": {
         "Type": "AWS::Bedrock::AgentAlias",
         "Properties": {
            "AgentId": "1234567890",
            "AgentAliasName": "TestAlias",
            "Description": "Alias for testing",
            "RoutingConfiguration": {
               "AgentVersion": "1"
            }
         }
      }
   }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SessionSummaryConfiguration

AgentAliasHistoryEvent

All content copied from https://docs.aws.amazon.com/.
