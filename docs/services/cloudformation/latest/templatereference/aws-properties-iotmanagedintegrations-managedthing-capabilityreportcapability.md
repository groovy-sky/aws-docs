This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTManagedIntegrations::ManagedThing CapabilityReportCapability

The capability used in capability report.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ String, ... ],
  "Events" : [ String, ... ],
  "Id" : String,
  "Name" : String,
  "Properties" : [ String, ... ],
  "Version" : String
}

```

### YAML

```yaml

  Actions:
    - String
  Events:
    - String
  Id: String
  Name: String
  Properties:
    - String
  Version: String

```

## Properties

`Actions`

The capability actions used in the capability report.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `128 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Events`

The capability events used in the capability report.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `128 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The id of the schema version.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9.\/]+(@(\d+\.\d+|\$latest))?$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the capability.

_Required_: Yes

_Type_: String

_Pattern_: `^[/a-zA-Z0-9\._ ]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Properties`

The capability properties used in the capability
report.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `128 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the capability.

_Required_: Yes

_Type_: String

_Pattern_: `^(0|[1-9][0-9]*)$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapabilityReport

CapabilityReportEndpoint

All content copied from https://docs.aws.amazon.com/.
