# Making requests using the AWS SDKs

###### Topics

- [Making requests using AWS account or IAM user credentials](authusingacctorusercredentials.md)

- [Making requests using IAM user temporary credentials](authusingtempsessiontoken.md)

- [Making requests using federated user temporary credentials](authusingtempfederationtoken.md)

You can send authenticated requests to Amazon S3 using either the AWS SDK or by making the REST
API calls directly in your application. The AWS SDK API uses the credentials that you
provide to compute the signature for authentication. If you use the REST API directly in
your applications, you must write the necessary code to compute the signature for
authenticating your request. For a list of available AWS SDKs go to, [Sample Code and Libraries](https://aws.amazon.com/code).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using dual-stack endpoints

Using AWS account or IAM user credentials

All content copied from https://docs.aws.amazon.com/.
