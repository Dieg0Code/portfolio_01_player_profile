package terraform

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func SkipTestTerraform(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",
	})

	defer terraform.Destroy(t, terraformOptions)

	// Inicializa y aplica las configuraciones de Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Obtén los valores de salida de Terraform
	instanceID := terraform.Output(t, terraformOptions, "instance_id")
	instancePublicIP := terraform.Output(t, terraformOptions, "instance_public_ip")
	dbInstanceEndpoint := terraform.Output(t, terraformOptions, "db_instance_endpoint")
	s3BucketName := terraform.Output(t, terraformOptions, "s3_bucket_name")
	loadBalancerDNSName := terraform.Output(t, terraformOptions, "load_balancer_dns_name")

	// Verifica que los recursos se hayan creado correctamente
	assert.NotEmpty(t, instanceID)
	assert.NotEmpty(t, instancePublicIP)
	assert.NotEmpty(t, dbInstanceEndpoint)
	assert.NotEmpty(t, s3BucketName)
	assert.NotEmpty(t, loadBalancerDNSName)

}
