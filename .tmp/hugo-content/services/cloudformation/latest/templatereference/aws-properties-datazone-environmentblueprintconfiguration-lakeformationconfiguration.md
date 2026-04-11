This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::EnvironmentBlueprintConfiguration LakeFormationConfiguration

The Lake Formation configuration of the Data Lake blueprint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LocationRegistrationExcludeS3Locations" : [ String, ... ],
  "LocationRegistrationRole" : String
}

```

### YAML

```yaml

  LocationRegistrationExcludeS3Locations:
    - String
  LocationRegistrationRole: String

```

## Properties

`LocationRegistrationExcludeS3Locations`

Specifies certain Amazon S3 locations if you do not want Amazon DataZone to
automatically register them in hybrid mode.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `1024 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationRegistrationRole`

The role that is used to manage read/write access to the chosen Amazon S3 bucket(s) for
Data Lake using AWS Lake Formation hybrid access mode.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataZone::EnvironmentBlueprintConfiguration

ProvisioningConfiguration

All content copied from https://docs.aws.amazon.com/.
