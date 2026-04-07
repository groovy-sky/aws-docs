# Synthetics runtime versions

When you create or update a canary, you choose a Synthetics runtime version for the
canary.
A Synthetics runtime is a combination of the Synthetics code that calls your script handler,
and the Lambda layers of bundled dependencies.

CloudWatch Synthetics currently supports runtimes that use Node.js, Python, or Java languages.
The frameworks supported are Puppeteer, Playwright, and Selenium.

We recommend that you always use the most recent runtime version for your canaries, to
be
able to use the latest features and updates made to the Synthetics library.

**Please note**: whenever you run a canary to use the new
version of the Synthetics runtime, all Synthetics library functions that your canary uses
are also automatically moved to the same version of NodeJS that the Synthetics runtime
supports.

###### Topics

- [Runtime versions using Java](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_Java.html)

- [Runtime versions using Node.js and Playwright](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_nodejs_playwright.html)

- [Runtime versions using Node.js and Puppeteer](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_nodejs_puppeteer.html)

- [Runtime versions using Python and Selenium Webdriver](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_python_selenium.html)

- [Runtime versions using Node.js](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_Nodejs.html)

- [Runtime versions support policy](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Runtime_Support_Policy.html)

- [Runtime versions update](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Runtime_Version_Update.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using the CloudWatch Synthetics Recorder for Google Chrome

Runtime versions using Java
