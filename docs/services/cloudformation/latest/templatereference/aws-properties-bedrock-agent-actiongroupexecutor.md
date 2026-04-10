This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent ActionGroupExecutor

Contains details about the Lambda function containing the business logic that is carried out upon invoking the action or the
custom control method for handling the information elicited from the user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomControl" : String,
  "Lambda" : String
}

```

### YAML

```yaml

  CustomControl: String
  Lambda: String

```

## Properties

`CustomControl`

To return the action group invocation results directly in the `InvokeInlineAgent` response, specify `RETURN_CONTROL`.

_Required_: No

_Type_: String

_Allowed values_: `RETURN_CONTROL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lambda`

The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:lambda:[a-z0-9-]{1,20}:\d{12}:function:[a-zA-Z0-9-_\.]+(:(\$LATEST|[a-zA-Z0-9-_]+))?$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Bedrock::Agent

AgentActionGroup

All content copied from https://docs.aws.amazon.com/.
