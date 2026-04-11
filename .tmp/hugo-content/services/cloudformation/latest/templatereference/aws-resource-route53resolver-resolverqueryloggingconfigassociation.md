This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation

The AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation resource is a configuration for DNS query logging. After you create a query logging configuration, Amazon Route 53 begins to publish log data to an Amazon CloudWatch Logs log group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation",
  "Properties" : {
      "ResolverQueryLogConfigId" : String,
      "ResourceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation
Properties:
  ResolverQueryLogConfigId: String
  ResourceId: String

```

## Properties

`ResolverQueryLogConfigId`

The ID of the query logging configuration that a VPC is associated with.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceId`

The ID of the Amazon VPC that is associated with the query logging configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the configuration for DNS query logging.

For example: `{ "Ref": "rqlca-1111222233334444" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The date and time that the VPC was associated with the query logging configuration, in Unix time format and Coordinated Universal Time (UTC).

`Error`

If the value of `Status` is `FAILED`, the value of
`Error` indicates the cause:

- `DESTINATION_NOT_FOUND`: The specified destination (for example, an Amazon S3 bucket) was deleted.

- `ACCESS_DENIED`: Permissions don't allow sending logs to the destination.

If the value of `Status` is a value other than `FAILED`, `Error` is null.

`ErrorMessage`

Contains additional information about the error. If the value or `Error` is null, the value of `ErrorMessage` is also
null.

`Id`

The ID of the query logging association.

`Status`

The status of the specified query logging association. Valid values include the following:

- `CREATING`: Resolver is creating an association between an Amazon Virtual Private Cloud (Amazon VPC) and a query logging configuration.

- `CREATED`: The association between an Amazon VPC and a query logging configuration
was successfully created. Resolver is logging queries that originate in the specified VPC.

- `DELETING`: Resolver is deleting this query logging association.

- `FAILED`: Resolver either couldn't create or couldn't delete the query logging association.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Route53Resolver::ResolverRule

All content copied from https://docs.aws.amazon.com/.
