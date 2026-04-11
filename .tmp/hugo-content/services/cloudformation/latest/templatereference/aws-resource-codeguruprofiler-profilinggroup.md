This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeGuruProfiler::ProfilingGroup

Creates a profiling group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeGuruProfiler::ProfilingGroup",
  "Properties" : {
      "AgentPermissions" : AgentPermissions,
      "AnomalyDetectionNotificationConfiguration" : [ Channel, ... ],
      "ComputePlatform" : String,
      "ProfilingGroupName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeGuruProfiler::ProfilingGroup
Properties:
  AgentPermissions:
    AgentPermissions
  AnomalyDetectionNotificationConfiguration:
    - Channel
  ComputePlatform: String
  ProfilingGroupName: String
  Tags:
    - Tag

```

## Properties

`AgentPermissions`

The agent permissions attached to this profiling group. This action group grants
`ConfigureAgent` and `PostAgentProfile` permissions to perform
actions required by the profiling agent. The Json consists of key
`Principals`.

_Principals_: A list of string ARNs for the roles and users you want
to grant access to the profiling group. Wildcards are not supported in the ARNs. You are
allowed to provide up to 50 ARNs. An empty list is not permitted. This is a required key.

For more information, see [Resource-based policies\
in CodeGuru Profiler](../../../codeguru/latest/profiler-ug/resource-based-policies.md) in the _Amazon CodeGuru Profiler user_
_guide_, [ConfigureAgent](../../../codeguru/latest/profiler-api/api-configureagent.md), and
[PostAgentProfile](../../../codeguru/latest/profiler-api/api-postagentprofile.md).

_Required_: No

_Type_: [AgentPermissions](aws-properties-codeguruprofiler-profilinggroup-agentpermissions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnomalyDetectionNotificationConfiguration`

Adds anomaly notifications for a profiling group.

_Required_: No

_Type_: Array of [Channel](aws-properties-codeguruprofiler-profilinggroup-channel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputePlatform`

The compute platform of the profiling group. Use `AWSLambda` if your
application runs on AWS Lambda. Use `Default` if your application runs on a
compute platform that is not AWS Lambda, such an Amazon EC2 instance, an on-premises
server, or a different platform. If not specified, `Default` is used. This
property is immutable.

_Required_: No

_Type_: String

_Allowed values_: `Default | AWSLambda`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProfilingGroupName`

The name of the profiling group.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags to add to the created profiling group.

_Required_: No

_Type_: Array of [Tag](aws-properties-codeguruprofiler-profilinggroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the profiling group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The full Amazon Resource Name (ARN) for that profiling group.

## Examples

- [CodeGuru Profiler profiling group resource configuration](#aws-resource-codeguruprofiler-profilinggroup--examples--CodeGuru_Profiler_profiling_group_resource_configuration)

- [CodeGuru Profiler profiling group resource with compute platform](#aws-resource-codeguruprofiler-profilinggroup--examples--CodeGuru_Profiler_profiling_group_resource_with_compute_platform)

- [CodeGuru Profiler profiling group resource with notifications](#aws-resource-codeguruprofiler-profilinggroup--examples--CodeGuru_Profiler_profiling_group_resource_with_notifications)

- [CodeGuru Profiler profiling group configuration](#aws-resource-codeguruprofiler-profilinggroup--examples--CodeGuru_Profiler_profiling_group_configuration)

### CodeGuru Profiler profiling group resource configuration

The following is an example of the profiling group resource with the profiling
group name and agent permissions properties.

#### JSON

```json

"MyProfilingGroupWithAgentPermissions": {
  "Type": "AWS::CodeGuruProfiler::ProfilingGroup",
  "Properties": {
    "ProfilingGroupName": "MyProfilingGroup",
    "AgentPermissions": {
      "Principals": [
          "arn:aws:iam::1233456789012:role/agent-permissions-role-1",
          "arn:aws:iam::1233456789012:role/agent-permissions-role-2"
      ]
    }
  }
}

```

#### YAML

```yaml

MyProfilingGroupWithAgentPermissions:
  Type: AWS::CodeGuruProfiler::ProfilingGroup
  Properties:
    ProfilingGroupName: "MyProfilingGroup"
    AgentPermissions:
      Principals:
        - "arn:aws:iam::1233456789012:role/agent-permissions-role-1"
        - "arn:aws:iam::1233456789012:role/agent-permissions-role-2"
```

### CodeGuru Profiler profiling group resource with compute platform

The following is an example of the profiling group resource that runs on AWS Lambda.

#### JSON

```json

"MyProfilingGroupWithComputePlatform": {
  "Type": "AWS::CodeGuruProfiler::ProfilingGroup",
  "Properties": {
    "ProfilingGroupName": "MyProfilingGroup",
    "ComputePlatform": "AWSLambda"
  }
}
```

#### YAML

```yaml

MyProfilingGroupWithComputePlatform:
  Type: AWS::CodeGuruProfiler::ProfilingGroup
  Properties:
    ProfilingGroupName: "MyProfilingGroup"
    ComputePlatform: "AWSLambda"
```

### CodeGuru Profiler profiling group resource with notifications

The following is an example of the a notification configuration for a profiling
group.

#### JSON

```json

"MyProfilingGroupWithNotificationChannelConfiguration": {
  "Type": "AWS::CodeGuruProfiler::ProfilingGroup",
  "Properties": {
    "ProfilingGroupName": "MyProfilingGroup",
    "AnomalyDetectionNotificationConfiguration": [
        {
            "channelUri": "SOME_SNS_TOPIC_ARN",
            "channelId": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
        }
    ]
  }
}
```

#### YAML

```yaml

MyProfilingGroupWithNotificationChannelConfiguration:
  Type: AWS::CodeGuruProfiler::ProfilingGroup
  Properties:
    ProfilingGroupName: MyProfilingGroup
    AnomalyDetectionNotificationConfiguration:
    - channelUri: SOME_SNS_TOPIC_ARN
      channelId: aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee
```

### CodeGuru Profiler profiling group configuration

The following is an example of a profiling group that runs on AWS Lambda. This
profiling group has enabled agent permissions. Notifications have also been
configured with `AnomalyDetectionConfiguration`.

#### JSON

```json

"MyProfilingGroupWithAgentPermissions": {
  "Type": "AWS::CodeGuruProfiler::ProfilingGroup",
  "Properties": {
    "ProfilingGroupName": "MyProfilingGroup",
    "ComputePlatform": "AWSLambda",
    "AgentPermissions": {
      "Principals": [
          "arn:aws:iam::1233456789012:role/agent-permissions-role-1",
          "arn:aws:iam::1233456789012:role/agent-permissions-role-2"
      ]
    },
    "AnomalyDetectionNotificationConfiguration": [
        {
            "channelUri": "SOME_SNS_TOPIC_ARN",
            "channelId": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
        }
    ]
  }
}
```

#### YAML

```yaml

MyProfilingGroup:
  Type: AWS::CodeGuruProfiler::ProfilingGroup
  Properties:
    ProfilingGroupName: "MyProfilingGroup"
    ComputePlatform: "AWSLambda"
    AgentPermissions:
      Principals:
        - "arn:aws:iam::1233456789012:role/agent-permissions-role-1"
        - "arn:aws:iam::1233456789012:role/agent-permissions-role-2"
    AnomalyDetectionNotificationConfiguration:
    - channelUri: SOME_SNS_TOPIC_ARN
      channelId: aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon CodeGuru Profiler

AgentPermissions

All content copied from https://docs.aws.amazon.com/.
