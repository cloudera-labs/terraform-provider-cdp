## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "terraform.cloudera.com/cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_profile = "default"
}

data "cdp_environments_aws_credential_prerequisites" "example" {}

output "cdp_environments_aws_credential_prerequisites" {
  value = data.cdp_environments_aws_credential_prerequisites.example
}
