This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway InterceptorInputConfiguration

The input configuration of the interceptor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PassRequestHeaders" : Boolean
}

```

### YAML

```yaml

  PassRequestHeaders: Boolean

```

## Properties

`PassRequestHeaders`

Indicates whether to pass request headers as input into the interceptor. When set to true, request headers will be passed.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InterceptorConfiguration

LambdaInterceptorConfiguration

All content copied from https://docs.aws.amazon.com/.
