This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::PromptVersion

Creates a static snapshot of your prompt that can be deployed to production. For more information, see [Deploy prompts using Prompt management by creating versions](../../../bedrock/latest/userguide/prompt-management-deploy.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::PromptVersion",
  "Properties" : {
      "Description" : String,
      "PromptArn" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::PromptVersion
Properties:
  Description: String
  PromptArn: String
  Tags:
    Key: Value

```

## Properties

`Description`

The description of the prompt version.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PromptArn`

The Amazon Resource Name (ARN) of the version of the prompt.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:prompt/[0-9a-zA-Z]{10})$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A map of tags attached to the prompt version and their values.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Number (ARN) of the prompt
version.

For example, `{ "Ref": "myPromptVersion" }` could return the value
`""arn:aws:bedrock:us-east-1:123456789012:prompt/PROMPT12345:1"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the prompt or the prompt version (if you specified a version in the request).

`CreatedAt`

The time at which the prompt was created.

`CustomerEncryptionKeyArn`

The Amazon Resource Name (ARN) of the KMS key that the prompt version is encrypted
with.

`DefaultVariant`

The name of the default variant for the prompt. This value must match the `name` field in the relevant [PromptVariant](../../../../reference/bedrock/latest/apireference/api-agent-promptvariant.md) object.

`Name`

The name of the prompt.

`PromptId`

The unique identifier of the prompt.

`UpdatedAt`

The time at which the prompt was last updated.

`Variants`

A list of objects, each containing details about a variant of the prompt.

`Version`

The version of the prompt that this summary applies to.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolSpecification

CachePointBlock

All content copied from https://docs.aws.amazon.com/.
