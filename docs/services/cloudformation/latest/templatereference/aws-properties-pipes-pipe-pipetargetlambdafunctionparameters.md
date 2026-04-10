This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetLambdaFunctionParameters

The parameters for using a Lambda function as a target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InvocationType" : String
}

```

### YAML

```yaml

  InvocationType: String

```

## Properties

`InvocationType`

Specify whether to invoke the function synchronously or asynchronously.

- `REQUEST_RESPONSE` (default) - Invoke synchronously. This corresponds
to the `RequestResponse` option in the `InvocationType`
parameter for the Lambda
[Invoke](../../../lambda/latest/dg/api-invoke.md#API_Invoke_RequestSyntax)
API.

- `FIRE_AND_FORGET` \- Invoke asynchronously. This corresponds to the
`Event` option in the `InvocationType` parameter for the
Lambda
[Invoke](../../../lambda/latest/dg/api-invoke.md#API_Invoke_RequestSyntax)
API.

For more information, see [Invocation\
types](../../../eventbridge/latest/userguide/eb-pipes.md#pipes-invocation) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `REQUEST_RESPONSE | FIRE_AND_FORGET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetKinesisStreamParameters

PipeTargetParameters

All content copied from https://docs.aws.amazon.com/.
