# PutGroupingConfiguration

Creates or updates the grouping configuration for this account. This operation allows you to define custom grouping attributes that determine how services are logically grouped based on telemetry attributes, AWS tags, or predefined mappings.
These grouping attributes can then be used to organize and filter services in the Application Signals console and APIs.

## Request Syntax

```nohighlight

PUT /grouping-configuration HTTP/1.1
Content-type: application/json

{
   "GroupingAttributeDefinitions": [
      {
         "DefaultGroupingValue": "string",
         "GroupingName": "string",
         "GroupingSourceKeys": [ "string" ]
      }
   ]
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[GroupingAttributeDefinitions](#API_PutGroupingConfiguration_RequestSyntax)**

An array of grouping attribute definitions that specify how services should be grouped. Each definition includes a friendly name, source keys to derive the grouping value from, and an optional default value.

Type: Array of [GroupingAttributeDefinition](api-groupingattributedefinition.md) objects

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "GroupingConfiguration": {
      "GroupingAttributeDefinitions": [
         {
            "DefaultGroupingValue": "string",
            "GroupingName": "string",
            "GroupingSourceKeys": [ "string" ]
         }
      ],
      "UpdatedAt": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[GroupingConfiguration](#API_PutGroupingConfiguration_ResponseSyntax)**

A structure containing the updated grouping configuration, including all grouping attribute definitions and the timestamp when it was last updated.

Type: [GroupingConfiguration](api-groupingconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 403

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

**ValidationException**

The resource is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/application-signals-2024-04-15/putgroupingconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/application-signals-2024-04-15/putgroupingconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListTagsForResource

StartDiscovery
