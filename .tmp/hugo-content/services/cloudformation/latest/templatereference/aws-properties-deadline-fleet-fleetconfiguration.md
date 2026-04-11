This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet FleetConfiguration

Fleet configuration details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerManaged" : CustomerManagedFleetConfiguration,
  "ServiceManagedEc2" : ServiceManagedEc2FleetConfiguration
}

```

### YAML

```yaml

  CustomerManaged:
    CustomerManagedFleetConfiguration
  ServiceManagedEc2:
    ServiceManagedEc2FleetConfiguration

```

## Properties

`CustomerManaged`

The customer managed fleets within a fleet configuration.

_Required_: No

_Type_: [CustomerManagedFleetConfiguration](aws-properties-deadline-fleet-customermanagedfleetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceManagedEc2`

The service managed Amazon EC2 instances for a fleet configuration.

_Required_: No

_Type_: [ServiceManagedEc2FleetConfiguration](aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FleetCapabilities

HostConfiguration

All content copied from https://docs.aws.amazon.com/.
