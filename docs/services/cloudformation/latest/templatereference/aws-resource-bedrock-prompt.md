This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt

Creates a prompt in your prompt library that you can add to a flow. For more information, see [Prompt management in Amazon Bedrock](../../../bedrock/latest/userguide/prompt-management.md), [Create a prompt using Prompt management](../../../bedrock/latest/userguide/prompt-management-create.md) and [Prompt flows in Amazon Bedrock](../../../bedrock/latest/userguide/flows.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::Prompt",
  "Properties" : {
      "CustomerEncryptionKeyArn" : String,
      "DefaultVariant" : String,
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...},
      "Variants" : [ PromptVariant, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::Prompt
Properties:
  CustomerEncryptionKeyArn: String
  DefaultVariant: String
  Description: String
  Name: String
  Tags:
    Key: Value
  Variants:
    - PromptVariant

```

## Properties

`CustomerEncryptionKeyArn`

The Amazon Resource Name (ARN) of the KMS key that the prompt is encrypted with.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultVariant`

The name of the default variant for the prompt. This value must match the `name` field in the relevant [PromptVariant](../../../../reference/bedrock/latest/apireference/api-agent-promptvariant.md) object.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the prompt.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the prompt.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that you can assign to a resource as key-value pairs. For more information,
see the following resources:

- [Tag naming\
limits and requirements](../../../tag-editor/latest/userguide/tagging.md#tag-conventions)

- [Tagging\
best practices](../../../tag-editor/latest/userguide/tagging.md#tag-best-practices)

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variants`

A list of objects, each containing details about a variant of the prompt.

_Required_: No

_Type_: Array of [PromptVariant](aws-properties-bedrock-prompt-promptvariant.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Number (ARN) of the prompt.

For example, `{ "Ref": "myPrompt" }` could return the value
`"arn:aws:bedrock:us-east-1:123456789012:prompt/PROMPT12345"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the prompt or the prompt version (if you specified a version in the request).

`CreatedAt`

The time at which the prompt was created.

`Id`

The unique identifier of the prompt.

`UpdatedAt`

The time at which the prompt was last updated.

`Version`

The version of the prompt that this summary applies to.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoSegmentationConfiguration

CachePointBlock

All content copied from https://docs.aws.amazon.com/.
