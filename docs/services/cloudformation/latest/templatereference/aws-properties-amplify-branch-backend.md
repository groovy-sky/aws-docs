This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::Branch Backend

Describes the backend associated with an Amplify
`Branch`.

This property is available to Amplify Gen 2 apps only. When you deploy
an application with Amplify Gen 2, you provision the app's backend infrastructure using
Typescript code.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StackArn" : String
}

```

### YAML

```yaml

  StackArn: String

```

## Properties

`StackArn`

The Amazon Resource Name (ARN) for the CloudFormation stack.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Amplify::Branch

BasicAuthConfig
