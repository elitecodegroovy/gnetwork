package pool_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/elitecodegroovy/gnetwork/apps/pattern/33-thread_pool/pool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConnPool(t *testing.T) {
	var connPool *pool.ConnPool
	connPool = pool.NewConnPool(&pool.Options{
		Dialer:             dummyDialer,
		PoolSize:           10,
		PoolTimeout:        time.Hour,
		IdleTimeout:        time.Millisecond,
		IdleCheckFrequency: time.Millisecond,
	})
	defer connPool.Close()

	// Reserve one connection.
	cn, err := connPool.Get()
	if err != nil {
		t.Fatal("failed to fetch a conn from the connection pool")
	}
	//t.Log(fmt.Sprintf("%v", cn))

	// Reserve all other connections.
	var cns []*pool.Conn
	for i := 0; i < 9; i++ {
		cn, err := connPool.Get()
		if err != nil {
			t.Fatal("failed to fetch a conn from the connection pool")

		}
		cns = append(cns, cn)
	}
	started := make(chan bool, 1)
	done := make(chan bool, 1)
	go func() {
		defer GinkgoRecover()

		started <- true
		_, err := connPool.Get()
		if err != nil {
			t.Fatal("failed to fetch a conn from the connection pool")
		}
		done <- true

		connPool.Put(cn)
	}()
	<-started

	// Check that Get is blocked.
	select {
	case <-done:
		t.Fatal("Get is not blocked")
	case <-time.After(time.Millisecond):
		// ok
	}

	connPool.Remove(cn, nil)

	// Check that Get is unblocked.
	select {
	case <-done:
		// ok
	case <-time.After(time.Second):
		t.Fatal("Get is not unblocked")
	}

	for _, cn := range cns {
		connPool.Put(cn)
	}
	t.Log("Idle connections: " + fmt.Sprintf("%d", connPool.Stats().IdleConns))
	t.Log("Total connections: " + fmt.Sprintf("%d", connPool.Stats().TotalConns))
	t.Log("Stale connections: " + fmt.Sprintf("%d", connPool.Stats().StaleConns))
}

func TestMinIdleConns(t *testing.T) {
	const poolSize = 100
	var minIdleConns int

	newConnPool := func() *pool.ConnPool {
		connPool := pool.NewConnPool(&pool.Options{
			Dialer:             dummyDialer,
			PoolSize:           poolSize,
			MinIdleConns:       minIdleConns,
			PoolTimeout:        100 * time.Millisecond,
			IdleTimeout:        -1,
			IdleCheckFrequency: -1,
		})
		assert.EqualValues(t, 0, connPool.Len())
		return connPool
	}

	dummyPool := newConnPool()
	defer dummyPool.Close()
	var cn *pool.Conn
	var err error
	_, err = dummyPool.Get()
	if err != nil {
		t.Fatal("failed to fetch a conn from the connection pool")
	}

	assert.EqualValues(t, 1, dummyPool.Len())

	cn, err = dummyPool.Get()
	if err != nil {
		t.Fatal("failed to fetch a conn from the connection pool")
	}
	assert.EqualValues(t, 2, dummyPool.Len())

	//subtract
	dummyPool.Remove(cn, nil)
	assert.EqualValues(t, 1, dummyPool.Len())
}

