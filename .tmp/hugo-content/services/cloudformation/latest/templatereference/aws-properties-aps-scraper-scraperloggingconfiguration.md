This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper ScraperLoggingConfiguration

The `ScraperLoggingConfiguration` property type specifies Property description not available. for an [AWS::APS::Scraper](aws-resource-aps-scraper.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LoggingDestination" : ScraperLoggingDestination,
  "ScraperComponents" : [ ScraperComponent, ... ]
}

```

### YAML

```yaml

  LoggingDestination:
    ScraperLoggingDestination
  ScraperComponents:
    - ScraperComponent

```

## Properties

`LoggingDestination`

Property description not available.

_Required_: Yes

_Type_: [ScraperLoggingDestination](aws-properties-aps-scraper-scraperloggingdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScraperComponents`

Property description not available.

_Required_: Yes

_Type_: Array of [ScraperComponent](aws-properties-aps-scraper-scrapercomponent.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScraperComponent

ScraperLoggingDestination

All content copied from https://docs.aws.amazon.com/.
