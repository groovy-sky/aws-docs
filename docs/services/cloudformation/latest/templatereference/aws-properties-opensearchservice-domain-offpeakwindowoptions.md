This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain OffPeakWindowOptions

Off-peak window settings for the domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "OffPeakWindow" : OffPeakWindow
}

```

### YAML

```yaml

  Enabled: Boolean
  OffPeakWindow:
    OffPeakWindow

```

## Properties

`Enabled`

Specifies whether off-peak window settings are enabled for the domain.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OffPeakWindow`

Off-peak window settings for the domain.

_Required_: No

_Type_: [OffPeakWindow](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchservice-domain-offpeakwindow.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OffPeakWindow

S3VectorsEngine
