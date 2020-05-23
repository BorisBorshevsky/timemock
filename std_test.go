package timemock

import (
	"time"

	mock_timemock "github.com/BorisBorshevsky/timemock/internal"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Std", func() {

	var mock *mock_timemock.MockClock
	var mockCtl *gomock.Controller

	BeforeEach(func() {
		mockCtl = gomock.NewController(GinkgoT())
		mock = mock_timemock.NewMockClock(mockCtl)
		std = mock
	})

	AfterEach(func() {
		mockCtl.Finish()
	})

	It("Now run clock.Now", func() {
		t := time.Now()
		mock.EXPECT().Now().Return(t)
		Ω(Now()).Should(Equal(t))
	})

	It("Since run clock.Since", func() {
		t := time.Now()
		dur := time.Duration(10)
		mock.EXPECT().Since(t).Return(dur)
		Ω(Since(t)).Should(Equal(dur))
	})

	It("Until run clock.Until", func() {
		t := time.Now()
		dur := time.Duration(10)
		mock.EXPECT().Until(t).Return(dur)
		Ω(Until(t)).Should(Equal(dur))
	})

	It("Freeze run clock.Freeze", func() {
		t := time.Now()
		mock.EXPECT().Freeze(t)
		Freeze(t)
	})

	It("Travel run clock.Travel", func() {
		t := time.Now()
		mock.EXPECT().Travel(t)
		Travel(t)
	})

	It("Scale run clock.Scale", func() {
		scale := float64(10)
		mock.EXPECT().Scale(scale)
		Scale(scale)
	})

	It("Now run clock.Return", func() {
		mock.EXPECT().Return()
		Return()
	})
})
