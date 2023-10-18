locals {
  function_name         = "transactions"
  src_path_transactions = "${path.module}/../"
  binary_name           = local.function_name
  binary_path           = "${path.module}/tf_generated/${local.binary_name}"
  archive_path          = "${path.module}/tf_generated/${local.function_name}.zip"
  stori_email           = "example@gmail.com"
  stori_password        = "super secret pass"
  stori_bucket          = "bucket name"
  stori_file            = "csvfile"
}
