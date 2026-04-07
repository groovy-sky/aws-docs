This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper Destination

Where to send the metrics from a scraper.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmpConfiguration" : AmpConfiguration
}

```

### YAML

```yaml

  AmpConfiguration:
    AmpConfiguration

```

## Properties

`AmpConfiguration`

The Amazon Managed Service for Prometheus workspace to send metrics to.

_Required_: No

_Type_: [AmpConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-ampconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ComponentConfig

EksConfiguration
