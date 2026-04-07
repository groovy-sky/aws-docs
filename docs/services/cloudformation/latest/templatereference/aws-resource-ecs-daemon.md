This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Daemon

Information about a daemon resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECS::Daemon",
  "Properties" : {
      "CapacityProviderArns" : [ String, ... ],
      "ClusterArn" : String,
      "DaemonName" : String,
      "DaemonTaskDefinitionArn" : String,
      "DeploymentConfiguration" : DaemonDeploymentConfiguration,
      "EnableECSManagedTags" : Boolean,
      "EnableExecuteCommand" : Boolean,
      "PropagateTags" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ECS::Daemon
Properties:
  CapacityProviderArns:
    - String
  ClusterArn: String
  DaemonName: String
  DaemonTaskDefinitionArn: String
  DeploymentConfiguration:
    DaemonDeploymentConfiguration
  EnableECSManagedTags: Boolean
  EnableExecuteCommand: Boolean
  PropagateTags: String
  Tags:
    - Tag

```

## Properties

`CapacityProviderArns`

The Amazon Resource Names (ARNs) of the capacity providers associated with the
daemon.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterArn`

The Amazon Resource Name (ARN) of the cluster that the daemon is running in.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DaemonName`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DaemonTaskDefinitionArn`

The Amazon Resource Name (ARN) of the daemon task definition used by this
revision.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentConfiguration`

The deployment configuration used for this daemon deployment.

_Required_: No

_Type_: [DaemonDeploymentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-daemon-daemondeploymentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableECSManagedTags`

Specifies whether Amazon ECS managed tags are turned on for the daemon tasks.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableExecuteCommand`

Specifies whether the execute command functionality is turned on for the daemon
tasks.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropagateTags`

Specifies whether tags are propagated from the daemon to the daemon tasks.

_Required_: No

_Type_: String

_Allowed values_: `DAEMON | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-daemon-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The Unix timestamp for the time when the daemon was created.

`DaemonArn`

The Amazon Resource Name (ARN) of the daemon.

`DaemonStatus`

The status of the daemon.

`DeploymentArn`

The Amazon Resource Name (ARN) of the most recent daemon deployment.

`UpdatedAt`

The Unix timestamp for the time when the daemon was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CapacityProviderStrategy

DaemonAlarmConfiguration
