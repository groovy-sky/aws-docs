This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SimSpaceWeaver::Simulation

Use the `AWS::SimSpaceWeaver::Simulation` resource to specify
a simulation that CloudFormation starts in the AWS Cloud, in
your AWS account. In the resource properties section of your template,
provide the name of an existing IAM role
configured with the proper permissions, and the name of an existing Amazon S3 bucket.
Your account must have permissions to read the Amazon S3 bucket.
The Amazon S3 bucket must contain a valid schema. The schema must refer to
simulation assets that are already uploaded to the AWS Cloud. For more information,
see the [detailed tutorial](https://docs.aws.amazon.com/simspaceweaver/latest/userguide/getting-started_detailed.html) in the _AWSSimSpace Weaver User Guide_.

Specify a `SnapshotS3Location` to start a simulation from a snapshot instead of from
a schema. When you start a simulation from a snapshot, SimSpace Weaver initializes the entity
data in the State Fabric with data saved in the snapshot, starts the spatial and service apps that
were running when the snapshot was created, and restores the clock to the appropriate tick. Your app
zip files must be in the same location in Amazon S3 as they were in for the original simulation.
You must start any custom apps separately. For more information about snapshots, see
[Snapshots](https://docs.aws.amazon.com/simspaceweaver/latest/userguide/working-with_snapshots.html)
in the _AWSSimSpace Weaver User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SimSpaceWeaver::Simulation",
  "Properties" : {
      "MaximumDuration" : String,
      "Name" : String,
      "RoleArn" : String,
      "SchemaS3Location" : S3Location,
      "SnapshotS3Location" : S3Location
    }
}

```

### YAML

```yaml

Type: AWS::SimSpaceWeaver::Simulation
Properties:
  MaximumDuration: String
  Name: String
  RoleArn: String
  SchemaS3Location:
    S3Location
  SnapshotS3Location:
    S3Location

```

## Properties

`MaximumDuration`

The maximum running time of the simulation,
specified as a number of minutes (m or M), hours (h or H), or days (d or D). The simulation
stops when it reaches this limit. The maximum value is `14D`, or its equivalent in the
other units. The default value is `14D`. A value equivalent to `0` makes the
simulation immediately transition to `STOPPING` as soon as it reaches `STARTED`.

_Required_: No

_Type_: String

_Minimum_: `2`

_Maximum_: `6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the simulation.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\-]{1,2048}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role
that the simulation assumes to perform actions.
For more information about ARNs, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
in the _AWS General Reference_.
For more information about IAM roles,
see [IAM\
roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the _AWS Identity and Access Management User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaS3Location`

The location of the simulation schema in Amazon Simple Storage Service (Amazon S3).
For more information about Amazon S3, see the
[_Amazon Simple Storage Service User Guide_](../../../s3/latest/userguide/welcome.md).

Provide a `SchemaS3Location` to start your simulation from a schema.

If you provide a `SchemaS3Location` then you can't provide a `SnapshotS3Location`.

_Required_: No

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-simspaceweaver-simulation-s3location.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotS3Location`

The location of the snapshot in Amazon Simple Storage Service (Amazon S3).
For more information about Amazon S3, see the
[_Amazon Simple Storage Service User Guide_](../../../s3/latest/userguide/welcome.md).

Provide a `SnapshotS3Location` to start your simulation from a snapshot.

If you provide a `SnapshotS3Location` then you can't provide a `SchemaS3Location`.

_Required_: No

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-simspaceweaver-simulation-s3location.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the `Simulation`. For example,
`MyTestSimulation_22-12-15_12_00_00`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DescribePayload`

The JSON blob that the [DescribeSimulation](https://docs.aws.amazon.com/simspaceweaver/latest/APIReference/API_DescribeSimulation.html) action returns.

## Examples

- [Start a simulation](#aws-resource-simspaceweaver-simulation--examples--Start_a_simulation)

- [Start a simulation with assets that you created with the app SDK scripts - version 1.12.x](#aws-resource-simspaceweaver-simulation--examples--Start_a_simulation_with_assets_that_you_created_with_the_app_SDK_scripts_-_version_1.12.x)

- [Start a simulation with assets that you created with the app SDK scripts - version 1.13.x](#aws-resource-simspaceweaver-simulation--examples--Start_a_simulation_with_assets_that_you_created_with_the_app_SDK_scripts_-_version_1.13.x)

- [Start a simulation with a 1 hour maximum duration](#aws-resource-simspaceweaver-simulation--examples--Start_a_simulation_with_a_1_hour_maximum_duration)

- [Start a simulation from a snapshot](#aws-resource-simspaceweaver-simulation--examples--Start_a_simulation_from_a_snapshot)

### Start a simulation

The following example specifies a simulation resource in the AWS Cloud.

#### JSON

```json

{
  "Description" : "Example - Start a simulation",
  "Resources" : {
    "MyTestSimulation" : {
      "Type" : "AWS::SimSpaceWeaver::Simulation",
      "Properties" : {
        "Name" : "MyTestSimulation",
        "RoleArn" : "arn:aws:iam::111122223333:role/my-test-simulation-app-role",
        "SchemaS3Location" : {
          "BucketName" : "MyTestSimulationBucket",
          "ObjectKey" : "MyTestSimulation-schema.yaml"
        }
      }
    }
  }
}

```

#### YAML

```yaml

Description: "Example - Start a simulation"
Resources:
  MyTestSimulation:
    Type: "AWS::SimSpaceWeaver::Simulation"
    Properties:
      Name: "MyTestSimulation"
      RoleArn: "arn:aws:iam::111122223333:role/my-test-simulation-app-role"
      SchemaS3Location:
        BucketName: "MyTestSimulationBucket"
        ObjectKey: "MyTestSimulation-schema.yaml"

```

### Start a simulation with assets that you created with the app SDK scripts - version 1.12.x

The following example specifies a simulation resource in the AWS Cloud
with assets that you created with the app SDK scripts (version 1.12.x).
The version 1.12.x scripts upload your app zips and schema to separate buckets for each project.
In this example, the `create-project` script used `--name MyTestSimulation`
and uploaded to `us-west-2`.
For more information, see the
[detailed\
tutorial](https://docs.aws.amazon.com/simspaceweaver/latest/userguide/getting-started_detailed.html) in the _AWS SimSpace Weaver User Guide_.

#### JSON

```json

{
  "Description" : "Example - Start an app SDK project simulation v.1.12.x",
  "Resources" : {
    "MyTestSimulation" : {
      "Type" : "AWS::SimSpaceWeaver::Simulation",
      "Properties" : {
        "Name" : "MyTestSimulation_22-12-15_12_00_00",
        "RoleArn" : "arn:aws:iam::111122223333:role/weaver-MyTestSimulation-app-role",
        "SchemaS3Location" : {
          "BucketName" : "weaver-mytestsimulation-111122223333-schemas-us-west-2",
          "ObjectKey" : "MyTestSimulation-schema.yaml"
        }
      }
    }
  }
}

```

#### YAML

```yaml

Description: "Example - Start an app SDK project simulation v.1.12.x"
Resources:
  MyTestSimulation:
    Type: "AWS::SimSpaceWeaver::Simulation"
    Properties:
      Name: "MyTestSimulation_22-12-15_12_00_00"
      RoleArn: "arn:aws:iam::111122223333:role/weaver-MyTestSimulation-app-role"
      SchemaS3Location:
        BucketName: "weaver-mytestsimulation-111122223333-schemas-us-west-2"
        ObjectKey: "MyTestSimulation-schema.yaml"

```

### Start a simulation with assets that you created with the app SDK scripts - version 1.13.x

The following example specifies a simulation resource in the AWS Cloud
with assets that you created with the app SDK scripts (version 1.13.x).
The version 1.13.x scripts upload your project artifacts to a single bucket for each project.
In this example, the `create-project` script used `--name MyTestSimulation`
and uploaded to `us-west-2`.
For more information, see the
[detailed \
tutorial](https://docs.aws.amazon.com/simspaceweaver/latest/userguide/getting-started_detailed.html) in the _AWS SimSpace Weaver User Guide_.

#### JSON

```json

{
  "Description" : "Example - Start an app SDK project simulation v.1.13.x",
  "Resources" : {
    "MyTestSimulation" : {
      "Type" : "AWS::SimSpaceWeaver::Simulation",
      "Properties" : {
        "Name" : "MyTestSimulation_22-12-15_12_00_00",
        "RoleArn" : "arn:aws:iam::111122223333:role/weaver-MyTestSimulation-app-role",
        "SchemaS3Location" : {
          "BucketName" : "weaver-mytestsimulation-111122223333-artifacts-us-west-2",
          "ObjectKey" : "MyTestSimulation-schema.yaml"
        }
      }
    }
  }
}

```

#### YAML

```yaml

Description: "Example - Start an app SDK project simulation v.1.13.x"
Resources:
  MyTestSimulation:
    Type: "AWS::SimSpaceWeaver::Simulation"
    Properties:
      Name: "MyTestSimulation_22-12-15_12_00_00"
      RoleArn: "arn:aws:iam::111122223333:role/weaver-MyTestSimulation-app-role"
      SchemaS3Location:
        BucketName: "weaver-mytestsimulation-111122223333-artifacts-us-west-2"
        ObjectKey: "MyTestSimulation-schema.yaml"

```

### Start a simulation with a 1 hour maximum duration

The following example adds a 1 hour `MaximumDuration` to
the previous example. This simulation will automatically stop after it
reaches the maximum duration.

#### JSON

```json

{
  "Description" : "Example - Start a simulation with 1H maximum duration",
  "Resources" : {
    "MyTestSimulation" : {
      "Type" : "AWS::SimSpaceWeaver::Simulation",
      "Properties" : {
        "MaximumDuration" : "1H",
        "Name" : "MyTestSimulation_22-12-15_12_00_00",
        "RoleArn" : "arn:aws:iam::111122223333:role/weaver-MyTestSimulation-app-role",
        "SchemaS3Location" : {
          "BucketName" : "weaver-mytestsimulation-111122223333-artifacts-us-west-2",
          "ObjectKey" : "MyTestSimulation-schema.yaml"
        }
      }
    }
  }
}

```

#### YAML

```yaml

Description: "Example - Start a simulation with 1H maximum duration"
Resources:
  MyTestSimulation:
    Type: "AWS::SimSpaceWeaver::Simulation"
    Properties:
      MaximumDuration: "1H"
      Name: "MyTestSimulation_22-12-15_12_00_00"
      RoleArn: "arn:aws:iam::111122223333:role/weaver-MyTestSimulation-app-role"
      SchemaS3Location:
        BucketName: "weaver-mytestsimulation-111122223333-artifacts-us-west-2"
        ObjectKey: "MyTestSimulation-schema.yaml"

```

### Start a simulation from a snapshot

The following example starts a simulation from a snapshot
instead of from a schema. The snapshot in this example was created
from the simulation in the previous example.
Note that the `MaximumDuration` setting isn't preserved
in the snapshot. This simulation is a new simulation
with a new `MaximumDuration` of 2 days.

We recommend that you make and use a copy of your original
simulation's app role. The original simulation's app role could
be deleted if you delete that simulation's CloudFormation
stack.

#### JSON

```json

{
  "Description" : "Example - Start a simulation from a snapshot",
  "Resources" : {
    "MyTestSimulation" : {
      "Type" : "AWS::SimSpaceWeaver::Simulation",
      "Properties" : {
        "MaximumDuration" : "2D",
        "Name" : "MyTestSimulation_from_snapshot",
        "RoleArn" : "arn:aws:iam::111122223333:role/weaver-MyTestSimulation-app-role-copy",
        "SnapshotS3Location" : {
          "BucketName" : "weaver-mytestsimulation-111122223333-artifacts-us-west-2",
          "ObjectKey" : "snapshot/MyTestSimulation_22-12-15_12_00_00-230428-1207-13.zip"
        }
      }
    }
  }
}

```

#### YAML

```yaml

Description: "Example - Start a simulation from a snapshot"
Resources:
  MyTestSimulation:
    Type: "AWS::SimSpaceWeaver::Simulation"
    Properties:
      MaximumDuration: "2D"
      Name: "MyTestSimulation_from_snapshot"
      RoleArn: "arn:aws:iam::111122223333:role/weaver-MyTestSimulation-app-role-copy"
      SnapshotS3Location:
        BucketName: "weaver-mytestsimulation-111122223333-artifacts-us-west-2"
        ObjectKey: "snapshot/MyTestSimulation_22-12-15_12_00_00-230428-1207-13.zip"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS SimSpace Weaver

S3Location
