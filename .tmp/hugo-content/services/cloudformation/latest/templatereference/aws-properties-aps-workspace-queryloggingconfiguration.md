This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace QueryLoggingConfiguration

The query logging configuration in an Amazon Managed Service for Prometheus
workspace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destinations" : [ LoggingDestination, ... ]
}

```

### YAML

```yaml

  Destinations:
    - LoggingDestination

```

## Properties

`Destinations`

Defines a destination and its associated filtering criteria for query logging.

_Required_: Yes

_Type_: Array of [LoggingDestination](aws-properties-aps-workspace-loggingdestination.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingFilter

Tag

All content copied from https://docs.aws.amazon.com/.
