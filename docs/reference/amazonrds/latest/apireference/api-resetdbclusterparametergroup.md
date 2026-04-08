# ResetDBClusterParameterGroup

Modifies the parameters of a DB cluster parameter group to the default value. To
reset specific parameters submit a list of the following: `ParameterName`
and `ApplyMethod`. To reset the
entire DB cluster parameter group, specify the `DBClusterParameterGroupName`
and `ResetAllParameters` parameters.

When resetting the entire group, dynamic parameters are updated immediately and static parameters
are set to `pending-reboot` to take effect on the next DB instance restart
or `RebootDBInstance` request. You must call `RebootDBInstance` for every
DB instance in your DB cluster that you want the updated static parameter to apply to.

For more information on Amazon Aurora DB clusters, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide._

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterParameterGroupName**

The name of the DB cluster parameter group to reset.

Type: String

Required: Yes

**Parameters.Parameter.N**

A list of parameter names in the DB cluster parameter group to reset to the default values. You can't use this
parameter if the `ResetAllParameters` parameter is enabled.

Type: Array of [Parameter](api-parameter.md) objects

Required: No

**ResetAllParameters**

Specifies whether to reset all parameters in the DB cluster parameter group
to their default values. You can't use this parameter if there
is a list of parameter names specified for the `Parameters` parameter.

Type: Boolean

Required: No

## Response Elements

The following element is returned by the service.

**DBClusterParameterGroupName**

The name of the DB cluster parameter group.

Constraints:

- Must be 1 to 255 letters or numbers.

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

###### Note

This value is stored as a lowercase string.

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

This example illustrates one usage of ResetDBClusterParameterGroup.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
    ?Action=ResetDBClusterParameterGroup
    &DBClusterParameterGroupName=sample-cluster-pg
    &Parameters.member.1.ApplyMethod=pending-reboot
    &Parameters.member.1.ParameterName=binlog_format
    &Parameters.member.2.ApplyMethod=pending-reboot
    &Parameters.member.2.ParameterName=innodb_support_xa
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20160913/us-west-2/rds/aws4_request
    &X-Amz-Date=20160913T230026Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=7cca4504082065e227696f2dd904fab2f39633bc7d120258c4bedd35da3ade7f
```

#### Sample Response

```

<ResetDBClusterParameterGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResetDBClusterParameterGroupResult>
    <DBClusterParameterGroupName>sample-cluster-pg</DBClusterParameterGroupName>
  </ResetDBClusterParameterGroupResult>
  <ResponseMetadata>
    <RequestId>dc2c61eb-7a05-11e6-b83b-cd70a540d79f</RequestId>
  </ResponseMetadata>
</ResetDBClusterParameterGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/resetdbclusterparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/resetdbclusterparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemoveTagsFromResource

ResetDBParameterGroup

All content copied from https://docs.aws.amazon.com/.
