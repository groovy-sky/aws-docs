This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Macie::Session

The `AWS::Macie::Session` resource represents the Amazon Macie
service and certain configuration settings for an Amazon Macie account in a
specific AWS Region. It enables Macie to become
operational for a specific account in a specific Region. An account can have only one
session in each Region.

You must create an `AWS::Macie::Session` resource for an account before you
can create other types of resources for the account. Use a [DependsOn\
attribute](../userguide/aws-attribute-dependson.md) to ensure that an `AWS::Macie::Session` resource is
created before other Macie resources are created for an account. For
example, `"DependsOn": "Session"`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Macie::Session",
  "Properties" : {
      "FindingPublishingFrequency" : String,
      "Status" : String
    }
}

```

### YAML

```yaml

Type: AWS::Macie::Session
Properties:
  FindingPublishingFrequency: String
  Status: String

```

## Properties

`FindingPublishingFrequency`

Specifies how often Amazon Macie publishes updates to policy findings for
the account. This includes publishing updates to AWS Security Hub CSPM and Amazon EventBridge (formerly Amazon CloudWatch Events). Valid values are:

- FIFTEEN\_MINUTES

- ONE\_HOUR

- SIX\_HOURS

_Required_: No

_Type_: String

_Allowed values_: `FIFTEEN_MINUTES | ONE_HOUR | SIX_HOURS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of Amazon Macie for the account. Valid values are:
`ENABLED`, start or resume Macie activities for the
account; and, `PAUSED`, suspend Macie activities for the
account.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | PAUSED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the account ID for the AWS account in
which the Amazon Macie session is created. For example, `{ "Ref":
                "Session" }`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AutomatedDiscoveryStatus`

The status of automated sensitive data discovery for the account. Possible values are: `ENABLED`, perform automated sensitive
data discovery activities for the account; and, `DISABLED`, don't perform automated sensitive data discovery activities for the
account.

`AwsAccountId`

The account ID for the AWS account in which the Amazon Macie session is created.

`ServiceRole`

The Amazon Resource Name (ARN) of the service-linked role that allows Amazon Macie to monitor and analyze data in
AWS resources for the account.

## Examples

The following example demonstrates how to declare an
`AWS::Macie::Session` resource.

### Creating a session

This example enables Amazon Macie for an account. It also
configures Macie to publish updated policy findings every hour
for the account.

#### JSON

```json

{
    "Type": "AWS::Macie::Session",
    "Properties": {
        "FindingPublishingFrequency": "ONE_HOUR",
        "Status": "ENABLED"
    }
}
```

#### YAML

```yaml

Type: 'AWS::Macie::Session'
Properties:
  FindingPublishingFrequency: ONE_HOUR
  Status: ENABLED
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
