This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper Source

The source of collected metrics for a scraper.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EksConfiguration" : EksConfiguration,
  "VpcConfiguration" : VpcConfiguration
}

```

### YAML

```yaml

  EksConfiguration:
    EksConfiguration
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`EksConfiguration`

The Amazon EKS cluster from which a scraper collects metrics.

_Required_: No

_Type_: [EksConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-eksconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConfiguration`

Property description not available.

_Required_: No

_Type_: [VpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-vpcconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ScraperLoggingDestination

Tag
