This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain OffPeakWindow

A custom 10-hour, low-traffic window during which OpenSearch Service can perform mandatory configuration changes on the domain.
These actions can include scheduled service software updates and blue/green Auto-Tune enhancements. OpenSearch Service will
schedule these actions during the window that you specify. If you don't specify a window start time, it defaults to 10:00 P.M. local time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WindowStartTime" : WindowStartTime
}

```

### YAML

```yaml

  WindowStartTime:
    WindowStartTime

```

## Properties

`WindowStartTime`

The desired start time for an off-peak maintenance window.

_Required_: No

_Type_: [WindowStartTime](aws-properties-opensearchservice-domain-windowstarttime.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeToNodeEncryptionOptions

OffPeakWindowOptions

All content copied from https://docs.aws.amazon.com/.
