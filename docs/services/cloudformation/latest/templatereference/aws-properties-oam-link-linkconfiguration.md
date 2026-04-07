This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Oam::Link LinkConfiguration

Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from
the source account to the monitoring account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupConfiguration" : LinkFilter,
  "MetricConfiguration" : LinkFilter
}

```

### YAML

```yaml

  LogGroupConfiguration:
    LinkFilter
  MetricConfiguration:
    LinkFilter

```

## Properties

`LogGroupConfiguration`

Use this structure to filter which log groups are to
share log events from this source account to the monitoring account.

_Required_: No

_Type_: [LinkFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-oam-link-linkfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricConfiguration`

Use this structure to filter which metric namespaces are to be shared from
the source account to the monitoring account.

_Required_: No

_Type_: [LinkFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-oam-link-linkfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Oam::Link

LinkFilter
