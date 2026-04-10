This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsGuru::ResourceCollection ResourceCollectionFilter

Information about a filter used to specify which AWS resources are analyzed for
anomalous behavior by DevOps Guru.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudFormation" : CloudFormationCollectionFilter,
  "Tags" : [ TagCollection, ... ]
}

```

### YAML

```yaml

  CloudFormation:
    CloudFormationCollectionFilter
  Tags:
    - TagCollection

```

## Properties

`CloudFormation`

Information about AWS CloudFormation stacks. You can use up to 1000
stacks to specify which AWS resources in your account to analyze. For more
information, see [Stacks](../userguide/stacks.md) in the
_AWS CloudFormation User Guide_.

_Required_: No

_Type_: [CloudFormationCollectionFilter](aws-properties-devopsguru-resourcecollection-cloudformationcollectionfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The AWS tags used to filter the resources in the resource collection.

Tags help you identify and organize your AWS resources. Many AWS services support
tagging, so you can assign the same tag to resources from different services to indicate
that the resources are related. For example, you can assign the same tag to an Amazon DynamoDB
table resource that you assign to an AWS Lambda function. For more information about
using tags, see the [Tagging\
best practices](../../../whitepapers/latest/tagging-best-practices/tagging-best-practices.md) whitepaper.

Each AWS tag has two parts.

- A tag _key_ (for example, `CostCenter`,
`Environment`, `Project`, or `Secret`). Tag
_keys_ are case-sensitive.

- A field known as a tag _value_ (for example,
`111122223333`, `Production`, or a team
name). Omitting the tag _value_ is the same as using an empty
string. Like tag _keys_, tag _values_ are
case-sensitive. The tag value is a required property when AppBoundaryKey is specified.

Together these are known as _key_- _value_ pairs.

###### Important

The string used for a _key_ in a tag that you use to define your resource coverage must begin with the
prefix `Devops-guru-`. The tag _key_ might be
`DevOps-Guru-deployment-application` or
`devops-guru-rds-application`. When you create a _key_, the case of characters in the _key_ can be whatever you choose. After you create a _key_, it is case-sensitive.
For example, DevOps Guru works with a
_key_ named `devops-guru-rds` and a _key_ named
`DevOps-Guru-RDS`, and these act as two different _keys_. Possible _key_/ _value_ pairs in your
application might be `Devops-Guru-production-application/RDS` or
`Devops-Guru-production-application/containers`.

_Required_: No

_Type_: Array of [TagCollection](aws-properties-devopsguru-resourcecollection-tagcollection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudFormationCollectionFilter

TagCollection

All content copied from https://docs.aws.amazon.com/.
