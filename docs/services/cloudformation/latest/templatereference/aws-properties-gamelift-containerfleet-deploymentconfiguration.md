This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet DeploymentConfiguration

Set of rules for processing a deployment for a container fleet update.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImpairmentStrategy" : String,
  "MinimumHealthyPercentage" : Integer,
  "ProtectionStrategy" : String
}

```

### YAML

```yaml

  ImpairmentStrategy: String
  MinimumHealthyPercentage: Integer
  ProtectionStrategy: String

```

## Properties

`ImpairmentStrategy`

Determines what actions to take if a deployment fails. If the fleet is multi-location,
this strategy applies across all fleet locations. With a rollback strategy, updated
fleet instances are rolled back to the last successful deployment. Alternatively, you
can maintain a few impaired containers for the purpose of debugging, while all other
tasks return to the last successful deployment.

_Required_: No

_Type_: String

_Allowed values_: `MAINTAIN | ROLLBACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumHealthyPercentage`

Sets a minimum level of healthy tasks to maintain during deployment activity.

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Maximum_: `75`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtectionStrategy`

Determines how fleet deployment activity affects active game sessions on the fleet.
With protection, a deployment honors game session protection, and delays actions that
would interrupt a protected active game session until the game session ends. Without
protection, deployment activity can shut down all running tasks, including active game
sessions, regardless of game session protection.

_Required_: No

_Type_: String

_Allowed values_: `WITH_PROTECTION | IGNORE_PROTECTION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionPortRange

DeploymentDetails

All content copied from https://docs.aws.amazon.com/.
