This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition

The properties that describe a container group resource. You can update all properties of
a container group definition properties. Updates to a container group definition are saved as
new versions.

**Used with:** [CreateContainerGroupDefinition](../../../../reference/gamelift/latest/apireference/api-createcontainergroupdefinition.md)

**Returned by:** [DescribeContainerGroupDefinition](../../../../reference/gamelift/latest/apireference/api-describecontainergroupdefinition.md),
[ListContainerGroupDefinitions](../../../../reference/gamelift/latest/apireference/api-listcontainergroupdefinitions.md),
[UpdateContainerGroupDefinition](../../../../reference/gamelift/latest/apireference/api-updatecontainergroupdefinition.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::ContainerGroupDefinition",
  "Properties" : {
      "ContainerGroupType" : String,
      "GameServerContainerDefinition" : GameServerContainerDefinition,
      "Name" : String,
      "OperatingSystem" : String,
      "SourceVersionNumber" : Integer,
      "SupportContainerDefinitions" : [ SupportContainerDefinition, ... ],
      "Tags" : [ Tag, ... ],
      "TotalMemoryLimitMebibytes" : Integer,
      "TotalVcpuLimit" : Number,
      "VersionDescription" : String
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::ContainerGroupDefinition
Properties:
  ContainerGroupType: String
  GameServerContainerDefinition:
    GameServerContainerDefinition
  Name: String
  OperatingSystem: String
  SourceVersionNumber: Integer
  SupportContainerDefinitions:
    - SupportContainerDefinition
  Tags:
    - Tag
  TotalMemoryLimitMebibytes: Integer
  TotalVcpuLimit: Number
  VersionDescription: String

```

## Properties

`ContainerGroupType`

The type of container group. Container group type determines how Amazon GameLift Servers deploys the
container group on each fleet instance.

_Required_: No

_Type_: String

_Allowed values_: `GAME_SERVER | PER_INSTANCE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GameServerContainerDefinition`

The definition for the game server container in this group. This property is used only
when the container group type is `GAME_SERVER`. This container definition specifies
a container image with the game server build.

_Required_: No

_Type_: [GameServerContainerDefinition](aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A descriptive identifier for the container group definition. The name value is unique in an AWS Region.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OperatingSystem`

The platform that all containers in the container group definition run on.

###### Note

Amazon Linux 2 (AL2) will reach end of support on 6/30/2026. See more details in the [Amazon Linux 2 FAQs](https://aws.amazon.com/amazon-linux-2/faqs). For game
servers that are hosted on AL2 and use server SDK version 4.x for Amazon GameLift Servers, first update the game
server build to server SDK 5.x, and then deploy to AL2023 instances. See [Migrate to\
server SDK version 5.](../../../../reference/gamelift/latest/developerguide/reference-serversdk5-migration.md)

_Required_: Yes

_Type_: String

_Allowed values_: `AMAZON_LINUX_2023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceVersionNumber`

The container group definition version to update. The new version starts with values from
the source version, and then updates values included in this request.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportContainerDefinitions`

The set of definitions for support containers in this group. A container group definition
might have zero support container definitions. Support container can be used in any type of
container group.

_Required_: No

_Type_: Array of [SupportContainerDefinition](aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of labels to assign to the container group definition resource. Tags are
developer-defined key-value pairs. Tagging AWS resources are useful for resource management,
access management and cost allocation. For more information, see [Tagging AWS Resources](../../../../general/latest/gr/aws-tagging.md) in the
_AWS General Reference_.

_Required_: No

_Type_: Array of [Tag](aws-properties-gamelift-containergroupdefinition-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalMemoryLimitMebibytes`

The amount of memory (in MiB) on a fleet instance to allocate for the container group. All
containers in the group share these resources.

You can set a limit for each container definition in the group. If individual containers
have limits, this total value must be greater than any individual container's memory
limit.

_Required_: Yes

_Type_: Integer

_Minimum_: `4`

_Maximum_: `1024000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalVcpuLimit`

The amount of vCPU units on a fleet instance to allocate for the container group (1 vCPU
is equal to 1024 CPU units). All containers in the group share these resources. You can set a
limit for each container definition in the group. If individual containers have limits, this
total value must be equal to or greater than the sum of the limits for each container in the
group.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionDescription`

An optional description that was provided for a container group definition update. Each
version can have a unique description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ContainerGroupDefinitionArn`

The Amazon Resource Name ( [ARN](../../../s3/latest/dev/s3-arn-format.md)) that is assigned to an Amazon GameLift Servers `ContainerGroupDefinition` resource. It uniquely identifies
the resource across all AWS Regions. Format is
`arn:aws:gamelift:[region]::containergroupdefinition/[container group definition name]:[version]`.

`CreationTime`

A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example `"1469498468.057"`).

`Status`

Current status of the container group definition resource. Values include:

- `COPYING` \-\- Amazon GameLift Servers is in the process of making copies of all container
images that are defined in the group. While in this state, the resource can't be used to
create a container fleet.

- `READY` \-\- Amazon GameLift Servers has copied the registry images for all containers that
are defined in the group. You can use a container group definition in this status to
create a container fleet.

- `FAILED` \-\- Amazon GameLift Servers failed to create a valid container group definition
resource. For more details on the cause of the failure, see `StatusReason`. A
container group definition resource in failed status will be deleted within a few
minutes.

`StatusReason`

Additional information about a container group definition that's in `FAILED`
status. Possible reasons include:

- An internal issue prevented Amazon GameLift Servers from creating
the container group definition resource. Delete the failed resource and call
[CreateContainerGroupDefinition](../../../../reference/gamelift/latest/apireference/api-createcontainergroupdefinition.md) again.

- An access-denied message means that you don't have permissions to access the container image on ECR. See
[IAM permission examples](../../../gamelift/latest/developerguide/gamelift-iam-policy-examples.md)
for help setting up required IAM permissions for Amazon GameLift Servers.

- The `ImageUri` value for at least one
of the containers in the container group definition was invalid or not found in the current
AWS account.

- At least one
of the container images referenced in the container group definition exceeds the
allowed size. For size limits, see [Amazon GameLift Servers endpoints and quotas](../../../../general/latest/gr/gamelift.md).

- At least one of the container images referenced in the
container group definition uses a different operating system than the one defined for the container group.

`VersionNumber`

Indicates the version of a particular container group definition. This number is
incremented automatically when you update a container group definition. You can view, update,
or delete individual versions or the entire container group definition.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetConfiguration

ContainerDependency

All content copied from https://docs.aws.amazon.com/.
