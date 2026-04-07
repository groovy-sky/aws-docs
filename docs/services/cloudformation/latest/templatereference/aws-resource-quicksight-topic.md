This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic

Creates a new Q topic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::Topic",
  "Properties" : {
      "AwsAccountId" : String,
      "ConfigOptions" : TopicConfigOptions,
      "CustomInstructions" : CustomInstructions,
      "DataSets" : [ DatasetMetadata, ... ],
      "Description" : String,
      "FolderArns" : [ String, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "TopicId" : String,
      "UserExperienceVersion" : String
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::Topic
Properties:
  AwsAccountId: String
  ConfigOptions:
    TopicConfigOptions
  CustomInstructions:
    CustomInstructions
  DataSets:
    - DatasetMetadata
  Description: String
  FolderArns:
    - String
  Name: String
  Tags:
    - Tag
  TopicId: String
  UserExperienceVersion: String

```

## Properties

`AwsAccountId`

The ID of the AWS account that you want to create a topic in.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigOptions`

Configuration options for a `Topic`.

_Required_: No

_Type_: [TopicConfigOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-topicconfigoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomInstructions`

Custom instructions for the topic.

_Required_: No

_Type_: [CustomInstructions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-custominstructions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSets`

The data sets that the topic is associated with.

_Required_: No

_Type_: Array of [DatasetMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-datasetmetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the topic.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FolderArns`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the topic.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicId`

The ID for the topic. This ID is unique per AWS Region for each AWS account.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9-_.\\+]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserExperienceVersion`

The user experience version of the topic.

_Required_: No

_Type_: String

_Allowed values_: `LEGACY | NEW_READER_EXPERIENCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the topic.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UIColorPalette

CellValueSynonym
