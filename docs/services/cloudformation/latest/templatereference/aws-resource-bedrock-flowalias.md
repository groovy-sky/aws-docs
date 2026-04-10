This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::FlowAlias

Creates an alias of a flow for deployment. For more information, see [Deploy a flow in Amazon Bedrock](../../../bedrock/latest/userguide/flows-deploy.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::FlowAlias",
  "Properties" : {
      "ConcurrencyConfiguration" : FlowAliasConcurrencyConfiguration,
      "Description" : String,
      "FlowArn" : String,
      "Name" : String,
      "RoutingConfiguration" : [ FlowAliasRoutingConfigurationListItem, ... ],
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::FlowAlias
Properties:
  ConcurrencyConfiguration:
    FlowAliasConcurrencyConfiguration
  Description: String
  FlowArn: String
  Name: String
  RoutingConfiguration:
    - FlowAliasRoutingConfigurationListItem
  Tags:
    Key: Value

```

## Properties

`ConcurrencyConfiguration`

The configuration that specifies how nodes in the flow are executed concurrently.

_Required_: No

_Type_: [FlowAliasConcurrencyConfiguration](aws-properties-bedrock-flowalias-flowaliasconcurrencyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the alias.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowArn`

The Amazon Resource Name (ARN) of the alias.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:flow/[0-9a-zA-Z]{10}$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the alias.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingConfiguration`

A list of configurations about the versions that the alias maps to. Currently, you can only specify one.

_Required_: Yes

_Type_: Array of [FlowAliasRoutingConfigurationListItem](aws-properties-bedrock-flowalias-flowaliasroutingconfigurationlistitem.md)

_Minimum_: `1`

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

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Number (ARN) of the flow and the
Amazon Resource Name (ARN) of its alias, separated by a pipe ( `|`).

For example, `{ "Ref": "myFlowAlias" }` could return the value
`"arn:aws:bedrock:us-east-1:123456789012:flow/FLOW12345|arn:aws:bedrock:us-east-1:123456789012:flow/FLOW12345/alias/ALIAS12345"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the alias.

`CreatedAt`

The time at which the alias was created.

`FlowId`

The unique identifier of the flow.

`Id`

The unique identifier of the alias of the flow.

`UpdatedAt`

The time at which the alias was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VectorSearchRerankingConfiguration

FlowAliasConcurrencyConfiguration

All content copied from https://docs.aws.amazon.com/.
