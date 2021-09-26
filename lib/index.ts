import * as cdk from '@aws-cdk/core';

export interface PasswordlessAuthProps {
  // Define construct properties here
}

export class PasswordlessAuth extends cdk.Construct {

  constructor(scope: cdk.Construct, id: string, props: PasswordlessAuthProps = {}) {
    super(scope, id);

    // Define construct contents here
  }
}
