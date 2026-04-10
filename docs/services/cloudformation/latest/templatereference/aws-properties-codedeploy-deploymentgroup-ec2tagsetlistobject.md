This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup EC2TagSetListObject

The `EC2TagSet` property type specifies information about groups of tags
applied to Amazon EC2 instances. The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same template
as EC2TagFilters.

For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment\
Groups in AWS CodeDeploy](../../../codedeploy/latest/userguide/instances-tagging.md) in the _AWS CodeDeploy User_
_Guide_.

`EC2TagSet` is a property of the [DeploymentGroup](../userguide/aws-resource-codedeploy-deploymentgroup.md) resource type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ec2TagGroup" : [ EC2TagFilter, ... ]
}

```

### YAML

```yaml

  Ec2TagGroup:
    - EC2TagFilter

```

## Properties

`Ec2TagGroup`

A list that contains other lists of Amazon EC2 instance tag groups. For an
instance to be included in the deployment group, it must be identified by all of the tag
groups in the list.

_Required_: No

_Type_: Array of [EC2TagFilter](aws-properties-codedeploy-deploymentgroup-ec2tagfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [EC2TagSet](../../../../reference/codedeploy/latest/apireference/api-ec2tagset.md) in the _AWS CodeDeploy API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EC2TagSet

ECSService

All content copied from https://docs.aws.amazon.com/.
