This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup Deployment

`Deployment` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource that specifies an AWS CodeDeploy application
revision to be deployed to instances in the deployment group. If you specify an application
revision, your target revision is deployed as soon as the provisioning process is complete.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "IgnoreApplicationStopFailures" : Boolean,
  "Revision" : RevisionLocation
}

```

### YAML

```yaml

  Description: String
  IgnoreApplicationStopFailures: Boolean
  Revision:
    RevisionLocation

```

## Properties

`Description`

A comment about the deployment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IgnoreApplicationStopFailures`

If true, then if an `ApplicationStop`, `BeforeBlockTraffic`, or
`AfterBlockTraffic` deployment lifecycle event to an instance fails, then the
deployment continues to the next deployment lifecycle event. For example, if
`ApplicationStop` fails, the deployment continues with DownloadBundle. If
`BeforeBlockTraffic` fails, the deployment continues with
`BlockTraffic`. If `AfterBlockTraffic` fails, the deployment continues
with `ApplicationStop`.

If false or not specified, then if a lifecycle event fails during a deployment to an
instance, that deployment fails. If deployment to that instance is part of an overall
deployment and the number of healthy hosts is not less than the minimum number of healthy
hosts, then a deployment to the next instance is attempted.

During a deployment, the AWS CodeDeploy agent runs the scripts specified for
`ApplicationStop`, `BeforeBlockTraffic`, and
`AfterBlockTraffic` in the AppSpec file from the previous successful deployment.
(All other scripts are run from the AppSpec file in the current deployment.) If one of these
scripts contains an error and does not run successfully, the deployment can fail.

If the cause of the failure is a script from the last successful deployment that will
never run successfully, create a new deployment and use
`ignoreApplicationStopFailures` to specify that the `ApplicationStop`,
`BeforeBlockTraffic`, and `AfterBlockTraffic` failures should be
ignored.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Revision`

Information about the location of stored application artifacts and the service from
which to retrieve them.

_Required_: Yes

_Type_: [RevisionLocation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentgroup-revisionlocation.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BlueInstanceTerminationOption

DeploymentReadyOption
