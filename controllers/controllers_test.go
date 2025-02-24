package controllers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("mergeMap", func() {
	When("the target map is populated", func() {
		var target map[string]string

		BeforeEach(func() {
			target = map[string]string{
				"test-key": "test-value",
			}
		})

		Context("and the desired map is populated with a new key/value pair", func() {
			var desired = map[string]string{
				"new-key": "new-value",
			}

			It("should add the new value to the target map", func() {
				Expect(mergeMap(target, desired)).To(Equal(true))
				Expect(target).To(Equal(map[string]string{
					"test-key": "test-value",
					"new-key":  "new-value",
				}))
			})
		})

		Context("and the desired map is populated with a new value for an existing key", func() {
			var desired = map[string]string{
				"test-key": "new-value",
			}

			It("should add the new value to the target map", func() {
				Expect(mergeMap(target, desired)).To(Equal(true))
				Expect(target).To(Equal(map[string]string{
					"test-key": "new-value",
				}))
			})
		})

		Context("and the desired map is populated with the same value for an existing key", func() {
			var desired = map[string]string{
				"test-key": "test-value",
			}

			It("should not change the target map", func() {
				Expect(mergeMap(target, desired)).To(Equal(false))
				Expect(target).To(Equal(map[string]string{
					"test-key": "test-value",
				}))
			})
		})

		Context("and the desired map is empty", func() {
			var desired = map[string]string{}

			It("should not change the target map", func() {
				Expect(mergeMap(target, desired)).To(Equal(false))
				Expect(target).To(Equal(map[string]string{
					"test-key": "test-value",
				}))
			})
		})

		Context("and the desired map is nil", func() {
			It("should not change the target map", func() {
				Expect(mergeMap(target, nil)).To(Equal(false))
				Expect(target).To(Equal(map[string]string{
					"test-key": "test-value",
				}))
			})
		})
	})

	When("the target map is empty", func() {
		var target map[string]string

		BeforeEach(func() {
			target = map[string]string{}
		})

		Context("and the desired map is populated with a new key/value pair", func() {
			var desired = map[string]string{
				"new-key": "new-value",
			}

			It("should add the new value to the target map", func() {
				Expect(mergeMap(target, desired)).To(Equal(true))
				Expect(target).To(Equal(map[string]string{
					"new-key": "new-value",
				}))
			})
		})

		Context("and the desired map is empty", func() {
			var desired = map[string]string{}

			It("should not change the target map", func() {
				Expect(mergeMap(target, desired)).To(Equal(false))
				Expect(target).To(Equal(map[string]string{}))
			})
		})

		Context("and the desired map is nil", func() {
			It("should not change the target map", func() {
				Expect(mergeMap(target, nil)).To(Equal(false))
				Expect(target).To(Equal(map[string]string{}))
			})
		})
	})

	When("the target map is nil", func() {
		var target map[string]string

		BeforeEach(func() {
			target = nil
		})

		Context("and the desired map is populated with a new key/value pair", func() {
			var desired = map[string]string{
				"new-key": "new-value",
			}

			It("should add the new value to the target map", func() {
				Expect(mergeMap(target, desired)).To(Equal(true))
				Expect(target).To(Equal(map[string]string{
					"new-key": "new-value",
				}))
			})
		})

		Context("and the desired map is empty", func() {
			var desired = map[string]string{
				"test-key": "new-value",
			}

			It("should not change the target map", func() {
				Expect(mergeMap(target, desired)).To(Equal(false))
				Expect(target).To(Equal(map[string]string{}))
			})
		})

		Context("and the desired map is nil", func() {
			It("should not change the target map", func() {
				Expect(mergeMap(target, nil)).To(Equal(false))
				Expect(target).To(BeNil())
			})
		})
	})
})
