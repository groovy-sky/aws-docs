This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::FlowAlias FlowAliasConcurrencyConfiguration

Determines how multiple nodes in a flow can run in parallel. Running nodes concurrently can improve your flow's performance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxConcurrency" : Number,
  "Type" : String
}

```

### YAML

```yaml

  MaxConcurrency: Number
  Type: String

```

## Properties

`MaxConcurrency`

The maximum number of nodes that can be executed concurrently in the flow.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of concurrency to use for parallel node execution. Specify one of the following options:

- `Automatic` \- Amazon Bedrock determines which nodes can be executed in parallel based on the flow definition and its dependencies.

- `Manual` \- You specify which nodes can be executed in parallel.

_Required_: Yes

_Type_: String

_Allowed values_: `Automatic | Manual`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Bedrock::FlowAlias

FlowAliasRoutingConfigurationListItem
