# Troubleshooting error response status codes in CloudFront

If CloudFront requests an object from your origin, and the origin returns an HTTP 4xx or 5xx
status code, there's a problem with communication between CloudFront and your origin.

This topic also includes troubleshooting steps for these status codes when using
Lambda@Edge or CloudFront Functions.

The following topics provide detailed explanations of the potential causes behind
these error responses and offers step-by-step guidance on how to diagnose and resolve
the underlying issues.

###### Topics

- [HTTP 400 status code (Bad Request)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-400-bad-request.html)

- [HTTP 401 status code (Unauthorized)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-401-unauthorized.html)

- [HTTP 403 status code (Invalid method)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-403-invalid-method.html)

- [HTTP 403 status code (Permission Denied)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-403-permission-denied.html)

- [HTTP 404 status code (Not Found)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-404-not-found.html)

- [HTTP 412 status code (Precondition Failed)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-412-precondition-failed.html)

- [HTTP 500 status code (Internal Server Error)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-500-internal-server-error.html)

- [HTTP 502 status code (Bad Gateway)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-502-bad-gateway.html)

- [HTTP 503 status code (Service Unavailable)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-503-service-unavailable.html)

- [HTTP 504 status code (Gateway Timeout)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/http-504-gateway-timeout.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Troubleshooting distribution issues

HTTP 400 status code (Bad Request)
