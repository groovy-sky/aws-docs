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

_Type_: [CustomerManagedFleetConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceManagedEc2`

The service managed Amazon EC2 instances for a fleet configuration.

_Required_: No

_Type_: [ServiceManagedEc2FleetConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FleetCapabilities

HostConfiguration
