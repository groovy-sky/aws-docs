This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper ScrapeConfiguration

A scrape configuration for a scraper, base 64 encoded. For more information, see
[Scraper configuration](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html#AMP-collector-configuration) in the _Amazon Managed Service for_
_Prometheus User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfigurationBlob" : String
}

```

### YAML

```yaml

  ConfigurationBlob: String

```

## Properties

`ConfigurationBlob`

The base 64 encoded scrape configuration file.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RoleConfiguration

ScraperComponent
