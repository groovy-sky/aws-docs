This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association

The `AWS::DevOpsAgent::Association` resource specifies an association between an Agent Space and a
service, defining how the Agent Space interacts with external services like GitHub, Slack, AWS
accounts, and others.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DevOpsAgent::Association",
  "Properties" : {
      "AgentSpaceId" : String,
      "Configuration" : ServiceConfiguration,
      "LinkedAssociationIds" : [ String, ... ],
      "ServiceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::DevOpsAgent::Association
Properties:
  AgentSpaceId: String
  Configuration:
    ServiceConfiguration
  LinkedAssociationIds:
    - String
  ServiceId: String

```

## Properties

`AgentSpaceId`

The unique identifier of the Agent Space.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`Configuration`

The configuration that directs how the Agent Space interacts with the given service. You can specify only one
configuration type per association.

_Allowed Values_: `SourceAws` \| `Aws` \| `GitHub` \|
`GitLab` \| `Slack` \| `Dynatrace` \| `ServiceNow` \| `MCPServer`
\| `MCPServerNewRelic` \| `MCPServerDatadog` \| `MCPServerSplunk` \|
`EventChannel`

_Required_: Yes

_Type_: [ServiceConfiguration](aws-properties-devopsagent-association-serviceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinkedAssociationIds`

Set of linked association IDs for parent-child relationships.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceId`

The identifier for the associated service. For `SourceAws` and `Aws` configurations, this
must be `aws`. For all other service types, this is a UUID generated from the RegisterService
command.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the association ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationId`

The unique identifier of the association.

`CreatedAt`

The timestamp when the association was created.

`UpdatedAt`

The timestamp when the association was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWSConfiguration

All content copied from https://docs.aws.amazon.com/.
