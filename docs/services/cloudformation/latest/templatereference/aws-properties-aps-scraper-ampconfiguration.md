This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper AmpConfiguration

The `AmpConfiguration` structure defines the Amazon Managed Service for
Prometheus instance a scraper should send metrics to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WorkspaceArn" : String
}

```

### YAML

```yaml

  WorkspaceArn: String

```

## Properties

`WorkspaceArn`

ARN of the Amazon Managed Service for Prometheus workspace.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z]*:aps:[-a-z0-9]+:[0-9]{12}:workspace/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::APS::Scraper

CloudWatchLogDestination

All content copied from https://docs.aws.amazon.com/.
