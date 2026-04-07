# Customize at the edge with CloudFront Functions

With CloudFront Functions, you can write lightweight functions in JavaScript for high-scale,
latency-sensitive CDN customizations. Your functions can manipulate the requests and
responses that flow through CloudFront, perform basic authentication and authorization, generate
HTTP responses at the edge, and more. The CloudFront Functions runtime environment offers
submillisecond startup times, scales immediately to handle millions of requests per second,
and is highly secure. CloudFront Functions is a native feature of CloudFront, which means you can build,
test, and deploy your code entirely within CloudFront.

When you associate a CloudFront function with a CloudFront distribution, CloudFront intercepts requests and
responses at CloudFront edge locations and passes them to your function. You can invoke
CloudFront Functions when the following events occur:

- When CloudFront receives a request from a viewer (viewer request)

- Before CloudFront returns the response to the viewer (viewer response)

- During TLS connection establishment (connection request) - currently available for mutual TLS (mTLS) connections

For more information about CloudFront Functions, see the following topics:

###### Topics

- [Tutorial: Create a simple function with CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/functions-tutorial.html)

- [Tutorial: Create a CloudFront function that includes key values](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/functions-tutorial-kvs.html)

- [Write function code](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html)

- [Create functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/create-function.html)

- [Test functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/test-function.html)

- [Update functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/update-function.html)

- [Publish functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/publish-function.html)

- [Associate functions with distributions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/associate-function.html)

- [Amazon CloudFront KeyValueStore](kvs-with-functions.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Differences between CloudFront Functions and Lambda@Edge

Tutorial: Create a simple CloudFront
function
