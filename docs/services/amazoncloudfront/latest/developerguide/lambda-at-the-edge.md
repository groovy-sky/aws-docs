# Customize at the edge with Lambda@Edge

Lambda@Edge is an extension of AWS Lambda. Lambda@Edge is a compute service that lets you
execute functions that customize the content that Amazon CloudFront delivers. You can author Node.js
or Python functions in the Lambda console in one AWS Region, US East (N. Virginia).

After you create the function, you can add triggers using the Lambda console or CloudFront
console so that the functions run in AWS locations that are closer to the viewer, without
provisioning or managing servers. Optionally, you can use Lambda and CloudFront API operations to
set up your functions and triggers programmatically.

Lambda@Edge scales automatically, from a few requests per day to thousands per second.
Processing requests at AWS locations closer to the viewer instead of on origin servers
significantly reduces latency and improves the user experience.

###### Note

Lambda@Edge isn't supported with gRPC requests. For more information. see [Using gRPC with CloudFront distributions](distribution-using-grpc.md).

###### Topics

- [How Lambda@Edge works with requests and responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-event-request-response.html)

- [Ways to use Lambda@Edge](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-ways-to-use.html)

- [Get started with Lambda@Edge functions (console)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-how-it-works.html)

- [Set up IAM permissions and roles for Lambda@Edge](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-permissions.html)

- [Write and create a Lambda@Edge function](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-create-function.html)

- [Add triggers for a Lambda@Edge function](lambda-edge-add-triggers.md)

- [Test and debug Lambda@Edge functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-testing-debugging.html)

- [Delete Lambda@Edge functions and replicas](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-delete-replicas.html)

- [Lambda@Edge event structure](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-event-structure.html)

- [Work with requests and responses](lambda-generating-http-responses.md)

- [Lambda@Edge example functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-examples.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Advanced revocation scenarios

How Lambda@Edge works with requests and responses
