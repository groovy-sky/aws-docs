# GetAccountBalance

## Description

The `GetAccountBalance` operation retrieves the Prepaid HITs balance in your Amazon Mechanical Turk account if you are a Prepaid Requester. Alternatively, this operation will retrieve the remaining available AWS Billing usage if you have enabled AWS Billing.

Note: If you have enabled AWS Billing and still have a remaining Prepaid HITs balance, this balance can be viewed on the My Account page in the Requester console.

## Request Syntax

```

{  }
```

## Request Parameters

The request accepts the following data in JSON format:

## Response Elements

A successful request returns a string representing your available balance details in US Dollars.

## Example

The following example shows how to use the `GetAccountBalance` operation:

### Sample Request

The following makes a GetAccountBalance request.

```

POST / HTTP/1.1
Host: mturk-requester.us-east-1.amazonaws.com
Content-Length: <PayloadSizeBytes>
X-Amz-Date: <Date>
{
}

```

### Sample Response

The following is an example response:

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
  AvailableBalance:10000.00
}

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DisassociateQualificationFromWorker

GetAssignment
