This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application CustomArtifactConfiguration

The configuration of connectors and user-defined functions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArtifactType" : String,
  "MavenReference" : MavenReference,
  "S3ContentLocation" : S3ContentLocation
}

```

### YAML

```yaml

  ArtifactType: String
  MavenReference:
    MavenReference
  S3ContentLocation:
    S3ContentLocation

```

## Properties

`ArtifactType`

Set this to either `UDF` or `DEPENDENCY_JAR`. `UDF`
stands for user-defined functions. This type of artifact must be in an S3 bucket. A
`DEPENDENCY_JAR` can be in either Maven or an S3 bucket.

_Required_: Yes

_Type_: String

_Allowed values_: `DEPENDENCY_JAR | UDF`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MavenReference`

The parameters required to fully specify a Maven reference.

_Required_: No

_Type_: [MavenReference](aws-properties-kinesisanalyticsv2-application-mavenreference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ContentLocation`

The location of the custom artifacts.

_Required_: No

_Type_: [S3ContentLocation](aws-properties-kinesisanalyticsv2-application-s3contentlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CSVMappingParameters

DeployAsApplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
