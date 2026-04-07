This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::ReplicationConfig

Creates a configuration that you can later provide to configure and start an AWS DMS
Serverless replication. You can also provide options to validate the configuration inputs
before you start the replication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::ReplicationConfig",
  "Properties" : {
      "ComputeConfig" : ComputeConfig,
      "ReplicationConfigIdentifier" : String,
      "ReplicationSettings" : Json,
      "ReplicationType" : String,
      "ResourceIdentifier" : String,
      "SourceEndpointArn" : String,
      "SupplementalSettings" : Json,
      "TableMappings" : Json,
      "Tags" : [ Tag, ... ],
      "TargetEndpointArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::DMS::ReplicationConfig
Properties:
  ComputeConfig:
    ComputeConfig
  ReplicationConfigIdentifier: String
  ReplicationSettings: Json
  ReplicationType: String
  ResourceIdentifier: String
  SourceEndpointArn: String
  SupplementalSettings: Json
  TableMappings: Json
  Tags:
    - Tag
  TargetEndpointArn: String

```

## Properties

`ComputeConfig`

Configuration parameters for provisioning an AWS DMS Serverless replication.

_Required_: Yes

_Type_: [ComputeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-replicationconfig-computeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationConfigIdentifier`

A unique identifier that you want to use to create a `ReplicationConfigArn`
that is returned as part of the output from this action. You can then pass this output
`ReplicationConfigArn` as the value of the `ReplicationConfigArn`
option for other actions to identify both AWS DMS Serverless replications and replication
configurations that you want those actions to operate on. For some actions, you can also
use either this unique identifier or a corresponding ARN in action filters to identify the
specific replication and replication configuration to operate on.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationSettings`

Optional JSON settings for AWS DMS Serverless replications that are provisioned using this
replication configuration. For example, see [Change processing tuning settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.ChangeProcessingTuning.html).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationType`

The type of AWS DMS Serverless replication to provision using this replication
configuration.

Possible values:

- `"full-load"`

- `"cdc"`

- `"full-load-and-cdc"`

_Required_: Yes

_Type_: String

_Allowed values_: `full-load | full-load-and-cdc | cdc`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdentifier`

Optional unique value or name that you set for a given resource that can be used to
construct an Amazon Resource Name (ARN) for that resource. For more information, see [Fine-grained access control using resource names and\
tags](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#CHAP_Security.FineGrainedAccess).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceEndpointArn`

The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless
replication configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupplementalSettings`

Optional JSON settings for specifying supplemental data. For more information, see
[Specifying supplemental data for task settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.TaskData.html).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableMappings`

JSON table mappings for AWS DMS Serverless replications that are provisioned using this
replication configuration. For more information, see [Specifying table selection and transformations rules using\
JSON](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.SelectionTransformation.html).

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more optional tags associated with resources used by the AWS DMS Serverless
replication. For more information, see [Tagging resources in AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tagging.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-replicationconfig-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetEndpointArn`

The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS serverless
replication configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ReplicationConfigArn`

The Amazon Resource Name (ARN) of this AWS DMS Serverless replication
configuration.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ComputeConfig
