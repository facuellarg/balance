// allow lambda service to assume (use) the role with such policy
data "aws_iam_policy_document" "assume_lambda_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

// create lambda role, that lambda function can assume (use)
resource "aws_iam_role" "lambda" {
  name               = "AssumeLambdaRole"
  description        = "Role for lambda to assume lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_lambda_role.json
}

//create document for policy that allow lambda logging
data "aws_iam_policy_document" "allow_lambda_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:*:*:*",
    ]
  }
}

// create a policy to allow writing into logs and create logs stream
resource "aws_iam_policy" "function_logging_policy" {
  name        = "AllowLambdaLoggingPolicy"
  description = "Policy for lambda cloudwatch logging"
  policy      = data.aws_iam_policy_document.allow_lambda_logging.json
}

// attach policy to out created lambda role
resource "aws_iam_role_policy_attachment" "lambda_logging_policy_attachment" {
  role       = aws_iam_role.lambda.id
  policy_arn = aws_iam_policy.function_logging_policy.arn
}

//create document to allow lambda to operate on dynamodb table
data "aws_iam_policy_document" "allow_dynamodb_table_operations" {
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
    ]

    resources = [
      aws_dynamodb_table.transactions.arn,
    ]
  }
}

//create policy to allow lambda to operate on dynamodb table
resource "aws_iam_policy" "dynamodb_lambda_policy" {
  name        = "DynamoDBLambdaPolicy"
  description = "Policy for lambda to operate on dynamodb table"
  policy      = data.aws_iam_policy_document.allow_dynamodb_table_operations.json
}


// attach policy to out created lambda role
resource "aws_iam_role_policy_attachment" "lambda_dynamodb_policy_attachment" {
  role       = aws_iam_role.lambda.id
  policy_arn = aws_iam_policy.dynamodb_lambda_policy.arn
  depends_on = [
    aws_iam_role.lambda
  ]
}


data "aws_iam_policy_document" "get_object_s3" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
    ]

    resources = [
      "${aws_s3_bucket.transactions.arn}/*",
    ]
  }
}

resource "aws_iam_policy" "s3_policy" {
  name        = "S3_Policy"
  description = "Allow lambda read from s3"
  policy      = data.aws_iam_policy_document.get_object_s3.json
}

resource "aws_iam_role_policy_attachment" "order_queue_policy_attachment" {
  role       = aws_iam_role.lambda.id
  policy_arn = aws_iam_policy.s3_policy.arn
  depends_on = [
    aws_iam_role.lambda
  ]
}
