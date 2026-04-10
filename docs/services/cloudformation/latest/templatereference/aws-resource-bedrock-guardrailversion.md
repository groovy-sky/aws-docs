This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::GuardrailVersion

Creates a version of the guardrail. Use this API to create a snapshot of the
guardrail when you are satisfied with a configuration, or to compare the configuration with another version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::GuardrailVersion",
  "Properties" : {
      "Description" : String,
      "GuardrailIdentifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::GuardrailVersion
Properties:
  Description: String
  GuardrailIdentifier: String

```

## Properties

`Description`

A description of the guardrail version.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GuardrailIdentifier`

The unique identifier of the guardrail. This can be an ID or the ARN.

_Required_: Yes

_Type_: String

_Pattern_: `^(([a-z0-9]+)|(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:guardrail/[a-z0-9]+))$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the guardrail ID and version, separated by a pipe
( `|`) in the format ( `guardrail-id|guardrail-version`).

For example, `{ "Ref": "myGuardrailVersion" }` would return the value
`"3mzzryddufkp|2"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GuardrailArn`

The ARN of the guardrail.

`GuardrailId`

The unique identifier of the guardrail.

`Version`

The version of the guardrail.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WordPolicyConfig

AWS::Bedrock::IntelligentPromptRouter

All content copied from https://docs.aws.amazon.com/.
