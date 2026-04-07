This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster

Creates an AWS PCS cluster resource. For more information, see
[Creating a cluster \
in AWS Parallel Computing Service](https://docs.aws.amazon.com/pcs/latest/userguide/working-with_clusters_create.html) in the _AWS PCS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCS::Cluster",
  "Properties" : {
      "Name" : String,
      "Networking" : Networking,
      "Scheduler" : Scheduler,
      "Size" : String,
      "SlurmConfiguration" : SlurmConfiguration,
      "Tags" : String
    }
}

```

### YAML

```yaml

Type: AWS::PCS::Cluster
Properties:
  Name: String
  Networking:
    Networking
  Scheduler:
    Scheduler
  Size: String
  SlurmConfiguration:
    SlurmConfiguration
  Tags: String

```

## Properties

`Name`

The name that identifies the cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Networking`

The networking configuration for the cluster's control plane.

_Required_: Yes

_Type_: [Networking](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pcs-cluster-networking.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scheduler`

The cluster management and job scheduling software associated with the cluster.

_Required_: Yes

_Type_: [Scheduler](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pcs-cluster-scheduler.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Size`

The size of the cluster.

- `SMALL`: 32 compute nodes and 256 jobs

- `MEDIUM`: 512 compute nodes and 8192 jobs

- `LARGE`: 2048 compute nodes and 16,384 jobs

_Required_: Yes

_Type_: String

_Allowed values_: `SMALL | MEDIUM | LARGE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SlurmConfiguration`

Additional options related to the Slurm scheduler.

_Required_: No

_Type_: [SlurmConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pcs-cluster-slurmconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

1 or more tags added to the resource. Each tag consists of a tag key and tag value. The
tag value is optional and can be an empty string.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The unique Amazon Resource Name (ARN) of the cluster.

`Endpoints`

The list of endpoints available for interaction with the scheduler.

`ErrorInfo`

The list of errors that occurred during cluster provisioning.

`Id`

The generated unique ID of the cluster.

`Status`

The provisioning status of the cluster.

###### Note

The provisioning status doesn't indicate the overall health of the cluster.

###### Important

The resource enters the `SUSPENDING` and `SUSPENDED` states
when the scheduler is beyond end of life and we have suspended the cluster. When in
these states, you can't use the cluster. The cluster controller is down and all compute
instances are terminated. The resources still count toward your service quotas. You can
delete a resource if its status is `SUSPENDED`. For more information, see
[Frequently asked questions about Slurm versions in AWS PCS](https://docs.aws.amazon.com/pcs/latest/userguide/slurm-versions_faq.html) in the
_AWS PCS User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS PCS

Accounting
