# Transactions

- [Transactions](#transactions)
  - [Deployment](#deployment)
    - [Variables](#variables)
  - [Specifications](#specifications)
    - [Architecture](#architecture)
    - [Pattern design](#pattern-design)
    - [Terraform](#terraform)
    - [Endpoints](#endpoints)
      - [/transactions](#transactions-1)
  - [Missing](#missing)
    - [testing](#testing)
    - [mailer](#mailer)


## Deployment

### Variables

In order to deploy the application you should define some variables that the project needs to run

```tf
  stori_email           = "example@gmail.com" //sender email
  stori_password        = "super secret pass" //sender password, your accont should be configured to allow that this works, usually it depends on a security configuration 
  stori_bucket          = "bucket name"  // the name for the bucket that will be created
  stori_file            = "csvfile" //a initial csv file
```

Those variables are locate in /terraform/local.tf

Additionally you should have installed [Terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli) in your local machine and an AWS account configure with aws cli.

```sh
aws configure
```

When you have installed Terraform move to the root of this project  and run the followings commands

``` sh
terraform init
terraform apply
```

Then type `yes` to apply the configuration and when it finished it will print the url to make the requests

## Specifications

### Architecture

To this project clean architecture was used aiming to keep easy make changes to it.

### Pattern design

For this project was used the pattern design `Dependency Injection` in order to make easy maintain the application and develop  tests.

### Terraform

Terraform was used in this project to create the cloud architecture and deploy id easily with a single command.

### Endpoints

There are two endpoints, one to create and order and other to make a payment and complete the order.

#### /transactions

This end point receive a `POST` request with this body

```json
{
    "to":string,
    "fileName":string,
}
```

This will execute the program and send an email with the balance information to the user set in the body. The balance will be make using the file specified.

## Missing

### testing

### mailer

The sender email was tested just with gmail.
