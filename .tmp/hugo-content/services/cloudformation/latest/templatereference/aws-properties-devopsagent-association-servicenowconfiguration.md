This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association ServiceNowConfiguration

Configuration for ServiceNow integration. Defines the ServiceNow instance URL, instance ID, and webhook update
settings required for the Agent Space to create, update, and manage incidents and change requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableWebhookUpdates" : Boolean,
  "InstanceId" : String
}

```

### YAML

```yaml

  EnableWebhookUpdates: Boolean
  InstanceId: String

```

## Properties

`EnableWebhookUpdates`

When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events
from the service.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceId`

ServiceNow instance ID.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceConfiguration

SlackChannel

All content copied from https://docs.aws.amazon.com/.
