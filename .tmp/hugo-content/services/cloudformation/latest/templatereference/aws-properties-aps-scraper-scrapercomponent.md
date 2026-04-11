This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper ScraperComponent

A component of a Amazon Managed Service for Prometheus scraper that can be configured for
logging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Config" : ComponentConfig,
  "Type" : String
}

```

### YAML

```yaml

  Config:
    ComponentConfig
  Type: String

```

## Properties

`Config`

The configuration settings for the scraper component.

_Required_: No

_Type_: [ComponentConfig](aws-properties-aps-scraper-componentconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the scraper component.

_Required_: Yes

_Type_: String

_Allowed values_: `SERVICE_DISCOVERY | COLLECTOR | EXPORTER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScrapeConfiguration

ScraperLoggingConfiguration

All content copied from https://docs.aws.amazon.com/.
