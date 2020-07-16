# goibcp
Go lang wrapper for interactive brokers Client Portal (CP) web gateway



FAQs
Could not connect to IB CP Gateway, Ensure the CP gateway is running and logged in before connecting
- Ensure the client portal gateway is running at the configured host and destination
- Ensure you have logged in to the portal 
- default link for the portal will be https://localhost:5000/.
- Follow the getting started instructions at https://interactivebrokers.github.io/cpwebapi/index.html

X509 Certificate error 
Please refer to the interactive brokers CP Web gateway FAQq in common question section at 
https://interactivebrokers.github.io/cpwebapi/faq.html
```Since the gateway is running on your premises the certificate needs to be created/self-signed by you, or officially signed by a 3rd party. The gateway is similar to another webserver such as Tomcat which doesn’t provide a certificate along with the release.```
https://www.sslshopper.com/article-how-to-create-a-self-signed-certificate-using-java-keytool.html
