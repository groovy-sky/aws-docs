This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Instance Networking

`Networking` is a property of the [AWS::Lightsail::Instance](../userguide/aws-resource-lightsail-instance.md) resource. It describes the public ports and the
monthly amount of data transfer allocated for the instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MonthlyTransfer" : MonthlyTransfer,
  "Ports" : [ Port, ... ]
}

```

### YAML

```yaml

  MonthlyTransfer:
    MonthlyTransfer
  Ports:
    - Port

```

## Properties

`MonthlyTransfer`

The monthly amount of data transfer, in GB, allocated for the instance

_Required_: No

_Type_: [MonthlyTransfer](aws-properties-lightsail-instance-monthlytransfer.md)

_Update requires_: Updates are not supported.

`Ports`

An array of ports to open on the instance.

_Required_: Yes

_Type_: Array of [Port](aws-properties-lightsail-instance-port.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonthlyTransfer

Port

All content copied from https://docs.aws.amazon.com/.
