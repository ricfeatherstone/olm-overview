package controllers

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	demov1beta1 "olm-overview/api/v1beta1"
	"time"
)

var _ = Describe("Noop Controller", func() {
	const (
		NoopName      = "test-noop"
		NoopNamespace = "default"
		NoopFoo       = "test"
		NoopBar       = "baz"

		timeout  = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When creating Noop", func() {
		It("Should be able to create", func() {
			By("Creating a new Noop")
			ctx := context.Background()
			noop := &demov1beta1.Noop{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name:      NoopName,
					Namespace: NoopNamespace,
				},
				Spec: demov1beta1.NoopSpec{
					Foo: NoopFoo,
				},
			}

			Expect(k8sClient.Create(ctx, noop)).Should(Succeed())

			createdNoop := &demov1beta1.Noop{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: NoopName, Namespace: NoopNamespace}, createdNoop)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())
		})

	})
})
