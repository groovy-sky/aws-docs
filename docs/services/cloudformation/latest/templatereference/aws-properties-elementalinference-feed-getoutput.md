This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElementalInference::Feed GetOutput

Contains configuration information about one output in a feed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Name" : String,
  "OutputConfig" : OutputConfig,
  "Status" : String
}

```

### YAML

```yaml

  Description: String
  Name: String
  OutputConfig:
    OutputConfig
  Status: String

```

## Properties

`Description`

The description that you assign to this output.

_Required_: No

_Type_: String

_Pattern_: `^[\w \-\.',@:;]*$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The user-friendly name that you assign to this output.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]([a-zA-Z0-9-_]{0,126}[a-zA-Z0-9])?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputConfig`

Contains one typed output. Including a specific typed output enables the corresponding Elemental Inference in the feed.

_Required_: Yes

_Type_: [OutputConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elementalinference-feed-outputconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status that you assign to this output.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClippingConfig

OutputConfig
