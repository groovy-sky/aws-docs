This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet ConnectionPortRange

The set of port numbers to open on each instance in a container fleet. Connection
ports are used by inbound traffic to connect with processes that are running in
containers on the fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FromPort" : Integer,
  "ToPort" : Integer
}

```

### YAML

```yaml

  FromPort: Integer
  ToPort: Integer

```

## Properties

`FromPort`

Starting value for the port range.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToPort`

Ending value for the port. Port numbers are end-inclusive. This value must be equal to
or greater than `FromPort`.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GameLift::ContainerFleet

DeploymentConfiguration
