This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association DynatraceConfiguration

Configuration for Dynatrace monitoring integration. Defines the Dynatrace environment ID, list of resources to
monitor, and webhook update settings required for the Agent Space to access metrics, traces, and logs from
Dynatrace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableWebhookUpdates" : Boolean,
  "EnvId" : String,
  "Resources" : [ String, ... ]
}

```

### YAML

```yaml

  EnableWebhookUpdates: Boolean
  EnvId: String
  Resources:
    - String

```

## Properties

`EnableWebhookUpdates`

When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events
from the service.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvId`

The Dynatrace environment ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resources`

List of Dynatrace resources to monitor.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSResource

EventChannelConfiguration

All content copied from https://docs.aws.amazon.com/.
