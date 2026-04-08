# CreateDBClusterParameterGroup

Creates a new DB cluster parameter group.

Parameters in a DB cluster parameter group apply to all of the instances in a DB cluster.

A DB cluster parameter group is initially created with the default parameters for the
database engine used by instances in the DB cluster. To provide custom values for any of the
parameters, you must modify the group after creating it using
`ModifyDBClusterParameterGroup`. Once you've created a DB cluster parameter group, you need to
associate it with your DB cluster using `ModifyDBCluster`.

When you associate a new DB cluster parameter group with a running Aurora DB cluster, reboot the DB
instances in the DB cluster without failover for the new DB cluster parameter group and
associated settings to take effect.

When you associate a new DB cluster parameter group with a running Multi-AZ DB cluster, reboot the DB
cluster without failover for the new DB cluster parameter group and associated settings to take effect.

###### Important

After you create a DB cluster parameter group, you should wait at least 5 minutes
before creating your first DB cluster that uses that DB cluster parameter group as
the default parameter group. This allows Amazon RDS to fully complete the create
action before the DB cluster parameter group is used as the default for a new DB
cluster. This is especially important for parameters that are critical when creating
the default database for a DB cluster, such as the character set for the default
database defined by the `character_set_database` parameter. You can use
the _Parameter Groups_ option of the [Amazon RDS console](https://console.aws.amazon.com/rds) or the
`DescribeDBClusterParameters` operation to verify that your DB
cluster parameter group has been created or modified.

For more information on Amazon Aurora, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterParameterGroupName**

The name of the DB cluster parameter group.

Constraints:

- Must not match the name of an existing DB cluster parameter group.

###### Note

This value is stored as a lowercase string.

Type: String

Required: Yes

**DBParameterGroupFamily**

The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster
parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.

**Aurora MySQL**

Example: `aurora-mysql5.7`, `aurora-mysql8.0`

**Aurora PostgreSQL**

Example: `aurora-postgresql14`

**RDS for MySQL**

Example: `mysql8.0`

**RDS for PostgreSQL**

Example: `postgres13`

To list all of the available parameter group families for a DB engine, use the following command:

`aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>`

For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:

`aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine aurora-postgresql`

###### Note

The output contains duplicates.

The following are the valid DB engine values:

- `aurora-mysql`

- `aurora-postgresql`

- `mysql`

- `postgres`

Type: String

Required: Yes

**Description**

The description for the DB cluster parameter group.

Type: String

Required: Yes

**Tags.Tag.N**

Tags to assign to the DB cluster parameter group.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**DBClusterParameterGroup**

Contains the details of an Amazon RDS DB cluster parameter group.

This data type is used as a response element in the `DescribeDBClusterParameterGroups` action.

Type: [DBClusterParameterGroup](api-dbclusterparametergroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBParameterGroupAlreadyExists**

A DB parameter group with the same name exists.

HTTP Status Code: 400

**DBParameterGroupQuotaExceeded**

The request would result in the user exceeding the allowed number of DB parameter
groups.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of CreateDBClusterParameterGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=CreateDBClusterParameterGroup
    &DBClusterParameterGroupName=samplegroup
    &DBParameterGroupFamily=aurora5.6
    &Description=Sample%20group
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20150318/us-east-1/rds/aws4_request
    &X-Amz-Date=20150318T183624Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=86d521a3a117a033df0aa381fde0cd8a5ab5c7ab87a29aa9154438c3790ba611

```

#### Sample Response

```

<CreateDBClusterParameterGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CreateDBClusterParameterGroupResult>
    <DBClusterParameterGroup>
      <DBParameterGroupFamily>aurora5.6</DBParameterGroupFamily>
      <Description>Sample group</Description>
      <DBClusterParameterGroupName>samplegroup</DBClusterParameterGroupName>
    </DBClusterParameterGroup>
  </CreateDBClusterParameterGroupResult>
  <ResponseMetadata>
    <RequestId>ae81a963-cd9d-11e4-8b88-8351746a4c92</RequestId>
  </ResponseMetadata>
</CreateDBClusterParameterGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/createdbclusterparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/createdbclusterparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDBClusterEndpoint

CreateDBClusterSnapshot

All content copied from https://docs.aws.amazon.com/.
