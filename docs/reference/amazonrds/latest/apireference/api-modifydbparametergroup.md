# ModifyDBParameterGroup

Modifies the parameters of a DB parameter group. To modify more than one parameter,
submit a list of the following: `ParameterName`, `ParameterValue`, and
`ApplyMethod`. A maximum of 20 parameters can be modified in a single request.

###### Important

After you modify a DB parameter group, you should wait at least 5 minutes
before creating your first DB instance that uses that DB parameter group as the default parameter
group. This allows Amazon RDS to fully complete the modify operation before the parameter
group is used as the default for a new DB instance. This is especially important for parameters
that are critical when creating the default database for a DB instance, such as the character set
for the default database defined by the `character_set_database` parameter. You can use the
_Parameter Groups_ option of the [Amazon RDS console](https://console.aws.amazon.com/rds) or the
_DescribeDBParameters_ command to verify
that your DB parameter group has been created or modified.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBParameterGroupName**

The name of the DB parameter group.

Constraints:

- If supplied, must match the name of an existing `DBParameterGroup`.

Type: String

Required: Yes

**Parameters.Parameter.N**

An array of parameter names, values, and the application methods for the parameter update. At least one parameter name, value, and
application method must be supplied; later arguments are optional. A maximum of 20 parameters can be modified in a single request.

Valid Values (for the application method): `immediate | pending-reboot`

You can use the `immediate` value with dynamic parameters only. You can use the `pending-reboot` value for both dynamic
and static parameters.

When the application method is `immediate`, changes to dynamic parameters are applied immediately to the DB instances associated with
the parameter group.

When the application method is `pending-reboot`, changes to dynamic and static parameters are applied after a reboot without failover
to the DB instances associated with the parameter group.

###### Note

You can't use `pending-reboot` with dynamic parameters on RDS for SQL Server DB instances. Use `immediate`.

For more information on modifying DB parameters, see [Working \
with DB parameter groups](../../../../services/amazonrds/latest/userguide/user-workingwithparamgroups.md) in the _Amazon RDS User Guide_.

Type: Array of [Parameter](api-parameter.md) objects

Required: Yes

## Response Elements

The following element is returned by the service.

**DBParameterGroupName**

The name of the DB parameter group.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBParameterGroupNotFound**

`DBParameterGroupName` doesn't refer to an
existing DB parameter group.

HTTP Status Code: 404

**InvalidDBParameterGroupState**

The DB parameter group is in use or is in an invalid state. If you are attempting
to delete the parameter group, you can't delete it when the parameter group is in
this state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of ModifyDBParameterGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=ModifyDBParameterGroup
   &DBParameterGroupName=mydbparametergroup01
   &Parameters.member.1.ApplyMethod=immediate
   &Parameters.member.1.ParameterName=binlog_cache_size
   &Parameters.member.1.ParameterValue=65536
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140425/us-east-1/rds/aws4_request
   &X-Amz-Date=20140425T193811Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=afd9acfee876360dd294189465aca26502343d405292dc6e43b1961ad4d1d7e2

```

#### Sample Response

```

<ModifyDBParameterGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ModifyDBParameterGroupResult>
    <DBParameterGroupName>mydbparametergroup01</DBParameterGroupName>
  </ModifyDBParameterGroupResult>
  <ResponseMetadata>
    <RequestId>12d7435e-bba0-11d3-fe11-33d33a9bb7e3</RequestId>
  </ResponseMetadata>
</ModifyDBParameterGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifydbparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifydbparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyDBInstance

ModifyDBProxy

All content copied from https://docs.aws.amazon.com/.
