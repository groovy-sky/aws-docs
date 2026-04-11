This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association GitLabConfiguration

Configuration for GitLab project integration. Defines the numeric project ID, full project path
(namespace/project-name), GitLab instance identifier, and webhook update settings required for the Agent Space to
access and interact with the GitLab project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableWebhookUpdates" : Boolean,
  "InstanceIdentifier" : String,
  "ProjectId" : String,
  "ProjectPath" : String
}

```

### YAML

```yaml

  EnableWebhookUpdates: Boolean
  InstanceIdentifier: String
  ProjectId: String
  ProjectPath: String

```

## Properties

`EnableWebhookUpdates`

When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events
from the service.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceIdentifier`

The GitLab instance identifier (for example, `gitlab.com`).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectId`

GitLab numeric project ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectPath`

The full GitLab project path (for example, `namespace/project-name`).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitHubConfiguration

KeyValuePair

All content copied from https://docs.aws.amazon.com/.
