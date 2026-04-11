This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet ReputationOptions

Enable or disable collection of reputation metrics for emails that you send using this
configuration set in the current AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReputationMetricsEnabled" : Boolean
}

```

### YAML

```yaml

  ReputationMetricsEnabled: Boolean

```

## Properties

`ReputationMetricsEnabled`

If `true`, tracking of reputation metrics is enabled for the configuration
set. If `false`, tracking of reputation metrics is disabled for the
configuration set.

_Required_: No

_Type_: Boolean

_Pattern_: `true|false`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OverallConfidenceThreshold

SendingOptions

All content copied from https://docs.aws.amazon.com/.
