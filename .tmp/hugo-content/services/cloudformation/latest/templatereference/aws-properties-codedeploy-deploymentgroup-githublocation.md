This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup GitHubLocation

`GitHubLocation` is a property of the [CodeDeploy DeploymentGroup Revision](../userguide/aws-properties-codedeploy-deploymentgroup-deployment-revision.md) property that specifies the
location of an application revision that is stored in GitHub.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommitId" : String,
  "Repository" : String
}

```

### YAML

```yaml

  CommitId: String
  Repository: String

```

## Properties

`CommitId`

The SHA1 commit ID of the GitHub commit that represents the bundled artifacts for the
application revision.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Repository`

The GitHub account and repository pair that stores a reference to the commit that
represents the bundled artifacts for the application revision.

Specify the value as `account/repository`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ELBInfo

GreenFleetProvisioningOption

All content copied from https://docs.aws.amazon.com/.
