This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTarget Targets

The `Targets` property type specifies adding a target to a maintenance
window target in AWS Systems Manager.

`Targets` is a property of the [AWS::SSM::MaintenanceWindowTarget](../userguide/aws-resource-ssm-maintenancewindowtarget.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Values:
    - String

```

## Properties

`Key`

User-defined criteria for sending commands that target managed nodes that meet the
criteria.

_Required_: Yes

_Type_: String

_Pattern_: `^[\p{L}\p{Z}\p{N}_.:/=\-@]*$|resource-groups:ResourceTypeFilters|resource-groups:Name`

_Minimum_: `1`

_Maximum_: `163`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

User-defined criteria that maps to `Key`. For example, if you specified
`tag:ServerRole`, you could specify `value:WebServer` to run a command on
instances that include EC2 tags of `ServerRole,WebServer`.

Depending on the type of target, the maximum number of values for a key might be lower than
the global maximum of 50.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSM::MaintenanceWindowTarget

AWS::SSM::MaintenanceWindowTask

All content copied from https://docs.aws.amazon.com/.
