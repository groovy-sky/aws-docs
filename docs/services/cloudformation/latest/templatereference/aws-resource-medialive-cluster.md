This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Cluster

The `AWS::MediaLive::Cluster` resource Property description not available. for MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::Cluster",
  "Properties" : {
      "ClusterType" : String,
      "InstanceRoleArn" : String,
      "Name" : String,
      "NetworkSettings" : ClusterNetworkSettings,
      "Tags" : [ Tags, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::Cluster
Properties:
  ClusterType: String
  InstanceRoleArn: String
  Name: String
  NetworkSettings:
    ClusterNetworkSettings
  Tags:
    - Tags

```

## Properties

`ClusterType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `ON_PREMISES | OUTPOSTS_RACK | OUTPOSTS_SERVER | EC2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceRoleArn`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^arn:.+:iam:.+:role/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkSettings`

Property description not available.

_Required_: No

_Type_: [ClusterNetworkSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-cluster-clusternetworksettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-cluster-tags.html) of [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-cluster-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN of the cluster.

`ChannelIds`

The MediaLive channels that are currently running on nodes in this cluster.

`Id`

The unique ID of the cluster.

`State`

The current state of the cluster.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MediaLive::CloudWatchAlarmTemplateGroup

ClusterNetworkSettings
