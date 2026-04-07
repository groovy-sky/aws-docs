This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup OnPremisesTagSet

The `OnPremisesTagSet` property type specifies a list containing other lists of
on-premises instance tag groups. In order for an instance to be included in the deployment
group, it must be identified by all the tag groups in the list.

For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment\
Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the _AWS CodeDeploy User_
_Guide_.

`OnPremisesTagSet` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnPremisesTagSetList" : [ OnPremisesTagSetListObject, ... ]
}

```

### YAML

```yaml

  OnPremisesTagSetList:
    - OnPremisesTagSetListObject

```

## Properties

`OnPremisesTagSetList`

A list that contains other lists of on-premises instance tag groups. For an instance to be
included in the deployment group, it must be identified by all of the tag groups in the
list.

Duplicates are not allowed.

_Required_: No

_Type_: Array of [OnPremisesTagSetListObject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentgroup-onpremisestagsetlistobject.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LoadBalancerInfo

OnPremisesTagSetListObject
