This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule AppSyncParameters

Contains the GraphQL operation to be parsed and executed, if the event target is an
AWS AppSync API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GraphQLOperation" : String
}

```

### YAML

```yaml

  GraphQLOperation: String

```

## Properties

`GraphQLOperation`

The GraphQL operation; that is, the query, mutation, or subscription to be parsed and
executed by the GraphQL service.

For more information, see [Operations](../../../appsync/latest/devguide/graphql-architecture.md#graphql-operations) in the _AWS AppSync User Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1048576`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Events::Rule

AwsVpcConfiguration

All content copied from https://docs.aws.amazon.com/.
