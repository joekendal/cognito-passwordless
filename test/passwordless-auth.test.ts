import { expect as expectCDK, countResources, haveResource, haveResourceLike, Capture } from '@aws-cdk/assert';
import * as cdk from '@aws-cdk/core';
import * as PasswordlessAuth from '../lib/index';

/**
 * Cognito user pool
 */
describe('AWS::Cognito::UserPool', () => {
  let app: cdk.App;
  let stack: cdk.Stack;

  // Setup
  beforeAll(() => {
    app = new cdk.App();
    stack = new cdk.Stack(app, "TestStack")
    new PasswordlessAuth.PasswordlessAuth(stack, 'MyTestConstruct')
  })

  test('UserPool Created', () => {
    expectCDK(stack).to(countResources("AWS::Cognito::UserPool", 1));
  });

  test('Signup attributes = given_name, family_name, phone_number', () => {
    expectCDK(stack).to(haveResource('AWS::Cognito::UserPool', {
      Schema: [
        { Mutable: true, Name: "phone_number", Required: true },
        { Mutable: true, Name: "given_name", Required: true },
        { Mutable: true, Name: "family_name", Required: true }
      ]
    }))
  })

  test('Password policy', () => {
    expectCDK(stack).to(haveResource('AWS::Cognito::UserPool', {
      Policies: {
        PasswordPolicy: {
          MinimumLength: 6,
          RequireLowercase: false,
          RequireNumbers: false,
          RequireSymbols: false,
          RequireUppercase: false,
        }
      }
    }))
  })

  test('Username attributes = phone_number', () => {
    expectCDK(stack).to(haveResource('AWS::Cognito::UserPool', {
      UsernameAttributes: ['phone_number']
    }))
  })

  test('Mfa off', () => {
    expectCDK(stack).to(haveResource('AWS::Cognito::UserPool', {
      MfaConfiguration: "OFF"
    }))
  })

  // test('Lambda triggers', () => {
  //   const preSignUpArn = Capture.aString()
  //   expectCDK(stack).to(haveResourceLike('AWS::Cognito::UserPool', {
  //     LambdaConfig: {
  //       PreSignUp: preSignUpArn.capture()
  //     }
  //   }))
  // })

})

test('AWS::Cognito::UserPoolClient', () => {
  const app = new cdk.App();
  const stack = new cdk.Stack(app, "TestStack")
  new PasswordlessAuth.PasswordlessAuth(stack, 'MyTestConstruct')

  expectCDK(stack).to(haveResource('AWS::Cognito::UserPoolClient', {
    ClientName: 'sms-auth-client',
    GenerateSecret: false,
    ExplicitAuthFlows: ["ALLOW_CUSTOM_AUTH", "ALLOW_REFRESH_TOKEN_AUTH"]
  }))
})


