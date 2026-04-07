# Library functions available for canary scripts

CloudWatch Synthetics includes several built-in classes and functions that you can call when
writing
Node.js scripts to use as canaries.

Some apply to both UI and API canaries. Others apply to UI canaries only. A UI canary is
a canary that uses the `getPage()` function and uses Puppeteer as a web driver to
navigate and interact with webpages.

###### Note

Whenever you upgrade a canary to use a new version of the the Synthetics runtime, all
Synthetics library functions that your canary uses are also automatically upgraded to the
same version of
NodeJS that the Synthetics runtime supports.

###### Topics

- [Library functions available for Node.js canary](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Library_function_Nodejs.html)

- [Library functions available for Java canary](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Java.html)

- [Library functions available for Node.js canary scripts using Playwright](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Nodejs_Playwright.html)

- [Library functions available for Node.js canary scripts using Puppeteer](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library_Nodejs.html)

- [Library functions available for Python canary scripts using Selenium](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library_Python.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Writing a JSON configuration for Node.js multi Checks blueprint

Library functions available for Node.js canary
