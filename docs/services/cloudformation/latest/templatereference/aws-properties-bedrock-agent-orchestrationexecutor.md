This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent OrchestrationExecutor

The structure of the executor invoking the actions in custom orchestration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Lambda" : String
}

```

### YAML

```yaml

  Lambda: String

```

## Properties

`Lambda`

The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:lambda:[a-z0-9-]{1,20}:\d{12}:function:[a-zA-Z0-9-_\.]+(:(\$LATEST|[a-zA-Z0-9-_]+))?$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MemoryConfiguration

ParameterDetail

All content copied from https://docs.aws.amazon.com/.
