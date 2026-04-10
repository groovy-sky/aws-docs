This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup RevisionLocation

`RevisionLocation` is a property that defines the location of the CodeDeploy application revision to deploy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GitHubLocation" : GitHubLocation,
  "RevisionType" : String,
  "S3Location" : S3Location
}

```

### YAML

```yaml

  GitHubLocation:
    GitHubLocation
  RevisionType: String
  S3Location:
    S3Location

```

## Properties

`GitHubLocation`

Information about the location of application artifacts stored in GitHub.

_Required_: No

_Type_: [GitHubLocation](aws-properties-codedeploy-deploymentgroup-githublocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RevisionType`

The type of application revision:

- S3: An application revision stored in Amazon S3.

- GitHub: An application revision stored in GitHub (EC2/On-premises deployments
only).

- String: A YAML-formatted or JSON-formatted string (AWS Lambda
deployments only).

- AppSpecContent: An `AppSpecContent` object that contains the
contents of an AppSpec file for an AWS Lambda or Amazon ECS
deployment. The content is formatted as JSON or YAML stored as a
RawString.

_Required_: No

_Type_: String

_Allowed values_: `S3 | GitHub | String | AppSpecContent | SystemsManagerPatchBaseline`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Location`

Information about the location of a revision stored in Amazon S3.

_Required_: No

_Type_: [S3Location](aws-properties-codedeploy-deploymentgroup-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnPremisesTagSetListObject

S3Location

All content copied from https://docs.aws.amazon.com/.
