This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Retriever RetrieverConfiguration

Provides information on how the retriever used for your Amazon Q Business application is
configured.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KendraIndexConfiguration" : KendraIndexConfiguration,
  "NativeIndexConfiguration" : NativeIndexConfiguration
}

```

### YAML

```yaml

  KendraIndexConfiguration:
    KendraIndexConfiguration
  NativeIndexConfiguration:
    NativeIndexConfiguration

```

## Properties

`KendraIndexConfiguration`

Provides information on how the Amazon Kendra index used as a retriever for your
Amazon Q Business application is configured.

_Required_: No

_Type_: [KendraIndexConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-retriever-kendraindexconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NativeIndexConfiguration`

Provides information on how a Amazon Q Business index used as a retriever for your
Amazon Q Business application is configured.

_Required_: No

_Type_: [NativeIndexConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-retriever-nativeindexconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NativeIndexConfiguration

Tag
