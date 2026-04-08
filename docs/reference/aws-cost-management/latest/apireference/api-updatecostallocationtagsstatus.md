# UpdateCostAllocationTagsStatus

Updates status for cost allocation tags in bulk, with maximum batch size of 20. If the tag
status that's updated is the same as the existing tag status, the request doesn't fail.
Instead, it doesn't have any effect on the tag status (for example, activating the active
tag).

## Request Syntax

```nohighlight

{
   "CostAllocationTagsStatus": [
      {
         "Status": "string",
         "TagKey": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CostAllocationTagsStatus](#API_UpdateCostAllocationTagsStatus_RequestSyntax)**

The list of `CostAllocationTagStatusEntry` objects that are used to update cost
allocation tags status for this request.

Type: Array of [CostAllocationTagStatusEntry](api-costallocationtagstatusentry.md) objects

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Required: Yes

## Response Syntax

```nohighlight

{
   "Errors": [
      {
         "Code": "string",
         "Message": "string",
         "TagKey": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Errors](#API_UpdateCostAllocationTagsStatus_ResponseSyntax)**

A list of `UpdateCostAllocationTagsStatusError` objects with error details
about each cost allocation tag that can't be updated. If there's no failure, an empty array
returns.

Type: Array of [UpdateCostAllocationTagsStatusError](api-updatecostallocationtagsstatuserror.md) objects

Array Members: Minimum number of 0 items. Maximum number of 20 items.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**LimitExceededException**

You made too many calls in a short period of time. Try again later.

HTTP Status Code: 400

## Examples

The following are sample requests and responses of the
`UpdateCostAllocationTagsStatus` operations.

### Example 1: Successfully updated all tag status

This example illustrates one usage of UpdateCostAllocationTagsStatus.

#### Sample Request

```

{
    "CostAllocationTagsStatus": [
       {
           "TagKey": "tagA",
           "Status": "Active"
       },
       {
           "TagKey": "tagB",
           "Status": "Inactive"
       }
    ]
}
```

#### Sample Response

```

{
    "Errors": []
}
```

### Example 2: Failed to update one of the tag statuses

This example illustrates one usage of UpdateCostAllocationTagsStatus.

#### Sample Request

```

{
    "CostAllocationTagsStatus": [
       {
           "TagKey": "tagC",
           "Status": "Active"
       },
       {
           "TagKey": "tagD",
           "Status": "Inactive"
       }
    ]
}
```

#### Sample Response

```

{
    "Errors": [
        {
            "TagKey": "tagC",
            "Code": "TagKeysNotFoundException",
            "Message": "Failed to update Cost Allocation Tag: tagC"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ce-2017-10-25/updatecostallocationtagsstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/updatecostallocationtagsstatus.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateAnomalySubscription

UpdateCostCategoryDefinition
