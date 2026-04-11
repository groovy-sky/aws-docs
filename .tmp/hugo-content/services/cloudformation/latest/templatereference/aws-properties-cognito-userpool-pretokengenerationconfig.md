This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool PreTokenGenerationConfig

The properties of a pre token generation Lambda trigger.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LambdaArn" : String,
  "LambdaVersion" : String
}

```

### YAML

```yaml

  LambdaArn: String
  LambdaVersion: String

```

## Properties

`LambdaArn`

The Amazon Resource Name (ARN) of the function that you want to assign to your Lambda trigger.

This parameter and the `PreTokenGeneration` property of
`LambdaConfig` have the same value. For new instances of pre token
generation triggers, set `LambdaArn`.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaVersion`

The user pool trigger version of the request that Amazon Cognito sends to your Lambda function. Higher-numbered versions add fields that support new features.

_Required_: No

_Type_: String

_Allowed values_: `V1_0 | V2_0 | V3_0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Policies

RecoveryOption

All content copied from https://docs.aws.amazon.com/.
