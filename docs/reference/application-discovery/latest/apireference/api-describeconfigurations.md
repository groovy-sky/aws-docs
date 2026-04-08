# DescribeConfigurations

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Retrieves attributes for a list of configuration item IDs.

###### Note

All of the supplied IDs must be for the same asset type from one of the
following:

- server

- application

- process

- connection

Output fields are specific to the asset type specified. For example, the output for a
_server_ configuration item includes a list of attributes about the
server, such as host name, operating system, number of network cards, etc.

For a complete list of outputs for each asset type, see [Using the DescribeConfigurations Action](../../../../services/application-discovery/latest/userguide/discovery-api-queries.md#DescribeConfigurations) in the _AWS Application_
_Discovery Service User Guide_.

## Request Syntax

```nohighlight

{
   "configurationIds": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[configurationIds](#API_DescribeConfigurations_RequestSyntax)**

One or more configuration IDs.

Type: Array of strings

Length Constraints: Maximum length of 200.

Pattern: `\S*`

Required: Yes

## Response Syntax

```nohighlight

{
   "configurations": [
      {
         "string" : "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[configurations](#API_DescribeConfigurations_ResponseSyntax)**

A key in the response map. The value is an array of data.

Type: Array of string to string maps

Key Length Constraints: Maximum length of 10000.

Key Pattern: `[\s\S]*`

Value Length Constraints: Maximum length of 10000.

Value Pattern: `[\s\S]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AuthorizationErrorException**

The user does not have permission to perform the action. Check the IAM policy
associated with this user.

HTTP Status Code: 400

**HomeRegionNotSetException**

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

The home Region is not set. Set the home Region to continue.

HTTP Status Code: 400

**InvalidParameterException**

One or more parameters are not valid. Verify the parameters and try again.

HTTP Status Code: 400

**InvalidParameterValueException**

The value of one or more parameters are either invalid or out of range. Verify the
parameter values and try again.

HTTP Status Code: 400

**ServerInternalErrorException**

The server experienced an internal error. Try again.

HTTP Status Code: 500

## Examples

### Describe configurations

The following example queries configuration descriptions of two applications as
specified by the values passed to the parameter `configurationIds` in the
request. Note that the configuration ID's passed meet the requirement that they represent
the same asset type, in this example, _application_.

#### Sample Request

```

{
    "configurationIds": [
	      "d-application-0282ccd1ba7c211ca",
	      "d-application-034219054305e4a34"
	  ]
}
```

#### Sample Response

```

{
    "configurations": [
        {
            "application.configurationId": "d-application-0282ccd1ba7c211ca",
            "application.description": "ads CreateApplication",
            "application.lastModifiedTime": "2018-02-27 00:48:16.0",
            "application.name": "adsdp2",
            "application.serverCount": "2",
            "application.timeOfCreation": "2018-02-27 00:48:16.0"
        },
        {
            "application.configurationId": "d-application-034219054305e4a34",
            "application.description": "ads CreateApplication place holder app",
            "application.lastModifiedTime": "2018-02-27 23:33:32.0",
            "application.name": "adsdp2_ph1",
            "application.serverCount": "0",
            "application.timeOfCreation": "2018-02-27 23:33:32.0"
        }
    ]
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/discovery-2015-11-01/describeconfigurations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/describeconfigurations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeBatchDeleteConfigurationTask

DescribeContinuousExports
