This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityScanConfiguration ContinuousIntegrationScanConfiguration

Configuration settings for continuous integration scans that run automatically when
code changes are made.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "supportedEvents" : [ String, ... ]
}

```

### YAML

```yaml

  supportedEvents:
    - String

```

## Properties

`supportedEvents`

The repository events that trigger continuous integration scans, such as pull requests
or commits.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeSecurityScanConfiguration

PeriodicScanConfiguration

All content copied from https://docs.aws.amazon.com/.
