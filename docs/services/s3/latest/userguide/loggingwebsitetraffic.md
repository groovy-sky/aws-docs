# (Optional) Logging web traffic

You can optionally enable Amazon S3 server access logging for a bucket that is configured
as a static website. Server access logging provides detailed records for the requests
that are made to your bucket. For more information, see [Logging requests with server access logging](serverlogs.md). If you plan to use Amazon CloudFront to [speed up your website](https://docs.aws.amazon.com/AmazonS3/latest/userguide/website-hosting-cloudfront-walkthrough.html), you
can also use CloudFront logging. For more information, see [Configuring and Using\
Access Logs](../../../amazoncloudfront/latest/developerguide/accesslogs.md) in the _Amazon CloudFront Developer Guide_.

###### To enable server access logging for your static website bucket

01. Open the Amazon S3 console at
     [https://console.aws.amazon.com/s3/](https://console.aws.amazon.com/s3).

02. In the same Region where you created the bucket that is configured as a static
     website, create a general purpose bucket for logging, for example
     `logs.example.com`.

03. Create a folder for the server access logging log files (for example,
     `logs`).

04. (Optional) If you want to use CloudFront to improve your website performance, create
     a folder for the CloudFront log files (for example, `cdn`).

    For more information, see [Speeding up your website with Amazon CloudFront](https://docs.aws.amazon.com/AmazonS3/latest/userguide/website-hosting-cloudfront-walkthrough.html).

05. In the **Buckets** list, choose your bucket.

06. Choose **Properties**.

07. Under **Server access logging**, choose
     **Edit**.

08. Choose **Enable**.

09. Under the **Target bucket**, choose the bucket and folder
     destination for the server access logs:

- Browse to the folder and bucket location:

1. Choose **Browse S3**.

2. Choose the bucket name, and then choose the logs folder.

3. Choose **Choose path**.

- Enter the S3 bucket path, for example,
`s3://logs.example.com/logs/`.

10. Choose **Save changes**.

    In your log bucket, you can now access your logs. Amazon S3 writes website access
     logs to your log bucket every 2 hours.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Setting permissions for website access

Configuring a redirect
