This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Keyspace ReplicationSpecification

You can use `ReplicationSpecification`
to configure the `ReplicationStrategy` of a keyspace in Amazon Keyspaces.

The `ReplicationSpecification` property applies automatically to all tables in the keyspace.

To review the permissions that are required to add a new Region to a single-Region keyspace,
see [Configure the IAM permissions required\
to add an AWS Region to a keyspace](https://docs.aws.amazon.com/keyspaces/latest/devguide/howitworks_replication_permissions_addReplica.html)
in the _Amazon Keyspaces Developer Guide_.

For more information about multi-Region replication, see [Multi-Region replication](https://docs.aws.amazon.com/keyspaces/latest/devguide/multiRegion-replication.html)
in the _Amazon Keyspaces Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RegionList" : [ String, ... ],
  "ReplicationStrategy" : String
}

```

### YAML

```yaml

  RegionList:
    - String
  ReplicationStrategy: String

```

## Properties

`RegionList`

Specifies the AWS Regions that the
keyspace is replicated in.
You must specify at least two Regions, including the Region that the keyspace is being created in.

To specify a Region [that's disabled by\
default](../../../accounts/latest/reference/manage-acct-regions.md#rande-manage-enable), you must first enable the Region. For more information, see
[Multi-Region replication in \
AWS Regions disabled by default](https://docs.aws.amazon.com/keyspaces/latest/devguide/multiRegion-replication_how-it-works.html#howitworks_mrr_opt_in) in the _Amazon Keyspaces Developer Guide_.

_Required_: No

_Type_: Array of String

_Allowed values_: `af-south-1 | ap-east-1 | ap-northeast-1 | ap-northeast-2 | ap-south-1 | ap-southeast-1 | ap-southeast-2 | ca-central-1 | eu-central-1 | eu-north-1 | eu-west-1 | eu-west-2 | eu-west-3 | me-central-1 | me-south-1 | sa-east-1 | us-east-1 | us-east-2 | us-west-1 | us-west-2`

_Minimum_: `2`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationStrategy`

The options are:

- `SINGLE_REGION` (optional)

- `MULTI_REGION`

If no value is specified, the default is `SINGLE_REGION`. If `MULTI_REGION` is specified, `RegionList` is required.

_Required_: No

_Type_: String

_Allowed values_: `SINGLE_REGION | MULTI_REGION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Cassandra::Keyspace

Tag
