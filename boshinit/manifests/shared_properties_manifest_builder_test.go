package manifests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf-experimental/bosh-bootloader/boshinit/manifests"
)

var _ = Describe("SharedPropertiesManifestBuilder", func() {
	var sharedPropertiesManifestBuilder *manifests.SharedPropertiesManifestBuilder

	BeforeEach(func() {
		sharedPropertiesManifestBuilder = manifests.NewSharedPropertiesManifestBuilder()
	})

	Describe("AWS", func() {
		It("returns job properties for AWS", func() {
			aws := sharedPropertiesManifestBuilder.AWS(manifests.ManifestProperties{
				ElasticIP:       "some-elastic-ip",
				AccessKeyID:     "some-access-key-id",
				SecretAccessKey: "some-secret-access-key",
				DefaultKeyName:  "some-key-name",
				Region:          "some-region",
				SecurityGroup:   "some-security-group",
			})

			Expect(aws).To(Equal(manifests.AWSProperties{
				AccessKeyId:           "some-access-key-id",
				SecretAccessKey:       "some-secret-access-key",
				DefaultKeyName:        "some-key-name",
				DefaultSecurityGroups: []string{"some-security-group"},
				Region:                "some-region",
			}))
		})
	})

	Describe("NTP", func() {
		It("returns job properties for NTP", func() {
			ntp := sharedPropertiesManifestBuilder.NTP()
			Expect(ntp).To(ConsistOf(
				[]string{"0.pool.ntp.org", "1.pool.ntp.org"},
			))
		})
	})
})
