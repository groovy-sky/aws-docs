This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsGuru::ResourceCollection TagCollection

A collection of AWS tags.

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
case-sensitive. The tag value is a required property when _AppBoundaryKey_ is specified.

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

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppBoundaryKey" : String,
  "TagValues" : [ String, ... ]
}

```

### YAML

```yaml

  AppBoundaryKey: String
  TagValues:
    - String

```

## Properties

`AppBoundaryKey`

An AWS tag _key_ that is used to identify the AWS resources that
DevOps Guru analyzes. All AWS resources in your account and Region tagged with this _key_ make
up your DevOps Guru application and analysis boundary.

###### Important

When you create a _key_, the case of characters in the _key_ can be whatever you choose. After you create a _key_, it is case-sensitive.
For example, DevOps Guru works with a
_key_ named `devops-guru-rds` and a _key_ named
`DevOps-Guru-RDS`, and these act as two different _keys_. Possible _key_/ _value_ pairs in your
application might be `Devops-Guru-production-application/RDS` or
`Devops-Guru-production-application/containers`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagValues`

The values in an AWS tag collection.

The tag's _value_ is a field used to associate a string with
the tag _key_ (for example, `111122223333`, `Production`, or a team
name). The _key_ and _value_ are the tag's _key_ pair.
Omitting the tag _value_ is the same as using an empty
string. Like tag _keys_, tag _values_ are
case-sensitive. You can specify a maximum of 256 characters for a tag value. The tag value is a required property when _AppBoundaryKey_ is specified.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceCollectionFilter

Next

All content copied from https://docs.aws.amazon.com/.
