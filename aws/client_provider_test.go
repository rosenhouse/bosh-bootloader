package aws_test

import (
	goaws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	awscloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pivotal-cf-experimental/bosh-bootloader/aws"
	"github.com/pivotal-cf-experimental/bosh-bootloader/aws/cloudformation"
	"github.com/pivotal-cf-experimental/bosh-bootloader/aws/ec2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClientProvider", func() {
	var provider aws.ClientProvider

	BeforeEach(func() {
		provider = aws.NewClientProvider()
	})

	Describe("CloudFormationClient", func() {
		It("returns a Client with the provided configuration", func() {
			client, err := provider.CloudFormationClient(aws.Config{
				AccessKeyID:      "some-access-key-id",
				SecretAccessKey:  "some-secret-access-key",
				Region:           "some-region",
				EndpointOverride: "some-endpoint-override",
			})
			Expect(err).NotTo(HaveOccurred())

			_, ok := client.(cloudformation.Client)
			Expect(ok).To(BeTrue())

			cloudFormationClient, ok := client.(*awscloudformation.CloudFormation)
			Expect(ok).To(BeTrue())

			Expect(cloudFormationClient.Config.Credentials).To(Equal(credentials.NewStaticCredentials("some-access-key-id", "some-secret-access-key", "")))
			Expect(cloudFormationClient.Config.Region).To(Equal(goaws.String("some-region")))
			Expect(cloudFormationClient.Config.Endpoint).To(Equal(goaws.String("some-endpoint-override")))
		})

		Context("failure cases", func() {
			It("returns an error when the credentials are not provided", func() {
				_, err := provider.CloudFormationClient(aws.Config{})
				Expect(err).To(MatchError("--aws-access-key-id must be provided"))
			})
		})
	})

	Describe("EC2Client", func() {
		It("returns a EC2Client with the provided configuration", func() {
			client, err := provider.EC2Client(aws.Config{
				AccessKeyID:      "some-access-key-id",
				SecretAccessKey:  "some-secret-access-key",
				Region:           "some-region",
				EndpointOverride: "some-endpoint-override",
			})
			Expect(err).NotTo(HaveOccurred())

			_, ok := client.(ec2.Client)
			Expect(ok).To(BeTrue())

			ec2Client, ok := client.(*awsec2.EC2)
			Expect(ok).To(BeTrue())

			Expect(ec2Client.Config.Credentials).To(Equal(credentials.NewStaticCredentials("some-access-key-id", "some-secret-access-key", "")))
			Expect(ec2Client.Config.Region).To(Equal(goaws.String("some-region")))
			Expect(ec2Client.Config.Endpoint).To(Equal(goaws.String("some-endpoint-override")))
		})

		Context("failure cases", func() {
			It("returns an error when the credentials are not provided", func() {
				_, err := provider.EC2Client(aws.Config{})
				Expect(err).To(MatchError("--aws-access-key-id must be provided"))
			})
		})
	})
})
