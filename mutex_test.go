package named

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mutexLocked = 1

func TestNamedMutex_Lock(t *testing.T) {
	m1 := Mutex("name1")
	m2 := Mutex("name2")

	i := 0
	m1.Lock()
	assert.True(t, MutexLocked(namedMutexes[m1]))
	i++
	m2.Lock()
	assert.True(t, MutexLocked(namedMutexes[m2]))
	i++

	assert.Equal(t, 2, i)

	m1.Unlock()
	i--
	m1.Unlock()
	i--

	assert.Equal(t, 0, i)
}

func MutexLocked(m *sync.Mutex) bool {
	state := reflect.ValueOf(m).Elem().FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}
