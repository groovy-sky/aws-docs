This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverQueryLoggingConfig

The AWS::Route53Resolver::ResolverQueryLoggingConfig resource is a complex type that contains settings for one query logging configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverQueryLoggingConfig",
  "Properties" : {
      "DestinationArn" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::ResolverQueryLoggingConfig
Properties:
  DestinationArn: String
  Name: String
  Tags:
    - Tag

```

## Properties

`DestinationArn`

The ARN of the resource that you want Resolver to send query logs: an Amazon S3 bucket, a CloudWatch Logs log group, or
a Kinesis Data Firehose delivery stream.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `600`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the query logging configuration.

_Required_: No

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-route53resolver-resolverqueryloggingconfig-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the resource that contains settings for one query logging configuration.

For example: `{ "Ref": "rqlc-1111222233334444" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the query logging configuration.

`AssociationCount`

The number of VPCs that are associated with the query logging configuration.

`CreationTime`

The date and time that the query logging configuration was created, in Unix time format and Coordinated Universal Time (UTC).

`CreatorRequestId`

A unique string that identifies the request that created the query logging configuration. The `CreatorRequestId` allows failed requests
to be retried without the risk of running the operation twice.

`Id`

The ID for the query logging configuration.

`OwnerId`

The AWS account ID for the account that created the query logging configuration.

`ShareStatus`

An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another
AWS account. Sharing is configured through AWS Resource Access Manager (AWS RAM).

`Status`

The status of the specified query logging configuration. Valid values include the following:

- `CREATING`: Resolver is creating the query logging configuration.

- `CREATED`: The query logging configuration was successfully created.
Resolver is logging queries that originate in the specified VPC.

- `DELETING`: Resolver is deleting this query logging configuration.

- `FAILED`: Resolver can't deliver logs to the location that is specified in the query logging configuration.
Here are two common causes:

- The specified destination (for example, an Amazon S3 bucket) was deleted.

- Permissions don't allow sending logs to the destination.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
