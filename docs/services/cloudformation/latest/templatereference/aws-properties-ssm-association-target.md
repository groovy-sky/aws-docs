This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::Association Target

`Target` is a property of the [AWS::SSM::Association](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html) resource that specifies the targets for an SSM
document in Systems Manager. You can target all instances in an AWS account by specifying the `InstanceIds` key with a value of
`*`. To view a JSON and a YAML example that targets all instances, see
the example "Create an association for all managed instances in an AWS account" later in this page.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Values:
    - String

```

## Properties

`Key`

User-defined criteria for sending commands that target managed nodes that meet the
criteria.

_Required_: Yes

_Type_: String

_Pattern_: `^[\p{L}\p{Z}\p{N}_.:/=+\-@]{1,128}$|resource-groups:Name`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

User-defined criteria that maps to `Key`. For example, if you specified
`tag:ServerRole`, you could specify `value:WebServer` to run a command on
instances that include EC2 tags of `ServerRole,WebServer`.

Depending on the type of target, the maximum number of values for a key might be lower than
the global maximum of 50.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3OutputLocation

AWS::SSM::Document
