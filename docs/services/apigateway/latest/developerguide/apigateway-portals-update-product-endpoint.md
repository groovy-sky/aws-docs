# Update a product REST endpoint in API Gateway

When you update a product REST endpoint, you can change all the settings of the product endpoint except for the path
and method of a REST API and the stage it's deployed to. Any changes to your product endpoints are updated
automatically, but you need to republish any portals that use your products to see the changes reflected in a
portal.

## Update a product REST endpoint

The following procedure shows how to update a product REST endpoint to overwrite the current API documentation with
custom API reference documentation.

###### To update a product REST endpoint

01. Sign in to the API Gateway console at [https://console.aws.amazon.com/apigateway](https://console.aws.amazon.com/apigateway).

02. In the main navigation pane, choose
     **Portal products**.

03. Choose a product.

04. In the **Documentation** tab, under **API reference pages**, choose the name of a product REST endpoint, such as
     **/dogs - GET**.

05. For **Preview**, choose **Edit page**.

06. For **Documentation source**, turn on **Override the documentation**.

07. For **Display content**, choose
     **Override the existing content**.

    If you choose **Remove all content**, the content is removed from the page, but you can
     choose **Override the existing content** to access the documentation again.

08. In **Page body**, enter custom API reference documentation using [GitHub Flavored Markdown](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax).

    This documentation will not propagate back to API Gateway.

09. Choose **Save changes**.

10. You must republish any portals that use this portal product for the changes to take effect.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Update a portal product

Update product page

All content copied from https://docs.aws.amazon.com/.
