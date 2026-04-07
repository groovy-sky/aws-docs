This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTManagedIntegrations::ManagedThing CapabilityReport

A report of the capabilities for the managed thing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Endpoints" : [ CapabilityReportEndpoint, ... ],
  "NodeId" : String,
  "Version" : String
}

```

### YAML

```yaml

  Endpoints:
    - CapabilityReportEndpoint
  NodeId: String
  Version: String

```

## Properties

`Endpoints`

The endpoints used in the capability report.

_Required_: Yes

_Type_: Array of [CapabilityReportEndpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotmanagedintegrations-managedthing-capabilityreportendpoint.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeId`

The numeric identifier of the node.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9=_.,@\+\-/]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the capability report.

_Required_: Yes

_Type_: String

_Pattern_: `^1\.0\.0$`

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoTManagedIntegrations::ManagedThing

CapabilityReportCapability
