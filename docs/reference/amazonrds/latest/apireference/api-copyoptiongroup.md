# CopyOptionGroup

Copies the specified option group.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceOptionGroupIdentifier**

The identifier for the source option group.

Constraints:

- Must specify a valid option group.

Type: String

Required: Yes

**TargetOptionGroupDescription**

The description for the copied option group.

Type: String

Required: Yes

**TargetOptionGroupIdentifier**

The identifier for the copied option group.

Constraints:

- Can't be null, empty, or blank

- Must contain from 1 to 255 letters, numbers, or hyphens

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

Example: `my-option-group`

Type: String

Required: Yes

**Tags.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

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

**OptionGroupNotFoundFault**

The specified option group could not be found.

HTTP Status Code: 404

**OptionGroupQuotaExceededFault**

The quota of 20 option groups was exceeded for this AWS account.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of CopyOptionGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=CopyOptionGroup
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &SourceOptionGroupIdentifier=my-option-group
   &TargetOptionGroupDescription=New%20option%20group
   &TargetOptionGroupIdentifier=new-option-group
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140429/us-east-1/rds/aws4_request
   &X-Amz-Date=20140429T175351Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=9164337efa99caf850e874a1cb7ef62f3cea29d0b448b9e0e7c53b288ddffed2

```

#### Sample Response

```

<CopyOptionGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CopyOptionGroupResult>
    <OptionGroup>
      <OptionGroupName>new-option-group</OptionGroupName>
      <MajorEngineVersion>5.6</MajorEngineVersion>
      <AllowsVpcAndNonVpcInstanceMemberships>false</AllowsVpcAndNonVpcInstanceMemberships>
      <EngineName>mysql</EngineName>
      <OptionGroupDescription>description</OptionGroupDescription>
      <Options>
        <Option>
          <Port>11211</Port>
          <OptionName>MEMCACHED</OptionName>
          <OptionDescription>Innodb Memcached for MySQL</OptionDescription>
          <Persistent>false</Persistent>
          <OptionSettings>
            <OptionSetting>
              <DataType>BOOLEAN</DataType>
              <IsModifiable>true</IsModifiable>
              <IsCollection>false</IsCollection>
              <Description>If enabled when there is no more memory to store items, memcached returns an error rather than evicting items.</Description>
              <Name>ERROR_ON_MEMORY_EXHAUSTED</Name>
              <Value>0</Value>
              <ApplyType>STATIC</ApplyType>
              <AllowedValues>0,1</AllowedValues>
              <DefaultValue>0</DefaultValue>
            </OptionSetting>
            <OptionSetting>
              <DataType>INTEGER</DataType>
              <IsModifiable>true</IsModifiable>
              <IsCollection>false</IsCollection>
              <Description>The backlog queue configures how many network connections can be waiting to be processed by memcached</Description>
              <Name>BACKLOG_QUEUE_LIMIT</Name>
              <Value>1024</Value>
              <ApplyType>STATIC</ApplyType>
              <AllowedValues>1-2048</AllowedValues>
              <DefaultValue>1024</DefaultValue>
            </OptionSetting>
          </OptionSettings>
          <VpcSecurityGroupMemberships/>
          <Permanent>false</Permanent>
          <DBSecurityGroupMemberships>
            <DBSecurityGroup>
              <Status>authorized</Status>
              <DBSecurityGroupName>default</DBSecurityGroupName>
            </DBSecurityGroup>
          </DBSecurityGroupMemberships>
        </Option>
      </Options>
    </OptionGroup>
  </CopyOptionGroupResult>
  <ResponseMetadata>
    <RequestId>2928d60e-beb6-11d3-8e5c-3ccda5460c46</RequestId>
  </ResponseMetadata>
</CopyOptionGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/copyoptiongroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/copyoptiongroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyDBSnapshot

CreateBlueGreenDeployment

All content copied from https://docs.aws.amazon.com/.
