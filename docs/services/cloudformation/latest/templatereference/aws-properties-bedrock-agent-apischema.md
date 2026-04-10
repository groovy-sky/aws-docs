This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent APISchema

Contains details about the OpenAPI schema for the action group. For more information, see [Action group OpenAPI schemas](../../../bedrock/latest/userguide/agents-api-schema.md).
You can either include the schema directly in the payload field or you can upload it to an S3 bucket and specify the S3 bucket location in the s3 field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Payload" : String,
  "S3" : S3Identifier
}

```

### YAML

```yaml

  Payload: String
  S3:
    S3Identifier

```

## Properties

`Payload`

The JSON or YAML-formatted payload defining the OpenAPI schema for the action group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

Contains details about the S3 object containing the OpenAPI schema for the action group.

_Required_: No

_Type_: [S3Identifier](aws-properties-bedrock-agent-s3identifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AgentKnowledgeBase

CustomOrchestration

All content copied from https://docs.aws.amazon.com/.