var _ = Describe("conns reaper", func() {
	const idleTimeout = time.Minute
	const maxAge = time.Hour

	var connPool *pool.ConnPool
	var conns, staleConns, closedConns []*pool.Conn

	assert := func(typ string) {
		BeforeEach(func() {
			closedConns = nil
			connPool = pool.NewConnPool(&pool.Options{
				Dialer:             dummyDialer,
				PoolSize:           10,
				IdleTimeout:        idleTimeout,
				MaxConnAge:         maxAge,
				PoolTimeout:        time.Second,
				IdleCheckFrequency: time.Hour,
				OnClose: func(cn *pool.Conn) error {
					closedConns = append(closedConns, cn)
					return nil
				},
			})

			conns = nil

			// add stale connections
			staleConns = nil
			for i := 0; i < 3; i++ {
				cn, err := connPool.Get()
				Expect(err).NotTo(HaveOccurred())
				switch typ {
				case "idle":
					cn.SetUsedAt(time.Now().Add(-2 * idleTimeout))
				case "aged":
					cn.SetCreatedAt(time.Now().Add(-2 * maxAge))
				}
				conns = append(conns, cn)
				staleConns = append(staleConns, cn)
			}

			// add fresh connections
			for i := 0; i < 3; i++ {
				cn, err := connPool.Get()
				Expect(err).NotTo(HaveOccurred())
				conns = append(conns, cn)
			}

			for _, cn := range conns {
				connPool.Put(cn)
			}

			Expect(connPool.Len()).To(Equal(6))
			Expect(connPool.IdleLen()).To(Equal(6))

			n, err := connPool.ReapStaleConns()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(3))
		})

		AfterEach(func() {
			_ = connPool.Close()
			Expect(connPool.Len()).To(Equal(0))
			Expect(connPool.IdleLen()).To(Equal(0))
			Expect(len(closedConns)).To(Equal(len(conns)))
			Expect(closedConns).To(ConsistOf(conns))
		})

		It("reaps stale connections", func() {
			Expect(connPool.Len()).To(Equal(3))
			Expect(connPool.IdleLen()).To(Equal(3))
		})

		It("does not reap fresh connections", func() {
			n, err := connPool.ReapStaleConns()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(0))
		})

		It("stale connections are closed", func() {
			Expect(len(closedConns)).To(Equal(len(staleConns)))
			Expect(closedConns).To(ConsistOf(staleConns))
		})

		It("pool is functional", func() {
			for j := 0; j < 3; j++ {
				var freeCns []*pool.Conn
				for i := 0; i < 3; i++ {
					cn, err := connPool.Get()
					Expect(err).NotTo(HaveOccurred())
					Expect(cn).NotTo(BeNil())
					freeCns = append(freeCns, cn)
				}

				Expect(connPool.Len()).To(Equal(3))
				Expect(connPool.IdleLen()).To(Equal(0))

				cn, err := connPool.Get()
				Expect(err).NotTo(HaveOccurred())
				Expect(cn).NotTo(BeNil())
				conns = append(conns, cn)

				Expect(connPool.Len()).To(Equal(4))
				Expect(connPool.IdleLen()).To(Equal(0))

				connPool.Remove(cn, nil)

				Expect(connPool.Len()).To(Equal(3))
				Expect(connPool.IdleLen()).To(Equal(0))

				for _, cn := range freeCns {
					connPool.Put(cn)
				}

				Expect(connPool.Len()).To(Equal(3))
				Expect(connPool.IdleLen()).To(Equal(3))
			}
		})
	}

	assert("idle")
	assert("aged")
})

var _ = Describe("race", func() {
	var connPool *pool.ConnPool
	var C, N int

	BeforeEach(func() {
		C, N = 10, 1000
		if testing.Short() {
			C = 4
			N = 100
		}
	})

	AfterEach(func() {
		connPool.Close()
	})

	It("does not happen on Get, Put, and Remove", func() {
		connPool = pool.NewConnPool(&pool.Options{
			Dialer:             dummyDialer,
			PoolSize:           10,
			PoolTimeout:        time.Minute,
			IdleTimeout:        time.Millisecond,
			IdleCheckFrequency: time.Millisecond,
		})

		perform(C, func(id int) {
			for i := 0; i < N; i++ {
				cn, err := connPool.Get()
				Expect(err).NotTo(HaveOccurred())
				if err == nil {
					connPool.Put(cn)
				}
			}
		}, func(id int) {
			for i := 0; i < N; i++ {
				cn, err := connPool.Get()
				Expect(err).NotTo(HaveOccurred())
				if err == nil {
					connPool.Remove(cn, nil)
				}
			}
		})
	})
})
