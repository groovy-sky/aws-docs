# Generate custom error responses

If an object that you’re serving through CloudFront is unavailable for some reason, your web
server typically returns a relevant HTTP status code to CloudFront to indicate this. For example,
if a viewer requests an invalid URL, your web server returns an HTTP 404 (Not Found) status
code to CloudFront, and then CloudFront returns that status code to the viewer. Instead of using this
default error response, you can create a custom one that CloudFront returns to the viewer.

If you configure CloudFront to return a custom error page for an HTTP status code but the custom
error page isn’t available, CloudFront returns to the viewer the status code that CloudFront received
from the origin that contains the custom error pages. For example, suppose your custom
origin returns a 500 status code and you have configured CloudFront to get a custom error page for
a 500 status code from an Amazon S3 bucket. However, someone accidentally deleted the custom
error page from your Amazon S3 bucket. CloudFront returns an HTTP 404 status code (Not Found) to the
viewer that requested the object.

When CloudFront returns a custom error page to a viewer, you pay the standard CloudFront charges for
the custom error page, not the charges for the requested object. For more information about
CloudFront charges, see [Amazon CloudFront\
Pricing](https://aws.amazon.com/cloudfront/pricing).

###### Topics

- [Configure error response behavior](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages-procedure.html)

- [Create a custom error page for specific HTTP status codes](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/creating-custom-error-pages.html)

- [Store objects and custom error pages in different locations](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages-different-locations.html)

- [Change response codes returned by CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages-response-code.html)

- [Control how long CloudFront caches errors](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages-expiration.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

How CloudFront processes HTTP 4xx and 5xx status codes from your origin

Configure error response behavior
