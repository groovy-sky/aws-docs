# DescribeOptionGroupOptions

Describes all available options for the specified engine.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**EngineName**

The name of the engine to describe options for.

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

**Filters.Filter.N**

This parameter isn't currently supported.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MajorEngineVersion**

If specified, filters the results to include only options for the specified major engine version.

Type: String

Required: No

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more records exist than the specified `MaxRecords` value,
a pagination token called a marker is included in the response so that
you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

**OptionGroupOptions.OptionGroupOption.N**

List of available option group options.

Type: Array of [OptionGroupOption](api-optiongroupoption.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### Example

This example illustrates one usage of DescribeOptionGroupOptions.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=DescribeOptionGroupOptions
   &EngineName=oracle-se1
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140421/us-west-2/rds/aws4_request
   &X-Amz-Date=20140421T194733Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=3792d1669ce65ba1ba6a85b2e4057235e46dd3d0072663c17f4b4439fd8af702

```

#### Sample Response

```

<DescribeOptionGroupOptionsResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeOptionGroupOptionsResult>
    <OptionGroupOptions>
      <OptionGroupOption>
        <MajorEngineVersion>11.2</MajorEngineVersion>
        <PortRequired>false</PortRequired>
        <Persistent>false</Persistent>
        <OptionsDependedOn>
          <OptionName>XMLDB</OptionName>
        </OptionsDependedOn>
        <OptionsConflictsWith/>
        <Permanent>false</Permanent>
        <Description>Oracle Application Express Runtime Environment</Description>
        <Name>APEX</Name>
        <OptionGroupOptionSettings/>
        <EngineName>oracle-se1</EngineName>
        <MinimumRequiredMinorEngineVersion>0.2.v4</MinimumRequiredMinorEngineVersion>
      </OptionGroupOption>
      <OptionGroupOption>
        <MajorEngineVersion>11.2</MajorEngineVersion>
        <PortRequired>false</PortRequired>
        <Persistent>false</Persistent>
        <OptionsDependedOn>
          <OptionName>APEX</OptionName>
        </OptionsDependedOn>
        <OptionsConflictsWith/>
        <Permanent>false</Permanent>
        <Description>Oracle Application Express Development Environment</Description>
        <Name>APEX-DEV</Name>
        <OptionGroupOptionSettings/>
        <EngineName>oracle-se1</EngineName>
        <MinimumRequiredMinorEngineVersion>0.2.v4</MinimumRequiredMinorEngineVersion>
      </OptionGroupOption>
      <OptionGroupOption>
        <MajorEngineVersion>11.2</MajorEngineVersion>
        <PortRequired>true</PortRequired>
        <Persistent>false</Persistent>
        <OptionsDependedOn/>
        <OptionsConflictsWith/>
        <Permanent>false</Permanent>
        <Description>Oracle Enterprise Manager (Database Control only)</Description>
        <DefaultPort>1158</DefaultPort>
        <Name>OEM</Name>
        <OptionGroupOptionSettings/>
        <EngineName>oracle-se1</EngineName>
        <MinimumRequiredMinorEngineVersion>0.2.v3</MinimumRequiredMinorEngineVersion>
      </OptionGroupOption>
    </OptionGroupOptions>
  </DescribeOptionGroupOptionsResult>
  <ResponseMetadata>
    <RequestId>b7b26a8f-b98c-11d3-f272-7cd6cce12cc5</RequestId>
  </ResponseMetadata>
</DescribeOptionGroupOptionsResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describeoptiongroupoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describeoptiongroupoptions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeIntegrations

DescribeOptionGroups

All content copied from https://docs.aws.amazon.com/.
