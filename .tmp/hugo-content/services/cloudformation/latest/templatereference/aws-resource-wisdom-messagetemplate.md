This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate

Creates an Amazon Q in Connect message template. The name of the message template has to be unique for each
knowledge base. The channel subtype of the message template is immutable and cannot be modified after creation. After
the message template is created, you can use the `$LATEST` qualifier to reference the created message
template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::MessageTemplate",
  "Properties" : {
      "ChannelSubtype" : String,
      "Content" : Content,
      "DefaultAttributes" : MessageTemplateAttributes,
      "Description" : String,
      "GroupingConfiguration" : GroupingConfiguration,
      "KnowledgeBaseArn" : String,
      "Language" : String,
      "MessageTemplateAttachments" : [ MessageTemplateAttachment, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::MessageTemplate
Properties:
  ChannelSubtype: String
  Content:
    Content
  DefaultAttributes:
    MessageTemplateAttributes
  Description: String
  GroupingConfiguration:
    GroupingConfiguration
  KnowledgeBaseArn: String
  Language: String
  MessageTemplateAttachments:
    - MessageTemplateAttachment
  Name: String
  Tags:
    - Tag

```

## Properties

`ChannelSubtype`

The channel subtype this message template applies to.

_Required_: Yes

_Type_: String

_Allowed values_: `EMAIL | SMS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Content`

The content of the message template.

_Required_: Yes

_Type_: [Content](aws-properties-wisdom-messagetemplate-content.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultAttributes`

An object that specifies the default values to use for variables in the message template. This object contains
different categories of key-value pairs. Each key defines a variable or placeholder in the message template. The
corresponding value defines the default value for that variable.

_Required_: No

_Type_: [MessageTemplateAttributes](aws-properties-wisdom-messagetemplate-messagetemplateattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the message template.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\\s_.,-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupingConfiguration`

The configuration information of the external data source.

_Required_: No

_Type_: [GroupingConfiguration](aws-properties-wisdom-messagetemplate-groupingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBaseArn`

The Amazon Resource Name (ARN) of the knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Language`

The language code value for the language in which the quick response is written. The supported language codes include `de_DE`, `en_US`, `es_ES`,
`fr_FR`, `id_ID`, `it_IT`, `ja_JP`, `ko_KR`, `pt_BR`,
`zh_CN`, `zh_TW`

_Required_: No

_Type_: String

_Minimum_: `2`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageTemplateAttachments`

List of message template attachments.

_Required_: No

_Type_: Array of [MessageTemplateAttachment](aws-properties-wisdom-messagetemplate-messagetemplateattachment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the message template.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\\s_.,-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-wisdom-messagetemplate-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`MessageTemplateArn`

The Amazon Resource Name (ARN) of the message template.

`MessageTemplateContentSha256`

The checksum value of the message template content that is referenced by the `$LATEST` qualifier. It
can be returned in `MessageTemplateData` or `ExtendedMessageTemplateData`. It’s calculated by
content, language, `defaultAttributes` and `Attachments` of the message template.

`MessageTemplateId`

The identifier of the message template.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerConfiguration

AgentAttributes

All content copied from https://docs.aws.amazon.com/.
