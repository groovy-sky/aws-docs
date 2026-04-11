This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule LogGroupNameConfiguration

Configuration that specifies a naming pattern for destination log groups created during centralization.
The pattern supports static text and dynamic variables that are replaced with source attributes
when log groups are created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupNamePattern" : String
}

```

### YAML

```yaml

  LogGroupNamePattern: String

```

## Properties

`LogGroupNamePattern`

The pattern used to generate destination log group names during centralization.
The pattern can contain static text and dynamic variables that are replaced with source attributes.
If a variable cannot be resolved, it inherits the value from its parent variable in the hierarchy.
The pattern must be between 1 and 512 characters.

Supported variables:

- **${source.logGroup}** — The original log group name from the source account.

- **${source.accountId}** — The AWS account ID where the log originated.

- **${source.region}** — The AWS Region where the log originated.

- **${source.org.id}** — The AWS Organization ID of the source account.

- **${source.org.ouId}** — The organizational unit ID of the source account.

- **${source.org.rootId}** — The organization Root ID.

- **${source.org.path}** — The organizational path from account to root.

_Required_: Yes

_Type_: String

_Pattern_: `^(?:[\._\-/#A-Za-z0-9]+|\$\{[A-Za-z]+(?:\.[A-Za-z]+){1,2}\})+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationLogsConfiguration

LogsBackupConfiguration

All content copied from https://docs.aws.amazon.com/.
