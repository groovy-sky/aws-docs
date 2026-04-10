This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper ComponentConfig

Configuration settings for a scraper component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Options" : {Key: Value, ...}
}

```

### YAML

```yaml

  Options:
    Key: Value

```

## Properties

`Options`

Configuration options for the scraper component.

_Required_: No

_Type_: Object of String

_Pattern_: `^[A-Za-z0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogDestination

Destination

All content copied from https://docs.aws.amazon.com/.
