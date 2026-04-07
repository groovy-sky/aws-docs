This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint DynamoDbSettings

Provides information, including the Amazon Resource Name (ARN) of the IAM
role used to define an Amazon DynamoDB target endpoint. This
information also includes the output format of records applied to the endpoint and details of
transaction and control table data information. For information about other available settings, see
[Using object mapping to migrate data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServiceAccessRoleArn" : String
}

```

### YAML

```yaml

  ServiceAccessRoleArn: String

```

## Properties

`ServiceAccessRoleArn`

The Amazon Resource Name (ARN) used by the service to access the IAM role. The role must allow the `iam:PassRole` action.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DocDbSettings

ElasticsearchSettings
