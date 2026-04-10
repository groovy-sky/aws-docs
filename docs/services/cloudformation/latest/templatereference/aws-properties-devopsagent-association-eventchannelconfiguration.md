This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association EventChannelConfiguration

Configuration for Event Channel integration. Defines webhook update settings to enable the Agent Space to
receive real-time event notifications from event channel integrations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableWebhookUpdates" : Boolean
}

```

### YAML

```yaml

  EnableWebhookUpdates: Boolean

```

## Properties

`EnableWebhookUpdates`

When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events
from the service.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynatraceConfiguration

GitHubConfiguration

All content copied from https://docs.aws.amazon.com/.
