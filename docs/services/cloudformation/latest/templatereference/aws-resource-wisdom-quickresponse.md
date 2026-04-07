This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::QuickResponse

Creates an Amazon Q in Connect quick response.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::QuickResponse",
  "Properties" : {
      "Channels" : [ String, ... ],
      "Content" : QuickResponseContentProvider,
      "ContentType" : String,
      "Description" : String,
      "GroupingConfiguration" : GroupingConfiguration,
      "IsActive" : Boolean,
      "KnowledgeBaseArn" : String,
      "Language" : String,
      "Name" : String,
      "ShortcutKey" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::QuickResponse
Properties:
  Channels:
    - String
  Content:
    QuickResponseContentProvider
  ContentType: String
  Description: String
  GroupingConfiguration:
    GroupingConfiguration
  IsActive: Boolean
  KnowledgeBaseArn: String
  Language: String
  Name: String
  ShortcutKey: String
  Tags:
    - Tag

```

## Properties

`Channels`

The Amazon Connect contact channels this quick response applies to. The supported contact channel types include `Chat`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Content`

The content of the quick response.

_Required_: Yes

_Type_: [QuickResponseContentProvider](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-quickresponse-quickresponsecontentprovider.html)

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentType`

The media type of the quick response content.

- Use `application/x.quickresponse;format=plain` for quick response written in plain text.

- Use `application/x.quickresponse;format=markdown` for quick response written in richtext.

_Required_: No

_Type_: String

_Pattern_: `^(application/x\.quickresponse;format=(plain|markdown))$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the quick response.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupingConfiguration`

The configuration information of the user groups that the quick response is accessible to.

_Required_: No

_Type_: [GroupingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-quickresponse-groupingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsActive`

Whether the quick response is active.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBaseArn`

The Amazon Resource Name (ARN) of the knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}){0,2}$`

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

`Name`

The name of the quick response.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShortcutKey`

The shortcut key of the quick response. The value should be unique across the
knowledge base.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-quickresponse-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`QuickResponseArn`

The Amazon Resource Name (ARN) of the quick response.

`QuickResponseId`

The identifier of the quick response.

`Status`

The status of the quick response data.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Wisdom::MessageTemplateVersion

GroupingConfiguration
