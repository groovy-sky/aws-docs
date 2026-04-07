This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::PartnerApp PartnerAppMaintenanceConfig

A collection of settings that specify the maintenance schedule for the PartnerApp.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaintenanceWindowStart" : String
}

```

### YAML

```yaml

  MaintenanceWindowStart: String

```

## Properties

`MaintenanceWindowStart`

The maintenance window start day and time for the PartnerApp.

_Required_: Yes

_Type_: String

_Pattern_: `(Mon|Tue|Wed|Thu|Fri|Sat|Sun):([01]\d|2[0-3]):([0-5]\d)`

_Maximum_: `9`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PartnerAppConfig

Tag
