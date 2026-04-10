This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Queue

Creates an AWS PCS queue resource. For more information, see
[Creating a queue \
in AWS PCS](../../../pcs/latest/userguide/working-with-queues-create.md) in the _AWS PCS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCS::Queue",
  "Properties" : {
      "ClusterId" : String,
      "ComputeNodeGroupConfigurations" : [ ComputeNodeGroupConfiguration, ... ],
      "Name" : String,
      "SlurmConfiguration" : SlurmConfiguration,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::PCS::Queue
Properties:
  ClusterId: String
  ComputeNodeGroupConfigurations:
    - ComputeNodeGroupConfiguration
  Name: String
  SlurmConfiguration:
    SlurmConfiguration
  Tags:
    Key: Value

```

## Properties

`ClusterId`

The ID of the cluster of the queue.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComputeNodeGroupConfigurations`

The list of compute node group configurations associated with the queue. Queues assign
jobs to associated compute node groups.

_Required_: No

_Type_: Array of [ComputeNodeGroupConfiguration](aws-properties-pcs-queue-computenodegroupconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name that identifies the queue.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SlurmConfiguration`

Additional options related to the Slurm scheduler.

_Required_: No

_Type_: [SlurmConfiguration](aws-properties-pcs-queue-slurmconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

1 or more tags added to the resource. Each tag consists of a tag key and tag value. The
tag value is optional and can be an empty string.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The unique Amazon Resource Name (ARN) of the queue.

`ErrorInfo`

The list of errors that occurred during queue provisioning.

`Id`

The generated unique ID of the queue.

`Status`

The provisioning status of the queue.

###### Note

The provisioning status doesn't indicate the overall health of the queue.

###### Important

The resource enters the `SUSPENDING` and `SUSPENDED` states
when the scheduler is beyond end of life and we have suspended the cluster. When in
these states, you can't use the cluster. The cluster controller is down and all compute
instances are terminated. The resources still count toward your service quotas. You can
delete a resource if its status is `SUSPENDED`. For more information, see
[Frequently asked questions about Slurm versions in AWS PCS](../../../pcs/latest/userguide/slurm-versions-faq.md) in the
_AWS PCS User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpotOptions

ComputeNodeGroupConfiguration

All content copied from https://docs.aws.amazon.com/.
