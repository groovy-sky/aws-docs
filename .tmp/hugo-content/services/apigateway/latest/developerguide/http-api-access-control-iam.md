# Control access to HTTP APIs with IAM authorization in API Gateway

You can enable IAM authorization for HTTP API routes. When IAM authorization is
enabled, clients must use
[Signature Version 4\
(SigV4)](../../../iam/latest/userguide/reference-sigv.md) to sign
their requests with AWS credentials. API Gateway invokes your API route only if the client has
`execute-api` permission for the route.

IAM authorization for HTTP APIs is similar to that for [REST\
APIs](api-gateway-control-access-using-iam-policies-to-invoke-api.md).

###### Note

Resource policies aren't currently supported for HTTP APIs.

For examples of IAM policies that grant clients the permission to invoke APIs, see [Control access for invoking an API](api-gateway-control-access-using-iam-policies-to-invoke-api.md).

## Enable IAM authorization for a route

The following [update-route](../../../cli/latest/reference/apigatewayv2/update-route.md) command enables
IAM authorization for an HTTP API route:

```nohighlight

aws apigatewayv2 update-route \
    --api-id abc123 \
    --route-id abcdef \
    --authorization-type AWS_IAM
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JWT authorizers

Integrations

All content copied from https://docs.aws.amazon.com/.
