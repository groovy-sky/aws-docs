# CreateOptionGroup

Creates a new option group. You can create up to 20 option groups.

This command doesn't apply to RDS Custom.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**EngineName**

The name of the engine to associate this option group with.

Valid Values:

- `db2-ae`

- `db2-se`

- `mariadb`

- `mysql`

- `oracle-ee`

- `oracle-ee-cdb`

- `oracle-se2`

- `oracle-se2-cdb`

- `postgres`

- `sqlserver-ee`

- `sqlserver-se`

- `sqlserver-ex`

- `sqlserver-web`

Type: String

Required: Yes

**MajorEngineVersion**

Specifies the major version of the engine that this option group should be associated with.

Type: String

Required: Yes

**OptionGroupDescription**

The description of the option group.

Type: String

Required: Yes

**OptionGroupName**

Specifies the name of the option group to be created.

Constraints:

- Must be 1 to 255 letters, numbers, or hyphens

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

Example: `myoptiongroup`

Type: String

Required: Yes

**Tags.Tag.N**

Tags to assign to the option group.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**OptionGroup**

Type: [OptionGroup](api-optiongroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**OptionGroupAlreadyExistsFault**

The option group you are trying to create already exists.

HTTP Status Code: 400

**OptionGroupQuotaExceededFault**

The quota of 20 option groups was exceeded for this AWS account.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of CreateOptionGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=CreateOptionGroup
   &EngineName=mysql
   &MajorEngineVersion=5.6
   &OptionGroupDescription=My%20Option%20Group
   &OptionGroupName=myawsuser-og00
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140425/us-east-1/rds/aws4_request
   &X-Amz-Date=20140425T174519Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=d3a89afa4511d0c4ecab046d6dc760a72bfe6bb15999cce053adeb2617b60384

```

#### Sample Response

```

<CreateOptionGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CreateOptionGroupResult>
    <OptionGroup>
      <AllowsVpcAndNonVpcInstanceMemberships>true</AllowsVpcAndNonVpcInstanceMemberships>
      <MajorEngineVersion>5.6</MajorEngineVersion>
      <OptionGroupName>myawsuser-og00</OptionGroupName>
      <EngineName>mysql</EngineName>
      <OptionGroupDescription>My Option Group</OptionGroupDescription>
      <Options/>
    </OptionGroup>
  </CreateOptionGroupResult>
  <ResponseMetadata>
    <RequestId>4d7f11f2-bbf0-11d3-ae4f-eec568ed6b36</RequestId>
  </ResponseMetadata>
</CreateOptionGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/createoptiongroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/createoptiongroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateIntegration

CreateTenantDatabase

All content copied from https://docs.aws.amazon.com/.
