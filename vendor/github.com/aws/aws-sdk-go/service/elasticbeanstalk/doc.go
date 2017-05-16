// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticbeanstalk provides the client and types for making API
// requests to AWS Elastic Beanstalk.
//
// AWS Elastic Beanstalk makes it easy for you to create, deploy, and manage
// scalable, fault-tolerant applications running on the Amazon Web Services
// cloud.
//
// For more information about this product, go to the AWS Elastic Beanstalk
// (http://aws.amazon.com/elasticbeanstalk/) details page. The location of the
// latest AWS Elastic Beanstalk WSDL is http://elasticbeanstalk.s3.amazonaws.com/doc/2010-12-01/AWSElasticBeanstalk.wsdl
// (http://elasticbeanstalk.s3.amazonaws.com/doc/2010-12-01/AWSElasticBeanstalk.wsdl).
// To install the Software Development Kits (SDKs), Integrated Development Environment
// (IDE) Toolkits, and command line tools that enable you to access the API,
// go to Tools for Amazon Web Services (http://aws.amazon.com/tools/).
//
// Endpoints
//
// For a list of region-specific endpoints that AWS Elastic Beanstalk supports,
// go to Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region)
// in the Amazon Web Services Glossary.
//
// See https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01 for more information on this service.
//
// See elasticbeanstalk package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/elasticbeanstalk/
//
// Using the Client
//
// To use the client for AWS Elastic Beanstalk you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := elasticbeanstalk.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Elastic Beanstalk client ElasticBeanstalk for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/elasticbeanstalk/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.AbortEnvironmentUpdate(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("AbortEnvironmentUpdate result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.AbortEnvironmentUpdateWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package elasticbeanstalk
