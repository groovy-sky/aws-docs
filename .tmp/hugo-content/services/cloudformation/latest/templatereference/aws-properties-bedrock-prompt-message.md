This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt Message

A message input, or returned from, a call to [Converse](../../../../reference/bedrock/latest/apireference/api-runtime-converse.md) or [ConverseStream](../../../../reference/bedrock/latest/apireference/api-runtime-conversestream.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : [ ContentBlock, ... ],
  "Role" : String
}

```

### YAML

```yaml

  Content:
    - ContentBlock
  Role: String

```

## Properties

`Content`

The message content. Note the following restrictions:

- You can include up to 20 images. Each image's size, height, and width must be no more than 3.75 MB, 8000 px, and 8000 px, respectively.

- You can include up to five documents. Each document's size must be no more than 4.5 MB.

- If you include a `ContentBlock` with a `document` field in the array, you must also include a `ContentBlock` with a `text` field.

- You can only include images and documents if the `role` is `user`.

_Required_: Yes

_Type_: Array of [ContentBlock](aws-properties-bedrock-prompt-contentblock.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The role that the message plays in the message.

_Required_: Yes

_Type_: String

_Allowed values_: `user | assistant`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContentBlock

PromptAgentResource

All content copied from https://docs.aws.amazon.com/.
